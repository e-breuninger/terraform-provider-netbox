# AGENTS.md — Tenstorrent fork of terraform-provider-netbox

> **Do not upstream this file.** It documents fork-management workflow and local dev quirks that have no place in `e-breuninger/terraform-provider-netbox`.

## What this fork is

`tenstorrent/terraform-provider-netbox` is a fork of `e-breuninger/terraform-provider-netbox`. We carry a small number of patches on top of upstream and rebase periodically.

It also depends on a **second** fork: `msollanych-tt/go-netbox` is a fork of `fbreckle/go-netbox`, used because we need at least one model field (`Tenant` on `WritableAvailableIP`) that upstream go-netbox does not expose. The dependency is wired in via a `replace` directive in `go.mod`:

```text
replace github.com/fbreckle/go-netbox => github.com/msollanych-tt/go-netbox vX.Y.Z
```

When you change anything that touches the API client (any code under `netbox/client/...` or `netbox/models/...`), you are working in **`go-netbox`**, not here. The fork's `master` branch is the canonical working branch — it is upstream `fbreckle/master` plus the patches we carry. See `Working in go-netbox` below.

## The patches we carry

Each carried delta is recorded as a per-delta block below. Blocks are not deleted when resolved — set `Status` to `Superseded by upstream <sha> on YYYY-MM` or `Removed on YYYY-MM` so the history stays intact for the next agent.

### Block template

When adding a new delta, copy this template:

```markdown
### `<short-id>` — <one-line description>

- **Type:** Bug fix | Feature | Workflow tweak | Compatibility shim
- **Introduced:** YYYY-MM (release tag if known)
- **Files:** `path/a.go`, `path/b.go`
- **Tests:** `TestAccX_y`, `TestAccX_z`
- **Why:** Plain-English problem statement.
- **What:** Plain-English summary of the change.
- **Upstream candidate:** Yes / No / Conditional (with prerequisite). Currently parked per user direction.
- **Related go-netbox change:** None | tag/SHA in msollanych-tt/go-netbox
- **Status:** Active | Superseded by upstream <sha> on YYYY-MM | Removed on YYYY-MM
```

### Carried deltas

#### `release-workflow-permissions` — Release workflow permissions tweak

- **Type:** Workflow tweak
- **Introduced:** 2026-04 (`v5.3.0-tenstorrent.0`)
- **Files:** `.github/workflows/release.yml`
- **Tests:** N/A (workflow-only)
- **Why:** Tenstorrent's GitHub org requires explicit `permissions:` blocks on workflow jobs that publish releases. Upstream's release workflow ran fine on `e-breuninger/...` but failed on `tenstorrent/...` until permissions were spelled out.
- **What:** Adds explicit `permissions:` blocks for the release jobs so goreleaser can write release artifacts and tags.
- **Upstream candidate:** No — internal only.
- **Related go-netbox change:** None
- **Status:** Active

#### `available-ip-tenant` — `tenant_id` on `netbox_available_ip_address`

- **Type:** Feature
- **Introduced:** 2026-04 (`v5.3.0-tenstorrent.0`)
- **Files:** `netbox/resource_netbox_available_ip_address.go`, `netbox/resource_netbox_available_ip_address_test.go`
- **Tests:** `TestAccNetboxAvailableIPAddress_withTenant`
- **Why:** Upstream `netbox_available_ip_address` never supported assigning a tenant at allocation time. We need it for our IPAM workflow where every leased address is owned by a tenant.
- **What:** Adds an optional `tenant_id` schema field, plumbs it through `Create`/`Read`/`Update` against the `WritableAvailableIP.Tenant` field. The `Tenant` field on `WritableAvailableIP` only exists in `msollanych-tt/go-netbox` (see related change), not in upstream `fbreckle/go-netbox`.
- **Upstream candidate:** Yes — clean candidate, but blocked on upstream go-netbox accepting the `Tenant` field on `WritableAvailableIP` first. Currently parked per user direction.
- **Related go-netbox change:** Commit "Add tenant field to WritableAvailableIP for available IP creation" on `msollanych-tt/go-netbox` master (tagged `v0.3.0-tenant-fix`).
- **Status:** Active

