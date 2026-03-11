package netbox

import (
	"strconv"
	"testing"

	"github.com/go-openapi/strfmt"
)

func TestPaginationHelper_ShouldContinuePaging(t *testing.T) {
	tests := []struct {
		name         string
		userLimit    int64
		currentCount int64
		next         *strfmt.URI
		want         bool
	}{
		{
			name:         "no limit, has next",
			userLimit:    0,
			currentCount: 50,
			next:         uriPtr("http://example.com/api?offset=50"),
			want:         true,
		},
		{
			name:         "no limit, no next",
			userLimit:    0,
			currentCount: 50,
			next:         nil,
			want:         false,
		},
		{
			name:         "no limit, empty next",
			userLimit:    0,
			currentCount: 50,
			next:         uriPtr(""),
			want:         false,
		},
		{
			name:         "at limit, has next",
			userLimit:    100,
			currentCount: 100,
			next:         uriPtr("http://example.com/api?offset=100"),
			want:         false,
		},
		{
			name:         "under limit, has next",
			userLimit:    100,
			currentCount: 50,
			next:         uriPtr("http://example.com/api?offset=50"),
			want:         true,
		},
		{
			name:         "over limit, has next",
			userLimit:    100,
			currentCount: 150,
			next:         uriPtr("http://example.com/api?offset=150"),
			want:         false,
		},
		{
			name:         "under limit, no next",
			userLimit:    100,
			currentCount: 50,
			next:         nil,
			want:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helper := NewPaginationHelper(tt.userLimit)
			got := helper.ShouldContinuePaging(tt.currentCount, tt.next)
			if got != tt.want {
				t.Errorf("ShouldContinuePaging() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaginationHelper_TrimToLimit(t *testing.T) {
	tests := []struct {
		name      string
		userLimit int64
		count     int
		want      int
	}{
		{
			name:      "no limit, no trim",
			userLimit: 0,
			count:     150,
			want:      150,
		},
		{
			name:      "under limit, no trim",
			userLimit: 100,
			count:     50,
			want:      50,
		},
		{
			name:      "at limit, no trim",
			userLimit: 100,
			count:     100,
			want:      100,
		},
		{
			name:      "over limit, trim to limit",
			userLimit: 100,
			count:     150,
			want:      100,
		},
		{
			name:      "way over limit, trim to limit",
			userLimit: 50,
			count:     500,
			want:      50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helper := NewPaginationHelper(tt.userLimit)
			got := helper.TrimToLimit(tt.count)
			if got != tt.want {
				t.Errorf("TrimToLimit() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestPaginationHelper_GetPageSize(t *testing.T) {
	helper := NewPaginationHelper(0)
	if got := helper.GetPageSize(); got != DefaultPageSize {
		t.Errorf("GetPageSize() = %d, want %d", got, DefaultPageSize)
	}
}

// TestPaginationHelper_ServerCapOffset verifies correct offset advancement when a
// NetBox server has MAX_PAGE_SIZE set below DefaultPageSize (100). In that case the
// server returns fewer items per page than requested, so the offset must advance by
// len(results) — not by the requested page size — to avoid skipping items.
//
// Example: MAX_PAGE_SIZE=50, DefaultPageSize=100, total=150 items
//
// Correct (offset += len(results)=50):
//
//	req offset=0   → server returns 50 items (0–49),  offset becomes 50
//	req offset=50  → server returns 50 items (50–99),  offset becomes 100
//	req offset=100 → server returns 50 items (100–149), done
//	collected: 150 ✓
//
// Incorrect (offset += pageSize=100):
//
//	req offset=0   → server returns 50 items (0–49),  offset becomes 100
//	req offset=100 → server returns 50 items (100–149), done
//	collected: 100 — items 50–99 silently skipped
func TestPaginationHelper_ServerCapOffset(t *testing.T) {
	// simulatePage mimics NetBox's OptionalLimitOffsetPagination:
	// returns min(serverCap, requestedLimit, remaining) items and whether a next page exists.
	simulatePage := func(totalItems, serverCap, offset, requestedLimit int) (returned int, next *strfmt.URI) {
		actualLimit := requestedLimit
		if serverCap < actualLimit {
			actualLimit = serverCap
		}
		remaining := totalItems - offset
		if remaining <= 0 {
			return 0, nil
		}
		returned = actualLimit
		if remaining < returned {
			returned = remaining
		}
		if offset+returned < totalItems {
			u := strfmt.URI("http://netbox/api/?offset=" + strconv.Itoa(offset+returned))
			return returned, &u
		}
		return returned, nil
	}

	tests := []struct {
		name          string
		totalItems    int
		serverCap     int // simulates NetBox MAX_PAGE_SIZE
		wantCollected int
	}{
		{
			// 50+50+30 — last page is partial
			name:          "server cap below page size",
			totalItems:    130,
			serverCap:     50,
			wantCollected: 130,
		},
		{
			// 100+100+10 — last page is partial
			name:          "server cap equal to page size",
			totalItems:    210,
			serverCap:     100,
			wantCollected: 210,
		},
		{
			// 100+30 — last page is partial
			name:          "server cap above page size",
			totalItems:    130,
			serverCap:     1000,
			wantCollected: 130,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pageSize := int(DefaultPageSize)
			helper := NewPaginationHelper(FetchAll)
			collected := 0
			for {
				returned, next := simulatePage(tt.totalItems, tt.serverCap, int(helper.CurrentOffset()), pageSize)
				collected += returned
				if returned == 0 {
					break
				}
				if !helper.ShouldContinuePaging(int64(collected), next) {
					break
				}
				helper.Advance(int64(returned))
			}
			if collected != tt.wantCollected {
				t.Errorf("collected %d, want %d", collected, tt.wantCollected)
			}
		})
	}
}

// Helper function to create URI pointers for tests
func uriPtr(s string) *strfmt.URI {
	uri := strfmt.URI(s)
	return &uri
}
