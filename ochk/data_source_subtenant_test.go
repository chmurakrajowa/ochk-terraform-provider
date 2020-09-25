package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

type SubtenantDataSourceTestData struct {
	ResourceName string
	Name         string
}

func (c *SubtenantDataSourceTestData) ToString() string {
	return executeTemplateToString(`
data "ochk_subtenant" "{{ .ResourceName}}" {
  name = "{{ .Name}}"
}
`, c)
}

func (c *SubtenantDataSourceTestData) FullResourceName() string {
	return "data.ochk_subtenant." + c.ResourceName
}

func TestAccSubtenantDataSource_read(t *testing.T) {

	subtenant := &SubtenantDataSourceTestData{
		ResourceName: "default",
		Name:         testDataSubtenant1Name,
	}

	resourceName := subtenant.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: subtenant.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", subtenant.Name),
					//FIXME waiting for backend fix
					//resource.TestCheckResourceAttr(resourceName, "description", subtenant.Description),
					resource.TestCheckResourceAttr(resourceName, "email", "test@example.com"),
					resource.TestCheckResourceAttr(resourceName, "memory_reserved_size_mb", "30000"),
					resource.TestCheckResourceAttr(resourceName, "storage_reserved_size_gb", "200"),
					resource.TestCheckResourceAttr(resourceName, "users.0", "bf5e40c6-191c-40f5-b4d1-9332a9e4ed48"),
					resource.TestCheckResourceAttr(resourceName, "networks.0", "bd814070-18f3-4182-b2af-edaa72a50fee"),
				),
			},
		},
	})
}