#### `service-43-parent` — NetBox 4.3+ parent object on `netbox_service`

- **Type:** Compatibility shim
- **Introduced:** 2026-04 (`v5.3.0-tenstorrent.0`)
- **Files:** `netbox/resource_netbox_service.go`, `netbox/resource_netbox_service_test.go`
- **Tests:** existing `netbox_service` accept tests
- **Why:** NetBox 4.3 replaced the `device`/`virtual_machine` fields on Service with a polymorphic `parent_object_type`/`parent_object_id` pair. Upstream provider had not yet caught up so `netbox_service` was broken on NetBox 4.3+.
- **What:** Maps the provider's existing `device_id` / `virtual_machine_id` schema fields onto the new parent object fields when talking to NetBox 4.3+. The model changes live in our go-netbox fork (the writable Service model now has the parent fields exposed).
- **Upstream candidate:** Yes — also a clean candidate, blocked on upstream go-netbox model alignment. Currently parked per user direction.
- **Related go-netbox change:** Service model updates carried on `msollanych-tt/go-netbox` master.
- **Status:** Active

#### `cf-null-clearing` — `custom_fields` clearing on IP-address-family resources

- **Type:** Bug fix
- **Introduced:** 2026-04 (`v5.3.1-tenstorrent.0`)
- **Files:** `netbox/custom_fields.go`, `netbox/resource_netbox_available_ip_address.go`, `netbox/resource_netbox_ip_address.go`, `netbox/resource_netbox_ip_range.go`
- **Tests:** `TestAccNetboxAvailableIPAddress_cf_clear`, `TestAccNetboxIPAddress_cf_clear`, `TestAccNetboxIpRange_cf_clear`
- **Why:** The `Update` path in all three IP-address-family resources used `if cf, ok := d.GetOk(customFieldsKey); ok { data.CustomFields = cf }`, which silently skipped the assignment whenever the user removed the field from their HCL. Combined with `json:"custom_fields,omitempty"` on the writable models, even when we did assign, an empty map was dropped before serialization, so NetBox never saw a clear and kept stale values. `Read` also gated `d.Set(customFieldsKey, ...)` on a non-empty map, so out-of-band clears never made it back into Terraform state.
- **What:** Adds a `customFieldsForUpdate(d)` helper in `netbox/custom_fields.go` that diffs `d.GetChange(customFieldsKey)` and emits a map containing every key present in new config (with its new value) plus every key dropped from old state (with an explicit `nil` so it marshals as JSON `null`). This is what NetBox's PATCH semantics require to actually clear a CF; sending `{}` is a no-op because `omitempty` strips it. All three resources call this helper in `Update`, the duplicate `if cf, ok := d.GetOk(...)` block in `resource_netbox_available_ip_address.go` is removed, and `Read` now unconditionally calls `d.Set(customFieldsKey, getCustomFields(...))` so state stays in sync with NetBox.
- **Upstream candidate:** Yes — clean candidate. Same gated-assignment pattern exists in upstream and almost certainly has the same bug. Currently parked per user direction.
- **Related go-netbox change:** None
- **Status:** Active

#### `device-type-nested-templates` — nested template lifecycle on `netbox_device_type`

