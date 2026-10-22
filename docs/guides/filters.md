---
page_title: "Data source filters"
description: |-
  How the filters of the plural data sources and the lookup inputs of the singular data sources map to NetBox list parameters.
---

# Data source filters

This document outlines some useful information on the way filters work in this provider.

Every NetBox object type has two data sources: a singular one such as `netbox_device`, which returns exactly one object and errors if no match is found, and a plural one such as `netbox_devices`, which returns every object matching a query or an empty list.
While singular data sources have accept their filters as top-level attributes, plural data sources take a `filters` set with a more flexible configuration.

See also the [NetBox filtering reference](https://netboxlabs.com/docs/netbox/reference/filtering/).

## Plural data sources

A plural data source takes a `filters` set. Each filter consists of a name and a value.
Entries with different names combine with a logical AND, so this returns the active devices at one site:

```terraform
data "netbox_site" "berlin" {
  name = "Berlin"
}

# Entries with different names combine with a logical AND: active devices at the site.
data "netbox_devices" "active_in_berlin" {
  filters = [
    { name = "site_id", value = data.netbox_site.berlin.id },
    { name = "status", value = "active" },
  ]
}
```

The supported names are listed on each data source's page.

### Multiple values

Repeating a filter name makes NetBox combine the values with a logical OR:

```terraform
# Repeating a name combines its values with a logical OR: devices in either status.
data "netbox_devices" "planned_or_staged" {
  filters = [
    { name = "status", value = "planned" },
    { name = "status", value = "staged" },
  ]
}
```

### Lookups

Most filters accept a lookup suffix that changes how the value is compared. Which suffixes a filter accepts depends on the type of its field; the data source page lists every combination.

| Suffix                           | Meaning                                                      | Fields          |
| -------------------------------- | ------------------------------------------------------------ | --------------- |
| `__n`                            | not equal. with several values: none of them                 | all             |
| `__ic`, `__nic`                  | contains / does not contain, case-insensitive                | string          |
| `__isw`, `__nisw`                | starts with / does not start with, case-insensitive          | string          |
| `__iew`, `__niew`                | ends with / does not end with, case-insensitive              | string          |
| `__ie`, `__nie`                  | equals / does not equal, case-insensitive                    | string          |
| `__regex`, `__iregex`            | matches a regular expression, case-sensitive / -insensitive  | string          |
| `__gt`, `__gte`, `__lt`, `__lte` | greater / less than, or equal                                | numeric         |
| `__empty`                        | is empty (`true`) or not empty (`false`)                     | string, numeric |

References to other objects (`site_id`, `role_id` and the like) only take `__n`.

```terraform
# Devices whose name contains "core" (case-insensitive) and which are not decommissioning.
data "netbox_devices" "core" {
  filters = [
    { name = "name__ic", value = "core" },
    { name = "status__n", value = "decommissioning" },
  ]
}

# VLANs with a VID of 100 or higher that have a name.
data "netbox_vlans" "numbered" {
  filters = [
    { name = "vid__gte", value = "100" },
    { name = "name__empty", value = "false" },
  ]
}
```

### Tags

`tag` is the exception to the OR rule: an object must carry every listed tag. `tag__any` matches objects carrying at least one of the tags, and `tag__n` matches objects carrying none of them. The value is the tag's slug.

```terraform
# Devices carrying both tags.
data "netbox_devices" "monitored_and_backed_up" {
  filters = [
    { name = "tag", value = "monitored" },
    { name = "tag", value = "backup" },
  ]
}

# Devices carrying at least one of the tags.
data "netbox_devices" "monitored_or_backed_up" {
  filters = [
    { name = "tag__any", value = "monitored" },
    { name = "tag__any", value = "backup" },
  ]
}

# Devices carrying neither tag.
data "netbox_devices" "neither" {
  filters = [
    { name = "tag__n", value = "monitored" },
    { name = "tag__n", value = "backup" },
  ]
}
```

### Custom fields

Custom fields filter as `cf_<field name>`, where the filter logic is defined by the custom field's "filter logic" field. `loose` is a case-insensitive substring match, `exact` an exact match. Lookup suffixes do not apply to custom fields. NetBox silently ignores a name it has no custom field for, so a misspelled entry does not narrow the result.

```terraform
# Devices whose custom field "tier" matches "gold": a substring match if the field's filter logic
# is loose, an exact match if it is exact.
data "netbox_devices" "gold" {
  filters = [
    { name = "cf_tier", value = "gold" },
  ]
}
```

### Regular expressions

`__regex` and `__iregex` match a string field against a regular expression. NetBox evaluates it, in PostgreSQL's regular expression dialect, so the match happens server-side and `limit` applies like with any other filter.

```terraform
# Up to 50 active devices named like "<site>-core-<number>". NetBox matches the name, so the
# limit applies to the matching devices.
data "netbox_devices" "core" {
  filters = [
    { name = "status", value = "active" },
    { name = "name__regex", value = "^[a-z]+-core-[0-9]+$" },
  ]
  limit = 50
}

output "core_device_names" {
  value = data.netbox_devices.core.devices[*].name
}
```

### Limits

`limit` caps the number of returned objects, 1000 by default. It is the page size NetBox is asked for, and the data source reads that one page.

Results come in NetBox's default order for the object type; there is no ordering input.

## Singular data sources

A singular data source takes typed lookup inputs instead of a filters set: `id`, `name`, `site_id`, `status` and so on, plus `tag_slug` for a tag the object carries. The set inputs combine with a logical AND, and the lookup must match exactly one object; zero or several matches fail the read.

```terraform
data "netbox_site" "berlin" {
  name = "Berlin"
}

# The one device named "core-01" at the site; zero or several matches fail the read.
data "netbox_device" "core_01" {
  name    = "core-01"
  site_id = data.netbox_site.berlin.id
}
```

Custom fields are matched through `custom_field_filters`, a map from field name to value with the same semantics as `cf_<field name>` above.

```terraform
# The one device whose custom field "asset_owner" matches "network-team".
data "netbox_device" "network_team" {
  custom_field_filters = {
    asset_owner = "network-team"
  }
}
```

Where a lookup needs more than the singular inputs offer, such as a substring or a regular expression, query the plural data source and take the single result, e.g. `one(data.netbox_devices.core.devices)`.
