package prismacloud

import (
	"fmt"
	"testing"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/integration"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIntegrationAmazonSqs(t *testing.T) {
	var o integration.Integration
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationConfig(name, "amazon_sqs", 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIntegrationExists("prismacloud_integration.test", &o),
					testAccCheckIntegrationAttributes(&o, name, "amazon_sqs", 1),
				),
			},
			{
				Config: testAccIntegrationConfig(name, "amazon_sqs", 2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIntegrationExists("prismacloud_integration.test", &o),
					testAccCheckIntegrationAttributes(&o, name, "amazon_sqs", 2),
				),
			},
		},
	})
}

func TestAccIntegrationQualys(t *testing.T) {
	var o integration.Integration
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationConfig(name, "qualys", 2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIntegrationExists("prismacloud_integration.test", &o),
					testAccCheckIntegrationAttributes(&o, name, "qualys", 2),
				),
			},
			{
				Config: testAccIntegrationConfig(name, "qualys", 3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIntegrationExists("prismacloud_integration.test", &o),
					testAccCheckIntegrationAttributes(&o, name, "qualys", 3),
				),
			},
		},
	})
}

func TestAccIntegrationServiceNow(t *testing.T) {
	var o integration.Integration
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationConfig(name, "service_now", 78005),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIntegrationExists("prismacloud_integration.test", &o),
					testAccCheckIntegrationAttributes(&o, name, "service_now", 78005),
				),
			},
			{
				Config: testAccIntegrationConfig(name, "service_now", 86411),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIntegrationExists("prismacloud_integration.test", &o),
					testAccCheckIntegrationAttributes(&o, name, "service_now", 86411),
				),
			},
		},
	})
}

func TestAccIntegrationWebhook(t *testing.T) {
	var o integration.Integration
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationConfig(name, "webhook", 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIntegrationExists("prismacloud_integration.test", &o),
					testAccCheckIntegrationAttributes(&o, name, "webhook", 1),
				),
			},
			{
				Config: testAccIntegrationConfig(name, "webhook", 2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIntegrationExists("prismacloud_integration.test", &o),
					testAccCheckIntegrationAttributes(&o, name, "webhook", 2),
				),
			},
		},
	})
}

func TestAccIntegrationPagerDuty(t *testing.T) {
	var o integration.Integration
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationConfig(name, "pager_duty", 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIntegrationExists("prismacloud_integration.test", &o),
					testAccCheckIntegrationAttributes(&o, name, "pager_duty", 1),
				),
			},
			{
				Config: testAccIntegrationConfig(name, "pager_duty", 2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIntegrationExists("prismacloud_integration.test", &o),
					testAccCheckIntegrationAttributes(&o, name, "pager_duty", 2),
				),
			},
		},
	})
}

func testAccCheckIntegrationExists(n string, o *integration.Integration) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label ID is not set")
		}

		client := testAccProvider.Meta().(*pc.Client)
		id := rs.Primary.ID
		lo, err := integration.Get(client, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}

		*o = lo

		return nil
	}
}

func testAccCheckIntegrationAttributes(o *integration.Integration, name, it string, num int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Name != name {
			return fmt.Errorf("Name is %s, expected %s", o.Name, name)
		}

		if o.IntegrationType != it {
			return fmt.Errorf("Integration type is %q, expected %q", o.IntegrationType, it)
		}

		desc := fmt.Sprintf("integration acctest for %s", it)
		if o.Description != desc {
			return fmt.Errorf("Description is %s, expected %s", o.Description, desc)
		}

		return nil
	}
}

func testAccIntegrationDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*pc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "prismacloud_integration" {
			continue
		}

		if rs.Primary.ID != "" {
			id := rs.Primary.ID
			if _, err := integration.Get(client, id); err == nil {
				return fmt.Errorf("Object %q still exists", rs.Primary.ID)
			}
		}
		return nil
	}

	return nil
}

func testAccIntegrationConfig(name, it string, num int) string {
	var ic string

	switch it {
	case "amazon_sqs":
		ic = fmt.Sprintf(`integration_config {
        queue_url = "https://sqs.us-east-1.amazonaws.com/12345678901%d/myintegration"
    }`, num)
	case "qualys":
		ic = fmt.Sprintf(`integration_config {
        login = "qualys%dlogin"
        password = "qualys%dpassword"
        base_url = "qualysapi.qg%d.apps.qualys.com"
    }`, num, num, num)
	case "service_now":
		ic = fmt.Sprintf(`integration_config {
        host_url = "dev%d.service-now.com"
        login = "servicenow%dlogin"
        password = "servicenow%dpassword"
        tables = {
            "incident": %t,
            "sn_si_incident": %t,
        }
        version = "LONDON"
    }`, num, num, num, num == 1, num == 2)
	case "webhook":
		ic = fmt.Sprintf(`integration_config {
        url = "https://webhook.site/4a40c4a7-d531-4190-a934-750e7fba4954"
        headers {
            key = "X-Do-Stuff"
            value = "numero %d"
        }
        headers {
            key = "Content-Type"
            value = "application/json"
            read_only = true
        }
    }`, num)
	case "pager_duty":
		ic = fmt.Sprintf(`integration_config {
        integration_key = "pagerduty%dkey"
        auth_token = "pagerduty-auth-token-%d"
    }`, num, num)
	}

	return fmt.Sprintf(`
resource "prismacloud_integration" "test" {
    name = %q
    description = "integration acctest for %s"
    enabled = false
    integration_type = %q
    %s
}`, name, it, it, ic)
}