- **Type:** Feature
- **Introduced:** 2026-04 (`v5.3.1-tenstorrent.0`)
- **Files:** `netbox/resource_netbox_device_type.go`, `netbox/device_type_templates.go`
- **Tests:** `TestAccNetboxDeviceType_templates_basic`, `TestAccNetboxDeviceType_templates_update`, `TestAccNetboxDeviceType_templates_destroy`, `TestAccNetboxDeviceType_templates_fk_ordering`, `TestAccNetboxDeviceType_templates_coexistence`
- **Why:** Upstream `netbox_device_type` is a plain wrapper around the device-type API, with no support for managing the component templates (interface, power_port, etc.) that get instantiated on every device of that type. Users had to manage every template via a separate top-level resource and string the `device_type_id` through manually, which made copy-pasting device-type definitions painful.
- **What:** Adds `power_port_templates`, `interface_templates`, `power_outlet_templates`, `front_port_templates`, `rear_port_templates`, `console_port_templates`, `console_server_port_templates`, `device_bay_templates`, `module_bay_templates`, and `inventory_item_templates` as nested `TypeSet` blocks on `netbox_device_type`. A single `syncDeviceTypeTemplates` helper reconciles state per type via list / create / partial-update / delete in dependency order (independents → power_outlet/front_port → inventory_item tree). Coexistence is enforced by an **ownership gate**: a template is "ours" iff its name appears in either the new desired set or the prior-state set (computed via `d.GetChange`). The reconciler only deletes templates that we previously owned, and the Read path only refreshes state for templates already tracked there. This is what makes it safe to use standalone `netbox_interface_template` / `netbox_device_bay_template` resources alongside the nested blocks: the nested-block code never touches templates it doesn't already own, even though they're attached to the same device_type. The contract is still "one ownership model per template object" — but the gate prevents the nested code from silently destroying templates owned by anything else.
- **Upstream candidate:** Conditional — design is upstream-friendly but it's a chunky feature. Would need its own discussion with upstream maintainers. Currently parked per user direction.
- **Related go-netbox change:** Depends on the `dcim-templates-list-filter-param` fix (see below) shipped in `msollanych-tt/go-netbox v0.4.0`. Without it, the per-device-type list filter is silently ignored and the reconciler will try to delete every template in the entire DCIM.
- **Status:** Active

#### `dcim-templates-list-filter-param` — go-netbox fix for `device_type_id` query param on `dcim/*-templates/` endpoints

- **Type:** Bug fix (in `msollanych-tt/go-netbox`)
- **Introduced:** 2026-04 (`v5.3.1-tenstorrent.0`, go-netbox `v0.4.0`)
- **Files:** `netbox/client/dcim/dcim_*_templates_list_parameters.go` in `msollanych-tt/go-netbox` (10 files, 30 query-param sites). No provider-side files.
- **Tests:** Implicitly covered by the `device-type-nested-templates` acceptance tests, which would otherwise mass-delete unrelated templates on every Update path.
- **Why:** The OpenAPI swagger fbreckle's go-netbox is generated from uses the legacy NetBox query-param names `devicetype_id` and `moduletype_id` for the `dcim/*-templates/` list endpoints. Modern NetBox (4.4+, possibly earlier) renamed them to `device_type_id` and `module_type_id` for consistency with the rest of the API. NetBox **silently ignores** unknown query params on these endpoints rather than erroring, so a list call with the old name returns the unfiltered global list of templates. Any reconciler using the result (notably the new nested-templates feature) would then read all templates in the DCIM and try to delete the ones not in the device_type's HCL config.
- **What:** In `msollanych-tt/go-netbox v0.4.0`, renames the wire-format query-param strings in the `SetQueryParam` calls from `devicetype_id` → `device_type_id` and `moduletype_id` → `module_type_id` (plus the `__n` "not equal" variants) across all 10 `dcim_*_templates_list_parameters.go` files. The Go-side struct field names (`DevicetypeID`, `ModuletypeID`) and the `WithDevicetypeID` / `WithModuletypeID` setter method names are left alone for source compatibility — only the wire param strings change. Provider's `go.mod` `replace` line bumps to `v0.4.0`.
- **Upstream candidate:** Yes — file an issue against `fbreckle/go-netbox` (or whatever upstream regeneration story they have). The right long-term fix is regenerating the client against the current NetBox swagger. Currently parked per user direction.
- **Related go-netbox change:** This is the go-netbox change itself. Tagged `v0.4.0` on `msollanych-tt/go-netbox` master.
- **Status:** Active

#### `device-type-templates-examples` — examples and docs for nested templates

