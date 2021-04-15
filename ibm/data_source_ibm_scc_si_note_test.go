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

func TestAccIBMNoteDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMNoteDataSourceConfig(provider_id, note_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_scc_si_note.note", "provider_id", provider_id),
					resource.TestCheckResourceAttr("data.ibm_scc_si_note.note", "note_id", note_id),
				),
			},
		},
	})
}

func testAccCheckIBMNoteDataSourceConfig(provider_id string, note_id string) string {
	return fmt.Sprintf(`
data "ibm_scc_si_note" "note" {
	provider_id = "%s"
	note_id = "%s"
}`, provider_id, note_id)
}
