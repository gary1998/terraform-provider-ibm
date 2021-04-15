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
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMNote_FINDING(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMNoteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMNote_FINDING(provider_id, note_id),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMNoteExists(provider_id, note_id),
					resource.TestCheckResourceAttr("ibm_scc_si_note.note", "kind", "FINDING"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.note", "short_description", "short"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.note", "long_description", "long"),
				),
			},
			{
				Config: testAccCheckIBMNote_FINDING_Update(provider_id, note_id),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMNoteExists(provider_id, note_id),
					resource.TestCheckResourceAttr("ibm_scc_si_note.note", "kind", "FINDING"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.note", "short_description", "short description"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.note", "long_description", "long description"),
				),
			},
		},
	})
}

func TestAccIBMNote_KPI(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMNoteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMNote_KPI(provider_id, note_id),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMNoteExists(provider_id, note_id),
					resource.TestCheckResourceAttr("ibm_scc_si_note.note", "kind", "KPI"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.note", "short_description", "short"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.note", "long_description", "long"),
				),
			},
			{
				Config: testAccCheckIBMNote_KPI_Update(provider_id, note_id),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMNoteExists(provider_id, note_id),
					resource.TestCheckResourceAttr("ibm_scc_si_note.note", "kind", "KPI"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.note", "short_description", "short description"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.note", "long_description", "long description"),
				),
			},
		},
	})
}

func TestAccIBMNote_CARD(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMNoteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMNote_CARD(provider_id, note_id),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMNoteExists(provider_id, note_id),
					resource.TestCheckResourceAttr("ibm_scc_si_note.card", "kind", "CARD"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.note", "short_description", "short"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.note", "long_description", "long"),
				),
			},
			{
				Config: testAccCheckIBMNote_CARD_Update(provider_id, note_id),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMNoteExists(provider_id, note_id),
					resource.TestCheckResourceAttr("ibm_scc_si_note.card", "kind", "CARD"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.note", "short_description", "short description"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.note", "long_description", "long description"),
				),
			},
		},
	})
}

func TestAccIBMNote_SECTION(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMNoteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMNote_SECTION(provider_id, note_id),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMNoteExists(provider_id, note_id),
					resource.TestCheckResourceAttr("ibm_scc_si_note.section", "kind", "SECTION"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.note", "short_description", "short"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.note", "long_description", "long"),
				),
			},
			{
				Config: testAccCheckIBMNote_SECTION_Update(provider_id, note_id),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMNoteExists(provider_id, note_id),
					resource.TestCheckResourceAttr("ibm_scc_si_note.section", "kind", "SECTION"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.note", "short_description", "short description"),
					resource.TestCheckResourceAttr("ibm_scc_si_note.note", "long_description", "long description"),
				),
			},
		},
	})
}

func TestAccIBMNote_MutualExclusion(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMNoteDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMNote_MutualExclusion(provider_id, note_id),
				ExpectError: regexp.MustCompile(fmt.Sprintf("only one of `card,finding,kpi,section` can be specified, but `finding,kpi` were specified.")),
			},
		},
	})
}

func testAccCheckIBMNote_FINDING(provider_id string, note_id string) string {
	return fmt.Sprintf(`
	resource "ibm_scc_si_note" "note" {
		provider_id                    = "%s"
		short_description              = "short"
		long_description               = "long"
		kind                           = "FINDING"
		related_url {
		  label   = "rel_label"
		  url     = "rel_url"
		}
		expiration_time                = "2006-01-02 15:04:11"
		create_time                    = "2006-01-02 15:04:00"
		update_time                    = "2006-01-02 15:04:11"
		note_id                        = "%s"
		shared                         = "true"
		reported_by {
		  id    = "rep_id"
		  title = "rep_title"
		  url   = "rep_url"
		}
		finding {
		  severity = "HIGH"
		  next_steps {
			url   = "next_url"
			title = "next_title"
		  }
		}
	  }
	`, provider_id, note_id)
}

