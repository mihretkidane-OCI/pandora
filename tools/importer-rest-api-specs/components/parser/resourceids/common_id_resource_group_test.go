// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourceids

import (
	"testing"

	"github.com/hashicorp/pandora/tools/data-api-sdk/v1/models"
	importerModels "github.com/hashicorp/pandora/tools/importer-rest-api-specs/models"
)

func TestCommonResourceID_ResourceGroup(t *testing.T) {
	valid := importerModels.ParsedResourceId{
		Constants: map[string]models.SDKConstant{},
		Segments: []models.ResourceIDSegment{
			models.NewStaticValueResourceIDSegment("subscriptions", "subscriptions"),
			models.NewSubscriptionIDResourceIDSegment("subscriptionId"),
			models.NewStaticValueResourceIDSegment("resourceGroups", "resourceGroups"),
			models.NewResourceGroupNameResourceIDSegment("resourceGroupName"),
		},
	}
	invalid := importerModels.ParsedResourceId{
		Constants: map[string]models.SDKConstant{},
		Segments: []models.ResourceIDSegment{
			models.NewStaticValueResourceIDSegment("subscriptions", "subscriptions"),
			models.NewSubscriptionIDResourceIDSegment("subscriptionId"),
			models.NewStaticValueResourceIDSegment("resourceGroups", "resourceGroups"),
			models.NewResourceGroupNameResourceIDSegment("resourceGroupName"),
			models.NewStaticValueResourceIDSegment("someResource", "someResource"),
			models.NewUserSpecifiedResourceIDSegment("resourceName", "resourceName"),
		},
	}
	input := []importerModels.ParsedResourceId{
		valid,
		invalid,
	}
	output := switchOutCommonResourceIDsAsNeeded(input)
	for _, actual := range output {
		if normalizedResourceId(actual.Segments) == normalizedResourceId(valid.Segments) {
			if actual.CommonAlias == nil {
				t.Fatalf("Expected `valid` to have the CommonAlias `ResourceGroup` but got nil")
			}
			if *actual.CommonAlias != "ResourceGroup" {
				t.Fatalf("Expected `valid` to have the CommonAlias `ResourceGroup` but got %q", *actual.CommonAlias)
			}

			continue
		}

		if normalizedResourceId(actual.Segments) == normalizedResourceId(invalid.Segments) {
			if actual.CommonAlias != nil {
				t.Fatalf("Expected `invalid` to have no CommonAlias but got %q", *actual.CommonAlias)
			}
			continue
		}

		t.Fatalf("unexpected Resource ID %q", normalizedResourceId(actual.Segments))
	}
}

func TestCommonResourceID_ResourceGroupIncorrectSegment(t *testing.T) {
	input := []importerModels.ParsedResourceId{
		{
			Constants: map[string]models.SDKConstant{},
			Segments: []models.ResourceIDSegment{
				models.NewStaticValueResourceIDSegment("subscriptions", "subscriptions"),
				models.NewSubscriptionIDResourceIDSegment("subscriptionId"),
				models.NewStaticValueResourceIDSegment("resourceGroups", "resourceGroups"),
				models.NewResourceGroupNameResourceIDSegment("resourceGroupName"),
			},
		},
		{
			Constants: map[string]models.SDKConstant{},
			Segments: []models.ResourceIDSegment{
				models.NewStaticValueResourceIDSegment("subscriptions", "subscriptions"),
				models.NewSubscriptionIDResourceIDSegment("subscriptionId"),
				models.NewStaticValueResourceIDSegment("resourceGroups", "resourceGroups"),
				models.NewResourceGroupNameResourceIDSegment("sourceResourceGroupName"),
			},
		},
	}
	output := switchOutCommonResourceIDsAsNeeded(input)
	for i, actual := range output {
		t.Logf("testing %d", i)
		if actual.CommonAlias == nil || *actual.CommonAlias != "ResourceGroup" {
			t.Fatalf("expected item %d to be detected as a ResourceGroup but it wasn't", i)
		}
	}
}
