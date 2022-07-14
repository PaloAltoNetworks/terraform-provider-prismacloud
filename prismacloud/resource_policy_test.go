package prismacloud

import (
	"bytes"
	"fmt"
	"testing"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/policy"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPolicyConfig(t *testing.T) {
	var o policy.Policy
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyConfig(name, "config", "Config", "desc1", "aws", "low", map[string]string{"savedSearch": "true"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyExists("prismacloud_policy.test", &o),
					testAccCheckPolicyAttributes(&o, name, "config", "Config", "desc1", "aws", "low", map[string]string{"savedSearch": "true"}),
				),
			},
			{
				Config: testAccPolicyConfig(name, "config", "Config", "desc2", "aws", "medium", map[string]string{"savedSearch": "true", "withIac": "true"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyExists("prismacloud_policy.test", &o),
					testAccCheckPolicyAttributes(&o, name, "config", "Config", "desc2", "aws", "medium", map[string]string{"savedSearch": "true", "withIac": "true"}),
				),
			},
		},
	})
}

func TestAccPolicyNetwork(t *testing.T) {
	var o policy.Policy
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyConfig(name, "network", "Network", "desc1", "all", "low", map[string]string{"savedSearch": "true"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyExists("prismacloud_policy.test", &o),
					testAccCheckPolicyAttributes(&o, name, "network", "Network", "desc1", "all", "low", map[string]string{"savedSearch": "true"}),
				),
			},
			{
				Config: testAccPolicyConfig(name, "network", "Network", "desc2", "all", "medium", map[string]string{"savedSearch": "true", "withIac": "true"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyExists("prismacloud_policy.test", &o),
					testAccCheckPolicyAttributes(&o, name, "network", "Network", "desc2", "all", "medium", map[string]string{"savedSearch": "true", "withIac": "true"}),
				),
			},
		},
	})
}

func TestAccPolicyAuditEvent(t *testing.T) {
	var o policy.Policy
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyConfig(name, "audit_event", "AuditEvent", "desc1", "all", "low", map[string]string{"savedSearch": "true"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyExists("prismacloud_policy.test", &o),
					testAccCheckPolicyAttributes(&o, name, "audit_event", "AuditEvent", "desc1", "all", "low", map[string]string{"savedSearch": "true"}),
				),
			},
			{
				Config: testAccPolicyConfig(name, "audit_event", "AuditEvent", "desc2", "all", "medium", map[string]string{"savedSearch": "true", "withIac": "true"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyExists("prismacloud_policy.test", &o),
					testAccCheckPolicyAttributes(&o, name, "audit_event", "AuditEvent", "desc2", "all", "medium", map[string]string{"savedSearch": "true", "withIac": "true"}),
				),
			},
		},
	})
}

func testAccCheckPolicyExists(n string, o *policy.Policy) resource.TestCheckFunc {
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
		lo, err := policy.Get(client, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}

		*o = lo

		return nil
	}
}

func testAccCheckPolicyAttributes(o *policy.Policy, name, pt, rt, desc, ct, sev string, params map[string]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Name != name {
			return fmt.Errorf("Name is %s, expected %s", o.Name, name)
		}

		if o.PolicyType != pt {
			return fmt.Errorf("Policy type is %q, expected %q", o.PolicyType, pt)
		}

		if o.Rule.Type != rt {
			return fmt.Errorf("Rule type is %q, expected %q", o.Rule.Type, rt)
		}

		if o.Description != desc {
			return fmt.Errorf("Description is %s, expected %s", o.Description, desc)
		}

		if o.CloudType != ct {
			return fmt.Errorf("Cloud type is %q, expected %q", o.CloudType, ct)
		}

		if o.Severity != sev {
			return fmt.Errorf("Severity is %q, expected %q", o.Severity, sev)
		}

		if len(o.Rule.Parameters) != len(params) {
			return fmt.Errorf("Rule params is len %d, not %d", len(o.Rule.Parameters), len(params))
		}

		for key := range params {
			if o.Rule.Parameters[key] != params[key] {
				return fmt.Errorf("Rule param %q is %q, expected %q", key, o.Rule.Parameters[key], params[key])
			}
		}

		return nil
	}
}

func testAccPolicyDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*pc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "prismacloud_policy" {
			continue
		}

		if rs.Primary.ID != "" {
			id := rs.Primary.ID
			if o, err := policy.Get(client, id); err == nil {
				if !o.Deleted {
					return fmt.Errorf("Object %q still exists", rs.Primary.ID)
				}
			}
		}
		return nil
	}

	return nil
}

func testAccPolicyConfig(name, pt, rt, desc, ct, sev string, params map[string]string) string {
	var buf bytes.Buffer
	buf.Grow(500)

	buf.WriteString(fmt.Sprintf(`
data "prismacloud_rql_historic_searches" "x" {}

locals {
    id_listing = [
        for inst in data.prismacloud_rql_historic_searches.x.listing :
        inst.search_id
        if inst.search_type == %q && inst.created_by == "Prisma Cloud System Admin"
    ]
}

resource "prismacloud_policy" "test" {
    name = %q
    description = %q
    cloud_type = %q
    policy_type = %q
    severity = %q
    enabled = false
    rule {
        name = "my rule"
        criteria = local.id_listing[0]
        rule_type = %q
        parameters = {`, pt, name, desc, ct, pt, sev, rt))

	for key, value := range params {
		buf.WriteString(fmt.Sprintf(`
            %q: %q,`, key, value))
	}

	buf.WriteString(`
        }
    }
}`)

	return buf.String()
}
