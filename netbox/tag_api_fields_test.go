package netbox

import (
	"io"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type tagTestResponse struct {
	code int
	body string
}

func (r tagTestResponse) Code() int {
	return r.code
}

func (r tagTestResponse) Message() string {
	return ""
}

func (r tagTestResponse) GetHeader(string) string {
	return ""
}

func (r tagTestResponse) GetHeaders(string) []string {
	return nil
}

func (r tagTestResponse) Body() io.ReadCloser {
	return io.NopCloser(strings.NewReader(r.body))
}

func TestTagReadReaderCapturesAPIFields(t *testing.T) {
	fields := tagAPIFields{}
	reader := tagReadReader{
		inner:  unusedTagResponseReader(t),
		fields: &fields,
	}
	response := tagTestResponse{
		code: 200,
		body: `{
			"id": 42,
			"name": "Production",
			"slug": "production",
			"description": "Production resources",
			"object_types": ["dcim.device", "virtualization.virtualmachine"],
			"weight": 2000
		}`,
	}

	result, err := reader.ReadResponse(response, runtime.JSONConsumer())

	require.NoError(t, err)
	payload := result.(*extras.ExtrasTagsReadOK).GetPayload()
	assert.Equal(t, int64(42), payload.ID)
	assert.Equal(t, "Production", *payload.Name)
	assert.Equal(
		t,
		[]string{"dcim.device", "virtualization.virtualmachine"},
		fields.ObjectTypes,
	)
	assert.Equal(t, int64(2000), fields.Weight)
}

func TestTagListReaderCapturesAPIFields(t *testing.T) {
	fieldsByID := make(map[int64]tagAPIFields)
	reader := tagListReader{
		inner:  unusedTagResponseReader(t),
		fields: fieldsByID,
	}
	response := tagTestResponse{
		code: 200,
		body: `{
			"count": 1,
			"next": null,
			"previous": null,
			"results": [{
				"id": 42,
				"name": "Production",
				"slug": "production",
				"description": "Production resources",
				"object_types": ["dcim.device"],
				"weight": 3000
			}]
		}`,
	}

	result, err := reader.ReadResponse(response, runtime.JSONConsumer())

	require.NoError(t, err)
	payload := result.(*extras.ExtrasTagsListOK).GetPayload()
	require.Len(t, payload.Results, 1)
	assert.Equal(t, int64(42), payload.Results[0].ID)
	assert.Equal(t, []string{"dcim.device"}, fieldsByID[42].ObjectTypes)
	assert.Equal(t, int64(3000), fieldsByID[42].Weight)
}

func TestTagAPIFieldSchema(t *testing.T) {
	tagSchema := resourceNetboxTag().Schema

	assert.Equal(t, schema.TypeSet, tagSchema["object_types"].Type)
	assert.True(t, tagSchema["object_types"].Optional)
	assert.Equal(t, defaultTagWeight, tagSchema["weight"].Default)

	warnings, errors := tagSchema["weight"].ValidateFunc(32768, "weight")
	assert.Empty(t, warnings)
	assert.NotEmpty(t, errors)
}

func unusedTagResponseReader(t *testing.T) runtime.ClientResponseReader {
	t.Helper()
	return runtime.ClientResponseReaderFunc(
		func(runtime.ClientResponse, runtime.Consumer) (any, error) {
			t.Fatal("fallback response reader was called")
			return nil, nil
		},
	)
}
