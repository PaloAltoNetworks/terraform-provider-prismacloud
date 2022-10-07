---
page_title: "Prisma Cloud: prismacloud_anomaly_setting"
---

# prismacloud_anomaly_setting

Retrieve information on a specific anomaly setting.

## Example Usage

```hcl
data "prismacloud_anomaly_setting" "example" {
    policy_id = "Policy id"
}
```

## Argument Reference

You must specify following: 

* `policy_id` - (Required) Policy ID

## Attribute Reference

* `alert_disposition` - Alert disposition
* `alert_disposition_description` - Alert disposition information [below](#alert-disposition-description)
* `policy_description` - Policy description
* `policy_name` - Policy name
* `training_model_description` - Training model information [below](#training-model-description)
* `training_model_threshold` - Training model threshold information

### Alert Disposition Description

* `aggressive` - Aggressive
* `moderate` - Moderate
* `conservative` - Conservative

### Training Model Description

* `low` - Low
* `medium` - Medium
* `high` - High


