---
page_title: "Prisma Cloud: prismacloud_anomaly_trusted_list"
---

# prismacloud_anomaly_trusted_list

Data source to return information on current anomaly trusted list in Prisma Cloud.

## Example Usage

```hcl
data "prismacloud_anomaly_trusted_list" "example" {
    atl_id = id
}
```

## Argument Reference

* `atl_id` - (int) Anomaly Trusted List ID

## Attribute Reference

* `name` - Anomaly Trusted List name
* `description` - Reason for trusted listing
* `trusted_list_type` - Anomaly Trusted List type
* `account_id` - Anomaly Trusted List account id
* `applicable_policies` - Applicable Policies
* `vpc` - VPC
* `created_by` - Created by
* `created_on` - Created on
* `trusted_list_entries` - List of network anomalies in the trusted list [below](#trusted-list-entries).

### Trusted List Entries

* `image_id` - Image ID
* `tag_key` - Tag key
* `tag_value` - Tag value
* `ip_cidr` -  IP CIDR
* `port` - Port
* `resource_id` - Resource ID
* `service` - Service
* `subject` - Subject
* `domain` - Domain
* `protocol` - Protocol