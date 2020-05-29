package github

import (
	"context"

	"github.com/google/go-github/v31/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGithubAdminOrganization() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubAdminOrganizationCreate,
		Update: resourceGithubAdminOrganizationUpdate,

		Schema: map[string]*schema.Schema{
			"admin": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceGithubAdminOrganizationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).v3client

	admin := d.Get("admin").(string)
	login := d.Get("login").(string)

	input := &github.Organization{
		Login: &login,
	}

	_, _, err := client.Admin.CreateOrg(context.Background(), input, admin)
	if err != nil {
		return err
	}

	d.SetId(admin)
	return nil
}

func resourceGithubAdminOrganizationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).v3client

	admin := d.Get("admin").(string)
	login := d.Get("login").(string)

	input := &github.Organization{
		Login: &login,
	}

	_, _, err := client.Admin.CreateOrg(context.Background(), input, admin)
	if err != nil {
		return err
	}

	d.SetId(admin)
	return nil
}
