package netbox

import (
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var genericObjectSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"object_type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"object_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
	},
}

func getGenericObjectsFromSchemaSet(schemaSet *schema.Set) []*models.GenericObject {
	retArr := make([]*models.GenericObject, 0, schemaSet.Len())
	for _, i := range schemaSet.List() {
		retArr = append(retArr, &models.GenericObject{
			ObjectID:   int64ToPtr(i.(map[string]interface{})["object_id"].(int64)),
			ObjectType: strToPtr(i.(map[string]interface{})["object_type"].(string)),
		})
	}
	return retArr
}

func getSchemaSetFromGenericObjects(objects []*models.GenericObject) []map[string]interface{} {
	retArr := make([]map[string]interface{}, 0, len(objects))
	for _, obj := range objects {
		mapping := make(map[string]interface{})
		mapping["object_type"] = obj.ObjectType
		mapping["object_id"] = obj.ObjectID

		retArr = append(retArr, mapping)
	}
	return retArr
}
