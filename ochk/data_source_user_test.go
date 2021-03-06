package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

type UserDataSourceTestData struct {
	ResourceName string
	Name         string
}

func (c *UserDataSourceTestData) ToString() string {
	return executeTemplateToString(`
data "ochk_user" "{{ .ResourceName}}" {
  name = "{{ .Name}}"
}
`, c)
}

func (c *UserDataSourceTestData) FullResourceName() string {
	return "data.ochk_user." + c.ResourceName
}

func TestAccUserDataSource_read(t *testing.T) {
	user := UserDataSourceTestData{
		ResourceName: "default",
		Name:         testData.User1Name,
	}

	resourceName := user.FullResourceName()

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: user.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", user.Name),
					resource.TestCheckResourceAttr(resourceName, "email_address", "terraform@ochk.pl"),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "first_name", "Terraform"),
					resource.TestCheckResourceAttr(resourceName, "last_name", "User"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "locked", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "user_principal_name"),
					resource.TestCheckResourceAttrSet(resourceName, "principal_id"),
					resource.TestCheckResourceAttrSet(resourceName, "principal_name"),
					resource.TestCheckResourceAttrSet(resourceName, "principal_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "principal_type"),
				),
			},
		},
	})
}
