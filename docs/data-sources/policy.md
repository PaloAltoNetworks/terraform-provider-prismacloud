---
page_title: "Prisma Cloud: prismacloud_policy"
---

# prismacloud_policy

Retrieve information on a specific policy.

## Example Usage

```hcl
data "prismacloud_policy" "example" {
    name = "My Policy"
}
```

## Argument Reference

You must specify at least one of the following:

* `policy_id` - Policy ID
* `name` - Policy name

## Attribute Reference

* `policy_type` - Policy type
* `system_default` - (bool) If policy is a system default policy or not
* `description` - Description
* `severity` - Severity
* `recommendation` - Remediation recommendation
* `cloud_type` - Cloud type
* `labels` - List of labels
* `enabled` - (bool) Enabled
* `created_on` - (int) Created on
* `created_by` - Created by
* `last_modified_on` - (int) Last modified on
* `last_modified_by` - Last modified by
* `rule_last_modified_on` - (int) Rule last modified on
* `overridden` - (bool) Overridden
* `deleted` - (bool) Deleted
* `restrict_alert_dismissal` - (bool) Restrict alert dismissal
* `open_alerts_count` - (int) Open alerts count
* `owner` - Owner
* `policy_mode` - Policy mode
* `remediable` - (bool) Is remediable or not
* `rule` - Model for the rule, as defined [below](#rule)
* `remediation` - Model for remediation, as defined [below](#remediation)
* `compliance_metadata` - List of compliance data.  Each item has compliance standard, requirement, and/or section information, as defined [below](#compliance-metadata)

### Rule

* `name` - Name
* `cloud_type` - Cloud type
* `cloud_account` - Cloud account
* `resource_type` - Resource type
* `api_name` - API name
* `resource_id_path` - Resource ID path
* `criteria` - Saved search ID that defines the rule criteria
* `parameters` - (map of strings) Parameters
* `rule_type` - Type of rule or RQL query

### Remediation

* `template_type` - Template type
* `description` - Description
* `cli_script_template` - CLI script template
* `cli_script_json_schema_string` - CLI script JSON schema

### Compliance Metadata

* `standard_name` - Compliance standard name
* `standard_description` - Compliance standard description
* `requirement_id` - Requirement ID
* `requirement_name` - Requirement name
* `requirement_description` - Requirement description
* `section_id` - Section ID
* `section_description` - Section description
* `policy_id` - Policy ID
* `compliance_id` - Compliance ID
* `section_label` - Section label
* `custom_assigned` - (bool) Custom assigned
