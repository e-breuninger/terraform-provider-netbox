package netbox

import (
	"net"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	maxUint16 = ^uint16(0)
	maxInt16  = int(maxUint16 >> 1)

	maxUint32 = ^uint32(0)
	maxInt32  = int(maxUint32 >> 1)
)

var (
	validatePositiveInt16 = validation.IntBetween(0, maxInt16)
	validatePositiveInt32 = validation.IntBetween(0, maxInt32)
)

func ValidationIPHasPrefixLenght(i interface{}, s string) ([]string, []error) {
	if _, _, err := net.ParseCIDR(s); err != nil {
		return nil, []error{err}
	}
	return nil, nil
}
