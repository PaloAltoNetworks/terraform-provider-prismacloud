package prismacloud

import (
	"os"
	"testing"

	_ "github.com/paloaltonetworks/prisma-cloud-go"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const (
	PrismacloudJsonConfigFileEnvVar = "PRISMACLOUD_JSON_CONFIG_FILE"
)

var (
	testAccProviders map[string]terraform.ResourceProvider
	testAccProvider  *schema.Provider
)

func init() {
	//var err error

	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"prismacloud": testAccProvider,
	}

	/*
	   client := pc.Client{}
	   if err = client.Initialize(os.Getenv(PrismacloudJsonConfigFileEnvVar)); err == nil {
	   }
	*/
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if os.Getenv(PrismacloudJsonConfigFileEnvVar) == "" {
		t.Fatalf("%s must be set for acceptance tests", PrismacloudJsonConfigFileEnvVar)
	}
}