- **Type:** Feature (docs)
- **Introduced:** 2026-04 (`v5.3.1-tenstorrent.0`)
- **Files:** `examples/resources/netbox_device_type/resource.tf`, regenerated files under `docs/resources/`
- **Tests:** Covered by feature acceptance tests above; doc regeneration is a `make docs` artifact.
- **Why:** With the nested-templates feature shipping, the example HCL had to grow beyond the previous 9-line minimal example so `make docs` would render meaningful guidance for the new fields.
- **What:** Updates `examples/resources/netbox_device_type/resource.tf` to showcase the nested template blocks, and commits the regenerated `docs/resources/netbox_device_type.md`.
- **Upstream candidate:** Yes — would go upstream alongside the feature itself.
- **Related go-netbox change:** None
- **Status:** Active

#### `device-type-extended-fields` — full metadata field coverage on `netbox_device_type`

- **Type:** Feature
- **Introduced:** 2026-04 (`v5.3.4`)
- **Files:** `netbox/resource_netbox_device_type.go`, `netbox/data_source_netbox_device_type.go`, `netbox/resource_netbox_device_type_test.go`, `examples/resources/netbox_device_type/resource.tf`, regenerated `docs/resources/device_type.md` and `docs/data-sources/device_type.md`
- **Tests:** `TestAccNetboxDeviceType_extendedFields`, `TestAccNetboxDeviceType_cf_clear`, `TestAccNetboxDeviceType_weightUnitRequiresWeight`
- **Why:** Upstream `netbox_device_type` only exposed `model`, `slug`, `manufacturer_id`, `part_number`, `u_height`, `is_full_depth`, and `subdevice_role`. NetBox itself supports a much richer set of metadata on device types — airflow direction, physical weight, free-form description/comments, a default platform pointer, an "exclude from utilization" flag, and arbitrary custom fields — and our IPAM workflow now wants all of them, including the ability to attach JSON-typed custom fields like `system_specs` (per the `TG-00002.yml` workflow).
- **What:** Adds `airflow`, `weight`, `weight_unit` (with `RequiredWith: ["weight"]` mirroring `netbox_rack`), `description` (200-char max), `comments`, `default_platform_id` (FK to `netbox_platform`), `exclude_from_utilization` (defaults to `false`), and the standard `custom_fields` schema to both the resource and the data source. The resource's Update path uses `customFieldsForUpdate(d)` for the same null-clearing behavior the IP-family resources got in `cf-null-clearing`. The two NetBox-4.x-only fields (`default_platform`, `exclude_from_utilization`) required a companion go-netbox change.
- **Upstream candidate:** Yes — clean candidate, but blocked on the related go-netbox change reaching `fbreckle/go-netbox` first. Currently parked per user direction.
- **Related go-netbox change:** Commit "Add default_platform and exclude_from_utilization to DeviceType models" on `msollanych-tt/go-netbox` master (tagged `v0.5.0`).
- **Status:** Active

If the customer asks about more features, they land here too. New patches should land as discrete commits with descriptive messages so the next rebase is bearable, and they should add a block above with `Status: Active`.

## Repo layout reminders

