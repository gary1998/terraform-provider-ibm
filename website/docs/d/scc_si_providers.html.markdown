---
subcategory: "Security and Compliance Center"
layout: "ibm"
page_title: "IBM : scc_si_provider"
sidebar_current: "docs-ibm-datasources-scc-si-provider"
description: |-
  Manages IBM Cloud Security Advisor Providers.
---

# ibm_scc_si_providers

Import the details of all existing IBM Cloud Security Advisor Providers as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_si_providers" "providers" {
  account_id = var.account_id
}
```

## Argument Reference

The following arguments are supported:

- `account_id` - (Optional, string) The GUID of account where providers reside.

## Attribute Reference

The following attributes are exported:

- `providers` - Object of Security Advisor Findings Note.
  - `name` - Full name of the provider.
  - `id` - Unique ID of the provider.
