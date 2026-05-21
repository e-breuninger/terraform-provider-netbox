package netbox

import "github.com/go-openapi/strfmt"

const (
	// DefaultPageSize balances API response time, request count, and memory usage
	DefaultPageSize = 100
	// FetchAll instructs the pagination helper to fetch all results (no user-imposed limit)
	FetchAll int64 = 0
)

// PaginatedListHelper manages automatic pagination for NetBox API list endpoints
type PaginatedListHelper struct {
	userLimit int64 // Maximum results to return (0 = fetch all)
	pageSize  int64 // Items per API request
	offset    int64 // Current offset, advanced by Advance()
}

// NewPaginationHelper creates a helper with the given user limit
func NewPaginationHelper(userLimit int64) *PaginatedListHelper {
	return &PaginatedListHelper{
		userLimit: userLimit,
		pageSize:  DefaultPageSize,
	}
}

// CurrentOffset returns the offset to use for the next API request
func (h *PaginatedListHelper) CurrentOffset() int64 {
	return h.offset
}

// Advance moves the offset forward by the number of items returned in the last page.
// Must be called with the actual returned count, not the requested page size, to correctly
// handle servers that cap responses below the requested limit (NetBox MAX_PAGE_SIZE).
func (h *PaginatedListHelper) Advance(returned int64) {
	h.offset += returned
}

// ShouldContinuePaging determines if another page should be fetched
func (h *PaginatedListHelper) ShouldContinuePaging(currentCount int64, next *strfmt.URI) bool {
	if h.userLimit > 0 && currentCount >= h.userLimit {
		return false
	}
	return next != nil && *next != ""
}

// TrimToLimit trims results to user's limit if specified
func (h *PaginatedListHelper) TrimToLimit(count int) int {
	if h.userLimit > 0 && int64(count) > h.userLimit {
		return int(h.userLimit)
	}
	return count
}

// GetPageSize returns the configured page size
func (h *PaginatedListHelper) GetPageSize() int64 {
	return h.pageSize
}
