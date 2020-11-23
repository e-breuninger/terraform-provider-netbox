---
page_title: "netbox_virtual_machines Data Source - terraform-provider-netbox"
subcategory: ""
description: |-
  
---

# Data Source `netbox_virtual_machines`





## Schema

### Optional

- **filter** (Block Set) (see [below for nested schema](#nestedblock--filter))
- **id** (String, Optional) The ID of this resource.
- **name_regex** (String, Optional)

### Read-only

- **vms** (List of Object, Read-only) (see [below for nested schema](#nestedatt--vms))

<a id="nestedblock--filter"></a>
### Nested Schema for `filter`

Required:

- **name** (String, Required)
- **value** (String, Required)


<a id="nestedatt--vms"></a>
### Nested Schema for `vms`

- **cluster_id** (Number)
- **comments** (String)
- **config_context** (String)
- **custom_fields** (Map of String)
- **disk_size_gb** (Number)
- **local_context_data** (String)
- **memory_mb** (Number)
- **name** (String)
- **platform_id** (Number)
- **primary_ip** (String)
- **primary_ip4** (String)
- **primary_ip6** (String)
- **role_id** (Number)
- **site_id** (Number)
- **status** (String)
- **tag_ids** (List of Number)
- **tenant_id** (Number)
- **vcpus** (Number)
- **vm_id** (Number)


