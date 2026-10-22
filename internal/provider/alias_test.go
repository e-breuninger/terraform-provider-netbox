package provider

// Locks the deprecated pre-6.0 alias attributes: every alias must exist on its resource,
// be optional+computed, and carry the "Use <canonical> instead." deprecation message; the
// canonical attribute stays non-required (alias_of only accepts optional targets). Hand-written;
// survives regeneration.

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// aliasAttributes maps resource type name -> alias -> canonical attribute. Only optional
// canonicals can be aliased, so the pre-6.0 names whose canonical here is required
// (circuit provider_id/type_id, contact_assignment content_type, custom_field and event_rule
// content_types, vpn_tunnel_termination tunnel_id, the primary-MAC interface_id) stay breaking.
var aliasAttributes = map[string]map[string]string{
	"netbox_ip_address":       {"nat_inside_address_id": "nat_inside_id"},
	"netbox_device_interface": {"mgmtonly": "mgmt_only", "tagged_vlans": "tagged_vlan_ids", "untagged_vlan": "untagged_vlan_id"},
	"netbox_circuit_termination": {
		"port_speed":     "port_speed_kbps",
		"upstream_speed": "upstream_speed_kbps",
	},
	"netbox_permission":            {"groups": "group_ids", "users": "user_ids"},
	"netbox_power_feed":            {"max_percent_utilization": "max_utilization_percent"},
	"netbox_rack_type":             {"mounting_depth_mm": "mounting_depth"},
	"netbox_power_outlet_template": {"power_port_id": "power_port_template_id"},
	"netbox_vpn_tunnel":            {"tunnel_group_id": "vpn_tunnel_group_id"},
	"netbox_user":                  {"active": "is_active"},
	"netbox_config_context": {
		"sites": "site_ids", "site_groups": "site_group_ids", "regions": "region_ids",
		"locations": "location_ids", "device_types": "device_type_ids", "roles": "device_role_ids",
		"platforms": "platform_ids", "cluster_types": "cluster_type_ids", "cluster_groups": "cluster_group_ids",
		"clusters": "cluster_ids", "tenant_groups": "tenant_group_ids", "tenants": "tenant_ids",
	},
}

func TestAliasAttributesDeclared(t *testing.T) {
	ctx := context.Background()
	p := New("test")()
	var meta provider.MetadataResponse
	p.Metadata(ctx, provider.MetadataRequest{}, &meta)

	seen := map[string]bool{}
	for _, newRes := range p.(interface {
		Resources(context.Context) []func() resource.Resource
	}).Resources(ctx) {
		r := newRes()
		var metadataResponse resource.MetadataResponse
		r.Metadata(ctx, resource.MetadataRequest{ProviderTypeName: meta.TypeName}, &metadataResponse)
		want, ok := aliasAttributes[metadataResponse.TypeName]
		if !ok {
			continue
		}
		seen[metadataResponse.TypeName] = true
		var schemaResponse resource.SchemaResponse
		r.Schema(ctx, resource.SchemaRequest{}, &schemaResponse)
		for alias, canonical := range want {
			attribute, ok := schemaResponse.Schema.Attributes[alias]
			if !ok {
				t.Errorf("%s: alias %s is missing", metadataResponse.TypeName, alias)
				continue
			}
			if !attribute.IsOptional() || !attribute.IsComputed() || attribute.IsRequired() {
				t.Errorf("%s.%s: alias must be optional+computed", metadataResponse.TypeName, alias)
			}
			if got, wantMsg := attribute.GetDeprecationMessage(), fmt.Sprintf("Use %s instead.", canonical); got != wantMsg {
				t.Errorf("%s.%s: deprecation message = %q, want %q", metadataResponse.TypeName, alias, got, wantMsg)
			}
			c, ok := schemaResponse.Schema.Attributes[canonical]
			if !ok {
				t.Errorf("%s: canonical %s of alias %s is missing", metadataResponse.TypeName, canonical, alias)
				continue
			}
			if c.IsRequired() || !c.IsOptional() || !c.IsComputed() {
				t.Errorf("%s.%s: aliased canonical must be optional+computed", metadataResponse.TypeName, canonical)
			}
			if c.GetDeprecationMessage() != "" {
				t.Errorf("%s.%s: the canonical attribute must not be deprecated", metadataResponse.TypeName, canonical)
			}
		}
	}
	for typeName := range aliasAttributes {
		if !seen[typeName] {
			t.Errorf("resource %s not found in the provider", typeName)
		}
	}
}
