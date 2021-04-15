/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSCCSIProvidersDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMProvidersDataSourceConfig(provider_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_scc_si_providers.providers", "providers.0.id", provider_id),
				),
			},
		},
	})
}

func testAccCheckIBMProvidersDataSourceConfig(provider_id string) string {
	return fmt.Sprintf(`
	resource "ibm_scc_si_note" "finding" {
		provider_id                      = %v
		note {
		  short_description              = "short_description"
		  long_description               = "long_description"
		  kind                           = "FINDING"
		  related_url {
			label                        = "rel_label"
			url                          = "rel_url"
		  }
		  expiration_time                = "2006-01-02 15:04:11"
		  create_time                    = "2006-01-02 15:04:00"
		  update_time                    = "2006-01-02 15:04:11"
		  id                             = "test-finding"
		  shared                         = "true"
		  reported_by {
			id                           = "rep_id"
			title                        = "rep_title"
			url                          = "rep_url"
		  }
		  finding {
			severity                     = "LOW"
			next_steps {
			  url                        = "next_url"
			  title                      = "next_title"
			}
		  }
		}
	  }`, provider_id)
}
