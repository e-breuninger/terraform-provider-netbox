# device/<device id>/<family> or vm/<virtual machine id>/<family>; family is 4 or 6.
# Use an explicit family when both primary IPv4 and IPv6 are set.
terraform import netbox_primary_ip.test device/123/4
