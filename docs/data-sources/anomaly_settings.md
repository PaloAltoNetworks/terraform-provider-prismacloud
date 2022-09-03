---
page_title: "Prisma Cloud: prismacloud_anomaly_settings"
---

# prismacloud_anomaly_settings

Data source to return information about all anomaly settings in Prisma Cloud.

## Example Usage

```hcl
resource "prismacloud_anomaly_settings" "example" {
   type = "Network"
}
```

## Argument Reference

* `type` - (Required) Type. Valid values are `Network`, `UEBA`.

## Attribute Reference

* `total` - (int) Total number of anomaly settings.
* `listing` - List of anomaly settings, as defined [below](#listing).

## Listing

* `alert_disposition` - Alert disposition
* `alert_disposition_description` - Alert disposition information [below](#alert-disposition-description)
* `policy_description` - Policy description
* `policy_id` - Policy ID
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

