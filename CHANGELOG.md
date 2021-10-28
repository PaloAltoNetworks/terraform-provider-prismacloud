## 1.1.4 (Sep 21, 2021)

* Added missing network RQL docs

## 1.1.3 (Sep 21, 2021)

* Added support for AWS orgs, anomaly policies, adding cloud accounts, numerous fixes to alert rules syntax, and bug fixes.  Documentation updated.

## 1.1.2 (April 19, 2021)

* Fixed self ref policy id in compliance section of policy

## 1.1.1 (April 9, 2021)

* Added new poller to handle master/replica delays in API

## 1.1.0 (February 5, 2021)

NEW RESOURCES:

* `prismacloud_rql_search`
* `prismacloud_saved_search`

## 1.0.8 (November 18, 2020)

* Fixed azure key to sensitive

## 1.0.7 (November 18, 2020)

* Fixed notification config in alert rules
* Fixed provider logging setting

## 1.0.6 (November 12, 2020)

* Documentation fixes

## 1.0.5 (November 11, 2020)

* Fixed AWS, Azure, GCP cloud account for API changes in Prisma Cloud

## 1.0.4 (October 29, 2020)

* Adding support for proxy

## 1.0.3 (October 21, 2020)

* Minor bug fix

## 1.0.2 (September 24, 2020)

* Documentation fixes

## 1.0.1 (August 28, 2020)

* Terraform Registry release

## 1.0.0 (July 07, 2020)

NEW DATA SOURCES:

* `prismacloud_account_group` / `prismacloud_account_groups`
* `prismacloud_alert_rule` / `prismacloud_alert_rules`
* `prismacloud_alerts`
* `prismacloud_cloud_account` / `prismacloud_cloud_accounts`
* `prismacloud_compliance_standard` / `prismacloud_compliance_standards`
* `prismacloud_compliance_standard_requirement` / `prismacloud_compliance_standard_requirements`
* `prismacloud_compliance_standard_requirement_section` / `prismacloud_compliance_standard_requirement_sections`
* `prismacloud_enterprise_settings`
* `prismacloud_integration` / `prismacloud_integrations`
* `prismacloud_policy` / `prismacloud_policies`
* `prismacloud_rql_historic_search` / `prismacloud_rql_historic_searches`

NEW RESOURCES:

* `prismacloud_account_group`
* `prismacloud_alert_rule`
* `prismacloud_cloud_account`
* `prismacloud_compliance_standard`
* `prismacloud_compliance_standard_requirement`
* `prismacloud_compliance_standard_requirement_section`
* `prismacloud_enterprise_settings`
* `prismacloud_integration`
* `prismacloud_policy`
* `prismacloud_user_role`
