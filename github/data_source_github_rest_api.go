package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRestApi() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubRestApiRead,

		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"code": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"headers": {
				Type:     schema.TypeMap,
				Computed: true,
			},
			"body": {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubRestApiRead(d *schema.ResourceData, meta interface{}) error {
	u := d.Get("endpoint").(string)

	client := meta.(*Owner).v3client
	ctx := context.Background()

	var body map[string]interface{}

	req, err := client.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}

	resp, _ := client.Do(ctx, req, &body)

	d.SetId(resp.Header.Get("x-github-request-id"))
	if err = d.Set("code", resp.StatusCode); err != nil {
		return err
	}
	if err = d.Set("status", resp.Status); err != nil {
		return err
	}
	if err = d.Set("headers", resp.Header); err != nil {
		return err
	}
	if err = d.Set("body", body); err != nil {
		return err
	}

	return nil
}
