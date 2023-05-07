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
