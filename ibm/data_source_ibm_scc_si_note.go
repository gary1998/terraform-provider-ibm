package ibm

import (
	"fmt"
	"log"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIBMNote() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMNoteRead,
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"provider_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"note_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"note": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"short_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"long_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"kind": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"related_url": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"label": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"url": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"expiration_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"create_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"update_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shared": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"reported_by": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"url": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"finding": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"next_steps": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"url": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"title": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"kpi": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aggregation_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"card": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"section": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"subtitle": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"order": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"finding_note_names": {
										Type:     schema.TypeList,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									"requires_configuration": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"badge_text": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"badge_image": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"elements": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"kind": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"text": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"value_type": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"kind": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"finding_note_names": {
																Type:     schema.TypeList,
																Computed: true,
																Elem:     &schema.Schema{Type: schema.TypeString},
															},
															"kpi_note_name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"text": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"value_types": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"kind": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"finding_note_names": {
																Type:     schema.TypeList,
																Computed: true,
																Elem:     &schema.Schema{Type: schema.TypeString},
															},
															"kpi_note_name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"text": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"default_time_range": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"default_interval": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"section": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"image": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMNoteRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).FindingsV1API()
	if err != nil {
		return err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	accountID := d.Get("account_id").(string)
	log.Println(fmt.Sprintf("[DEBUG] using specified AccountID %s", accountID))
	if accountID == "" {
		accountID = userDetails.userAccount
		log.Println(fmt.Sprintf("[DEBUG] AccountID not spedified, using %s", accountID))
	}

	providerID := d.Get("provider_id").(string)
	noteID := d.Get("note_id").(string)

	getNoteOptions := sess.NewGetNoteOptions(accountID, providerID, noteID)
	apiResp, _, err := sess.GetNote(getNoteOptions)

	if err != nil {
		return fmt.Errorf("[ERROR] error occurred while getting note: %v", err)
	}

	if apiResp == nil {
		return fmt.Errorf("[ERROR] no such note found: %v", noteID)
	}

	var noteCreateTime strfmt.DateTime
	var noteUpdateTime strfmt.DateTime
	var noteExpirationTime strfmt.DateTime
	noteShared := "false"
	noteFinding := make([]map[string]interface{}, 0)
	noteKpi := make([]map[string]interface{}, 0)
	noteCard := make([]map[string]interface{}, 0)
	noteSection := make([]map[string]interface{}, 0)

	if apiResp.CreateTime != nil {
		noteCreateTime, err = core.ParseDateTime(apiResp.CreateTime.String())
	}

	if apiResp.UpdateTime != nil {
		noteUpdateTime, err = core.ParseDateTime(apiResp.UpdateTime.String())
	}

	if apiResp.ExpirationTime != nil {
		noteExpirationTime, err = core.ParseDateTime(apiResp.ExpirationTime.String())
	}

	if err != nil {
		return fmt.Errorf("[ERROR] error occurred while validate datetime: %v", err)
	}

	if apiResp.Shared != nil {
		noteShared = "true"
	}

	noteRelatedURLList := make([]map[string]interface{}, 0)
	for _, noteRelatedURL := range apiResp.RelatedURL {
		noteRelatedURLObject := map[string]interface{}{}
		noteRelatedURLObject["label"] = noteRelatedURL.Label
		noteRelatedURLObject["url"] = noteRelatedURL.URL
		noteRelatedURLList = append(noteRelatedURLList, noteRelatedURLObject)
	}

	noteReportedByObjectList := make([]map[string]interface{}, 1)
	noteReportedByObject := map[string]interface{}{}
	noteReportedByObject["id"] = apiResp.ReportedBy.ID
	noteReportedByObject["title"] = apiResp.ReportedBy.Title
	noteReportedByObject["url"] = apiResp.ReportedBy.URL
	noteReportedByObjectList[0] = noteReportedByObject

	if apiResp.Finding != nil {
		noteFindingTypeObjectList := make([]map[string]interface{}, 1)
		noteFindingTypeObject := map[string]interface{}{}
		noteFindingTypeObject["severity"] = apiResp.Finding.Severity
		noteFindingTypeNextStepsList := make([]map[string]interface{}, 0)
		for _, noteFindingNextStep := range apiResp.Finding.NextSteps {
			noteFindingTypeNextStepObject := map[string]interface{}{}
			noteFindingTypeNextStepObject["title"] = noteFindingNextStep.Title
			noteFindingTypeNextStepObject["url"] = noteFindingNextStep.URL
			noteFindingTypeNextStepsList = append(noteFindingTypeNextStepsList, noteFindingTypeNextStepObject)
		}
		noteFindingTypeObject["next_steps"] = noteFindingTypeNextStepsList
		noteFindingTypeObjectList[0] = noteFindingTypeObject
		noteFinding = noteFindingTypeObjectList
	}

	if apiResp.Kpi != nil {
		noteKpiTypeObjectList := make([]map[string]interface{}, 1)
		noteKpiTypeObject := map[string]interface{}{}
		noteKpiTypeObject["aggregation_type"] = apiResp.Kpi.AggregationType
		noteKpiTypeObjectList[0] = noteKpiTypeObject
		noteKpi = noteKpiTypeObjectList
	}

	if apiResp.Card != nil {
		noteCardObjectList := make([]map[string]interface{}, 1)
		noteCardObject := map[string]interface{}{}
		noteCardObject["section"] = apiResp.Card.Section
		noteCardObject["title"] = apiResp.Card.Title
		noteCardObject["subtitle"] = apiResp.Card.Subtitle
		noteCardObject["order"] = apiResp.Card.Order
		noteCardObject["finding_note_names"] = apiResp.Card.FindingNoteNames
		noteCardObject["requires_configuration"] = "false"
		if apiResp.Card.RequiresConfiguration != nil && *apiResp.Card.RequiresConfiguration {
			noteCardObject["requires_configuration"] = "true"
		}
		noteCardObject["badge_text"] = apiResp.Card.BadgeText
		noteCardObject["badge_image"] = apiResp.Card.BadgeImage
		noteCardElementsList := make([]map[string]interface{}, 0)
		for _, noteCardElement := range apiResp.Card.Elements {
			noteCardElementObject := map[string]interface{}{}
			noteCardElementObject["kind"] = noteCardElement.Kind
			noteCardElementObject["text"] = noteCardElement.Text
			noteCardElementObject["default_time_range"] = noteCardElement.DefaultTimeRange
			noteCardElementObject["default_interval"] = noteCardElement.DefaultInterval
			if noteCardElement.ValueType != nil {
				noteCardElementValueTypeObjectList := make([]map[string]interface{}, 1)
				noteCardElementValueTypeObject := map[string]interface{}{}
				noteCardElementValueTypeObject["kind"] = noteCardElement.ValueType.Kind
				noteCardElementValueTypeObject["finding_note_names"] = noteCardElement.ValueType.FindingNoteNames
				noteCardElementValueTypeObject["kpi_note_name"] = noteCardElement.ValueType.KpiNoteName
				noteCardElementValueTypeObject["text"] = noteCardElement.ValueType.Text
				noteCardElementValueTypeObjectList[0] = noteCardElementValueTypeObject
				noteCardElementObject["value_type"] = noteCardElementValueTypeObjectList
			}
			if noteCardElement.ValueTypes != nil {
				noteCardElementValueTypesList := make([]map[string]interface{}, 0)
				for _, noteCardElementValueTypes := range noteCardElement.ValueTypes {
					noteCardElementValueTypesObject := map[string]interface{}{}
					noteCardElementValueTypesObject["kind"] = noteCardElementValueTypes.Kind
					noteCardElementValueTypesObject["finding_note_names"] = noteCardElementValueTypes.FindingNoteNames
					noteCardElementValueTypesObject["kpi_note_name"] = noteCardElementValueTypes.KpiNoteName
					noteCardElementValueTypesObject["text"] = noteCardElementValueTypes.Text
				}
				noteCardElementObject["value_types"] = noteCardElementValueTypesList
				noteCardElementsList = append(noteCardElementsList, noteCardElementObject)
			}
		}
		noteCardObject["elements"] = noteCardElementsList
		noteCardObjectList[0] = noteCardObject
		noteCard = noteCardObjectList
	}

	if apiResp.Section != nil {
		noteSectionObjectList := make([]map[string]interface{}, 1)
		noteSectionObject := map[string]interface{}{}
		noteSectionObject["image"] = apiResp.Section.Image
		noteSectionObject["title"] = apiResp.Section.Title
		noteSectionObjectList[0] = noteSectionObject
		noteSection = noteSectionObjectList
	}

	resp := make([]map[string]interface{}, 1)
	note := make(map[string]interface{})

	note["short_description"] = apiResp.ShortDescription
	note["long_description"] = apiResp.LongDescription
	note["kind"] = apiResp.Kind
	note["expiration_time"] = noteExpirationTime
	note["shared"] = noteShared
	note["related_url"] = noteRelatedURLList
	note["reported_by"] = noteReportedByObjectList
	note["finding"] = noteFinding
	note["kpi"] = noteKpi
	note["card"] = noteCard
	note["section"] = noteSection
	note["create_time"] = noteCreateTime
	note["update_time"] = noteUpdateTime
	note["id"] = apiResp.ID
	resp[0] = note

	d.Set("note", resp)
	d.SetId(fmt.Sprintf("%v/providers/%v/notes/%v", accountID, providerID, noteID))
	return nil
}
