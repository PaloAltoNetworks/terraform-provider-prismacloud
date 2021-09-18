---
page_title: "Prisma Cloud: prismacloud_reports"
---

# prismacloud_reports

Retrieve a list of alert reports and compliance reports.

## Example Usage

```hcl
data "prismacloud_reports" "example" {}
```

## Attribute Reference

* `total` - (int) Total number of reports.
* `listing` - List of reports returned, as defined [below](#listing).

### Listing

Each report has the following attributes:

* `report_id` - Report ID
* `name` - Report name
* `report_type` - Report type
* `cloud_type` - Cloud type
* `status` - Report status
