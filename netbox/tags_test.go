package netbox

import (
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTagListFromNestedTagList(t *testing.T) {

	tags := []*models.NestedTag{
		&models.NestedTag{
			Name: strToPtr("Foo"),
			Slug: strToPtr("foo"),
		},
		&models.NestedTag{
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
