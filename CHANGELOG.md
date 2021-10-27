## 0.3.1 (Oct 27, 2021)

FEATURES

* **New Resource:** `netbox_vlan` ([#83](https://github.com/e-breuninger/terraform-provider-netbox/pull/83) by [@Sanverik](https://github.com/Sanverik))
* **New Resource:** `netbox_ipam_role` ([#86](https://github.com/e-breuninger/terraform-provider-netbox/pull/86) by [@Sanverik](https://github.com/Sanverik))

ENHANCEMENTS

* resource/prefix: Add `site_id`, `vlan_id` and `role_id` attributes ([#85](https://github.com/e-breuninger/terraform-provider-netbox/pull/85) and [#85](https://github.com/e-breuninger/terraform-provider-netbox/pull/85)) by [@Sanverik](https://github.com/Sanverik))

## 0.3.0 (Oct 19, 2021)

FEATURES

* provider: Now supports NetBox v2.11.12

BREAKING CHANGES

* resource/virtual_machine: `vcpus` is now a string to match upstream API

## 0.2.5 (Oct 8, 2021)

ENHANCEMENTS

* **New Resource:** `netbox_site` ([#78](https://github.com/e-breuninger/terraform-provider-netbox/pull/78))

BUG FIXES

* resource/cluster: Properly set tags when updating ([#69](https://github.com/e-breuninger/terraform-provider-netbox/issues/69))

## 0.2.4 (Sep 20, 2021)

CHANGES

* Use go 1.17 to fix some builds

## 0.2.3 (Sep 20, 2021)

ENHANCEMENTS

* Add arm64 builds ([#71](https://github.com/e-breuninger/terraform-provider-netbox/pull/71) by [@richardklose](https://github.com/richardklose))

## 0.2.2 (Aug 23, 2021)

ENHANCEMENTS

* resource/interface: Add `mac_address` attribute ([#65](https://github.com/e-breuninger/terraform-provider-netbox/pull/65) by [@holmesb](https://github.com/holmesb))

## 0.2.1 (Jul 26, 2021)

ENHANCEMENTS

* resource/prefix: Add `vrf` and `tenant` attribute ([#61](https://github.com/e-breuninger/terraform-provider-netbox/pull/61) by [@jeansebastienh](https://github.com/jeansebastienh))

BUG FIXES

* resource/prefix: Correctly read `prefix` and `status` ([#60](https://github.com/e-breuninger/terraform-provider-netbox/pull/60) by [@jeansebastienh](https://github.com/jeansebastienh))

## 0.2.0 (May 31, 2021)

FEATURES

* provider: Now supports NetBox v2.10.10

CHANGES

* resource/service: `port` field is now deprecated in favor of `ports` field.

## 0.1.3 (May 17, 2021)

ENHANCEMENTS

* **New Resource:** `netbox_tenant_group` ([#48](https://github.com/e-breuninger/terraform-provider-netbox/pull/48) by [@pezhore](https://github.com/pezhore))
* **New Data Source:** `netbox_tenant_group` ([#48](https://github.com/e-breuninger/terraform-provider-netbox/pull/48) by [@pezhore](https://github.com/pezhore))
* data-source/tenant: Add `group_id` attribute ([#48](https://github.com/e-breuninger/terraform-provider-netbox/pull/48) by [@pezhore](https://github.com/pezhore))
* resource/tenant: Add `group_id` attribute ([#48](https://github.com/e-breuninger/terraform-provider-netbox/pull/48) by [@pezhore](https://github.com/pezhore))
* Documentation ([#46](https://github.com/e-breuninger/terraform-provider-netbox/pull/46) by [@pezhore](https://github.com/pezhore))

## 0.1.2 (May 4, 2021)

ENHANCEMENTS

* **New Resource:** `netbox_prefix` ([#43](https://github.com/e-breuninger/terraform-provider-netbox/pull/43) by [@pezhore](https://github.com/pezhore))
* **New Data Source:** `netbox_prefix` ([#43](https://github.com/e-breuninger/terraform-provider-netbox/pull/43) by [@pezhore](https://github.com/pezhore))
* **New Resource:** `netbox_available_ip_address` ([#43](https://github.com/e-breuninger/terraform-provider-netbox/pull/43) by [@pezhore](https://github.com/pezhore))

## 0.1.1 (February 15, 2021)

ENHANCEMENTS

* data-source/netbox_virtual_machines: Add `limit` attribute ([#33](https://github.com/e-breuninger/terraform-provider-netbox/pull/33) by [@jake2184](https://github.com/jake2184))

## 0.1.0 Ho-Ho-Ho (December 24, 2020)

FEATURES

* **New Resource:** `netbox_vrf` ([#26](https://github.com/e-breuninger/terraform-provider-netbox/pull/26) by [@rthomson](https://github.com/rthomson))
* **New Data Source:** `netbox_vrf` ([#26](https://github.com/e-breuninger/terraform-provider-netbox/pull/26) by [@rthomson](https://github.com/rthomson))
* **New Resource:** `netbox_cluster_group`
* **New Data Source:** `netbox_cluster_group`

ENHANCEMENTS

* resource/netbox_ip_address: Add `tenant_id` attribute
* resource/netbox_cluster: Add `cluster_group_id` attribute

## 0.0.9 (November 20, 2020)

FEATURES

* **New Data Source:** `netbox_interfaces` ([#9](https://github.com/e-breuninger/terraform-provider-netbox/pull/9) by [@jake2184](https://github.com/jake2184))

BUG FIXES

* provider: Honor Sub-Paths in netbox URL ([#15](https://github.com/e-breuninger/terraform-provider-netbox/pull/15) by [@kasimon](https://github.com/kasimon))

## 0.0.8 (November 19, 2020)

FEATURES

* **New Data Source:** `netbox_virtual_machines` ([#8](https://github.com/e-breuninger/terraform-provider-netbox/pull/8) by [@jake2184](https://github.com/jake2184))
