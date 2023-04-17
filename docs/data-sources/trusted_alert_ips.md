---
page_title: "Prisma Cloud: prismacloud_trusted_alert_ips"
---

# prismacloud_trusted_alert_ips

Retrieves trusted alert ips information

## Example Usage

```hcl
data "prismacloud_trusted_alert_ips" "example" {

}
```

### Attribute Reference

* `name` - Name of the trusted alert ip.
* `uuid` - UUID.
* `cidrs` - CIDRs, as defined [below](#CIDRs).
* `cidr_count` - CIDR count.


### CIDRs

* `cidr` - (int) Last modified timestamp.
* `description` - Description.
* `uuid` - List of cloud account IDs.
* `created_on` - List of child account group IDs.