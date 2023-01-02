package netbox

import (
	"testing"

	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/stretchr/testify/assert"
)

func TestGetTagListFromNestedTagList(t *testing.T) {

	tags := []*models.NestedTag{
		{
			Name: strToPtr("Foo"),
			Slug: strToPtr("foo"),
		},
		{
			Name: strToPtr("Bar"),
			Slug: strToPtr("bar"),
		},
	}

	flat := getTagListFromNestedTagList(tags)
	expected := []string{
		"Foo",
		"Bar",
	}
	assert.Equal(t, flat, expected)
}
