package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-azuread/internal/provider"
)

func Provider() *schema.Provider {
	return provider.AzureADProvider()
}
