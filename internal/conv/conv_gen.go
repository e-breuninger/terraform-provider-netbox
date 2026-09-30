// Code generated automatically. DO NOT EDIT.

// Package conv converts between terraform-plugin-framework values and the
// pointer-typed fields of the generated DTOs (the <X>RequestDTO/<X>ResponseDTO
// structs the model's expand and flatten produce and consume). It never
// touches the backend's client library, which has request and response types
// of its own: converting a DTO to and from those is the backend-generated
// adapters' job, one layer further out.
//
// Direction out (expand): null or unknown becomes nil, so the field is
// omitted from the request DTO. Direction in (flatten): nil becomes null.
// Collections follow the same rule with nil slices.
package conv

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/netip"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// ValidRegexp validates that a string attribute is a compilable Go regular expression (the
// plural data sources' name_regex input).
func ValidRegexp() validator.String { return validRegexp{} }

type validRegexp struct{}

func (validRegexp) Description(context.Context) string {
	return "must be a valid Go regular expression"
}

func (v validRegexp) MarkdownDescription(ctx context.Context) string { return v.Description(ctx) }

func (validRegexp) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	if _, err := regexp.Compile(req.ConfigValue.ValueString()); err != nil {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid regular expression", err.Error())
	}
}

// ObjectAsOptions decodes nested objects leniently: null and unknown
// members become their zero value instead of an error.
var ObjectAsOptions = basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true}

// ListTo converts a list of primitives to a slice; null or unknown
// becomes nil.
func ListTo[T any](ctx context.Context, v types.List, diags *diag.Diagnostics) []T {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	var out []T
	diags.Append(v.ElementsAs(ctx, &out, false)...)
	return out
}

// SetTo is ListTo for sets.
func SetTo[T any](ctx context.Context, v types.Set, diags *diag.Diagnostics) []T {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	var out []T
	diags.Append(v.ElementsAs(ctx, &out, false)...)
	return out
}

// MapTo converts a map of primitives to a Go map; null or unknown becomes
// nil.
func MapTo[T any](ctx context.Context, v types.Map, diags *diag.Diagnostics) map[string]T {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	var out map[string]T
	diags.Append(v.ElementsAs(ctx, &out, false)...)
	return out
}

// MapFrom converts a Go map to a map of elem; nil (and, with emptyIsNull,
// empty) becomes null.
func MapFrom[T any](ctx context.Context, elem attr.Type, m map[string]T, emptyIsNull bool, diags *diag.Diagnostics) types.Map {
	if m == nil || (emptyIsNull && len(m) == 0) {
		return types.MapNull(elem)
	}
	v, d := types.MapValueFrom(ctx, elem, m)
	diags.Append(d...)
	return v
}

// ListFrom converts a slice to a list of elem; nil (and, with emptyIsNull,
// empty) becomes null.
func ListFrom[T any](ctx context.Context, elem attr.Type, xs []T, emptyIsNull bool, diags *diag.Diagnostics) types.List {
	if xs == nil || (emptyIsNull && len(xs) == 0) {
		return types.ListNull(elem)
	}
	v, d := types.ListValueFrom(ctx, elem, xs)
	diags.Append(d...)
	return v
}

// SetFrom is ListFrom for sets.
func SetFrom[T any](ctx context.Context, elem attr.Type, xs []T, emptyIsNull bool, diags *diag.Diagnostics) types.Set {
	if xs == nil || (emptyIsNull && len(xs) == 0) {
		return types.SetNull(elem)
	}
	v, d := types.SetValueFrom(ctx, elem, xs)
	diags.Append(d...)
	return v
}

func StringPtr(v types.String) *string {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	s := v.ValueString()
	return &s
}

func Int64Ptr(v types.Int64) *int64 {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	i := v.ValueInt64()
	return &i
}

func Float64Ptr(v types.Float64) *float64 {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	f := v.ValueFloat64()
	return &f
}

func BoolPtr(v types.Bool) *bool {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	b := v.ValueBool()
	return &b
}

