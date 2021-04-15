# # #################################

# # List providers
# data "ibm_scc_si_providers" "providers" {
#   account_id                         = var.account_id
# }
# output "ibm_scc_si_providers" {
#   value                              = data.ibm_scc_si_providers.providers
# }

# # #################################

# # List notes
# data "ibm_scc_si_notes" "notes" {
#   account_id                         = var.account_id
#   provider_id                        = var.provider_id
# }
# output "ibm_scc_si_notes" {
#   value                              = data.ibm_scc_si_notes.notes
# }

# # #################################

# # Get note 
# data "ibm_scc_si_note" "note" {
#   account_id                         = var.account_id
#   provider_id                        = var.provider_id
#   note_id                            = var.note_id
# }
# output "ibm_scc_si_note" {
#   value                              = data.ibm_scc_si_note.note
# }

# # #################################

# # Create Card
# resource "ibm_scc_si_note" "card" {
#   account_id                         = var.account_id
#   provider_id                        = var.provider_id
#   note {
#     short_description                = "short_description"
#     long_description                 = "long_description"
#     kind                             = "CARD"
#     related_url {
#       label                          = "rel_label"
#       url                            = "rel_url"
#     }
#     expiration_time                  = "2006-01-02 15:04:11"
#     create_time                      = "2006-01-02 15:04:00"
#     update_time                      = "2006-01-02 15:04:11"
#     id                               = "test-card"
#     shared                           = "true"
#     reported_by {
#       id                             = "rep_id"
#       title                          = "rep_title"
#       url                            = "rep_url"
#     }
#     card {
#       section                        = "test-section"
#       title                          = "test-card"
#       subtitle                       = "test"
#       order                          = 1
#       finding_note_names             = ["providers/terraform-test/notes/test-finding-note"]
#       requires_configuration         = false
#       badge_text                     = "string"
#       badge_image                    = "string"
#       elements {
#           kind                       = "NUMERIC"
#           default_time_range         = "1d"
#           text                       = "text"
#           value_type {
#             finding_note_names       = [ "providers/terraform-test/notes/test-finding-note" ]
#             kind                     = "FINDING_COUNT"
#           }
#       }
#     }
#   }
# }

# # #################################

# # Create Finding
# resource "ibm_scc_si_note" "finding" {
#   account_id                       = var.account_id
#   provider_id                      = var.provider_id
#   note {
#     short_description              = "short_description"
#     long_description               = "long_description"
#     kind                           = "FINDING"
#     related_url {
#       label                        = "rel_label"
#       url                          = "rel_url"
#     }
#     expiration_time                = "2006-01-02 15:04:11"
#     create_time                    = "2006-01-02 15:04:00"
#     update_time                    = "2006-01-02 15:04:11"
#     id                             = "test-finding"
#     shared                         = "true"
#     reported_by {
#       id                           = "rep_id"
#       title                        = "rep_title"
#       url                          = "rep_url"
#     }
#     finding {
#       severity                     = "LOW"
#       next_steps {
#         url                        = "next_url"
#         title                      = "next_title"
#       }
#     }
#   }
# }

# # #################################

# # Create KPI
# resource "ibm_scc_si_note" "kpi" {
#   account_id                       = var.account_id
#   provider_id                      = var.provider_id
#   note {
#     short_description              = "short_description"
#     long_description               = "long_description"
#     kind                           = "KPI"
#     related_url {
#       label                        = "rel_label"
#       url                          = "rel_url"
#     }
#     expiration_time                = "2006-01-02 15:04:11"
#     create_time                    = "2006-01-02 15:04:00"
#     update_time                    = "2006-01-02 15:04:11"
#     id                             = "test-kpi"
#     shared                         = "true"
#     reported_by {
#       id                           = "rep_id"
#       title                        = "rep_title"
#       url                          = "rep_url"
#     }
#     kpi {
#       aggregation_type             = "SUM"
#     }
#   }
# }

# # #################################

# # Create Section
# resource "ibm_scc_si_note" "section" {
#   account_id                       = var.account_id
#   provider_id                      = var.provider_id
#   note {
#     short_description              = "short_description"
#     long_description               = "long_description"
#     kind                           = "SECTION"
#     related_url {
#       label                        = "rel_label"
#       url                          = "rel_url"
#     }
#     expiration_time                = "2006-01-02 15:04:11"
#     create_time                    = "2006-01-02 15:04:00"
#     update_time                    = "2006-01-02 15:04:11"
#     id                             = "test-section"
#     shared                         = "true"
#     reported_by {
#       id                           = "rep_id"
#       title                        = "rep_title"
#       url                          = "rep_url"
#     }
#     section {
#       image                        = "image"
#       title                        = "title"
#     }
#   }
# }

# # #################################