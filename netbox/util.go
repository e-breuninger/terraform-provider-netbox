package netbox

import (
	sp "github.com/davecgh/go-spew/spew"
	"strconv"
)

func spew(obj interface{}) string {
	return sp.Sdump(obj)
}

func getInt64FromString(s string) int64 {
	res, _ := strconv.ParseInt(s, 10, 64)
	return res
}