func testAccCheckIBMNote_FINDING_Update(provider_id string, note_id string) string {
	return fmt.Sprintf(`
	resource "ibm_scc_si_note" "note" {
		provider_id                    = "%s"
		short_description              = "short description"
		long_description               = "long description"
		kind                           = "FINDING"
		related_url {
		  label   = "rel_label"
		  url     = "rel_url"
		}
		expiration_time                = "2006-01-02 15:04:11"
		create_time                    = "2006-01-02 15:04:00"
		update_time                    = "2006-01-02 15:04:11"
		note_id                        = "%s"
		shared                         = "true"
		reported_by {
		  id    = "rep_id"
		  title = "rep_title"
		  url   = "rep_url"
		}
		finding {
		  severity = "HIGH"
		  next_steps {
			url   = "next_url"
			title = "next_title"
		  }
		}
	  }
	`, provider_id, note_id)
}

func testAccCheckIBMNote_KPI(provider_id string, note_id string) string {
	return fmt.Sprintf(`
	resource "ibm_scc_si_note" "note" {
		provider_id                    = "%s"
		short_description              = "short"
		long_description               = "long"
		kind                           = "KPI"
		related_url {
		  label   = "rel_label"
		  url     = "rel_url"
		}
		expiration_time                = "2006-01-02 15:04:11"
		create_time                    = "2006-01-02 15:04:00"
		update_time                    = "2006-01-02 15:04:11"
		note_id                        = "%s"
		shared                         = "true"
		reported_by {
		  id    = "rep_id"
		  title = "rep_title"
		  url   = "rep_url"
		}
		kpi {
		  aggregation_type = "SUM"
		}
	  }
	`, provider_id, note_id)
}

func testAccCheckIBMNote_KPI_Update(provider_id string, note_id string) string {
	return fmt.Sprintf(`
	resource "ibm_scc_si_note" "note" {
		provider_id                    = "%s"
		short_description              = "short description"
		long_description               = "long description"
		kind                           = "KPI"
		related_url {
		  label   = "rel_label"
		  url     = "rel_url"
		}
		expiration_time                = "2006-01-02 15:04:11"
		create_time                    = "2006-01-02 15:04:00"
		update_time                    = "2006-01-02 15:04:11"
		note_id                        = "%s"
		shared                         = "true"
		reported_by {
		  id    = "rep_id"
		  title = "rep_title"
		  url   = "rep_url"
		}
		kpi {
		  aggregation_type = "SUM"
		}
	  }
	`, provider_id, note_id)
}

func testAccCheckIBMNote_CARD(provider_id string, note_id string) string {
	return fmt.Sprintf(`
	resource "ibm_scc_si_note" "card" {
		provider_id                    = "%s"
		short_description              = "short"
		long_description               = "long"
		kind                           = "CARD"
		related_url {
		  label   = "rel_label"
		  url     = "rel_url"
		}
		expiration_time                = "2006-01-02 15:04:11"
		create_time                    = "2006-01-02 15:04:00"
		update_time                    = "2006-01-02 15:04:11"
		note_id                        = "%s"
		shared                         = "true"
		reported_by {
		  id    = "rep_id"
		  title = "rep_title"
		  url   = "rep_url"
		}
		card {
		  section = "test-section"
		  title = "test-card"
		  subtitle = "test"
		  order = 1
		  finding_note_names = ["providers/terraform-test/notes/test-finding-note"]
		  requires_configuration = false
		  badge_text = "string"
		  badge_image = "string"
		  elements {
			  kind = "NUMERIC"
			  default_time_range = "1d"
			  text = "text"
			  value_type {
						  finding_note_names = [ "providers/terraform-test/notes/test-finding-note" ]
						  kind = "FINDING_COUNT"
					  }
		  }
		}
	  }
	`, provider_id, note_id)
}

func testAccCheckIBMNote_CARD_Update(provider_id string, note_id string) string {
	return fmt.Sprintf(`
	resource "ibm_scc_si_note" "card" {
		provider_id                    = "%s"
		short_description              = "short description"
		long_description               = "long description"
		kind                           = "CARD"
		related_url {
		  label   = "rel_label"
		  url     = "rel_url"
		}
		expiration_time                = "2006-01-02 15:04:11"
		create_time                    = "2006-01-02 15:04:00"
		update_time                    = "2006-01-02 15:04:11"
		note_id                        = "%s"
		shared                         = "true"
		reported_by {
		  id    = "rep_id"
		  title = "rep_title"
		  url   = "rep_url"
		}
		card {
		  section = "test-section"
		  title = "test-card"
		  subtitle = "test"
		  order = 1
		  finding_note_names = ["providers/terraform-test/notes/test-finding-note"]
		  requires_configuration = false
		  badge_text = "string"
		  badge_image = "string"
		  elements {
			  kind = "NUMERIC"
			  default_time_range = "1d"
			  text = "text"
			  value_type {
						  finding_note_names = [ "providers/terraform-test/notes/test-finding-note" ]
						  kind = "FINDING_COUNT"
					  }
		  }
		}
	  }
	`, provider_id, note_id)
}

