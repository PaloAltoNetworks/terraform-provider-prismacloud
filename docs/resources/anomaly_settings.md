---
page_title: "Prisma Cloud: prismacloud_anomaly_settings"
---

# prismacloud_anomaly_settings

Manage an anomaly setting.

## Example Usage

```hcl
resource "prismacloud_anomaly_settings" "example" {
    policy_id = "policy ID"
}
```

## Argument Reference

You must specify following: 

* `policy_id` - (Required) Policy ID
* `alert_disposition` - (Optional) Alert disposition. Valid values are `aggressive`, `moderate`, or `conservative`.
* `training_model_threshold` - (Optional) Training model threshold information. Valid values are `low`, `medium`, or `high`.

## Attribute Reference

* `alert_disposition_description` - Alert disposition information [below](#alert-disposition-description)
* `policy_description` - Policy description
* `policy_name` - Policy name
* `training_model_description` - Training model info [below](#training-model-description)

### Alert Disposition Description

* `aggressive` - Aggressive
* `moderate` - Moderate
* `conservative` - Conservative

### Training Model Description

* `low` - Low
* `medium` - Medium 
* `high` - High

```
$ terraform import prismacloud_anomaly_settings.example 11111111-2222-3333-4444-555555555555
```