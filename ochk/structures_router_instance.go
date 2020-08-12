package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
)

func flattenRouterInstancesFromIDs(m []*models.RouterInstance) *schema.Set {
	s := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range m {
		s.Add(v.RouterID)
	}
	return s
}

func expandRouterInstancesFromIDs(in []interface{}) []*models.RouterInstance {
	if len(in) == 0 {
		return nil
	}

	var out = make([]*models.RouterInstance, len(in))

	for i, v := range in {
		securityGroup := &models.RouterInstance{
			RouterID: v.(string),
		}

		out[i] = securityGroup
	}

	return out
}