func FromStringPtr(p *string) types.String {
	if p == nil {
		return types.StringNull()
	}
	return types.StringValue(*p)
}

// EmptyStringAsNil maps an empty API string to nil. A JSON text attribute (spec normalize json)
// takes it: an empty string is not a JSON document, and null is the only value its state can hold.
func EmptyStringAsNil(p *string) *string {
	if p == nil || *p == "" {
		return nil
	}
	return p
}

func FromInt64Ptr(p *int64) types.Int64 {
	if p == nil {
		return types.Int64Null()
	}
	return types.Int64Value(*p)
}

func FromFloat64Ptr(p *float64) types.Float64 {
	if p == nil {
		return types.Float64Null()
	}
	return types.Float64Value(*p)
}

func FromBoolPtr(p *bool) types.Bool {
	if p == nil {
		return types.BoolNull()
	}
	return types.BoolValue(*p)
}

// IDOnlyObject returns an object of the given type with every attribute null except the named
// id. The generated UpgradeState uses it: a model struct's zero value would carry collections
// without element types, which the framework rejects on State.Set.
func IDOnlyObject(ctx context.Context, t types.ObjectType, idName string, id attr.Value) types.Object {
	vals := map[string]attr.Value{}
	for name, attrType := range t.AttrTypes {
		v, err := attrType.ValueFromTerraform(ctx, tftypes.NewValue(attrType.TerraformType(ctx), nil))
		if err != nil {
			return types.ObjectNull(t.AttrTypes)
		}
		vals[name] = v
	}
	vals[idName] = id
	obj, diags := types.ObjectValue(t.AttrTypes, vals)
	if diags.HasError() {
		return types.ObjectNull(t.AttrTypes)
	}
	return obj
}

// The custom string types of spec normalize cidr, ip and mac: two spellings of one network or
// address are equal, so the state keeps the configured text when the API re-serialises it, and a
// value of the wrong shape fails at plan.

var (
	_ basetypes.StringTypable                    = CIDRType{}
	_ basetypes.StringValuableWithSemanticEquals = CIDR{}
	_ xattr.ValidateableAttribute                = CIDR{}
	_ basetypes.StringTypable                    = IPAddressType{}
	_ basetypes.StringValuableWithSemanticEquals = IPAddress{}
	_ xattr.ValidateableAttribute                = IPAddress{}
	_ basetypes.StringTypable                    = MACAddressType{}
	_ basetypes.StringValuableWithSemanticEquals = MACAddress{}
	_ xattr.ValidateableAttribute                = MACAddress{}
)

// parseMACAddress parses an EUI-48 or EUI-64 address in colon, dash or dot notation.
func parseMACAddress(text string) (net.HardwareAddr, error) {
	address, err := net.ParseMAC(text)
	if err != nil {
		return nil, err
	}
	if len(address) != 6 && len(address) != 8 {
		return nil, fmt.Errorf("%d bytes, want an EUI-48 or EUI-64 address", len(address))
	}
	return address, nil
}

// parseIPAddress parses an IP address with an optional prefix length; without one the address is a
// host (/32 or /128).
func parseIPAddress(text string) (netip.Prefix, error) {
	if prefix, err := netip.ParsePrefix(text); err == nil {
		return prefix, nil
	}
	address, err := netip.ParseAddr(text)
	if err != nil {
		return netip.Prefix{}, err
	}
	return netip.PrefixFrom(address, address.BitLen()), nil
}

// CIDRType is the attr.Type of CIDR.
type CIDRType struct{ basetypes.StringType }

func (CIDRType) Equal(other attr.Type) bool { _, isCIDR := other.(CIDRType); return isCIDR }
func (CIDRType) String() string             { return "conv.CIDRType" }
func (CIDRType) ValueType(context.Context) attr.Value {
	return CIDR{}
}
func (CIDRType) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return CIDR{StringValue: in}, nil
}
func (cidrType CIDRType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	value, err := cidrType.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}
	stringValue, isString := value.(basetypes.StringValue)
	if !isString {
		return nil, fmt.Errorf("unexpected value type %T", value)
	}
	return CIDR{StringValue: stringValue}, nil
}

