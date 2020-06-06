package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccGithubAdminOrg_basic(t *testing.T) {

	orgResourceName := "resource_github_admin_organization.test_org"
	adminValue := "super_secret_value"
	updatedAdminValue := "updated_super_secret_value"
	t.Log(testAccGithubAdminOrgConfig(adminValue))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubAdminOrgConfig(adminValue),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubAdminOrgExists(orgResourceName, "test_admin_name", t),
					resource.TestCheckResourceAttr("github_admin_org.test_org", "admin", adminValue),
				),
			},
			{
				Config: testAccGithubAdminOrgConfig(updatedAdminValue),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubAdminOrgExists(orgResourceName, "test_admin_name", t),
					resource.TestCheckResourceAttr("github_admin_org.test_org", "admin", updatedAdminValue),
				),
			},
		},
	})
}

func testAccGithubAdminOrgConfig(admin string) string {
	return fmt.Sprintf(`
resource "github_admin_organization" "test_org" {
  admin       = "%s"
}
`, admin)
}

func testAccCheckGithubAdminOrgExists(resourceName, adminName string, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// TODO: Make sure admin org exists

		return nil
	}
}
