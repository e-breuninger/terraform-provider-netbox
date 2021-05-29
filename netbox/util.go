package netbox

import (
	"strconv"

	sp "github.com/davecgh/go-spew/spew"
)

func spew(obj interface{}) string {
	return sp.Sdump(obj)
}

func getInt64FromString(s string) int64 {
	res, _ := strconv.ParseInt(s, 10, 64)
	return res
}

func strToPtr(str string) *string {
	return &str
}

func intToPtr(i int) *int {
	return &i
}

func int64ToPtr(i int64) *int64 {
	return &i
}
