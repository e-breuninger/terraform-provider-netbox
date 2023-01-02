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
