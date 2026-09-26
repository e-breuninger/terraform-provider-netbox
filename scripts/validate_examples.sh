#!/bin/bash
# Validates the docs examples against the provider built from this checkout.
#
# Every .tf file under examples/resources/*/ and examples/data-sources/*/ (plus
# examples/provider/provider.tf) is validated on its own with `terraform validate`, so an example
# must be self-contained: it declares every resource it references. On top of that, every
# registered resource needs resource.tf and import.sh, every registered data source needs
# data-source.tf, every example directory must name a registered type, and examples/ must be
# `terraform fmt` clean. Type names are read from the provider sources (internal/provider), so
# the check needs no NetBox and no registry access; terraform must be on PATH.

set -u
cd "$(dirname "$0")/.." || exit 1

work="$(mktemp -d)"
trap 'rm -rf "$work"' EXIT
mkdir -p "$work/bin"

go build -o "$work/bin/terraform-provider-netbox" . || exit 1
cat >"$work/terraformrc" <<RC
provider_installation {
  dev_overrides {
    "e-breuninger/netbox" = "$work/bin"
  }
  direct {}
}
RC
export TF_CLI_CONFIG_FILE="$work/terraformrc" TF_IN_AUTOMATION=1

failures=0
fail() {
    echo "FAIL: $1"
    failures=$((failures + 1))
}

# Registered types, from the sources: the generated files and the companion files set
# resp.TypeName = req.ProviderTypeName + "_<name>"; the two interface primary MAC resources
# build theirs from a typeSuffix constant.
type_names() {
    local kind="$1" # resource | data_source
    local -a files
    case "$kind" in
    resource) files=(internal/provider/*_resource_gen.go internal/provider/*_resource.go) ;;
    data_source) files=(internal/provider/*_data_source_gen.go internal/provider/*_data_source.go) ;;
    esac
    {
        grep -hoE 'ProviderTypeName \+ "_[a-z0-9_]+"' "${files[@]}" | sed -E 's/.*"_([a-z0-9_]+)"/\1/'
        if [ "$kind" = resource ]; then
            # The interface primary MAC resources build their type name from a kind struct whose
            # first field is the suffix: primaryMACKind{"_device_interface_primary_mac_address", ...}.
            grep -hoE 'Kind\{"_[a-z0-9_]+",' "${files[@]}" | sed -E 's/.*"_([a-z0-9_]+)",/\1/'
        fi
    } | sort -u
}

validate_file() {
    local file="$1"
    local dir
    dir="$work/run/$(echo "$file" | tr '/' '_')"
    mkdir -p "$dir"
    cp "$file" "$dir/main.tf"
    # The provider example declares required_providers itself; every other example gets it added.
    if ! grep -q 'required_providers' "$file"; then
        cat >"$dir/versions.tf" <<TF
terraform {
  required_providers {
    netbox = {
      source = "e-breuninger/netbox"
    }
  }
}
TF
    fi
    local out
    if ! out=$(cd "$dir" && terraform validate -no-color 2>&1); then
        fail "$file"
        echo "$out" | grep -vE 'Provider development overrides|^\s*$|development overrides are set|behavior may therefore|applying changes may cause|releases\.$|- e-breuninger/netbox in' | sed 's/^/    /'
    fi
}

# With paths as arguments (files or example directories) only those are validated and the
# completeness and fmt checks are skipped.
if [ "$#" -gt 0 ]; then
    echo "== validating $# path(s)"
    for path in "$@"; do
        if [ -d "$path" ]; then
            for file in "$path"/*.tf; do validate_file "$file"; done
        else
            validate_file "$path"
        fi
    done
    echo "== $failures failure(s)"
    [ "$failures" = 0 ]
    exit
fi

echo "== validating example files"
for file in examples/resources/*/*.tf examples/data-sources/*/*.tf examples/provider/provider.tf; do
    [ -e "$file" ] || continue
    validate_file "$file"
done

echo "== completeness"
for name in $(type_names resource); do
    dir="examples/resources/netbox_$name"
    [ -e "$dir/resource.tf" ] || [ -n "$(ls "$dir"/*.tf 2>/dev/null)" ] || fail "resource netbox_$name has no example"
    [ -e "$dir/import.sh" ] || fail "resource netbox_$name has no import.sh"
done
for name in $(type_names data_source); do
    [ -e "examples/data-sources/netbox_$name/data-source.tf" ] || fail "data source netbox_$name has no example"
done
for dir in examples/resources/*/; do
    name=$(basename "$dir")
    name=${name#netbox_}
    type_names resource | grep -qx "$name" || fail "$dir names no registered resource"
done
for dir in examples/data-sources/*/; do
    name=$(basename "$dir")
    name=${name#netbox_}
    type_names data_source | grep -qx "$name" || fail "$dir names no registered data source"
done

echo "== terraform fmt"
unformatted=$(terraform fmt -check -recursive -list=true examples/ 2>&1)
if [ -n "$unformatted" ]; then
    fail "terraform fmt: $(echo "$unformatted" | tr '\n' ' ')"
fi

echo "== $failures failure(s)"
[ "$failures" = 0 ]