func testAccCheckIBMNote_SECTION(provider_id string, note_id string) string {
	return fmt.Sprintf(`
	resource "ibm_scc_si_note" "section" {
		provider_id                    = "%s"
		short_description              = "short"
		long_description               = "long"
		kind                           = "SECTION"
		related_url {
		  label   = "rel_label"
		  url     = "rel_url"
		}
		expiration_time                = "2006-01-02 15:04:11"
		create_time                    = "2006-01-02 15:04:00"
		update_time                    = "2006-01-02 15:04:11"
		note_id                        = "%s"
		shared                         = "true"
		reported_by {
		  id    = "rep_id"
		  title = "rep_title"
		  url   = "rep_url"
		}
		section {
		  title = "title"
		  image = "image"
		}
	  }
	`, provider_id, note_id)
}

func testAccCheckIBMNote_SECTION_Update(provider_id string, note_id string) string {
	return fmt.Sprintf(`
	resource "ibm_scc_si_note" "section" {
		provider_id                    = "%s"
		short_description              = "short description"
		long_description               = "long description"
		kind                           = "SECTION"
		related_url {
		  label   = "rel_label"
		  url     = "rel_url"
		}
		expiration_time                = "2006-01-02 15:04:11"
		create_time                    = "2006-01-02 15:04:00"
		update_time                    = "2006-01-02 15:04:11"
		note_id                        = "%s"
		shared                         = "true"
		reported_by {
		  id    = "rep_id"
		  title = "rep_title"
		  url   = "rep_url"
		}
		section {
		  title = "title"
		  image = "image"
		}
	  }
	`, provider_id, note_id)
}

func testAccCheckIBMNote_MutualExclusion(provider_id string, note_id string) string {
	return fmt.Sprintf(`
	resource "ibm_scc_si_note" "note" {
		provider_id                    = "%s"
		short_description              = "short"
		long_description               = "long"
		kind                           = "FINDING"
		related_url {
		  label   = "rel_label"
		  url     = "rel_url"
		}
		expiration_time                = "2006-01-02 15:04:11"
		create_time                    = "2006-01-02 15:04:00"
		update_time                    = "2006-01-02 15:04:11"
		note_id                        = "%s"
		shared                         = "true"
		reported_by {
		  id    = "rep_id"
		  title = "rep_title"
		  url   = "rep_url"
		}
		finding {
		  severity = "HIGH"
		  next_steps {
			url   = "next_url"
			title = "next_title"
		  }
		}
		kpi {
			aggregation_type = "SUM"
		}
	  }
	`, provider_id, note_id)
}

func testAccCheckIBMNoteExists(provider_id string, note_id string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		sess, err := testAccProvider.Meta().(ClientSession).FindingsV1API()
		if err != nil {
			return err
		}

		userDetails, err := testAccProvider.Meta().(ClientSession).BluemixUserDetails()
		if err != nil {
			return err
		}
		accountID := userDetails.userAccount

		getNoteOptions := sess.NewGetNoteOptions(accountID, provider_id, note_id)
		_, resp, err := sess.GetNote(getNoteOptions)

		if err != nil {
			return fmt.Errorf("[ERROR] error occurred while reading note: %v", err)
		}

		if resp.StatusCode == 404 {
			return fmt.Errorf("Not found: %s", note_id)
		}

		return nil
	}
}

func testAccCheckIBMNoteDestroy(s *terraform.State) error {
	sess, err := testAccProvider.Meta().(ClientSession).FindingsV1API()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_si_note" {
			continue
		}

		providerID := rs.Primary.Attributes["provider_id"]
		noteID := rs.Primary.ID

		userDetails, err := testAccProvider.Meta().(ClientSession).BluemixUserDetails()
		if err != nil {
			return err
		}
		accountID := userDetails.userAccount

		getNoteOptions := sess.NewGetNoteOptions(accountID, providerID, noteID)
		_, resp, err := sess.GetNote(getNoteOptions)
		if resp.StatusCode != 404 {
			return fmt.Errorf("Resource key still exists: %s", noteID)
		}
	}

	return nil
}
