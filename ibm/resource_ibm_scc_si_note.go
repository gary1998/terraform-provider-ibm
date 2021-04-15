package ibm

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ibm-cloud-security/security-advisor-sdk-go/v2/findingsapiv1"
)

func resourceIBMNote() *schema.Resource {
	return &schema.Resource{
		Create: resourceIBMNoteCreate,
		Read:   resourceIBMNoteRead,
		Update: resourceIBMNoteUpdate,
		Delete: resourceIBMNoteDelete,
		Exists: resourceIBMNoteExists,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"provider_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"note": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"short_description": {
							Type:     schema.TypeString,
							Required: true,
						},
						"long_description": {
							Type:     schema.TypeString,
							Required: true,
						},
						"kind": {
							Type:     schema.TypeString,
							Required: true,
						},
						"related_url": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"label": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"url": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"expiration_time": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: applyOnce,
						},
						"create_time": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: applyOnce,
						},
						"update_time": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: applyOnce,
						},
						"shared": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"reported_by": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"title": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"url": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"finding": {
							Type:     schema.TypeList,
							Optional: true,
							// ExactlyOneOf: []string{"finding", "kpi", "card", "section"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"severity": {
										Type:     schema.TypeString,
										Required: true,
									},
									"next_steps": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"url": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"title": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"kpi": {
							Type: schema.TypeList,
							// ExactlyOneOf: []string{"finding", "kpi", "card", "section"},
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aggregation_type": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"card": {
							Type: schema.TypeList,
							// ExactlyOneOf: []string{"finding", "kpi", "card", "section"},
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"section": {
										Type:     schema.TypeString,
										Required: true,
									},
									"title": {
										Type:     schema.TypeString,
										Required: true,
									},
									"subtitle": {
										Type:     schema.TypeString,
										Required: true,
									},
									"order": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"finding_note_names": {
										Type:     schema.TypeList,
										Required: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									"requires_configuration": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"badge_text": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"badge_image": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"elements": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"kind": {
													Type:     schema.TypeString,
													Required: true,
												},
												"text": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"value_type": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"kind": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"finding_note_names": {
																Type:     schema.TypeList,
																Optional: true,
																Elem:     &schema.Schema{Type: schema.TypeString},
															},
															"kpi_note_name": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"text": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"value_types": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"kind": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"finding_note_names": {
																Type:     schema.TypeList,
																Optional: true,
																Elem:     &schema.Schema{Type: schema.TypeString},
															},
															"kpi_note_name": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"text": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"default_time_range": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"default_interval": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"section": {
							Type: schema.TypeList,
							// ExactlyOneOf: []string{"finding", "kpi", "card", "section"},
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"image": {
										Type:     schema.TypeString,
										Required: true,
									},
									"title": {
										Type:     schema.TypeString,
										Required: true,
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

func resourceIBMNoteKindValidator() *ResourceValidator {
	validateSchema := make([]ValidateSchema, 1)
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "kind",
			ValidateFunctionIdentifier: ValidateAllowedStringValue,
			Type:                       TypeString,
			Required:                   true,
			AllowedValues:              "CARD, CARD_CONFIGURED, FINDING, KPI, SECTION",
		})

	resourceValidator := ResourceValidator{ResourceName: "ibm_scc_si_note", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMNoteCreate(d *schema.ResourceData, meta interface{}) error {
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

	noteDef := d.Get("note").([]interface{})[0].(map[string]interface{})
	createNoteOptions := findingsapiv1.CreateNoteOptions{}

	var noteCreateTime strfmt.DateTime
	var noteUpdateTime strfmt.DateTime
	var noteExpirationTime strfmt.DateTime
	noteShared := false

	if noteDef["create_time"] != nil {
		noteCreateTime, err = core.ParseDateTime(noteDef["create_time"].(string))
	}

	if noteDef["update_time"] != nil {
		noteUpdateTime, err = core.ParseDateTime(noteDef["update_time"].(string))
	}

	if noteDef["expiration_time"] != nil {
		noteExpirationTime, err = core.ParseDateTime(noteDef["expiration_time"].(string))
	}

	if err != nil {
		return fmt.Errorf("[ERROR] error occurred while validate datetime: %v", err)
	}

	if noteDef["shared"] != nil {
		shared, err := strconv.ParseBool(noteDef["shared"].(string))
		if err != nil {
			return fmt.Errorf("[ERROR] error occurred while validating boolean: %v", err)
		}
		noteShared = shared
	}

	if noteDef["related_url"] != nil && len(noteDef["related_url"].([]interface{})) > 0 {
		relatedURLObjectList := make([]findingsapiv1.ApiNoteRelatedURL, 0)
		for _, relatedURL := range noteDef["related_url"].([]interface{}) {
			relatedURLObject := findingsapiv1.ApiNoteRelatedURL{}
			label := relatedURL.(map[string]interface{})["label"].(string)
			url := relatedURL.(map[string]interface{})["url"].(string)
			relatedURLObject.Label = &label
			relatedURLObject.URL = &url
			relatedURLObjectList = append(relatedURLObjectList, relatedURLObject)
		}
		createNoteOptions.RelatedURL = relatedURLObjectList
	}

	if noteDef["reported_by"] != nil && len(noteDef["reported_by"].([]interface{})) > 0 {
		reportedByObject := findingsapiv1.Reporter{}
		reporterID := noteDef["reported_by"].([]interface{})[0].(map[string]interface{})["id"].(string)
		reportedTitle := noteDef["reported_by"].([]interface{})[0].(map[string]interface{})["title"].(string)
		reporterURL := noteDef["reported_by"].([]interface{})[0].(map[string]interface{})["url"].(string)
		reportedByObject.ID = &reporterID
		reportedByObject.Title = &reportedTitle
		reportedByObject.URL = &reporterURL
		createNoteOptions.ReportedBy = &reportedByObject
	}

	if noteDef["finding"] != nil && len(noteDef["finding"].([]interface{})) > 0 {
		findingObject := findingsapiv1.FindingType{}
		findingSeverity := noteDef["finding"].([]interface{})[0].(map[string]interface{})["severity"].(string)
		findingObject.Severity = &findingSeverity
		findingNextSteps := make([]findingsapiv1.RemediationStep, 0)
		for _, nextStep := range noteDef["finding"].([]interface{})[0].(map[string]interface{})["next_steps"].([]interface{}) {
			nextStepObject := findingsapiv1.RemediationStep{}
			remStepURL := nextStep.(map[string]interface{})["url"].(string)
			remStepTitle := nextStep.(map[string]interface{})["title"].(string)
			nextStepObject.URL = &remStepURL
			nextStepObject.Title = &remStepTitle
			findingNextSteps = append(findingNextSteps, nextStepObject)
		}
		findingObject.NextSteps = findingNextSteps
		createNoteOptions.Finding = &findingObject
	}

	if noteDef["kpi"] != nil && len(noteDef["kpi"].([]interface{})) > 0 {
		kpiObject := findingsapiv1.KpiType{}
		kpiAggregationType := noteDef["kpi"].([]interface{})[0].(map[string]interface{})["aggregation_type"].(string)
		kpiObject.AggregationType = &kpiAggregationType
		createNoteOptions.Kpi = &kpiObject
	}

	if noteDef["card"] != nil && len(noteDef["card"].([]interface{})) > 0 {
		cardObject := findingsapiv1.Card{}
		cardSection := noteDef["card"].([]interface{})[0].(map[string]interface{})["section"].(string)
		cardTitle := noteDef["card"].([]interface{})[0].(map[string]interface{})["title"].(string)
		cardSubtitle := noteDef["card"].([]interface{})[0].(map[string]interface{})["subtitle"].(string)
		cardOrder := int64(noteDef["card"].([]interface{})[0].(map[string]interface{})["order"].(int))
		cardRequiresConf, _ := strconv.ParseBool(noteDef["card"].([]interface{})[0].(map[string]interface{})["requires_configuration"].(string))
		cardBadgeText := noteDef["card"].([]interface{})[0].(map[string]interface{})["badge_text"].(string)
		cardBadgeImage := noteDef["card"].([]interface{})[0].(map[string]interface{})["badge_image"].(string)
		cardObject.Section = &cardSection
		cardObject.Title = &cardTitle
		cardObject.Subtitle = &cardSubtitle
		cardObject.Order = &cardOrder
		cardObject.RequiresConfiguration = &cardRequiresConf
		cardObject.BadgeText = &cardBadgeText
		cardObject.BadgeImage = &cardBadgeImage
		cardObjectFindingNoteNamesList := make([]string, 0)
		for _, findingNoteName := range noteDef["card"].([]interface{})[0].(map[string]interface{})["finding_note_names"].([]interface{}) {
			cardObjectFindingNoteNamesList = append(cardObjectFindingNoteNamesList, findingNoteName.(string))
		}
		cardObject.FindingNoteNames = cardObjectFindingNoteNamesList
		cardElementsList := make([]findingsapiv1.CardElement, 0)
		for _, cardElement := range noteDef["card"].([]interface{})[0].(map[string]interface{})["elements"].([]interface{}) {
			cardElementObject := findingsapiv1.CardElement{}
			cardElementKind := cardElement.(map[string]interface{})["kind"].(string)
			cardElementText := cardElement.(map[string]interface{})["text"].(string)
			cardElementDefTime := cardElement.(map[string]interface{})["default_time_range"].(string)
			cardElementDefInterval := cardElement.(map[string]interface{})["default_interval"].(string)
			cardElementObject.Kind = &cardElementKind
			cardElementObject.Text = &cardElementText
			cardElementObject.DefaultTimeRange = &cardElementDefTime
			cardElementObject.DefaultInterval = &cardElementDefInterval

			cardValueTypeObject := findingsapiv1.CardValueType{}
			cardValueTypeInterface := cardElement.(map[string]interface{})["value_type"].([]interface{})
			cardValueTypeKind := cardValueTypeInterface[0].(map[string]interface{})["kind"].(string)
			cardValueTypeKpiNoteName := cardValueTypeInterface[0].(map[string]interface{})["kpi_note_name"].(string)
			cardValueTypeText := cardValueTypeInterface[0].(map[string]interface{})["text"].(string)
			cardValueTypeObject.Kind = &cardValueTypeKind
			cardValueTypeObject.KpiNoteName = &cardValueTypeKpiNoteName
			cardValueTypeObject.Text = &cardValueTypeText
			cardValueTypeFindingNoteNames := make([]string, 0)
			for _, cardValueTypeFindingNoteName := range cardValueTypeInterface[0].(map[string]interface{})["finding_note_names"].([]interface{}) {
				cardValueTypeFindingNoteNames = append(cardValueTypeFindingNoteNames, cardValueTypeFindingNoteName.(string))
			}
			cardValueTypeObject.FindingNoteNames = cardValueTypeFindingNoteNames
			cardElementObject.ValueType = &cardValueTypeObject

			cardValueTypesObjectList := make([]findingsapiv1.CardValueType, 0)
			for _, cardValueType := range cardElement.(map[string]interface{})["value_types"].([]interface{}) {
				cardValueTypeObject := findingsapiv1.CardValueType{}
				cardValueTypeKind := cardValueType.(map[string]interface{})["kind"].(string)
				cardValueTypeKpiNoteName := cardValueType.(map[string]interface{})["kpi_note_name"].(string)
				cardValueTypeText := cardValueType.(map[string]interface{})["text"].(string)
				cardValueTypeObject.Kind = &cardValueTypeKind
				cardValueTypeObject.KpiNoteName = &cardValueTypeKpiNoteName
				cardValueTypeObject.Text = &cardValueTypeText
				cardValueTypeFindingNoteNames := make([]string, 0)
				for _, cardValueTypeFindingNoteName := range cardValueType.(map[string]interface{})["finding_note_names"].([]interface{}) {
					cardValueTypeFindingNoteNames = append(cardValueTypeFindingNoteNames, cardValueTypeFindingNoteName.(string))
				}
				cardValueTypeObject.FindingNoteNames = cardValueTypeFindingNoteNames
				cardValueTypesObjectList = append(cardValueTypesObjectList, cardValueTypeObject)
			}
			cardElementsList = append(cardElementsList, cardElementObject)
			cardElementObject.ValueTypes = cardValueTypesObjectList
		}
		cardObject.Elements = cardElementsList
		createNoteOptions.Card = &cardObject
	}

	if noteDef["section"] != nil && len(noteDef["section"].([]interface{})) > 0 {
		sectionObject := findingsapiv1.Section{}
		sectionTitle := noteDef["section"].([]interface{})[0].(map[string]interface{})["title"].(string)
		sectionImage := noteDef["section"].([]interface{})[0].(map[string]interface{})["image"].(string)
		sectionObject.Title = &sectionTitle
		sectionObject.Image = &sectionImage
		createNoteOptions.Section = &sectionObject
	}

	createNoteOptions.AccountID = &accountID
	createNoteOptions.ProviderID = &providerID
	noteID := noteDef["id"].(string)
	noteShortDescription := noteDef["short_description"].(string)
	noteLongDescription := noteDef["long_description"].(string)
	noteKind := noteDef["kind"].(string)
	createNoteOptions.ID = &noteID
	createNoteOptions.ShortDescription = &noteShortDescription
	createNoteOptions.LongDescription = &noteLongDescription
	createNoteOptions.Kind = &noteKind
	createNoteOptions.ExpirationTime = &noteExpirationTime
	createNoteOptions.CreateTime = &noteCreateTime
	createNoteOptions.UpdateTime = &noteUpdateTime
	createNoteOptions.Shared = &noteShared

	note, _, err := sess.CreateNote(&createNoteOptions)

	if err != nil {
		return fmt.Errorf("[ERROR] error occurred while creating note: %v", err)
	}
	if note.ID != nil {
		d.SetId(fmt.Sprintf("%v/providers/%v/notes/%v", accountID, providerID, noteDef["id"].(string)))
		return resourceIBMNoteRead(d, meta)
	}

	return nil

}

func resourceIBMNoteRead(d *schema.ResourceData, meta interface{}) error {
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
	idParts := strings.Split(d.Id(), "/")
	noteID := idParts[len(idParts)-1]

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
	noteShared := false
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
		noteShared = *apiResp.Shared
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

func resourceIBMNoteExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).FindingsV1API()
	if err != nil {
		return false, err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return false, err
	}
	accountID := d.Get("account_id").(string)
	log.Println(fmt.Sprintf("[DEBUG] using specified AccountID %s", accountID))
	if accountID == "" {
		accountID = userDetails.userAccount
		log.Println(fmt.Sprintf("[DEBUG] AccountID not spedified, using %s", accountID))
	}
	providerID := d.Get("provider_id").(string)
	idParts := strings.Split(d.Id(), "/")
	noteID := idParts[len(idParts)-1]

	getNoteOptions := sess.NewGetNoteOptions(accountID, providerID, noteID)
	_, resp, err := sess.GetNote(getNoteOptions)

	if err != nil {
		return false, fmt.Errorf("error occurred while reading note: %v", err)
	}

	if resp.StatusCode == 404 {
		return false, nil
	}

	return true, nil
}

func resourceIBMNoteUpdate(d *schema.ResourceData, meta interface{}) error {
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

	noteDef := d.Get("note").([]interface{})[0].(map[string]interface{})
	updateNoteOptions := findingsapiv1.UpdateNoteOptions{}

	var noteCreateTime strfmt.DateTime
	var noteUpdateTime strfmt.DateTime
	var noteExpirationTime strfmt.DateTime
	noteShared := false

	if noteDef["create_time"] != nil {
		noteCreateTime, err = core.ParseDateTime(noteDef["create_time"].(string))
	}

	if noteDef["update_time"] != nil {
		noteUpdateTime, err = core.ParseDateTime(noteDef["update_time"].(string))
	}

	if noteDef["expiration_time"] != nil {
		noteExpirationTime, err = core.ParseDateTime(noteDef["expiration_time"].(string))
	}

	if err != nil {
		return fmt.Errorf("[ERROR] error occurred while validate datetime: %v", err)
	}

	if noteDef["shared"] != nil {
		shared, err := strconv.ParseBool(noteDef["shared"].(string))
		if err != nil {
			return fmt.Errorf("[ERROR] error occurred while validating boolean: %v", err)
		}
		noteShared = shared
	}

	if noteDef["related_url"] != nil && len(noteDef["related_url"].([]interface{})) > 0 {
		relatedURLObjectList := make([]findingsapiv1.ApiNoteRelatedURL, 0)
		for _, relatedURL := range noteDef["related_url"].([]interface{}) {
			relatedURLObject := findingsapiv1.ApiNoteRelatedURL{}
			label := relatedURL.(map[string]interface{})["label"].(string)
			url := relatedURL.(map[string]interface{})["url"].(string)
			relatedURLObject.Label = &label
			relatedURLObject.URL = &url
			relatedURLObjectList = append(relatedURLObjectList, relatedURLObject)
		}
		updateNoteOptions.RelatedURL = relatedURLObjectList
	}

	if noteDef["reported_by"] != nil && len(noteDef["reported_by"].([]interface{})) > 0 {
		reportedByObject := findingsapiv1.Reporter{}
		reporterID := noteDef["reported_by"].([]interface{})[0].(map[string]interface{})["id"].(string)
		reportedTitle := noteDef["reported_by"].([]interface{})[0].(map[string]interface{})["title"].(string)
		reporterURL := noteDef["reported_by"].([]interface{})[0].(map[string]interface{})["url"].(string)
		reportedByObject.ID = &reporterID
		reportedByObject.Title = &reportedTitle
		reportedByObject.URL = &reporterURL
		updateNoteOptions.ReportedBy = &reportedByObject
	}

	if noteDef["finding"] != nil && len(noteDef["finding"].([]interface{})) > 0 {
		findingObject := findingsapiv1.FindingType{}
		findingSeverity := noteDef["finding"].([]interface{})[0].(map[string]interface{})["severity"].(string)
		findingObject.Severity = &findingSeverity
		findingNextSteps := make([]findingsapiv1.RemediationStep, 0)
		for _, nextStep := range noteDef["finding"].([]interface{})[0].(map[string]interface{})["next_steps"].([]interface{}) {
			nextStepObject := findingsapiv1.RemediationStep{}
			remStepURL := nextStep.(map[string]interface{})["url"].(string)
			remStepTitle := nextStep.(map[string]interface{})["title"].(string)
			nextStepObject.URL = &remStepURL
			nextStepObject.Title = &remStepTitle
			findingNextSteps = append(findingNextSteps, nextStepObject)
		}
		findingObject.NextSteps = findingNextSteps
		updateNoteOptions.Finding = &findingObject
	}

	if noteDef["kpi"] != nil && len(noteDef["kpi"].([]interface{})) > 0 {
		kpiObject := findingsapiv1.KpiType{}
		kpiAggregationType := noteDef["kpi"].([]interface{})[0].(map[string]interface{})["aggregation_type"].(string)
		kpiObject.AggregationType = &kpiAggregationType
		updateNoteOptions.Kpi = &kpiObject
	}

	if noteDef["card"] != nil && len(noteDef["card"].([]interface{})) > 0 {
		cardObject := findingsapiv1.Card{}
		cardSection := noteDef["card"].([]interface{})[0].(map[string]interface{})["section"].(string)
		cardTitle := noteDef["card"].([]interface{})[0].(map[string]interface{})["title"].(string)
		cardSubtitle := noteDef["card"].([]interface{})[0].(map[string]interface{})["subtitle"].(string)
		cardOrder := int64(noteDef["card"].([]interface{})[0].(map[string]interface{})["order"].(int))
		cardRequiresConf, _ := strconv.ParseBool(noteDef["card"].([]interface{})[0].(map[string]interface{})["requires_configuration"].(string))
		cardBadgeText := noteDef["card"].([]interface{})[0].(map[string]interface{})["badge_text"].(string)
		cardBadgeImage := noteDef["card"].([]interface{})[0].(map[string]interface{})["badge_image"].(string)
		cardObject.Section = &cardSection
		cardObject.Title = &cardTitle
		cardObject.Subtitle = &cardSubtitle
		cardObject.Order = &cardOrder
		cardObject.RequiresConfiguration = &cardRequiresConf
		cardObject.BadgeText = &cardBadgeText
		cardObject.BadgeImage = &cardBadgeImage
		cardObjectFindingNoteNamesList := make([]string, 0)
		for _, findingNoteName := range noteDef["card"].([]interface{})[0].(map[string]interface{})["finding_note_names"].([]interface{}) {
			cardObjectFindingNoteNamesList = append(cardObjectFindingNoteNamesList, findingNoteName.(string))
		}
		cardObject.FindingNoteNames = cardObjectFindingNoteNamesList
		cardElementsList := make([]findingsapiv1.CardElement, 0)
		for _, cardElement := range noteDef["card"].([]interface{})[0].(map[string]interface{})["elements"].([]interface{}) {
			cardElementObject := findingsapiv1.CardElement{}
			cardElementKind := cardElement.(map[string]interface{})["kind"].(string)
			cardElementText := cardElement.(map[string]interface{})["text"].(string)
			cardElementDefTime := cardElement.(map[string]interface{})["default_time_range"].(string)
			cardElementDefInterval := cardElement.(map[string]interface{})["default_interval"].(string)
			cardElementObject.Kind = &cardElementKind
			cardElementObject.Text = &cardElementText
			cardElementObject.DefaultTimeRange = &cardElementDefTime
			cardElementObject.DefaultInterval = &cardElementDefInterval

			cardValueTypeObject := findingsapiv1.CardValueType{}
			cardValueTypeInterface := cardElement.(map[string]interface{})["value_type"].([]interface{})
			cardValueTypeKind := cardValueTypeInterface[0].(map[string]interface{})["kind"].(string)
			cardValueTypeKpiNoteName := cardValueTypeInterface[0].(map[string]interface{})["kpi_note_name"].(string)
			cardValueTypeText := cardValueTypeInterface[0].(map[string]interface{})["text"].(string)
			cardValueTypeObject.Kind = &cardValueTypeKind
			cardValueTypeObject.KpiNoteName = &cardValueTypeKpiNoteName
			cardValueTypeObject.Text = &cardValueTypeText
			cardValueTypeFindingNoteNames := make([]string, 0)
			for _, cardValueTypeFindingNoteName := range cardValueTypeInterface[0].(map[string]interface{})["finding_note_names"].([]interface{}) {
				cardValueTypeFindingNoteNames = append(cardValueTypeFindingNoteNames, cardValueTypeFindingNoteName.(string))
			}
			cardValueTypeObject.FindingNoteNames = cardValueTypeFindingNoteNames
			cardElementObject.ValueType = &cardValueTypeObject

			cardValueTypesObjectList := make([]findingsapiv1.CardValueType, 0)
			for _, cardValueType := range cardElement.(map[string]interface{})["value_types"].([]interface{}) {
				cardValueTypeObject := findingsapiv1.CardValueType{}
				cardValueTypeKind := cardValueType.(map[string]interface{})["kind"].(string)
				cardValueTypeKpiNoteName := cardValueType.(map[string]interface{})["kpi_note_name"].(string)
				cardValueTypeText := cardValueType.(map[string]interface{})["text"].(string)
				cardValueTypeObject.Kind = &cardValueTypeKind
				cardValueTypeObject.KpiNoteName = &cardValueTypeKpiNoteName
				cardValueTypeObject.Text = &cardValueTypeText
				cardValueTypeFindingNoteNames := make([]string, 0)
				for _, cardValueTypeFindingNoteName := range cardValueType.(map[string]interface{})["finding_note_names"].([]interface{}) {
					cardValueTypeFindingNoteNames = append(cardValueTypeFindingNoteNames, cardValueTypeFindingNoteName.(string))
				}
				cardValueTypeObject.FindingNoteNames = cardValueTypeFindingNoteNames
				cardValueTypesObjectList = append(cardValueTypesObjectList, cardValueTypeObject)
			}
			cardElementsList = append(cardElementsList, cardElementObject)
			cardElementObject.ValueTypes = cardValueTypesObjectList
		}
		cardObject.Elements = cardElementsList
		updateNoteOptions.Card = &cardObject
	}

	if noteDef["section"] != nil && len(noteDef["section"].([]interface{})) > 0 {
		sectionObject := findingsapiv1.Section{}
		sectionTitle := noteDef["section"].([]interface{})[0].(map[string]interface{})["title"].(string)
		sectionImage := noteDef["section"].([]interface{})[0].(map[string]interface{})["image"].(string)
		sectionObject.Title = &sectionTitle
		sectionObject.Image = &sectionImage
		updateNoteOptions.Section = &sectionObject
	}

	updateNoteOptions.AccountID = &accountID
	updateNoteOptions.ProviderID = &providerID
	noteID := noteDef["id"].(string)
	noteShortDescription := noteDef["short_description"].(string)
	noteLongDescription := noteDef["long_description"].(string)
	noteKind := noteDef["kind"].(string)
	updateNoteOptions.ID = &noteID
	updateNoteOptions.NoteID = &noteID
	updateNoteOptions.ShortDescription = &noteShortDescription
	updateNoteOptions.LongDescription = &noteLongDescription
	updateNoteOptions.Kind = &noteKind
	updateNoteOptions.ExpirationTime = &noteExpirationTime
	updateNoteOptions.CreateTime = &noteCreateTime
	updateNoteOptions.UpdateTime = &noteUpdateTime
	updateNoteOptions.Shared = &noteShared

	note, _, err := sess.UpdateNote(&updateNoteOptions)

	if err != nil {
		return fmt.Errorf("[ERROR] error occurred while updating note: %v", err)
	}
	if note.ID != nil {
		d.SetId(fmt.Sprintf("%v/providers/%v/notes/%v", accountID, providerID, noteDef["id"].(string)))
		return resourceIBMNoteRead(d, meta)
	}

	return nil
}

func resourceIBMNoteDelete(d *schema.ResourceData, meta interface{}) error {
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
	idParts := strings.Split(d.Id(), "/")
	noteID := idParts[len(idParts)-1]

	deleteNoteOptions := sess.NewDeleteNoteOptions(accountID, providerID, noteID)
	_, err = sess.DeleteNote(deleteNoteOptions)

	if err != nil {
		return fmt.Errorf("[ERROR] error occurred while deleting note: %v", err)
	}

	d.SetId("")

	return nil
}
