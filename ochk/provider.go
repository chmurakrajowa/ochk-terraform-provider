package ochk

import (
	"context"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validateHost,
				DefaultFunc:      schema.EnvDefaultFunc("OCHK_HOST", nil),
				Description:      "host value",
			},
			"tenant": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OCHK_TENANT", nil),
				Description: "tenant value",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OCHK_USERNAME", nil),
				Description: "username value",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OCHK_PASSWORD", nil),
				Description: "password value",
				Sensitive:   true,
			},
			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "if set uses http scheme instead of https",
				Default:     false,
			},
			"debug_log_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "path to debug log",
				Sensitive:   true,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"ochk_security_group":  dataSourceSecurityGroup(),
			"ochk_security_policy": dataSourceSecurityPolicy(),
			"ochk_gateway_policy":  dataSourceGatewayPolicy(),
			"ochk_service":         dataSourceService(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"ochk_security_group":   resourceSecurityGroup(),
			"ochk_firewall_ew_rule": resourceFirewallEWRule(),
			"ochk_firewall_sn_rule": resourceFirewallSNRule(),
		},
		ConfigureContextFunc: func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
			client, err := sdk.NewClient(
				ctx,
				d.Get("host").(string),
				d.Get("tenant").(string),
				d.Get("username").(string),
				d.Get("password").(string),
				d.Get("insecure").(bool),
				d.Get("debug_log_file").(string),
			)
			if err != nil {
				return nil, diag.FromErr(err)
			}

			return client, nil
		},
	}
}

func validateHost(val interface{}, path cty.Path) diag.Diagnostics {
	if val == nil || val.(string) == "" {
		diag.Errorf("%s value is not valid: %s", path, val.(string))
	}

	return nil
}
