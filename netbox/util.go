package netbox

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func strToPtr(str string) *string {
	return &str
}

func int64ToPtr(i int64) *int64 {
	return &i
}

func float64ToPtr(i float64) *float64 {
	return &i
}

func toInt64List(a interface{}) []int64 {
	intList := []int64{}
	for _, number := range a.(*schema.Set).List() {
		if n, ok := number.(int); ok {
			intList = append(intList, int64(n))
		} else if n, ok := number.(int64); ok {
			intList = append(intList, n)
		}
	}
	return intList
}

func toInt64PtrList(a interface{}) []*int64 {
	intList := []*int64{}
	if set, ok := a.(*schema.Set); ok {
		for _, number := range set.List() {
			if n, ok := number.(int); ok {
				intList = append(intList, int64ToPtr(int64(n)))
			} else if n, ok := number.(int64); ok {
				intList = append(intList, int64ToPtr(n))
			}
		}
	}
	return intList
}

func joinStringWithFinalConjunction(elems []string, sep, con string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	}

	var b strings.Builder
	b.WriteString(strings.Join(elems[0:len(elems)-1], sep))
	b.WriteString(fmt.Sprintf(" %s %s", con, elems[len(elems)-1]))
	return b.String()
}

func getOptionalStr(d *schema.ResourceData, key string, useSpace bool) string {
	strVal := ""
	// check if key is set
	strValInterface, ok := d.GetOk(key)
	if ok || d.HasChange(key) {
		if !ok && useSpace {
			// Setting an space string deletes the value
			// Ideally we would have the ability to determine if a struct value was set to the zero-value or not
			// The API often supports setting null to clear a value, but this is a Go/Swagger limitation
			strVal = " "
		} else if ok {
			strVal = strValInterface.(string)
		}
	}
	return strVal
}

func getOptionalVal[SchemaT int | float64, ApiT int64 | float64](d *schema.ResourceData, key string) *ApiT {
	var apiPtr *ApiT
	schemaValIf, ok := d.GetOk(key)
	if ok {
		schemaVal, _ := schemaValIf.(SchemaT)
		apiVal := ApiT(schemaVal)
		apiPtr = &apiVal
	}
	return apiPtr
}

func getOptionalInt(d *schema.ResourceData, key string) *int64 {
	return getOptionalVal[int, int64](d, key)
}

func getOptionalFloat(d *schema.ResourceData, key string) *float64 {
	return getOptionalVal[float64, float64](d, key)
}
