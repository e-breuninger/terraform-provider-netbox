package netbox

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

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
