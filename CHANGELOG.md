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
