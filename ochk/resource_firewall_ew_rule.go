package ochk

import (
	"context"
	"fmt"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	FirewallEWRuleRetryTimeout = 1 * time.Minute
)

func resourceFirewallEWRule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFirewallEWRuleCreate,
		ReadContext:   resourceFirewallEWRuleRead,
		UpdateContext: resourceFirewallEWRuleUpdate,
		DeleteContext: resourceFirewallEWRuleDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(FirewallEWRuleRetryTimeout),
			Update: schema.DefaultTimeout(FirewallEWRuleRetryTimeout),
			Delete: schema.DefaultTimeout(FirewallEWRuleRetryTimeout),
		},

		Schema: map[string]*schema.Schema{
			"security_policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
				Default:  "ALLOW",
			},
			"direction": {
				Type:     schema.TypeString,
				Required: true,
				Default:  "IN_OUT",
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ip_protocol": {
				Type:     schema.TypeString,
				Required: true,
				Default:  "IPV4_IPV6",
			},
			"services": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"source": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"destination": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"position": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"revise_operation": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceFirewallEWRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallEWFRules

	securityPolicyID := d.Get("security_policy_id").(string)

	firewallEWRule := &models.DFWRule{
		DisplayName: d.Get("display_name").(string),
		Action:      d.Get("action").(string),
		Direction:   d.Get("direction").(string),
	}

	if disabled, ok := d.GetOk("disabled"); ok && disabled.(bool) {
		firewallEWRule.Disabled = true
	}

	if ipProtocol, ok := d.GetOk("ip_protocol"); ok {
		firewallEWRule.IPProtocol = ipProtocol.(string)
	}

	if services, ok := d.GetOk("services"); ok {
		firewallEWRule.DefaultServices = expandServicesFromIDs(services.(*schema.Set).List())
	}

	if source, ok := d.GetOk("source"); ok {
		firewallEWRule.Source = expandSecurityGroupFromIDs(source.(*schema.Set).List())
	}

	if destination, ok := d.GetOk("destination"); ok {
		firewallEWRule.Destination = expandSecurityGroupFromIDs(destination.(*schema.Set).List())
	}

	if position, ok := d.GetOk("position"); ok {
		firewallEWRule.Position = expandFirewallRulePosition(position.([]interface{}))
	}

	created, err := proxy.Create(ctx, securityPolicyID, firewallEWRule)
	if err != nil {
		return diag.Errorf("error while creating firewall EW rule: %+v", err)
	}

	d.SetId(created.RuleID)

	return resourceFirewallEWRuleRead(ctx, d, meta)
}

func resourceFirewallEWRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallEWFRules

	securityPolicyID := d.Get("security_policy_id").(string)

	firewallEWRule, err := proxy.Read(ctx, securityPolicyID, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			d.SetId("")
			return diag.Errorf("firewall EW rule with id %s not found: %+v", d.Id(), err)
		}

		return diag.Errorf("error while reading firewall EW rule: %+v", err)
	}

	if err := d.Set("display_name", firewallEWRule.DisplayName); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("action", firewallEWRule.Action); err != nil {
		return diag.Errorf("error setting action: %+v", err)
	}

	if err := d.Set("direction", firewallEWRule.Direction); err != nil {
		return diag.Errorf("error setting direction: %+v", err)
	}

	if err := d.Set("disabled", firewallEWRule.Disabled); err != nil {
		return diag.Errorf("error setting disabled: %+v", err)
	}

	if err := d.Set("ip_protocol", firewallEWRule.IPProtocol); err != nil {
		return diag.Errorf("error setting ip_protocol: %+v", err)
	}

	if err := d.Set("services", flattenServicesFromIDs(firewallEWRule.DefaultServices)); err != nil {
		return diag.Errorf("error setting services: %+v", err)
	}

	if err := d.Set("source", flattenSecurityGroupFromIDs(firewallEWRule.Source)); err != nil {
		return diag.Errorf("error setting source: %+v", err)
	}

	if err := d.Set("destination", flattenSecurityGroupFromIDs(firewallEWRule.Destination)); err != nil {
		return diag.Errorf("error setting destination: %+v", err)
	}

	if err := d.Set("position", flattenFirewallRulePosition(firewallEWRule.Position)); err != nil {
		return diag.Errorf("error setting position: %+v", err)
	}

	return nil
}

func resourceFirewallEWRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.FromErr(fmt.Errorf("updating firewall EW rule is not implemented"))
}

func resourceFirewallEWRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).FirewallEWFRules

	securityPolicyID := d.Get("security_policy_id").(string)

	err := proxy.Delete(ctx, securityPolicyID, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			d.SetId("")
			return diag.Errorf("firewall EW rule with id %s not found: %+v", d.Id(), err)
		}

		return diag.Errorf("error while reading deleting firewall EW rule: %+v", err)
	}

	return nil
}