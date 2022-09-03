---
page_title: "Prisma Cloud: prismacloud_anomaly_trusted_list"
---

# prismacloud_anomaly_trusted_list

Manage an anomaly trusted list.

## Example Usage

```hcl
resource "prismacloud_anomaly_trusted_list" "example" {
    atl_id = id
}
```

## Argument Reference

* `name` - (Required) Anomaly Trusted List name
* `description` - (Optional) Reason for trusted listing
* `trusted_list_type` - (Required) Anomaly Trusted List type. Valid values : `ip`, `resource`, `image`, `tag`, `service`, `port`, `subject`, `domain` or `protocol`,
* `account_id` - (Optional) Anomaly Trusted List account id. Default value is `any`.
* `applicable_policies` - (Required) Applicable Policies
* `vpc` - (Optional) VPC. Default value is `any`.
* `trusted_list_entries` - (Required) List of network anomalies in the trusted list [below](#trusted-list-entries).

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

```
$ terraform import prismacloud_anomaly_trusted_list.example 11111111-2222-3333-4444-555555555555
```