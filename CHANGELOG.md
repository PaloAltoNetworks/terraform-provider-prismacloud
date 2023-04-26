## 1.3.6 (April 26, 2023)

NEW RESOURCES: 

* `prismacloud_trusted_login_ip_status`


## 1.3.5 (April 26, 2023)

* Enterprise settings audit log support.
* The following new data sources are added, mainly related to GCP and IBM Cloud account onboarding with Security Capabilities and Permissions and to generate terraform template for GCP and IBM.

  * data 'prismacloud_gcp_template' - For GCP terraform template generation.
  * data 'prismacloud_ibm_template' - For IBM terraform template generation.
 
NEW DATA SOURCES:

* `prismacloud_trusted_alert_ip` / `prismacloud_trusted_alert_ips`

NEW RESOURCES: 

* `prismacloud_trusted_alert_ip`


## 1.3.4 (April 18, 2023)

* Alert Rule Policy Filter for Compute Access Group IDs.
* Bug fix for integration rate limit issue.
* The following new data sources are added, mainly related to AZURE Cloud account onboarding with Security Capabilities and Permissions and to generate terraform template for azure.

  * data 'prismacloud_azure_template' - For AZURE terraform template generation.
 
NEW DATA SOURCES:

* `prismacloud_trusted_login_ip` / `prismacloud_trusted_login_ips`

NEW RESOURCES: 

* `prismacloud_trusted_login_ip`



## 1.3.3 (March 29, 2023)

* terraform provider bug fix for `prismacloud_permission_group`


## 1.3.2 (March 28, 2023)

* Enterprise settings bug resolution
* Prismacloud resources state update when deleted outside of terraform fixed for account_group, alert_rule and compliance_standard_requirement

NEW DATA SOURCES:

* `prismacloud_permission_group` / `prismacloud_permission_groups`

NEW RESOURCES: 

* `prismacloud_permission_group`


## 1.3.1 (March 10, 2023)

* Documentation fix

## 1.3.0 (March 10, 2023)

* The following new data sources and resources are added, mainly related to AWS Cloud account onboarding with Security Capabilities and Permissions and to generate External ID in the CFT for IAM Role creation.

  * data 'prismacloud_aws_cft_generator' - For AWS CFT Generation 
  * data 'prismacloud_account_supported_features' - For Fetching supported features for cloud type and account type.
  * data 'prismacloud_cloud_account_v2' 
  * data 'prismacloud_org_cloud_account_v2'
  * resource 'prismacloud_cloud_account_v2' - For Onboarding AWS Cloud Account
  * resource 'prismacloud_org_cloud_account_v2' - For Onboarding AWS Cloud Organization
* Added support for WIF credentials for gcp org account
* Support for updating saved_search query

  
## 1.2.11 (Dec 23, 2022)

* Added support for 'Critical' and 'Informational' severity for policy
* Added support for policy rule type 'NetworkConfig'

## 1.2.10 (Oct 28, 2022)

* Added support for 'Developer' and 'NetSecOps' user roles
* Added support for 'skipResult' flag in RQL search resource

## 1.2.9 (September 22, 2022)

NEW DATA SOURCES:

* `anomaly_setting` / `anomaly_settings`
* `anomaly_trusted_list`/ `anomaly_trusted_lists`

NEW RESOURCES: 

* `anomaly_settings`
* `anomaly_trusted_list`

* Custom build policy and code security policy support.

## 1.2.8 (August 17, 2022)

* Terraform SDK upgrade

## 1.2.7 (July 11, 2022)

* Azure ORG hierarchy support
* Documentation update for lookahead notice for AWS account onboarding

## 1.2.6 (April 26, 2022)

* Fix for os/arch support

## 1.2.4 (Apr 26, 2022)

* Bug fixes
* AWS ORG hierarchy support

## 1.2.3 (Feb 22, 2022)

* Bug fixes
* Documentation updated

## 1.2.2 (Feb 3, 2022)

* Documentation updated

## 1.2.1 (Jan 28, 2022)

* Bug fixes

## 1.2.0 (Jan 27, 2022)

NEW DATA SOURCES:

* `prismacloud_user_profile` / `prismacloud_user_profiles`
* `prismacloud_user_role` / `prismacloud_user_roles`

NEW RESOURCES:

* `prismacloud_user_profile`

## 1.1.10 (Jan 13, 2022)

* Bug fixes

## 1.1.9 (Jan 13, 2022)

* Bug fixes

## 1.1.8 (Jan 12, 2022)

* Bug fixes

## 1.1.7 (Jan 12, 2022)

* Bug fixes

## 1.1.6 (Dec 9, 2021)

* Bug fixes

## 1.1.5 (Oct 29, 2021)

* Multiple integrations added as well as data security

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
