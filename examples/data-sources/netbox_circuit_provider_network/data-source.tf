// Get Provider Network by Name
data "netbox_circuilt_provider_network" "name" {
    name = "provider_network"
}

// Get Provider Network by Regex
data "netbox_circuit_provider_network" "nameregex" {
    name_regex = "provider_.*"
}

// Get Provider Network by Tag
data "netbox_circuit_provider_network" "tag" {
    filter {
        name    = "tag"
        value   = "service-a"
    }
}

// Get Provider Network by Tags
data "netbox_circuit_provider_network" "tags" {
    filter {
        name    = "tag"
        value   = "service-a"
    }
    filter {
        name    = "tag"
        value   = "service-b"
    }
}