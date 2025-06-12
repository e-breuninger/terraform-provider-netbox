## 3.11.1 (June 12th, 2025)

ENHANCEMENTS

* **New Data Source:** `netbox_device_power_ports` ([#721](https://github.com/e-breuninger/terraform-provider-netbox/pull/721) by [@mraerino](https://github.com/mraerino))
* **New Resource:** `netbox_available_vlan` ([#717](https://github.com/e-breuninger/terraform-provider-netbox/pull/717) by [@MacherelR](https://github.com/MacherelR))

## 3.11.0 (May 27th, 2025)

* provider: Add `default_tags` attribute ([#711](https://github.com/e-breuninger/terraform-provider-netbox/pull/711) by [@mraerino](https://github.com/mraerino))

ENHANCEMENTS

* **New Data Source:** `netbox_ip_ranges` ([#719](https://github.com/e-breuninger/terraform-provider-netbox/pull/719) by [@waza-ari](https://github.com/waza-ari))
* resource/netbox_device: Allow `decommissioning` value in `status` attribute ([#767](https://github.com/e-breuninger/terraform-provider-netbox/pull/767) by [@sboschman](https://github.com/sboschman))
* resource/netbox_site: Add `comments` attribute ([#710](https://github.com/e-breuninger/terraform-provider-netbox/pull/710) by [@mraerino](https://github.com/mraerino))
* resource/netbox_user: Add `email`, `first_name` and `last_name` attributes ([#693](https://github.com/e-breuninger/terraform-provider-netbox/pull/693) by [@mraerino](https://github.com/mraerino))
* data-source/netbox_device_interfaces: Add `limit` attribute ([#695](https://github.com/e-breuninger/terraform-provider-netbox/pull/695) by [@sempervictus](https://github.com/sempervictus))
* resource/netbox_group: Add `description` attribute ([#694](https://github.com/e-breuninger/terraform-provider-netbox/pull/694) by [@mraerino](https://github.com/mraerino))
* resource/netbox_location: Add `facility` attribute ([#718](https://github.com/e-breuninger/terraform-provider-netbox/pull/718) by [@mraerino](https://github.com/mraerino))
* data-source/netbox_locations: Add `facility` attribute ([#718](https://github.com/e-breuninger/terraform-provider-netbox/pull/718) by [@mraerino](https://github.com/mraerino))

## 3.10.0 (January 9th, 2025)

**BREAKING CHANGES**

NetBox 4.1 came with some breaking changes and these are reflected in the provider.

* resource/netbox_event_rule: Replace `trigger_on_X` attributes with `event_types` list attribute
* resource/netbox_racks: Remove `type` attribute
* resource/netbox_virtual_disk: Change `size_gb` attribute to `size_mb`
* resource/netbox_virtual_machine: Change `disk_size_gb` attribute to `disk_size_mb`
* resource/netbox_vlan_group: Remove `min_vid` and `max_vid` attributes in favor of `vid_ranges` attribute

ENHANCEMENTS

provider: Now supports NetBox 4.1.x
* **New Resource:** `netbox_rack_type`
* resource/netbox_racks: Add `form_factor` attribute

## 3.9.3 (January 9th, 2025)

ENHANCEMENTS

* resource/netbox_custom_field: Add `default` attribute ([#647](https://github.com/e-breuninger/terraform-provider-netbox/pull/647) by [@jenxie](https://github.com/jenxie))
* data-source/netbox_vlans: Allow filtering by `site_id` ([#654](https://github.com/e-breuninger/terraform-provider-netbox/pull/654) by [@i-am-smolli](https://github.com/i-am-smolli))
* data-source/netbox_vlan_group: Make `name` and `slug` definitions optional when `scope_type` is defined ([#657](https://github.com/e-breuninger/terraform-provider-netbox/pull/657) by [@TGM](https://github.com/TGM))
* resource/netbox_asn: Add `description` and `comments` attributes ([#664](https://github.com/e-breuninger/terraform-provider-netbox/pull/664) by [@ymylei](https://github.com/ymylei))
* data-source/netbox_prefix: Allow searching by `tenant_id` and `status` ([#666](https://github.com/e-breuninger/terraform-provider-netbox/pull/666) by [@xabinapal](https://github.com/xabinapal))
* data-source/netbox_prefixes: Allow searching by `tenant_id` ([#666](https://github.com/e-breuninger/terraform-provider-netbox/pull/666) by [@xabinapal](https://github.com/xabinapal))

## 3.9.2 (October 10th, 2024)

ENHANCEMENTS

* provider: Include 4.0.11 in supported versions
* resource/netbox_ip_address: Add `custom_fields` attribute ([#638](https://github.com/e-breuninger/terraform-provider-netbox/pull/638) by [@greatman](https://github.com/greatman))
* resource/netbox_service: Add `device_id`, `description` and `tags` attributes ([#637](https://github.com/e-breuninger/terraform-provider-netbox/pull/637) by [@STANIAC](https://github.com/STANIAC))
* data-source/netbox_vrf: Fix a bug where `tenant_id` was not used ([#643](https://github.com/e-breuninger/terraform-provider-netbox/pull/643) by [@c3JpbmkK](https://github.com/c3JpbmkK))

## 3.9.1 (September 2nd, 2024)

ENHANCEMENTS

provider: Include 4.0.9 and 4.0.10 in supported versions

## 3.9.0 (August 10th, 2024)

ENHANCEMENTS

provider: Now is tested against (= supports) the NetBox 4.0.x range

## 3.8.9 (July 31st, 2024)

ENHANCEMENTS

* data-source/netbox_virtual_machines: Add `status` attribute ([#612](https://github.com/e-breuninger/terraform-provider-netbox/pull/612) by [@twink0r](https://github.com/twink0r))
* data-source/netbox_vlans: Add `tag_ids` attribute ([#621](https://github.com/e-breuninger/terraform-provider-netbox/pull/621) by [@Piethan](https://github.com/Piethan))
* data-source/netbox_vlans: Add `status` attribute ([#622](https://github.com/e-breuninger/terraform-provider-netbox/pull/622) by [@Piethan](https://github.com/Piethan))
* data-source/netbox_devices: Add `device_type_id` attribute ([#624](https://github.com/e-breuninger/terraform-provider-netbox/pull/624) by [@Piethan](https://github.com/Piethan))

## 3.8.8 (July 22th, 2024)

ENHANCEMENTS

* data-source/netbox_prefixes: Add `contains` and `site_id` attributes ([#617](https://github.com/e-breuninger/terraform-provider-netbox/pull/617) by [@tagur87](https://github.com/tagur87))

BUG FIXES

* resource/netbox_vpn_tunnel_termination: Fix a interface conversion panic when updating tunnel terminations ([#616](https://github.com/e-breuninger/terraform-provider-netbox/pull/616) by [@mraerino](https://github.com/mraerino))

## 3.8.7 (June 28th, 2024)

ENHANCEMENTS

* **New Resource:** `netbox_interface_template` ([#588](https://github.com/e-breuninger/terraform-provider-netbox/pull/588) by [@thibaultbustarret-ovhcloud](https://github.com/thibaultbustarret-ovhcloud))
* **New Resource:** `netbox_config_context` ([#590](https://github.com/e-breuninger/terraform-provider-netbox/pull/590) by [@diogenxs](https://github.com/diogenxs))
* **New Data Source:** `netbox_config_context` ([#590](https://github.com/e-breuninger/terraform-provider-netbox/pull/590) by [@diogenxs](https://github.com/diogenxs))
* data-source/netbox_devices: Add `config_context` and `local_context_data` attributes ([#590](https://github.com/e-breuninger/terraform-provider-netbox/pull/590) by [@diogenxs](https://github.com/diogenxs))
* resource/netbox_device_interface: Add `label` attribute ([#605](https://github.com/e-breuninger/terraform-provider-netbox/pull/605) by [@thibaultbustarret-ovhcloud](https://github.com/thibaultbustarret-ovhcloud))
* **New Resource:** `netbox_config_template` ([#604](https://github.com/e-breuninger/terraform-provider-netbox/pull/604) by [@thibaultbustarret-ovhcloud](https://github.com/thibaultbustarret-ovhcloud))
* resource/netbox_device: Add `config_template_id` attribute ([#604](https://github.com/e-breuninger/terraform-provider-netbox/pull/604) by [@thibaultbustarret-ovhcloud](https://github.com/thibaultbustarret-ovhcloud))
* data-source/netbox_prefix: Add `role_id` and `custom_fields` attributes ([#607](https://github.com/e-breuninger/terraform-provider-netbox/pull/607) by [@ad8lmondy](https://github.com/ad8lmondy))
* resource/netbox_platform: Add `manufacturer_id` attribute ([#608](https://github.com/e-breuninger/terraform-provider-netbox/pull/608) by [@ad8lmondy](https://github.com/ad8lmondy))
* data-source/netbox_platform: Add `manufacturer_id` attribute ([#608](https://github.com/e-breuninger/terraform-provider-netbox/pull/608) by [@ad8lmondy](https://github.com/ad8lmondy))

## 3.8.6 (May 17th, 2024)

ENHANCEMENTS

* resource/netbox_rir: Add `is_private` attribute ([#594](https://github.com/e-breuninger/terraform-provider-netbox/pull/594) by [@thibaultbustarret-ovhcloud](https://github.com/thibaultbustarret-ovhcloud))
* resource/netbox_vrf: Add `rd` and `enforce_unique` attributes ([#585](https://github.com/e-breuninger/terraform-provider-netbox/pull/585) by [@thibaultbustarret-ovhcloud](https://github.com/thibaultbustarret-ovhcloud))
* **New Resource:** `netbox_group` ([#584](https://github.com/e-breuninger/terraform-provider-netbox/pull/584) by [@thibaultbustarret-ovhcloud](https://github.com/thibaultbustarret-ovhcloud))
* resource/netbox_user: Add `group_ids` attribute ([#584](https://github.com/e-breuninger/terraform-provider-netbox/pull/584) by [@thibaultbustarret-ovhcloud](https://github.com/thibaultbustarret-ovhcloud))

## 3.8.5 (March 18th, 2024)

BUG FIXES

* All resources with `slug` attributes now properly allow for up to 100 characters in that attribute

## 3.8.4 (March 11th, 2024)

ENHANCEMENTS

* data-source/netbox_interfaces: Add `limit` attribute

## 3.8.3 (March 8th, 2024)

ENHANCEMENTS

* **New Resource:** `netbox_vpn_tunnel_termination`

## 3.8.2 (March 4th, 2024)

ENHANCEMENTS

* **New Resource:** `netbox_virtual_disk` ([#558](https://github.com/e-breuninger/terraform-provider-netbox/pull/558) by [@Ikke](https://github.com/Ikke))
* resource/netbox_prefix: Add `custom_fields` attribute ([#553](https://github.com/e-breuninger/terraform-provider-netbox/pull/553) by [@nothinux](https://github.com/nothinux))

## 3.8.1 (February 16th, 2024)

ENHANCEMENTS

* **New Resource:** `netbox_vpn_tunnel_group`
* **New Resource:** `netbox_vpn_tunnel`
* data-source/netbox_virtual_machines: Add `platform_slug` attribute
* data-source/netbox_locations: Add `parent_id` attribute ([#548](https://github.com/e-breuninger/terraform-provider-netbox/pull/548) by [@GennadySpb](https://github.com/GennadySpb))
* data-source/netbox_location: Add `parent_id` attribute ([#548](https://github.com/e-breuninger/terraform-provider-netbox/pull/548) by [@GennadySpb](https://github.com/GennadySpb))
* resource/netbox_location: Add `parent_id` attribute ([#548](https://github.com/e-breuninger/terraform-provider-netbox/pull/548) by [@GennadySpb](https://github.com/GennadySpb))
* resource/device_type: Add `is_full_depth` attribute

## 3.8.0 (January 30th, 2024)

**BREAKING CHANGES**
Due to a change in NetBox 3.7's behavior regarding Webhooks and the corresponding changes in the API,
the `netbox_webhook` resource might cause problems with NetBox versions older than 3.7.0.
For all other resources and data sources, the provider should still perform fine with older NetBox versions.

* resource/netbox_webhook: Removed `enabled`, `trigger_on_create`, `trigger_on_update`, `trigger_on_delete`, `content_types` and `conditions`

ENHANCEMENTS

* provider: Now officially supports NetBox 3.7
* **New Resource:** `netbox_event_rule`

## 3.7.7 (January 30th, 2024)

BUG FIXES

* resource/netbox_device: Fix Virtual Chassis Master Update Function for Tag Input ([#532](https://github.com/e-breuninger/terraform-provider-netbox/pull/532) by [@adelekanley](https://github.com/adelekanley))

## 3.7.6 (January 2nd, 2024)

ENHANCEMENTS

* resource/netbox_webhook: Add `additional_headers` and `conditions` attributes ([#505](https://github.com/e-breuninger/terraform-provider-netbox/pull/505) by [@Ikke](https://github.com/Ikke))
* data-source/netbox_vrfs: Allow filtering by `tag` ([#513](https://github.com/e-breuninger/terraform-provider-netbox/pull/513) by [@sjurtf](https://github.com/sjurtf))
* data-source/netbox_virtual_machines: Allow filtering by `tenant_id` ([#511](https://github.com/e-breuninger/terraform-provider-netbox/pull/511) by [@sjurtf](https://github.com/sjurtf))
* data-source/netbox_ip_addresses: Allow filtering by `tag` ([#510](https://github.com/e-breuninger/terraform-provider-netbox/pull/510) by [@sjurtf](https://github.com/sjurtf))
* data-source/netbox_cluster: Allow filtering by `cluster_group_id` ([#528](https://github.com/e-breuninger/terraform-provider-netbox/pull/528) by [@Ikke](https://github.com/Ikke))

BUG FIXES

* resources/netbox_webhook: Fix a bug where JSON encoding would break drift detection ([#505](https://github.com/e-breuninger/terraform-provider-netbox/pull/505) by [@Ikke](https://github.com/Ikke))
* data-source/netbox_site: Mark optionally searchable attributes as `computed` as well ([#520](https://github.com/e-breuninger/terraform-provider-netbox/pull/520) by [@tagur87](https://github.com/tagur87))

## 3.7.5 (November 27th, 2023)

* **New Data Source:** `netbox_locations` ([#503](https://github.com/e-breuninger/terraform-provider-netbox/pull/503) by [@Ikke](https://github.com/Ikke))

## 3.7.4 (November 22nd, 2023)

ENHANCEMENTS

* **New Resource:** `netbox_virtual_chassis` ([#497](https://github.com/e-breuninger/terraform-provider-netbox/pull/497) by [@Ikke](https://github.com/Ikke))
* resource/netbox_device: Add `virtual_chassis_id`, `virtual_chassis_master`, `virtual_chassis_position` and `virtual_chassis_priority` attributes ([#500](https://github.com/e-breuninger/terraform-provider-netbox/pull/500) by [@Ikke](https://github.com/Ikke))

## 3.7.3 (November 3rd, 2023)

ENHANCEMENTS

* resource/netbox_site: Allow unsetting the `latitude` and `longitude` attributes ([#480](https://github.com/e-breuninger/terraform-provider-netbox/pull/480) by [@haipersuccor02](https://github.com/haipersuccor02))
* **New Data Source:** `netbox_tags` ([#484](https://github.com/e-breuninger/terraform-provider-netbox/pull/484) by [@zeddD1abl0](https://github.com/zeddD1abl0))
* data-source/netbox_ip_addresses: Allow filtering by `parent_prefix` ([#485](https://github.com/e-breuninger/terraform-provider-netbox/pull/485) by [@sjurtf](https://github.com/sjurtf))
* data-source/netbox_devices: Allow filtering by `tags` and `status` ([#491](https://github.com/e-breuninger/terraform-provider-netbox/pull/491) by [@Kenterfie](https://github.com/Kenterfie))
* **New Data Source:** `netbox_available_prefixes` ([#489](https://github.com/e-breuninger/terraform-provider-netbox/pull/489) by [@theochita](https://github.com/theochita))
* resource/netbox_device: Add `local_context_data` attribute ([#493](https://github.com/e-breuninger/terraform-provider-netbox/pull/493) by [@RickyRajinder](https://github.com/RickyRajinder))

## 3.7.2 (October 10th, 2023)

ENHANCEMENTS

* data-source/netbox_location: Allow searching by `site_id` ([#482](https://github.com/e-breuninger/terraform-provider-netbox/pull/482) by [@w87x](https://github.com/w87x))
* data-source/netbox_ip_addresses: Allow searching by `role`, `status`, `vrf` and `tenant` ([#479](https://github.com/e-breuninger/terraform-provider-netbox/pull/479) by [@sjurtf](https://github.com/sjurtf))
* data-source/netbox_site: Allow Searching by `facility` ([#483](https://github.com/e-breuninger/terraform-provider-netbox/pull/483) by [@ikke](https://github.com/ikke))

## 3.7.1 (September 25th, 2023)

ENHANCEMENTS

* **New Data Source:** `netbox_device_interfaces` ([#476](https://github.com/e-breuninger/terraform-provider-netbox/pull/476) by [@w87x](https://github.com/w87x))
* resource/netbox_token: Add `description` attribute ([#473](https://github.com/e-breuninger/terraform-provider-netbox/pull/473) by [@twink0r](https://github.com/twink0r))
* data-source/netbox_virtual_machines: Include `device_id` and `device_name` attributes in result ([#477](https://github.com/e-breuninger/terraform-provider-netbox/pull/477) by [@zeddD1abl0](https://github.com/zeddD1abl0))

## 3.7.0 (September 14th, 2023)

**BREAKING CHANGES**

* resource/netbox_custom_field: Replace `choices` attribute with `choice_set_id` attribute

ENHANCEMENTS

* provider: Now officially supports NetBox 3.6
* **New Resource:** `netbox_custom_field_choice_set`

## 3.6.2 (September 14th, 2023)

FEATURES

* **New Data Source:** `netbox_location` ([#467](https://github.com/e-breuninger/terraform-provider-netbox/pull/467) by [@w87x](https://github.com/w87x))
* data-source/netbox_virtual_machines: Allow searching by tag ([#466](https://github.com/e-breuninger/terraform-provider-netbox/pull/466) by [@twink0r](https://github.com/twink0r))
* resource/netbox_device: Add `asset_tag` attribute ([#470](https://github.com/e-breuninger/terraform-provider-netbox/pull/470) by [@bebehei](https://github.com/bebehei))
* resource/netbox_device_interface: Add `speed`, `lag_device_interface_id` and `parent_device_interface_id` attributes ([#469](https://github.com/e-breuninger/terraform-provider-netbox/pull/469) by [@bebehei](https://github.com/bebehei))

BUG FIXES

* resource/netbox_custom_field: Allow correct value `json` instead of `JSON` ([#459](https://github.com/e-breuninger/terraform-provider-netbox/pull/459) by [@menselman](https://github.com/menselman))

## 3.6.1 (August 31th, 2023)

FEATURES

* **New Resource:** `netbox_webhook` ([#438](https://github.com/e-breuninger/terraform-provider-netbox/pull/438) by [@haipersuccor02](https://github.com/haipersuccor02))

## 3.6.0 (August 18th, 2023)

**BREAKING CHANGES**
Due to a change in NetBox 3.5's behavior regarding ASN and the corresponding required changes in the go library that is used in this provider,
the `netbox_asn` resource and the `netbox_asns` data source are no longer supported in versions older than 3.5.
For all other resources and data sources, the provider should still perform fine with older NetBox versions.

ENHANCEMENTS

* provider: Now officially supports NetBox 3.5

## 3.5.5 (August 18th, 2023)

ENHANCEMENTS

* data-source/netbox_cluster: Allow searching by `id` field and include `custom_fields` attribute in output ([#457](https://github.com/e-breuninger/terraform-provider-netbox/pull/457) by [@fred-clement-91](https://github.com/fred-clement-91))

BUG FIXES

* resource/netbox_device_interface: Changing `mac_address` no longer needlessly forces a recreate of the resource ([#454](https://github.com/e-breuninger/terraform-provider-netbox/pull/454) by [@hamzazaman](https://github.com/hamzazaman))

## 3.5.4 (August 7th, 2023)

FEATURES

* **New Resource:** `netbox_cable` ([#450](https://github.com/e-breuninger/terraform-provider-netbox/pull/450) by [@joeyberkovitz](https://github.com/joeyberkovitz))
* **New Resource:** `netbox_device_console_port` ([#450](https://github.com/e-breuninger/terraform-provider-netbox/pull/450) by [@joeyberkovitz](https://github.com/joeyberkovitz))
* **New Resource:** `netbox_device_console_server_port` ([#450](https://github.com/e-breuninger/terraform-provider-netbox/pull/450) by [@joeyberkovitz](https://github.com/joeyberkovitz))
* **New Resource:** `netbox_device_power_port` ([#450](https://github.com/e-breuninger/terraform-provider-netbox/pull/450) by [@joeyberkovitz](https://github.com/joeyberkovitz))
* **New Resource:** `netbox_device_power_outlet` ([#450](https://github.com/e-breuninger/terraform-provider-netbox/pull/450) by [@joeyberkovitz](https://github.com/joeyberkovitz))
* **New Resource:** `netbox_device_front_port` ([#450](https://github.com/e-breuninger/terraform-provider-netbox/pull/450) by [@joeyberkovitz](https://github.com/joeyberkovitz))
* **New Resource:** `netbox_device_rear_port` ([#450](https://github.com/e-breuninger/terraform-provider-netbox/pull/450) by [@joeyberkovitz](https://github.com/joeyberkovitz))
* **New Resource:** `netbox_device_module_bay` ([#450](https://github.com/e-breuninger/terraform-provider-netbox/pull/450) by [@joeyberkovitz](https://github.com/joeyberkovitz))
* **New Resource:** `netbox_module` ([#450](https://github.com/e-breuninger/terraform-provider-netbox/pull/450) by [@joeyberkovitz](https://github.com/joeyberkovitz))
* **New Resource:** `netbox_module_type` ([#450](https://github.com/e-breuninger/terraform-provider-netbox/pull/450) by [@joeyberkovitz](https://github.com/joeyberkovitz))
* **New Resource:** `netbox_power_feed` ([#450](https://github.com/e-breuninger/terraform-provider-netbox/pull/450) by [@joeyberkovitz](https://github.com/joeyberkovitz))
* **New Resource:** `netbox_power_panel` ([#450](https://github.com/e-breuninger/terraform-provider-netbox/pull/450) by [@joeyberkovitz](https://github.com/joeyberkovitz))
* **New Resource:** `netbox_inventory_item_role` ([#450](https://github.com/e-breuninger/terraform-provider-netbox/pull/450) by [@joeyberkovitz](https://github.com/joeyberkovitz))
* **New Resource:** `netbox_inventory_item` ([#450](https://github.com/e-breuninger/terraform-provider-netbox/pull/450) by [@joeyberkovitz](https://github.com/joeyberkovitz))

## 3.5.3 (August 4th, 2023)

ENHANCEMENTS

* resource/netbox_service: Add `custom_fields` attribute ([#448](https://github.com/e-breuninger/terraform-provider-netbox/pull/448) by [@sebastianreloaded](https://github.com/sebastianreloaded))
* data-source/netbox_site: Allow searching by `id` field

BUG FIXES

* resource/netbox_interface: Allow setting `enabled` to `false`
* resource/netbox_device_interface: Allow setting `enabled` to `false`

## 3.5.2 (August 3rd, 2023)

FEATURES

* **New Data Source:** `netbox_vrfs` ([#441](https://github.com/e-breuninger/terraform-provider-netbox/pull/441) by [@robvand](https://github.com/robvand))

ENHANCEMENTS

* Added `description` attribute to
  - data-source/netbox_cluster
  - data-source/netbox_devices
  - data-source/netbox_virtual_machines
  - resource/netbox_cluster
  - resource/netbox_device
  - resource/virtual_machine ([#401](https://github.com/e-breuninger/terraform-provider-netbox/pull/401) by [@tagur87](https://github.com/tagur87))
* data-source/netbox_devices: Return `tags` attribute
* resource/netbox_interface: Ignore drift when providing lowercase MAC addresses ([#446](https://github.com/e-breuninger/terraform-provider-netbox/pull/446) by [@bebehei](https://github.com/bebehei))
* resource/netbox_device_interface: Ignore drift when providing lowercase MAC addresses ([#446](https://github.com/e-breuninger/terraform-provider-netbox/pull/446) by [@bebehei](https://github.com/bebehei))

## 3.5.1 (July 24th, 2023)

BUG FIXES

* resource/netbox_ip_address: Use correct attribute when using the `device_interface_id` attribute ([#437](https://github.com/e-breuninger/terraform-provider-netbox/pull/437) by [@switchcorp](https://github.com/switchcorp))
* resource/netbox_primary_ip: Fix a bug where setting a primary IP unsets the `local_context_data` attribute ([#435](https://github.com/e-breuninger/terraform-provider-netbox/pull/435) by [@tagur87](https://github.com/tagur87))

## 3.5.0 (July 20th, 2023)

**BREAKING CHANGES**
Historically, this provider primarily handled virtual machines, so when linking a `netbox_ip_address` resource to an interface, the interface was initially assumed to always be a virtual machine interface. In [v3.1.0](https://github.com/e-breuninger/terraform-provider-netbox/commit/76f11292a162d88eb1616d9a5b7d70d986b2db3f), support was added for device interfaces by setting the newly introduced `object_type` attribute, once again defaulting to virtual machine interfaces. The valid values for `object_type` directly reflect the API values of NetBox, which are very unintuitive.

In this version, we make the type of connection between IP addresses and interfaces explicit: We introduce two new attributes: `virtual_machine_interface_id` and `device_interface_id` to the `netbox_ip_address` resource. These fields are easier to use and convey their meaning directly to the user. The `object_type` and `interface_id` method is still supported, but `object_type` no longer has a default value and is now mandatory when `interface_id` is used.

**Migration guide**

In your existing codebase:

* replace `interface_id` with `virtual_machine_interface_id` if `object_type` is currently unset or set to `virtualization.vminterface`
* replace `interface_id` with `device_interface_id` if `object_type` is currently set to `dcim.interface`

ENHANCEMENTS

* resource/netbox_ip_address: Add `virtual_machine_interface_id` and `device_interface_id` attributes 
* resource/netbox_ip_address: Add `slaac` to the list of valid statuses
* resource/netbox_ip_address: Add `nat_inside_address_id` and `nat_outside_addresses` attributes

BUG FIXES

* resource/netbox_permission: Fix perpetual drift when `constraints` is nil ([#432](https://github.com/e-breuninger/terraform-provider-netbox/pull/432) by [@tagur87](https://github.com/tagur87))

## 3.4.1 (July 19th, 2023)

ENHANCEMENTS

* resource/netbox_cluster: Add `comments` attribute ([#429](https://github.com/e-breuninger/terraform-provider-netbox/pull/429) by [@edwin-bruurs](https://github.com/edwin-bruurs))
* data-source/netbox_prefix: Add `family` attribute ([#431](https://github.com/e-breuninger/terraform-provider-netbox/pull/431) by [@tagur87](https://github.com/tagur87))

BUG FIXES

* resource/netbox_virtual_machine: Fix `local_context_data` attribute ([#430](https://github.com/e-breuninger/terraform-provider-netbox/pull/430) by [@zeddD1abl0](https://github.com/zeddD1abl0))

## 3.4.0 (July 10th, 2023)

ENHANCEMENTS

* **New Resource:** `netbox_device_primary_ip` ([#424](https://github.com/e-breuninger/terraform-provider-netbox/pull/424) by [@Ikke](https://github.com/Ikke))
* resource/netbox_virtual_machine: Add `local_context_data` attribute ([#421](https://github.com/e-breuninger/terraform-provider-netbox/pull/421) by [@zeddD1abl0](https://github.com/zeddD1abl0))

BUG FIXES

* resource/netbox_primary_ip: Fix a bug where setting the primary ip of a VM unsets the device id

## 3.3.3 (June 28th, 2023)

ENHANCEMENTS

* **New Data Source:** `netbox_vlans` ([#420](https://github.com/e-breuninger/terraform-provider-netbox/pull/420) by [@danischm](https://github.com/danischm))
* resource/netbox_contact_assignment: Add `priority` attribute ([#418](https://github.com/e-breuninger/terraform-provider-netbox/pull/418) by [@Ikke](https://github.com/Ikke))

## 3.3.2 (June 12th, 2023)

ENHANCEMENTS

* **New Data Source:** `netbox_contact_role` ([#414](https://github.com/e-breuninger/terraform-provider-netbox/pull/414) by [@Ikke](https://github.com/Ikke))

## 3.3.1 (May 31th, 2023)

ENHANCEMENTS

* data-source/netbox_prefixes: Allow filtering by `site_id` ([#397](https://github.com/e-breuninger/terraform-provider-netbox/pull/397) by [@tagur87](https://github.com/tagur87))
* data-source/netbox_ip_addresses: Add `tags` attributes to output ([#406](https://github.com/e-breuninger/terraform-provider-netbox/pull/406) by [@pier-nl](https://github.com/pier-nl))
* Improved error handling for tags ([#400](https://github.com/e-breuninger/terraform-provider-netbox/pull/400) by [@tagur87](https://github.com/tagur87))

## 3.3.0 (May 7th, 2023)

ENHANCEMENTS

* **New Resource:** `netbox_permission` ([#390](https://github.com/e-breuninger/terraform-provider-netbox/pull/390) by [@tagur87](https://github.com/tagur87))
* **New Resource:** `netbox_contact_group` ([#366](https://github.com/e-breuninger/terraform-provider-netbox/pull/366) by [@leasley199](https://github.com/leasley199))
* **New Data Source:** `netbox_contact_group` ([#366](https://github.com/e-breuninger/terraform-provider-netbox/pull/366) by [@leasley199](https://github.com/leasley199))
* **New Data Source:** `netbox_contact` ([#366](https://github.com/e-breuninger/terraform-provider-netbox/pull/366) by [@leasley199](https://github.com/leasley199))
* data-source/netbox_cluster: Allow searching by `site_id`

BUG FIXES

* resource/netbox_prefix: Allow unsetting `description` attribute ([#382](https://github.com/e-breuninger/terraform-provider-netbox/pull/382) by [@DevOpsFu](https://github.com/DevOpsFu))

## 3.2.1 (April 27th, 2023)

ENHANCEMENTS

* **New Resource:** `netbox_vlan_group` ([#377](https://github.com/e-breuninger/terraform-provider-netbox/pull/377) by [@zeddD1abl0](https://github.com/zeddD1abl0))
* **New Data Source:** `netbox_vlan_group` ([#377](https://github.com/e-breuninger/terraform-provider-netbox/pull/377) by [@zeddD1abl0](https://github.com/zeddD1abl0))
* resource/netbox_vlan: Add `group_id` attribute ([#377](https://github.com/e-breuninger/terraform-provider-netbox/pull/377) by [@zeddD1abl0](https://github.com/zeddD1abl0))

BUG FIXES

* data-source/netbox_prefixes: Fix error when filtering by `vlan_vid` ([#381](https://github.com/e-breuninger/terraform-provider-netbox/pull/381) by [@zeddD1abl0](https://github.com/zeddD1abl0))

## 3.2.0 (March 26th, 2023)

ENHANCEMENTS

* **New Resource:** `netbox_route_target` ([#344](https://github.com/e-breuninger/terraform-provider-netbox/pull/344) by [@imdhruva](https://github.com/imdhruva))
* **New Resource:** `netbox_rack` ([#358](https://github.com/e-breuninger/terraform-provider-netbox/pull/358) by [@joeyberkovitz](https://github.com/joeyberkovitz))
* **New Resource:** `netbox_rack_reservation` ([#358](https://github.com/e-breuninger/terraform-provider-netbox/pull/358) by [@joeyberkovitz](https://github.com/joeyberkovitz))
* **New Resource:** `netbox_rack_role` ([#358](https://github.com/e-breuninger/terraform-provider-netbox/pull/358) by [@joeyberkovitz](https://github.com/joeyberkovitz))
* **New Data Source:** `netbox_ipam_role` ([#344](https://github.com/e-breuninger/terraform-provider-netbox/pull/344) by [@imdhruva](https://github.com/imdhruva))
* **New Data Source:** `netbox_route_target` ([#344](https://github.com/e-breuninger/terraform-provider-netbox/pull/344) by [@imdhruva](https://github.com/imdhruva))
* **New Data Source:** `netbox_racks` ([#358](https://github.com/e-breuninger/terraform-provider-netbox/pull/358) by [@joeyberkovitz](https://github.com/joeyberkovitz))
* **New Data Source:** `netbox_rack_role` ([#358](https://github.com/e-breuninger/terraform-provider-netbox/pull/358) by [@joeyberkovitz](https://github.com/joeyberkovitz))
* resource/netbox_device: Add `rack_face`,  `rack_id` and `rack_position` attributes ([#358](https://github.com/e-breuninger/terraform-provider-netbox/pull/358) by [@joeyberkovitz](https://github.com/joeyberkovitz))
* data-source/netbox_device: Add `rack_face`,  `rack_id` and `rack_position` attributes ([#358](https://github.com/e-breuninger/terraform-provider-netbox/pull/358) by [@joeyberkovitz](https://github.com/joeyberkovitz))
* data-source/netbox_prefixes: Add support for filtering by `status` and `tag` ([#367](https://github.com/e-breuninger/terraform-provider-netbox/pull/367) by [@kyle-burnett](https://github.com/kyle-burnett))
* resource/netbox_location: Add `description` attribute
* resource/netbox_rir: Add `description` attribute
* resource/netbox_vrf: Add `description` attribute
* data-source/netbox_prefixes: Include `description` attribute in search results
* data-source/netbox_ip_addresses: Add `limit` attribute (default 1000)

## 3.1.0 (February 19th, 2023)

CHANGES

* provider: `slug` fields are now generated to match the netbox GUI behavior

ENHANCEMENTS

* resource/netbox_interface: Updating `mac_address` no longer forces resource recreation ([#336](https://github.com/e-breuninger/terraform-provider-netbox/pull/336) by [@johann8384](https://github.com/johann8384))
* resource/netbox_site: Add `physical_address` and `shipping_address` ([#337](https://github.com/e-breuninger/terraform-provider-netbox/pull/337) by [@Ikke](https://github.com/Ikke))
* resource/netbox_ip_address: IP addresses can now be assigned to devices via the `object_type` field ([#341](https://github.com/e-breuninger/terraform-provider-netbox/pull/341) by [@arjenvri](https://github.com/arjenvri))

## 3.0.13 (January 24th, 2023)

ENHANCEMENTS

* data-source/netbox_prefix: Add `site_id` attribute ([#320](https://github.com/e-breuninger/terraform-provider-netbox/pull/320) by [@TGM](https://github.com/TGM))

## 3.0.12 (January 3rd, 2023)

ENHANCEMENTS

* resource/netbox_token: Add `write_enabled` attribute ([#309](https://github.com/e-breuninger/terraform-provider-netbox/pull/309) by [@keshy7](https://github.com/keshy7))
* data-source/netbox_interfaces: The resulting interfaces now have their interface ID set

BUG FIXES

* resource/site: Allow unsetting `description` attribute ([#314](https://github.com/e-breuninger/terraform-provider-netbox/pull/314) by [@keshy7](https://github.com/keshy7))
* resource/site: Set max length of the `slug` attribute to 100 ([#317](https://github.com/e-breuninger/terraform-provider-netbox/pull/317) by [@keshy7](https://github.com/keshy7))

## 3.0.11 (December 13th, 2022)

ENHANCEMENTS

* resource/netbox_available_ip_address: Add `role` attribute

BUG FIXES

* resource/netbox_location: Fix updates of the `site_id` attribute ([#307](https://github.com/e-breuninger/terraform-provider-netbox/pull/307) by [@nneul](https://github.com/nneul))

## 3.0.10 (November 24th, 2022)

ENHANCEMENTS

* provider: Add `strip_trailing_slashes_from_url` attribute

BUG FIXES

* data-source/netbox_region: Use correct field for slug attribute ([#302](https://github.com/e-breuninger/terraform-provider-netbox/pull/302) by [@paulexyz](https://github.com/paulexyz))

## 3.0.9 (November 17th, 2022)

ENHANCEMENTS

* data-source/netbox_vlan: Allow querying by `group_id`, `role` and `tenant` ([#287](https://github.com/e-breuninger/terraform-provider-netbox/pull/287) by [@tstarck](https://github.com/tstarck))
* data-source/netbox_prefix: Allow querying by `description` ([#298](https://github.com/e-breuninger/terraform-provider-netbox/pull/298) by [@luispcoutinho](https://github.com/luispcoutinho))

## 3.0.8 (November 9th, 2022)

ENHANCEMENTS

* **New Resource:** `netbox_device_interface` ([#286](https://github.com/e-breuninger/terraform-provider-netbox/pull/286) by [@arjenvri](https://github.com/arjenvri))
* **New Data Source:** `netbox_asn` ([#285](https://github.com/e-breuninger/terraform-provider-netbox/pull/285) by [@kyle-burnett](https://github.com/kyle-burnett))
* **New Data Source:** `netbox_asns` ([#292](https://github.com/e-breuninger/terraform-provider-netbox/pull/292) by [@kyle-burnett](https://github.com/kyle-burnett))
* data-source/netbox_prefix: Add `tags` and tag filter attributes ([#284](https://github.com/e-breuninger/terraform-provider-netbox/pull/284) by [@kyle-burnett](https://github.com/kyle-burnett))

BUG FIXES

* data-source/netbox_prefixes: Fix kernel panic when finding prefixes without vlan or vrf

## 3.0.7 (November 3rd, 2022)

ENHANCEMENTS

* **New Resource:** `netbox_contact_role` ([#279](https://github.com/e-breuninger/terraform-provider-netbox/pull/279) by [@arjenvri](https://github.com/arjenvri))
* **New Resource:** `netbox_contact_assignment` ([#279](https://github.com/e-breuninger/terraform-provider-netbox/pull/279) by [@arjenvri](https://github.com/arjenvri))
* resource/netbox_device: Add `primary_ipv6` attribute ([#282](https://github.com/e-breuninger/terraform-provider-netbox/pull/282) by [@arjenvri](https://github.com/arjenvri))
* resource/netbox_virtual_machine: Add `primary_ipv6` attribute ([#283](https://github.com/e-breuninger/terraform-provider-netbox/pull/283) by [@arjenvri](https://github.com/arjenvri))
* resource/netbox_custom_field: Add `group_name` atribute ([#280](https://github.com/e-breuninger/terraform-provider-netbox/pull/280) by [@arjenvri](https://github.com/arjenvri))

## 3.0.6 (October 21st, 2022)

ENHANCEMENTS

* **New Resource:** `netbox_contact` ([#273](https://github.com/e-breuninger/terraform-provider-netbox/pull/273) by [@arjenvri](https://github.com/arjenvri))
* data-source/netbox_prefix: Add `description` attribute ([#277](https://github.com/e-breuninger/terraform-provider-netbox/pull/277) by [@holmesb](https://github.com/holmesb))
* resource/netbox_cluster: Add `tenant_id` attribute ([#275](https://github.com/e-breuninger/terraform-provider-netbox/pull/275) by [@arjenvri](https://github.com/arjenvri))

## 3.0.5 (October 18th, 2022)

ENHANCEMENTS

* resource/netbox_device_role: Add `tags` attribute ([#269](https://github.com/e-breuninger/terraform-provider-netbox/pull/269) by [@hollow](https://github.com/hollow))
* data-source/netbox_device_role: Add `tags` attribute ([#269](https://github.com/e-breuninger/terraform-provider-netbox/pull/269) by [@hollow](https://github.com/hollow))

CHANGES

* resource/netbox_service: Implement provider-side validation on allowed values. Valid values are `tcp`, `udp` and `sctp`.

## 3.0.4 (October 11th, 2022)

ENHANCEMENTS

* resource/netbox_device: Add `platform_id` attribute ([#264](https://github.com/e-breuninger/terraform-provider-netbox/pull/264) by [@mifrost](https://github.com/mifrost))
* **New Data Source:** `netbox_prefixes ` ([#253](https://github.com/e-breuninger/terraform-provider-netbox/pull/253) by [@ironashram](https://github.com/ironashram))
* data-source/netbox_prefix: Add `prefix`, `status`, `vlan_id`, `vlan_vid` attributes ([#253](https://github.com/e-breuninger/terraform-provider-netbox/pull/253) by [@ironashram](https://github.com/ironashram))
* resource/netbox_device: Add `status` attribute [#266](https://github.com/e-breuninger/terraform-provider-netbox/pull/266) by [@mifrost](https://github.com/mifrost))

CHANGES

* resource/netbox_prefix: Deprecate `cidr` attribute in favor of new canonical `prefix` attribute

## 3.0.3 (October 4th, 2022)

ENHANCEMENTS

* resource/netbox_site: Add `group_id` attribute ([#255](https://github.com/e-breuninger/terraform-provider-netbox/pull/255) by [@arjenvri](https://github.com/arjenvri))

## 3.0.2 (September 30th, 2022)

ENHANCEMENTS

* data-source/netbox_cluster: Add `site_id`, `cluster_type_id`, `cluster_group_id` and `tags` attribute ([#251](https://github.com/e-breuninger/terraform-provider-netbox/pull/251) by [@ns1pelle](https://github.com/ns1pelle))

## 3.0.1 (September 25th, 2022)

This is a re-release of 3.0.0 because there seem to be some issues with the checksums in the 3.0.0 version.

## 3.0.0 (September 23th, 2022)

FEATURES

* provider: Now supports NetBox v3.3

ENHANCEMENTS

* resource/netbox_virtual_machine: In accordance with upstream API changes, VMs can now have `site_id` set directly
* resource/netbox_virtual_machine: Add `device_id` attribute ([#238](https://github.com/e-breuninger/terraform-provider-netbox/pull/238) by [@ns1pelle](https://github.com/ns1pelle))
* resource/netbox_circuit_termination: Add `tags` and `custom_fields` attributes ([#238](https://github.com/e-breuninger/terraform-provider-netbox/pull/238) by [@ns1pelle](https://github.com/ns1pelle))
* resource/netbox_token: Add `allowed_ips`, `last_used` and `expires` attributes ([#238](https://github.com/e-breuninger/terraform-provider-netbox/pull/238) by [@ns1pelle](https://github.com/ns1pelle))
* resource/netbox_device: Add `cluster_id` attribute ([#238](https://github.com/e-breuninger/terraform-provider-netbox/pull/238) by [@ns1pelle](https://github.com/ns1pelle))

## 2.0.7 (September 23th, 2022)

ENHANCEMENTS

* **New Data Source:** `netbox_devices` ([#236](https://github.com/e-breuninger/terraform-provider-netbox/pull/236) by [@dipeshsharma](https://github.com/dipeshsharma))
* provider: Add `request_timeout` attribute ([#227](https://github.com/e-breuninger/terraform-provider-netbox/pull/227) by [@twink0r](https://github.com/twink0r))
* data-source/netbox_tenants: Add `limit` attribute to allow for larger queries

## 2.0.6 (September 9th, 2022)

ENHANCEMENTS

* **New Resource:** `netbox_site_group`
* **New Data Source:** `netbox_site_group`
* resource/netbox_virtual_machine: Add `status` attribute. The `status` attribute will default to `active`, which matches the implicit behavior of NetBox. If you manually changed the status of your terraform-managed NetBox VMs, be cautious
* data-source/netbox_tenant: Allow searching by `slug` attribute

## 2.0.5 (August 10th, 2022)

ENHANCEMENTS

* provider: Update list of supported versions
* docs: Add all missing docs and update existing ones
* docs: Add subcategories

## 2.0.4 (August 1st, 2022)

ENHANCEMENTS

* resource/netbox_ip_address: Add `role` attribute
* resource/netbox_available_ip_address: improve documentation ([#220](https://github.com/e-breuninger/terraform-provider-netbox/pull/220) by [@holmesb](https://github.com/holmesb))

BUG FIXES

* resource/netbox_device: Set correct attribute for `device_type_id` ([#219](https://github.com/e-breuninger/terraform-provider-netbox/pull/219) by [@BegBlev](https://github.com/BegBlev))
* data-source/netbox_ip_addresses: Use correct attribute for `role` ([#217](https://github.com/e-breuninger/terraform-provider-netbox/pull/217) by [@twink0r](https://github.com/twink0r))

## 2.0.3 (July 8th, 2022)

ENHANCEMENTS

* data-source/netbox_prefix: Add `vrf_id` attribute

BUG FIXES

* resource/netbox_prefix: Allow unsetting mark_utilized and is_pool attributes
* resource/netbox_available_prefix: Allow unsetting mark_utilized and is_pool attributes

## 2.0.2 (July 6th, 2022)

BUG FIXES

* resource/netbox_device: Make `role_id` and `site_id` attributes mandatory
* resource/netbox_device_type: Make `manufacturer_id` attribute mandatory

## 2.0.1 (June 25th, 2022)

ENHANCEMENTS

* **New Resource:** `netbox_location` ([#195](https://github.com/e-breuninger/terraform-provider-netbox/pull/195) by [@arjenvri](https://github.com/arjenvri))
* resource/netbox_device: Add `location_id` attribute ([#195](https://github.com/e-breuninger/terraform-provider-netbox/pull/195) by [@arjenvri](https://github.com/arjenvri))

## 2.0.0 (June 16th, 2022)

**BREAKING CHANGES**
NetBox 3.2.0 came with [breaking changes](https://docs.netbox.dev/en/stable/release-notes/version-3.2/#breaking-changes). In accordance with the upstream API, the `netbox_site` resource and data source now have an `asn_ids` attribute that replaces the `asn` attriute. Note that `asn_ids` contains **IDs** of ASN objects, not numbers.

ENHANCEMENTS

* **New Resource:** `netbox_asn`

## 1.6.7 (June 14th, 2022)

ENHANCEMENTS

* resource/netbox_site: Make `status` attribute optional and default to `active` ([#187](https://github.com/e-breuninger/terraform-provider-netbox/pull/187) by [@tstarck](https://github.com/tstarck))
* data-source/netbox_site: Add `slug` parameter to allow searching for a slug ([#187](https://github.com/e-breuninger/terraform-provider-netbox/pull/187) by [@tstarck](https://github.com/tstarck))
* data-source/netbox_site: Include `asn`, `slug`, `comments`, `description`, `group_id`, `status`, `region_id`, `tenant_id` and `time_zone` attributes in the search result ([#187](https://github.com/e-breuninger/terraform-provider-netbox/pull/187) by [@tstarck](https://github.com/tstarck))
* resource/netbox_vlan: Add default values to `status` and `description` attributes ([#184](https://github.com/e-breuninger/terraform-provider-netbox/pull/184) by [@tstarck](https://github.com/tstarck))
* resource/netbox_interface: Add `enabled`, `mtu`, `mode`, `tagged_vlans` and `untagged_vlans` attributes ([#183](https://github.com/e-breuninger/terraform-provider-netbox/pull/183) by [@tstarck](https://github.com/tstarck))

## 1.6.6 (May 27th, 2022)

ENHANCEMENTS

* **New Data Source:** `netbox_device_type` ([#179](https://github.com/e-breuninger/terraform-provider-netbox/pull/179) by [@tstarck](https://github.com/tstarck))
* **New Data Source:** `netbox_vlan` ([#180](https://github.com/e-breuninger/terraform-provider-netbox/pull/180) by [@tstarck](https://github.com/tstarck))
* provider: Add `skip_version_check` attribute
* provider: Update list of officially supported versions
* resource/netbox_device_type: Add `part_number` attribute ([#179](https://github.com/e-breuninger/terraform-provider-netbox/pull/179) by [@tstarck](https://github.com/tstarck))

BUG FIXES

* resource/netbox_circuit: Fix bug that prevented updates from being made
* resource/netbox_circuit_provider: Fix bug that prevented updates from being made

## 1.6.5 (May 18th, 2022)

ENHANCEMENTS

* docs: Fix critical error in usage documentation

## 1.6.4 (May 18th, 2022)

FEATURES

* **New Resource:** `netbox_user` ([#169](https://github.com/e-breuninger/terraform-provider-netbox/pull/169) by [@arjenvri](https://github.com/arjenvri))
* **New Resource:** `netbox_token` ([#169](https://github.com/e-breuninger/terraform-provider-netbox/pull/169) by [@arjenvri](https://github.com/arjenvri))

ENHANCEMENTS

* resource/netbox_site: Add `timezone`, `latitude`, `longitude` and `custom_fields` attributes ([#168](https://github.com/e-breuninger/terraform-provider-netbox/pull/168) by [@arjenvri](https://github.com/arjenvri))
* docs: Regenerate docs with updated tooling ([#165](https://github.com/e-breuninger/terraform-provider-netbox/pull/165) by [@d-strobel](https://github.com/d-strobel))

## 1.6.3 (May 6th, 2022)

FEATURES

* **New Data Source:** `netbox_ip_addresses` ([#159](https://github.com/e-breuninger/terraform-provider-netbox/pull/159) by [@twink0r](https://github.com/twink0r))
* **New Resource:** `netbox_circuit` ([#160](https://github.com/e-breuninger/terraform-provider-netbox/pull/160) by [@arjenvri](https://github.com/arjenvri))
* **New Resource:** `netbox_circuit_provider` ([#160](https://github.com/e-breuninger/terraform-provider-netbox/pull/160) by [@arjenvri](https://github.com/arjenvri))
* **New Resource:** `netbox_circuit_termination` ([#160](https://github.com/e-breuninger/terraform-provider-netbox/pull/160) by [@arjenvri](https://github.com/arjenvri))
* **New Resource:** `netbox_circuit_type` ([#160](https://github.com/e-breuninger/terraform-provider-netbox/pull/160) by [@arjenvri](https://github.com/arjenvri))
* **New Resource:** `netbox_custom_field` ([#158](https://github.com/e-breuninger/terraform-provider-netbox/pull/158) by [@chapsuk](https://github.com/chapsuk))

ENHANCEMENTS

* resource/netbox_ip_address: Add `description` attribute ([#156](https://github.com/e-breuninger/terraform-provider-netbox/pull/156) by [@fbreckle](https://github.com/fbreckle))
* resource/netbox_virtual_machine: Add `custom_fields` attribute ([#158](https://github.com/e-breuninger/terraform-provider-netbox/pull/158) by [@chapsuk](https://github.com/chapsuk))

## 1.6.2 (Apr 11, 2022)

FEATURES

* **New Resource:** `netbox_rir` ([#153](https://github.com/e-breuninger/terraform-provider-netbox/pull/153) by [@arjenvri](https://github.com/arjenvri))
* **New Resource:** `netbox_aggregate` ([#153](https://github.com/e-breuninger/terraform-provider-netbox/pull/153) by [@arjenvri](https://github.com/arjenvri))

## 1.6.1 (Apr 8, 2022)

ENHANCEMENTS

* resource/netbox_site: Add `tags` and `tenant_id` attributes ([#149](https://github.com/e-breuninger/terraform-provider-netbox/pull/149) by [@arjenvri](https://github.com/arjenvri))

## 1.6.0 (Apr 8, 2022)

FEATURES

* **New Resource:** `netbox_device` ([#142](https://github.com/e-breuninger/terraform-provider-netbox/pull/142) by [@arjenvri](https://github.com/arjenvri))
* **New Resource:** `netbox_device_type` ([#142](https://github.com/e-breuninger/terraform-provider-netbox/pull/142) by [@arjenvri](https://github.com/arjenvri))
* **New Resource:** `netbox_manufacturer` ([#142](https://github.com/e-breuninger/terraform-provider-netbox/pull/142) by [@arjenvri](https://github.com/arjenvri))

## 1.5.2 (Mar 4, 2022)

ENHANCEMENTS

* data-source/netbox_tenants: Add `tenant_group` attribute ([#129](https://github.com/e-breuninger/terraform-provider-netbox/pull/129) by [@twink0r](https://github.com/twink0r))

## 1.5.1 (Feb 24, 2022)

ENHANCEMENTS

* No longer crashes if netbox is unreachable when initialising the provider [#126](https://github.com/e-breuninger/terraform-provider-netbox/pull/126) by [@twink0r](https://github.com/twink0r)

## 1.5.0 (Feb 23, 2022)

FEATURES

* **New Data Source:** `netbox_tenants` [#124](https://github.com/e-breuninger/terraform-provider-netbox/pull/124) by [@twink0r](https://github.com/twink0r)

## 1.4.0 (Feb 21, 2022)

FEATURES

* **New Data Source:** `netbox_cluster_type` [#122](https://github.com/e-breuninger/terraform-provider-netbox/pull/122) by [@madnutter56](https://github.com/madnutter56)
* **New Data Source:** `netbox_site` [#122](https://github.com/e-breuninger/terraform-provider-netbox/pull/122) by [@madnutter56](https://github.com/madnutter56)

## 1.3.0 (Feb 17, 2022)

FEATURES

* **New Resource:** `netbox_region` ([#121](https://github.com/e-breuninger/terraform-provider-netbox/pull/121) by [@gerl1ng](https://github.com/gerl1ng))
* **New Data Source:** `netbox_region` [#121](https://github.com/e-breuninger/terraform-provider-netbox/pull/121) by [@gerl1ng](https://github.com/gerl1ng)

## 1.2.2 (Feb 9, 2022)

ENHANCEMENTS

* resource/netbox_virtual_machine: Now has a state migration for the `vcpus` attribute ([#120](https://github.com/e-breuninger/terraform-provider-netbox/pull/120) by [@pascal-hofmann](https://github.com/pascal-hofmann))

## 1.2.1 (Jan 31, 2022)

FEATURES

* provider: Can now optionally pass custom HTTP headers for every request ([#116](https://github.com/e-breuninger/terraform-provider-netbox/pull/116) by [@mariuskiessling](https://github.com/mariuskiessling))

## 1.2.0 (Jan 20, 2022)

FEATURES

* resource/netbox_available_ip_address: Can now be created in netbox_ip_ranges ([#106](https://github.com/e-breuninger/terraform-provider-netbox/pull/106) by [@holmesb](https://github.com/holmesb))

ENHANCEMENTS

* resource/netbox_available_ip_address: fixed duplicates [#59](https://github.com/e-breuninger/terraform-provider-netbox/issues/59) ([#106](https://github.com/e-breuninger/terraform-provider-netbox/pull/106) by [@holmesb](https://github.com/holmesb))
* resource/netbox_available_ip_address: Add `description` argument ([#106](https://github.com/e-breuninger/terraform-provider-netbox/pull/106) by [@holmesb](https://github.com/holmesb))
* resource/netbox_available_ip_address: `status` argument is now optional ([#106](https://github.com/e-breuninger/terraform-provider-netbox/pull/106) by [@holmesb](https://github.com/holmesb))
* resource/netbox_vrf: Add `tenant_id` attribute ([#112](https://github.com/e-breuninger/terraform-provider-netbox/pull/112) by [@cova-fe](https://github.com/cova-fe))
* data-source/netbox_vrf: Add `tenant_id` attribute ([#112](https://github.com/e-breuninger/terraform-provider-netbox/pull/112) by [@cova-fe](https://github.com/cova-fe))
* resource/available_prefix: Add `mark_utilized` attribute ([#111](https://github.com/e-breuninger/terraform-provider-netbox/pull/111) by [@cova-fe](https://github.com/cova-fe))

## 1.1.0 (Jan 3, 2022)

FEATURES

* provider: Now supports NetBox v3.1.3

## 1.0.2 Ho-Ho-Ho (Dec 24, 2021)

ENHANCEMENTS

* resource/tag: Add `description` attribute ([#98](https://github.com/e-breuninger/terraform-provider-netbox/pull/98)) by [@lu1as](https://github.com/lu1as))

## 1.0.1 (Dec 23, 2021)

FEATURES

* **New Resource:** `netbox_ip_range` ([#101](https://github.com/e-breuninger/terraform-provider-netbox/pull/100) by [@holmesb](https://github.com/holmesb))
* **New Data Source:** `netbox_ip_range` ([#101](https://github.com/e-breuninger/terraform-provider-netbox/pull/100) by [@holmesb](https://github.com/holmesb))
 
## 1.0.0 (Nov 8, 2021)

FEATURES

* provider: Now supports NetBox v3.0.9

BREAKING CHANGES

* resource/virtual_machine: `vcpus` is now a float to match upstream API

## 0.3.2 (Nov 2, 2021)

ENHANCEMENTS

* resource/primary_ip: Support both v4 and v6 primary IP ([#87](https://github.com/e-breuninger/terraform-provider-netbox/pull/87) by [@t-tran](https://github.com/t-tran))

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

## 1.5.1 (Feb 24, 2022)

ENHANCEMENTS

* No longer crashes if netbox is unreachable when initialising the provider [#126](https://github.com/e-breuninger/terraform-provider-netbox/pull/126) by [@twink0r](https://github.com/twink0r)

## 1.5.0 (Feb 23, 2022)

FEATURES

* **New Data Source:** `netbox_tenants` [#124](https://github.com/e-breuninger/terraform-provider-netbox/pull/124) by [@twink0r](https://github.com/twink0r)

## 1.4.0 (Feb 21, 2022)

FEATURES

* **New Data Source:** `netbox_cluster_type` [#122](https://github.com/e-breuninger/terraform-provider-netbox/pull/122) by [@madnutter56](https://github.com/madnutter56)
* **New Data Source:** `netbox_site` [#122](https://github.com/e-breuninger/terraform-provider-netbox/pull/122) by [@madnutter56](https://github.com/madnutter56)

## 1.3.0 (Feb 17, 2022)

FEATURES

* **New Resource:** `netbox_region` ([#121](https://github.com/e-breuninger/terraform-provider-netbox/pull/121) by [@gerl1ng](https://github.com/gerl1ng))
* **New Data Source:** `netbox_region` [#121](https://github.com/e-breuninger/terraform-provider-netbox/pull/121) by [@gerl1ng](https://github.com/gerl1ng)

## 1.2.2 (Feb 9, 2022)

ENHANCEMENTS

* resource/netbox_virtual_machine: Now has a state migration for the `vcpus` attribute ([#120](https://github.com/e-breuninger/terraform-provider-netbox/pull/120) by [@pascal-hofmann](https://github.com/pascal-hofmann))

## 1.2.1 (Jan 31, 2022)

FEATURES

* provider: Can now optionally pass custom HTTP headers for every request ([#116](https://github.com/e-breuninger/terraform-provider-netbox/pull/116) by [@mariuskiessling](https://github.com/mariuskiessling))

## 1.2.0 (Jan 20, 2022)

FEATURES

* resource/netbox_available_ip_address: Can now be created in netbox_ip_ranges ([#106](https://github.com/e-breuninger/terraform-provider-netbox/pull/106) by [@holmesb](https://github.com/holmesb))

ENHANCEMENTS

* resource/netbox_available_ip_address: fixed duplicates [#59](https://github.com/e-breuninger/terraform-provider-netbox/issues/59) ([#106](https://github.com/e-breuninger/terraform-provider-netbox/pull/106) by [@holmesb](https://github.com/holmesb))
* resource/netbox_available_ip_address: Add `description` argument ([#106](https://github.com/e-breuninger/terraform-provider-netbox/pull/106) by [@holmesb](https://github.com/holmesb))
* resource/netbox_available_ip_address: `status` argument is now optional ([#106](https://github.com/e-breuninger/terraform-provider-netbox/pull/106) by [@holmesb](https://github.com/holmesb))
* resource/netbox_vrf: Add `tenant_id` attribute ([#112](https://github.com/e-breuninger/terraform-provider-netbox/pull/112) by [@cova-fe](https://github.com/cova-fe))
* data-source/netbox_vrf: Add `tenant_id` attribute ([#112](https://github.com/e-breuninger/terraform-provider-netbox/pull/112) by [@cova-fe](https://github.com/cova-fe))
* resource/available_prefix: Add `mark_utilized` attribute ([#111](https://github.com/e-breuninger/terraform-provider-netbox/pull/111) by [@cova-fe](https://github.com/cova-fe))

## 1.1.0 (Jan 3, 2022)

FEATURES

* provider: Now supports NetBox v3.1.3

## 1.0.2 Ho-Ho-Ho (Dec 24, 2021)

ENHANCEMENTS

* resource/tag: Add `description` attribute ([#98](https://github.com/e-breuninger/terraform-provider-netbox/pull/98)) by [@lu1as](https://github.com/lu1as))

## 1.0.1 (Dec 23, 2021)

FEATURES

* **New Resource:** `netbox_ip_range` ([#101](https://github.com/e-breuninger/terraform-provider-netbox/pull/100) by [@holmesb](https://github.com/holmesb))
* **New Data Source:** `netbox_ip_range` ([#101](https://github.com/e-breuninger/terraform-provider-netbox/pull/100) by [@holmesb](https://github.com/holmesb))
 
## 1.0.0 (Nov 8, 2021)

FEATURES

* provider: Now supports NetBox v3.0.9

BREAKING CHANGES

* resource/virtual_machine: `vcpus` is now a float to match upstream API

## 0.3.2 (Nov 2, 2021)

ENHANCEMENTS

* resource/primary_ip: Support both v4 and v6 primary IP ([#87](https://github.com/e-breuninger/terraform-provider-netbox/pull/87) by [@t-tran](https://github.com/t-tran))

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
