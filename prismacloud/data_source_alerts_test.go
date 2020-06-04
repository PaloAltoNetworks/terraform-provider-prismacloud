package prismacloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDsAbsoluteAlerts(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsAlertsConfig("absolute"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_alerts.test", "total"),
				),
			},
		},
	})
}

func TestAccDsRelativeAlerts(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsAlertsConfig("relative"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_alerts.test", "total"),
				),
			},
		},
	})
}

func TestAccDsToNowAlerts(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsAlertsConfig("to_now"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_alerts.test", "total"),
				),
			},
		},
	})
}

func testAccDsAlertsConfig(ct string) string {
	switch ct {
	case "absolute":
		return `
data "prismacloud_alerts" "test" {
    limit = 2
    time_range {
        absolute {
            start = 1504448933000
            end = 1504794533000
        }
    }
}
`
	case "relative":
		return `
data "prismacloud_alerts" "test" {
    limit = 2
    time_range {
        relative {
            amount = 48
            unit = "hour"
        }
    }
}
`
	case "to_now":
		return `
data "prismacloud_alerts" "test" {
    limit = 2
    time_range {
        to_now {
            unit = "login"
        }
    }
}
`
	}

	return ""
}
