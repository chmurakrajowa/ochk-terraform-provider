package ochk

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk"
)

func dataSourceGatewayPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGatewayPolicyRead,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceGatewayPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).GatewayPolicy

	displayName := d.Get("display_name").(string)

	services, err := proxy.ListByDisplayName(ctx, displayName)
	if err != nil {
		return diag.Errorf("error while listing gateway policies: %+v", err)
	}

	if len(services) < 1 {
		return diag.Errorf("no gateway policy found for display_name: %s", displayName)
	}

	if len(services) > 1 {
		return diag.Errorf("more than one gateway policy with display_name: %s found!", displayName)
	}

	d.SetId(services[0].GatewayPolicyID)

	return nil
}