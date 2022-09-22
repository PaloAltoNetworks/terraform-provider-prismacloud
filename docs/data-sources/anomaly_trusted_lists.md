---
page_title: "Prisma Cloud: prismacloud_anomaly_trusted_lists"
---

# prismacloud_anomaly_trusted_lists

Data source to return information on all anomaly trusted lists in Prisma Cloud.

## Example Usage

```hcl
data "prismacloud_anomaly_trusted_lists" "example" {}
```

## Attribute Reference

* `total` - (int) Total number of anomaly trusted lists
* `listing` - List of anomaly trusted list, as defined [below](#listing).

### Listing

* `atl_id` - Anomaly Trusted List ID
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