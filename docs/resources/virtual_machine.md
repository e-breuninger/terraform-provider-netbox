---
page_title: "netbox_virtual_machine Resource - terraform-provider-netbox"
subcategory: ""
description: |-
  
---

# Resource `netbox_virtual_machine`





## Schema

### Required

- **cluster_id** (Number, Required)
- **name** (String, Required)

### Optional

- **comments** (String, Optional)
- **disk_size_gb** (Number, Optional)
- **id** (String, Optional) The ID of this resource.
- **memory_mb** (Number, Optional)
- **platform_id** (Number, Optional)
- **role_id** (Number, Optional)
- **tags** (Set of String, Optional)
- **tenant_id** (Number, Optional)
- **vcpus** (Number, Optional)

### Read-only

- **primary_ipv4** (Number, Read-only)


