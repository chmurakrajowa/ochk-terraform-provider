package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

func generateRandName() string {
	return fmt.Sprintf("tf-acc_test-%s", acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum))
}