// CIDR is a string holding an IP network in CIDR notation (spec normalize cidr). Host bits must be
// zero (an API storing networks rejects "10.0.0.5/24"); what may differ between two spellings of
// one network is the IPv6 case and compression.
type CIDR struct{ basetypes.StringValue }

func (cidr CIDR) Equal(other attr.Value) bool {
	otherCIDR, isCIDR := other.(CIDR)
	return isCIDR && cidr.StringValue.Equal(otherCIDR.StringValue)
}
func (CIDR) Type(context.Context) attr.Type { return CIDRType{} }
func (cidr CIDR) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics
	newCIDR, isCIDR := newValuable.(CIDR)
	if !isCIDR {
		diags.AddError("Semantic Equality Check Error", fmt.Sprintf("Expected conv.CIDR, got %T.", newValuable))
		return false, diags
	}
	current, errCurrent := netip.ParsePrefix(cidr.ValueString())
	next, errNext := netip.ParsePrefix(newCIDR.ValueString())
	return errCurrent == nil && errNext == nil && current == next, diags
}
func (cidr CIDR) ValidateAttribute(_ context.Context, req xattr.ValidateAttributeRequest, resp *xattr.ValidateAttributeResponse) {
	if cidr.IsNull() || cidr.IsUnknown() {
		return
	}
	prefix, err := netip.ParsePrefix(cidr.ValueString())
	if err != nil {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid CIDR String Value", fmt.Sprintf("Expected an IP network in CIDR notation such as 10.0.0.0/24 or 2001:db8::/32: %s.", err))
		return
	}
	if prefix != prefix.Masked() {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid CIDR String Value", fmt.Sprintf("%s has host bits set; did you mean %s?", cidr.ValueString(), prefix.Masked()))
	}
}

func NewCIDRNull() CIDR              { return CIDR{StringValue: basetypes.NewStringNull()} }
func NewCIDRUnknown() CIDR           { return CIDR{StringValue: basetypes.NewStringUnknown()} }
func NewCIDRValue(value string) CIDR { return CIDR{StringValue: basetypes.NewStringValue(value)} }
func NewCIDRPointerValue(value *string) CIDR {
	return CIDR{StringValue: basetypes.NewStringPointerValue(value)}
}

// IPAddressType is the attr.Type of IPAddress.
type IPAddressType struct{ basetypes.StringType }

func (IPAddressType) Equal(other attr.Type) bool {
	_, isIPAddress := other.(IPAddressType)
	return isIPAddress
}
func (IPAddressType) String() string { return "conv.IPAddressType" }
func (IPAddressType) ValueType(context.Context) attr.Value {
	return IPAddress{}
}
func (IPAddressType) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return IPAddress{StringValue: in}, nil
}
func (ipAddressType IPAddressType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	value, err := ipAddressType.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}
	stringValue, isString := value.(basetypes.StringValue)
	if !isString {
		return nil, fmt.Errorf("unexpected value type %T", value)
	}
	return IPAddress{StringValue: stringValue}, nil
}

// IPAddress is a string holding an IP address with an optional prefix length (spec normalize ip).
// The parsed address and length count, not the spelling: "2001:DB8:0:0::1/64" is "2001:db8::1/64",
// and a bare address is a host ("10.0.0.1" is "10.0.0.1/32").
type IPAddress struct{ basetypes.StringValue }

func (ipAddress IPAddress) Equal(other attr.Value) bool {
	otherIPAddress, isIPAddress := other.(IPAddress)
	return isIPAddress && ipAddress.StringValue.Equal(otherIPAddress.StringValue)
}
func (IPAddress) Type(context.Context) attr.Type { return IPAddressType{} }
func (ipAddress IPAddress) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics
	newIPAddress, isIPAddress := newValuable.(IPAddress)
	if !isIPAddress {
		diags.AddError("Semantic Equality Check Error", fmt.Sprintf("Expected conv.IPAddress, got %T.", newValuable))
		return false, diags
	}
	current, errCurrent := parseIPAddress(ipAddress.ValueString())
	next, errNext := parseIPAddress(newIPAddress.ValueString())
	return errCurrent == nil && errNext == nil && current == next, diags
}
func (ipAddress IPAddress) ValidateAttribute(_ context.Context, req xattr.ValidateAttributeRequest, resp *xattr.ValidateAttributeResponse) {
	if ipAddress.IsNull() || ipAddress.IsUnknown() {
		return
	}
	if _, err := parseIPAddress(ipAddress.ValueString()); err != nil {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid IP Address String Value", fmt.Sprintf("Expected an IP address with an optional prefix length such as 10.0.0.1/24 or 2001:db8::1: %s.", err))
	}
}

