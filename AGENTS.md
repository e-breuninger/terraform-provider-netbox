# AGENTS.md — Tenstorrent fork of terraform-provider-netbox

> **Do not upstream this file.** It documents fork-management workflow and local dev quirks that have no place in `e-breuninger/terraform-provider-netbox`.

## What this fork is

`tenstorrent/terraform-provider-netbox` is a fork of `e-breuninger/terraform-provider-netbox`. We carry a small number of patches on top of upstream and rebase periodically.

It also depends on a **second** fork: `msollanych-tt/go-netbox` is a fork of `fbreckle/go-netbox`, used because we need at least one model field (`Tenant` on `WritableAvailableIP`) that upstream go-netbox does not expose. The dependency is wired in via a `replace` directive in `go.mod`:

```
replace github.com/fbreckle/go-netbox => github.com/msollanych-tt/go-netbox vX.Y.Z
```

When you change anything that touches the API client (any code under `netbox/client/...` or `netbox/models/...`), you are working in **`go-netbox`**, not here. The fork's `master` branch is the canonical working branch — it is upstream `fbreckle/master` plus the patches we carry. See `Working in go-netbox` below.

## The patches we carry

As of the last rebase, these are the substantive deltas vs. upstream:

| Area | Why | Upstreamable? |
|---|---|---|
| `netbox_available_ip_address` — `tenant_id` field | Upstream resource never supported assigning a tenant when allocating from a prefix. | **Yes** — clean candidate. Requires the `Tenant` field added in our go-netbox fork to also be upstreamed, or for upstream go-netbox to add it. |
| `netbox_service` — NetBox 4.3+ `parent_object_type`/`parent_object_id` | NetBox 4.3 changed the service parent schema. Upstream provider doesn't yet handle it. | **Yes** — also a clean candidate. Requires go-netbox model changes to be in upstream go-netbox first. |
| `.github/workflows/release.yml` permissions tweak | Tenstorrent-fork-specific GH Actions permissions. | **No** — internal only. |

If the customer asks about more features, those will land here too. New patches should land as discrete commits with descriptive messages so the next rebase is bearable.

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

| Target | What it does |
|---|---|
| `make test` | Unit tests (no NetBox needed). |
| `make testacc` | Brings up a NetBox in Docker (`make docker-up`), then runs the full acceptance suite. Slow. |
| `make testacc-specific-test TEST_FUNC=TestAccNetboxAvailableIPAddress_withTenant` | Run a single acceptance test against the dockerized NetBox. |
| `make docker-up` / `make docker-down` | Start / tear down the NetBox container. |
| `make docs` | Regenerate provider docs from schema. Run before committing if you've touched a schema. |
| `make fmt` | `go fmt` over the package. |

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

```
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

```
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

```
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

```
# Edit go.mod replace line to new tag
go mod tidy
go vet ./...
go build ./...
go test -count=1 -short ./netbox/...
```

If anything fails, fix it on this branch, not by going back and amending cherry-picks. Land the fix as a "rebase fixup" amend into the last cherry-pick or as a new commit.

### Step 4: prerelease tag, verify CI build

The release pipeline is the only place we have full cross-platform build coverage. Tags are immutable once published — we cannot reuse a real version tag if the build fails. So:

```
git push -u origin rebase-onto-upstream-<MMMYYYY>
git tag -a vX.Y.Z-tenstorrent-rc1 -m "Prerelease, rebased on upstream <sha>"
git push origin vX.Y.Z-tenstorrent-rc1
```

Watch the `release` workflow. It takes ~6 minutes. On success, the GitHub release will have all platform zips + `SHA256SUMS` + signature. Pull it locally and smoke-test against a real NetBox if at all in doubt.

### Step 5: promote to real tag

Once the prerelease has been validated:

```
git checkout master
git merge --ff-only rebase-onto-upstream-<MMMYYYY>
git push origin master
git tag -a vX.Y.Z-tenstorrent.0 -m "Rebased on upstream vX.Y.Z + tenstorrent patches"
git push origin vX.Y.Z-tenstorrent.0
```

**Tag scheme**: `v<upstream_version>-tenstorrent.<n>`. Anchor to whatever upstream tag (or sha-equivalent) we're sitting on top of. Bump the trailing `.<n>` if we add more tenstorrent commits without a fresh upstream rebase.

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

Last rebase: **Apr 2026**.

- Provider rebased onto upstream master at `8257f4d` ("test: add acceptance test for dns_name case drift"), which is upstream's tip 4 commits past tag `v5.3.0`.
- go-netbox `master` carries one commit (`Add tenant field to WritableAvailableIP for available IP creation`) on top of upstream `53bc6c52`. Tagged `v0.3.0-tenant-fix`. The provider's `go.mod` `replace` line points to that tag. The previously-separate `tenant-fix` branch was retired; `master` is now the canonical branch.
- Provider tagged as `v5.3.0-tenstorrent.0` after a successful `v5.3.0-tenstorrent-rc1` prerelease build.
- Conflicts encountered during the rebase: `go.sum` (every cherry-pick — resolved with `--ours` then `go mod tidy`), and one cherry-pick artifact in `netbox/resource_netbox_available_ip_address_test.go` (stray closing braces, caught by `go vet`).

Tag scheme going forward: `v<upstream_anchor>-tenstorrent.<n>`. Note: previous releases used a bare `v5.3.1` / `v5.3.2` scheme that collided with upstream's tag namespace; we no longer do that.

When you start a new rebase, update this section with the new state at the end. Don't make the next agent grovel through commit logs to figure out where things are.
