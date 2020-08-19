package netbox

import (
	"bytes"
	"fmt"
	sp "github.com/davecgh/go-spew/spew"
	"github.com/go-openapi/runtime"
	"log"
	"strconv"
)

func spew(obj interface{}) string {
	return sp.Sdump(obj)
}

func getMessageFromError(e error) string {
	apiError := e.(*runtime.APIError)
	clientResponse := apiError.Response.(runtime.ClientResponse)
	buf := new(bytes.Buffer)
	buf.ReadFrom(clientResponse.Body())
	newStr := buf.String()
	log.Printf("[FABI] Body: %s\n", newStr)
	log.Printf("[FABI] Code: %v\n", clientResponse.Code())
	return fmt.Sprintf("%s", clientResponse.Message())
}

func getInt64FromString(s string) int64 {
	res, _ := strconv.ParseInt(s, 10, 64)
	return res
}
