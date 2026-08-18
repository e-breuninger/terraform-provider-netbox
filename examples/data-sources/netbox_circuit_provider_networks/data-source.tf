// Get All Provider Networks
data "netbox_circuit_provider_networks" "all" {
}

// Get Provider Network by Name
data "netbox_circuilt_provider_networks" "name" {
    name = "provider_network"
}

// Get Provider Network by Regex
data "netbox_circuit_provider_networks" "nameregex" {
    name_regex = "provider_.*"
}

// Get Provider Network by Tag
data "netbox_circuit_provider_networks" "tag" {
    filter {
        name    = "tag"
        value   = "service-a"
    }
}

// Get Provider Network by Tags
data "netbox_circuit_provider_networks" "tags" {
    filter {
        name    = "tag"
        value   = "service-a"
    }
    filter {
        name    = "tag"
        value   = "service-b"
    }
}

// Get Provider Networks with Limit
data "netbox_circuit_provider_networks" "limit" {
    limit   = 10
    filter {
        name    = "tag"
        value   = "service-a"
    }
}