- Code that maps Terraform resources/data sources lives in `netbox/`.
- `GNUmakefile` is the entry point for everything (don't run `go test` directly for acceptance tests — use `make`).
- Acceptance test infra (NetBox in Docker) is in `docker/docker-compose.yml`.
- Release build is `goreleaser` driven by `.github/workflows/release.yml` on any pushed `v*` tag.

## Local development — `dev_overrides` is the standard workflow

**This is how all iterative work on the provider gets done.** Don't reach for any other testing path (custom plugin-dir installs, building zips, etc.) unless `dev_overrides` is genuinely insufficient for what you're doing.

The user's `~/.terraformrc` contains a `dev_overrides` block:

```hcl
provider_installation {
  dev_overrides {
    "tenstorrent/netbox" = "/Users/msollanych/git/terraform-plugin-work/terraform-provider-netbox"
  }
  direct {}
}
```

The override key is `tenstorrent/netbox` (not `e-breuninger/netbox`), so consuming Terraform code must declare:

```hcl
terraform {
  required_providers {
    netbox = {
      source = "tenstorrent/netbox"
    }
  }
}
```

If the override entry is commented out (lines starting with `#`), uncomment it before working — it's left commented when the user wants Terraform to use a real registered provider for some unrelated task.

### The cycle

1. Edit code.
2. Build in place: `go build -o terraform-provider-netbox` from this directory.
3. `cd` to the consuming Terraform project and run `terraform plan` / `terraform apply`. It picks up the new binary immediately.
4. Repeat.

### Hard rules

- **Do not** install to `~/.terraform.d/plugins/`. The override bypasses it; mixing the two leads to confusing version-resolution errors.
- **Do not** run `terraform init -upgrade` after rebuilding. With dev_overrides active, `terraform init` is unnecessary and prints a warning. Just go straight to `plan` or `apply`.
- **Do not** rename the output binary. It must be exactly `terraform-provider-netbox` (no `_v...` suffix), at the path the override points to.
- The override prints a "Provider development overrides are in effect" warning on every command — that is the signal it is wired up correctly. If the warning is missing, the override is not loaded.

### When you're done iterating

Either leave the override on (cheap, but the warning every command may annoy a real workflow), or comment the line out in `~/.terraformrc` to fall back to whatever real provider source the consuming code references.

### Useful Make targets

| Target                                                                            | What it does                                                                                |
| --------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `make test`                                                                       | Unit tests (no NetBox needed).                                                              |
| `make testacc`                                                                    | Brings up a NetBox in Docker (`make docker-up`), then runs the full acceptance suite. Slow. |
| `make testacc-specific-test TEST_FUNC=TestAccNetboxAvailableIPAddress_withTenant` | Run a single acceptance test against the dockerized NetBox.                                 |
| `make docker-up` / `make docker-down`                                             | Start / tear down the NetBox container.                                                     |
| `make docs`                                                                       | Regenerate provider docs from schema. Run before committing if you've touched a schema.     |
| `make fmt`                                                                        | `go fmt` over the package.                                                                  |

`NETBOX_VERSION` is pinned in `GNUmakefile` (currently `v4.4.10`); override per invocation if needed: `NETBOX_VERSION=v4.5.0 make testacc`.

## Working in go-netbox

When a feature requires changes to the OpenAPI-generated client (most "the field doesn't exist" or "the parent type changed" symptoms), work happens in `../go-netbox`:

- Remote layout: `origin = msollanych-tt/go-netbox`, `upstream = fbreckle/go-netbox`.
- **`master` is the canonical working branch** — it is upstream's master plus the patches we carry, kept linear. Develop directly on `master` (or short-lived feature branches off it) for ongoing work.
- The pre-cleanup state of `master` (with various intermediate merges from initial development) is archived at the tag `master-archive-pre-cleanup-2026-04-27` in case anything in those commits is ever needed.
- Cut a tag like `vX.Y.Z` from `master` after a change and push it to `origin`. Earlier tags used a `-tenant-fix` suffix back when there was a separate working branch named that; clean tags going forward.
- Then bump the `replace` line in this repo's `go.mod` to the new tag and `go mod tidy`.

`go-netbox` is OpenAPI-generated. If you're adding a field, the source of truth is the swagger file under `netbox/swagger.json` (or whatever the upstream layout is at the time). Ideally regenerate; in practice past patches have edited the generated Go directly because regen is fragile. If you do edit generated code by hand, keep the change minimal and isolated to the model/operation you need.

When rebasing onto upstream:

```bash
cd ../go-netbox
git fetch upstream
git checkout master
git rebase upstream/master
git push --force-with-lease origin master
git tag -a vX.Y.Z -m "Rebased onto upstream <sha>"
git push origin vX.Y.Z
```

Force-pushing `master` is acceptable because (a) the provider pins by tag not branch, and (b) this is a fork only consumed by the provider in this checkout.

## Rebasing onto upstream — the standard procedure

Do this whenever upstream has tagged a new release, or every few months, whichever comes first.

### Step 1: rebase go-netbox first

The order matters. If upstream provider ships a feature that depends on a recent go-netbox change (e.g. v2 API token support depended on the `WritableToken.Expires` `omitempty` removal), the provider rebase will fail or behave wrong without the underlying dep being current.

```bash
cd ../go-netbox
git fetch upstream
git checkout master
git rebase upstream/master
# resolve any conflicts on our patches
git push --force-with-lease origin master
git tag -a vX.Y.Z -m "Rebased onto upstream <sha>"
git push origin vX.Y.Z
```

Force-pushing `master` is fine — only this repo consumes it, and it's pinned by tag, not branch.

### Step 2: rebase the provider

Create a fresh branch from upstream, cherry-pick our patches in order. Don't rebase `master` onto `upstream/master` directly — branches first, validate, then fast-forward `master`.

```bash
cd ../terraform-provider-netbox
git fetch upstream
git checkout -b rebase-onto-upstream-<MMMYYYY> upstream/master
git cherry-pick <our-commits-in-order>
# resolve conflicts as they come
```

Conflicts that have shown up historically:

- `go.sum` always conflicts. Take `--ours` and let `go mod tidy` regenerate it after step 3.
- `netbox/resource_netbox_available_ip_address.go` is the file most likely to conflict because both we and upstream have been adding things to it (custom_fields, DNS case suppression, our tenant work).
- `netbox/resource_netbox_service.go` may conflict if upstream has touched the parent-object handling.

Cherry-picks sometimes leave whitespace or stray brace artifacts. **Always run `go vet ./...` after cherry-picks** — it catches the dumb stuff before you waste time on `go test`.

### Step 3: bump go-netbox dep and verify

```bash
# Edit go.mod replace line to new tag
go mod tidy
go vet ./...
go build ./...
go test -count=1 -short ./netbox/...
```

If anything fails, fix it on this branch, not by going back and amending cherry-picks. Land the fix as a "rebase fixup" amend into the last cherry-pick or as a new commit.

### Step 4: prerelease tag, verify CI build (optional)

The release pipeline is the only place we have full cross-platform build coverage. If you want a low-risk dry run of the build before cutting a real tag — particularly the first time after a big rebase — push a prerelease tag:

```bash
git push -u origin rebase-onto-upstream-<MMMYYYY>
git tag -a vX.Y.Z-rc1 -m "Prerelease, rebased on upstream <sha>"
git push origin vX.Y.Z-rc1
```

Watch the `release` workflow. It takes ~6 minutes. On success, the GitHub release will have all platform zips + `SHA256SUMS` + signature.

Caveat with prereleases: per SemVer, `X.Y.Z-rc1 < X.Y.Z` and Terraform `~>` constraints don't auto-resolve to prereleases. If you want to test the rc binary from the registry in a downstream deployment, pin the constraint exactly (`version = ">= X.Y.Z-rc1, < X.Y.Z"` or similar). Most of the time the dev_overrides workflow has already validated the same code, so the rc step is skippable.

### Step 5: cut the real tag

```bash
git checkout master
git merge --ff-only rebase-onto-upstream-<MMMYYYY>
git push origin master
git tag -a vX.Y.Z -m "Rebased on upstream <sha> + tenstorrent patches"
git push origin vX.Y.Z
```

**Tag scheme**: plain `vX.Y.Z` semver. Pick a patch version above whatever the latest tenstorrent tag is — don't anchor to upstream's version. Upstream tags (e.g. `e-breuninger v5.3.1`) live as bare commit refs in our repo and shouldn't be confused with our releases. Going forward, our scheme is monotonically increasing semver from our own history (`v5.3.0` → `v5.3.1` → `v5.3.2` → `v5.3.3` → ...). This avoids `-tenstorrent.N` prerelease quirks in downstream `~>` constraints.

If we ever need a non-rebase patch release (just tenstorrent-side bug fixes / feature additions on top of the same upstream anchor), still bump the patch normally.

## Branch hygiene

- `master` — should always be `upstream/master` + our patches in order. Force-pushing here is acceptable after a rebase, since downstream consumers pin to tags not branches. Coordinate with the user before doing it.
- `rebase-onto-upstream-<MMMYYYY>` — temporary, kept around until promoted. Delete a couple weeks after merge.
- `fixes` — historical scratch branch with assorted in-flight changes. Don't develop on it. Cherry-pick anything still useful onto a clean branch off `master` before touching it.
- `netbox_service_improvement`, similar — old work-in-progress. Confirm with user before deleting.
- Stale `rebase-onto-upstream-<old-month>` branches can be deleted once their successor has been merged.

## Upstream contributions

**Currently parked.** The user has decided not to invest in upstreaming patches at this time. Two of the three carried patches are technically good candidates if that decision changes later:

- Tenant support on `netbox_available_ip_address` — would need go-netbox model change upstreamed to `fbreckle/go-netbox` first.
- NetBox 4.3+ `netbox_service` parent fix — same prerequisite.

The release-workflow tweak and `AGENTS.md` are fork-internal forever; they should never go upstream.

If/when upstreaming resumes: branch any upstream PR off `upstream/master` directly — don't branch off our `master` and try to strip our extra commits later. And don't include the `replace` directive in `go.mod`.

## What to never do

- **Never** push to `upstream` remote on either repo. We don't have permission and you'll get a misleading auth error.
- **Never** edit `go.sum` by hand — always regenerate via `go mod tidy`.
- **Never** reuse a published version tag. Make a new prerelease (`-rc2`, `-rc3`, ...) or bump the patch.
- **Never** break the `dev_overrides` workflow — keep the produced binary at the path the override points to.
- **Never** let `master` diverge from upstream by an unmanageable number of commits. If we're more than ~5 commits ahead, we're probably carrying patches that should be upstreamed.
- **Never** create a parallel working branch in go-netbox (no more `tenant-fix`, `feature/...`, etc. as long-lived dev branches). All ongoing development is on `master` or short-lived feature branches off it.

## Quick reference: state at last rebase

Last rebase: **Apr 2026** (carried forward into `v5.3.4`, no new upstream rebase since `v5.3.3`).

- Provider's `master` is upstream master at `8257f4d` ("test: add acceptance test for dns_name case drift") — upstream's tip 4 commits past tag `v5.3.0` — plus the carried tenstorrent patches.
- go-netbox `master` carries three commits on top of upstream `53bc6c52`: the `Tenant` field on `WritableAvailableIP` (`ad4a0111`), the `device_type_id` / `module_type_id` query-param fix on the dcim templates list endpoints (`af097a32`), and the `default_platform` + `exclude_from_utilization` fields on `WritableDeviceType` / `DeviceType` (`cc70b0e9`). Tagged `v0.5.0`. The provider's `go.mod` `replace` line points to `v0.5.0`. The previously-separate `tenant-fix` branch was retired; `master` is now the canonical branch.
- Tagged releases on the provider: `v5.3.0-tenstorrent.0` and `v5.3.0-tenstorrent-rc1` (legacy scheme, kept on origin), then `v5.3.1`, `v5.3.2`, `v5.3.3`, `v5.3.4` (current scheme). `v5.3.4` is the "extended device_type fields + go-netbox v0.5.0" release.
- Carried deltas active at `v5.3.4`: `release-workflow-permissions`, `available-ip-tenant`, `service-43-parent`, `cf-null-clearing`, `device-type-nested-templates`, `dcim-templates-list-filter-param` (in go-netbox `v0.4.0`+), `device-type-templates-examples`, `device-type-extended-fields`. See "The patches we carry" above for details.
- Conflicts encountered during the original Apr 2026 rebase: `go.sum` (every cherry-pick — resolved with `--ours` then `go mod tidy`), and one cherry-pick artifact in `netbox/resource_netbox_available_ip_address_test.go` (stray closing braces, caught by `go vet`).

Tag scheme going forward: plain `vX.Y.Z` semver, monotonically increasing from our own release history. Don't try to anchor patch numbers to upstream's version — keep ours self-contained so downstream `~> 5` constraints resolve cleanly. The `-tenstorrent.<n>` prerelease scheme used in `v5.3.0-tenstorrent.0` was retired because it sorts below `v5.3.0` per SemVer prerelease rules, which is the opposite of what we want.

When you start a new rebase, update this section with the new state at the end. Don't make the next agent grovel through commit logs to figure out where things are.
