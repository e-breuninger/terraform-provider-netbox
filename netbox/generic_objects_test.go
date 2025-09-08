package netbox

import (
	"testing"

	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestGetGenericObjectsFromSchemaSet(t *testing.T) {
	set := schema.NewSet(schema.HashResource(&schema.Resource{
		Schema: map[string]*schema.Schema{
			"object_type": {Type: schema.TypeString},
			"object_id":   {Type: schema.TypeInt},
		},
	}), []interface{}{
		map[string]interface{}{
			"object_type": "dcim.interface",
			"object_id":   1,
		},
		map[string]interface{}{
			"object_type": "dcim.powerport",
			"object_id":   2,
		},
	})

	result := getGenericObjectsFromSchemaSet(set)

	// Since sets are unordered, we need to check that all expected objects are present
	expectedMap := map[string]int64{
		"dcim.interface": 1,
		"dcim.powerport": 2,
	}

	if len(result) != len(expectedMap) {
		t.Fatalf("expected length %d, got %d", len(expectedMap), len(result))
	}

	for _, obj := range result {
		expectedID, exists := expectedMap[*obj.ObjectType]
		if !exists {
			t.Fatalf("unexpected object type %s", *obj.ObjectType)
		}
		if *obj.ObjectID != expectedID {
			t.Fatalf("expected object ID %d for type %s, got %d", expectedID, *obj.ObjectType, *obj.ObjectID)
		}
	}
}

func TestGetSchemaSetFromGenericObjects(t *testing.T) {
	objects := []*models.GenericObject{
		{
			ObjectType: strToPtr("dcim.interface"),
			ObjectID:   int64ToPtr(1),
		},
		{
			ObjectType: strToPtr("dcim.powerport"),
			ObjectID:   int64ToPtr(2),
		},
	}

	result := getSchemaSetFromGenericObjects(objects)

	// Since the function returns a slice of maps, we need to check that all expected values are present
	expectedMap := map[string]int64{
		"dcim.interface": 1,
		"dcim.powerport": 2,
	}

	if len(result) != len(expectedMap) {
		t.Fatalf("expected length %d, got %d", len(expectedMap), len(result))
	}

	for _, item := range result {
		objTypePtr := item["object_type"].(*string)
		objIDPtr := item["object_id"].(*int64)
		objType := *objTypePtr
		objID := *objIDPtr
		expectedID, exists := expectedMap[objType]
		if !exists {
			t.Fatalf("unexpected object type %s", objType)
		}
		if objID != expectedID {
			t.Fatalf("expected object ID %d for type %s, got %d", expectedID, objType, objID)
		}
	}
}

func TestGetGenericObjectsFromSchemaSet_Empty(t *testing.T) {
	set := schema.NewSet(schema.HashResource(&schema.Resource{
		Schema: map[string]*schema.Schema{
			"object_type": {Type: schema.TypeString},
			"object_id":   {Type: schema.TypeInt},
		},
	}), []interface{}{})

	result := getGenericObjectsFromSchemaSet(set)

	if len(result) != 0 {
		t.Fatalf("expected empty result, got %d items", len(result))
	}
}

func TestGetSchemaSetFromGenericObjects_Empty(t *testing.T) {
	objects := []*models.GenericObject{}

	result := getSchemaSetFromGenericObjects(objects)

	if len(result) != 0 {
		t.Fatalf("expected empty result, got %d items", len(result))
	}
}
