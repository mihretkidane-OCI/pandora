// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourceids

import (
	"fmt"
	"sort"

	"github.com/hashicorp/pandora/tools/data-api-sdk/v1/models"
)

// Parse takes a list of Swagger Resources and returns a ParseResult, containing
// a list of ResourceIDs found within the Swagger Resources.
func (p *Parser) Parse() (*ParseResult, error) {
	// 1. Go through and map the Operation IDs to the parsed Resource ID
	// (which includes the Resource ID and any UriSuffix as needed)
	p.logger.Trace("Parsing the segments for each operation..")
	operationIdsToSegments, err := p.parseSegmentsForEachOperation()
	if err != nil {
		return nil, fmt.Errorf("parsing the segments for each operation: %+v", err)
	}

	// 2. Process the list of parsed segments to obtain a unique list of Resource IDs
	p.logger.Trace("Determining the list of unique Resource IDs from the parsed input")
	uniqueResourceIds, distinctConstants := p.distinctResourceIds(*operationIdsToSegments)

	// 3. Then we need to find any Common Resource IDs and switch those references out
	p.logger.Trace("Generating Names for Resource IDs..")
	resourceIds := switchOutCommonResourceIDsAsNeeded(uniqueResourceIds)

	// 4. We then need to generate a unique Resource ID name for each of the Resource IDs
	p.logger.Trace("Generating Names for Resource IDs..")
	namesToResourceIds, err := p.generateNamesForResourceIds(resourceIds, nil)
	if err != nil {
		return nil, fmt.Errorf("generating Names for Resource IDs: %+v", err)
	}

	// 5. Then we need to work through the list of Resource IDs and Operation IDs to map the data across
	p.logger.Trace("Updating the Parsed Operations with the Processed ResourceIds..")
	operationIdsToResourceIds, err := p.updateParsedOperationsWithProcessedResourceIds(*operationIdsToSegments, *namesToResourceIds)
	if err != nil {
		return nil, fmt.Errorf("updating the parsed Operations with the Processed Resource ID information: %+v", err)
	}

	return &ParseResult{
		OperationIdsToParsedResourceIds: *operationIdsToResourceIds,
		NamesToResourceIDs:              *namesToResourceIds,
		Constants:                       distinctConstants,
	}, nil
}

func (p *Parser) updateParsedOperationsWithProcessedResourceIds(operationIdsToSegments map[string]processedResourceId, namesToResourceIds map[string]models.ResourceID) (*map[string]ParsedOperation, error) {
	output := make(map[string]ParsedOperation)

	for operationId, operation := range operationIdsToSegments {
		p.logger.Trace(fmt.Sprintf("Processing Operation ID %q", operationId))
		if operation.segments == nil {
			if operation.uriSuffix == nil {
				return nil, fmt.Errorf("the Operation ID %q had no Segments and no UriSuffix", operationId)
			}

			output[operationId] = ParsedOperation{
				UriSuffix: operation.uriSuffix,
			}
			continue
		}

		constantNames := make([]string, 0)
		for k := range operation.constants {
			constantNames = append(constantNames, k)
		}
		sort.Strings(constantNames)

		placeholder := models.ResourceID{
			ConstantNames: constantNames,
			Segments:      *operation.segments,
		}

		found := false
		for name, resourceId := range namesToResourceIds {
			// NOTE: we intentionally use an empty `id` here to avoid comparing on the Alias
			other := models.ResourceID{
				ConstantNames: resourceId.ConstantNames,
				Segments:      resourceId.Segments,
			}
			if ResourceIdsMatch(placeholder, other) {
				output[operationId] = ParsedOperation{
					ResourceId:     &resourceId,
					ResourceIdName: &name,
					UriSuffix:      operation.uriSuffix,
				}
				found = true
				break
			}
		}

		if !found {
			return nil, fmt.Errorf("couldn't find the Processed Resource Id for the Operation Id %q", operationId)
		}
	}

	return &output, nil
}
