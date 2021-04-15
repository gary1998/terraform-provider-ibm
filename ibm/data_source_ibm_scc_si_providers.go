package ibm

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIBMProviders() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMProvidersRead,
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"providers": {
				Type:        schema.TypeList,
				Description: "Collection of SA Notification Channels",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMProvidersRead(d *schema.ResourceData, meta interface{}) error {
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

	listProvidersOptions := sess.NewListProvidersOptions(accountID)
	providers, _, err := sess.ListProviders(listProvidersOptions)

	if err != nil {
		return fmt.Errorf("[ERROR] error occurred while listing providers: %v", err)
	}

	providersList := make([]map[string]interface{}, 0)
	for _, provider := range providers.Providers {
		providerObject := map[string]interface{}{}
		providerObject["id"] = provider.ID
		providerObject["name"] = provider.Name
		providersList = append(providersList, providerObject)
	}

	d.Set("providers", providersList)
	d.SetId(fmt.Sprintf("%v/providers", accountID))

	return nil
}
