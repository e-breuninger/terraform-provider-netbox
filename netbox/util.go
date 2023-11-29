package netbox

import (
	"encoding/json"
	"fmt"
	"reflect"
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

func toStringList(a interface{}) []string {
	strList := []string{}
	for _, str := range a.(*schema.Set).List() {
		strList = append(strList, str.(string))
	}
	return strList
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

func buildValidValueDescription(options []string) string {
	var quoted []string
	for _, option := range options {
		quoted = append(quoted, fmt.Sprintf("`%s`", option))
	}
	return "Valid values are " + joinStringWithFinalConjunction(quoted, ", ", "and")
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

// jsonSemanticCompare returns true when 2 json strings encode the same
// structure, regardless of whitespace differences. This can be used in
// DiffSuppressFunc implementations to prevent terraform showing whitespace
// changes as differences on refresh.
func jsonSemanticCompare(a, b string) (equal bool, err error) {
	var aDecoded, bDecoded any

	err = json.Unmarshal([]byte(a), &aDecoded)
	if err != nil {
		return false, fmt.Errorf("could not decode a: %w", err)
	}

	err = json.Unmarshal([]byte(b), &bDecoded)
	if err != nil {
		return false, fmt.Errorf("could not decode b: %w", err)
	}

	return reflect.DeepEqual(aDecoded, bDecoded), nil
}