func NewIPAddressNull() IPAddress    { return IPAddress{StringValue: basetypes.NewStringNull()} }
func NewIPAddressUnknown() IPAddress { return IPAddress{StringValue: basetypes.NewStringUnknown()} }
func NewIPAddressValue(value string) IPAddress {
	return IPAddress{StringValue: basetypes.NewStringValue(value)}
}
func NewIPAddressPointerValue(value *string) IPAddress {
	return IPAddress{StringValue: basetypes.NewStringPointerValue(value)}
}

// MACAddressType is the attr.Type of MACAddress.
type MACAddressType struct{ basetypes.StringType }

func (MACAddressType) Equal(other attr.Type) bool {
	_, isMACAddress := other.(MACAddressType)
	return isMACAddress
}
func (MACAddressType) String() string { return "conv.MACAddressType" }
func (MACAddressType) ValueType(context.Context) attr.Value {
	return MACAddress{}
}
func (MACAddressType) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return MACAddress{StringValue: in}, nil
}
func (macAddressType MACAddressType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	value, err := macAddressType.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}
	stringValue, isString := value.(basetypes.StringValue)
	if !isString {
		return nil, fmt.Errorf("unexpected value type %T", value)
	}
	return MACAddress{StringValue: stringValue}, nil
}

// MACAddress is a string holding a MAC address (spec normalize mac). The bytes count, not the
// spelling: "aa-bb-cc-dd-ee-ff", "aabb.ccdd.eeff" and "AA:BB:CC:DD:EE:FF" are one address.
type MACAddress struct{ basetypes.StringValue }

func (macAddress MACAddress) Equal(other attr.Value) bool {
	otherMACAddress, isMACAddress := other.(MACAddress)
	return isMACAddress && macAddress.StringValue.Equal(otherMACAddress.StringValue)
}
func (MACAddress) Type(context.Context) attr.Type { return MACAddressType{} }
func (macAddress MACAddress) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics
	newMACAddress, isMACAddress := newValuable.(MACAddress)
	if !isMACAddress {
		diags.AddError("Semantic Equality Check Error", fmt.Sprintf("Expected conv.MACAddress, got %T.", newValuable))
		return false, diags
	}
	current, errCurrent := parseMACAddress(macAddress.ValueString())
	next, errNext := parseMACAddress(newMACAddress.ValueString())
	return errCurrent == nil && errNext == nil && bytes.Equal(current, next), diags
}
func (macAddress MACAddress) ValidateAttribute(_ context.Context, req xattr.ValidateAttributeRequest, resp *xattr.ValidateAttributeResponse) {
	if macAddress.IsNull() || macAddress.IsUnknown() {
		return
	}
	if _, err := parseMACAddress(macAddress.ValueString()); err != nil {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid MAC Address String Value", fmt.Sprintf("Expected a MAC address such as aa:bb:cc:dd:ee:ff: %s.", err))
	}
}

func NewMACAddressNull() MACAddress    { return MACAddress{StringValue: basetypes.NewStringNull()} }
func NewMACAddressUnknown() MACAddress { return MACAddress{StringValue: basetypes.NewStringUnknown()} }
func NewMACAddressValue(value string) MACAddress {
	return MACAddress{StringValue: basetypes.NewStringValue(value)}
}
func NewMACAddressPointerValue(value *string) MACAddress {
	return MACAddress{StringValue: basetypes.NewStringPointerValue(value)}
}
