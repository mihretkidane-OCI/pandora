package models

import (
	"fmt"
	"strings"

	"github.com/hashicorp/pandora/tools/sdk/resourcemanager"
)

type SchemaDefinition struct {
	Fields map[string]FieldDefinition
}

type FieldDefinition struct {
	Definition resourcemanager.ApiObjectDefinition
	Required   bool
	ForceNew   bool
	Optional   bool
	Computed   bool
	Validation ValidationDefinition

	// WriteOnly specifies if this field is Write-Only, that is, setable but not returned
	WriteOnly bool
}

func (d FieldDefinition) String() string {
	return strings.Join([]string{
		fmt.Sprintf("Definition %+v", d.Definition),
		fmt.Sprintf("Required %t", d.Required),
		fmt.Sprintf("ForceNew %t", d.ForceNew),
		fmt.Sprintf("Optional %t", d.Optional),
		fmt.Sprintf("Computed %t", d.Computed),
		fmt.Sprintf("Validation %+v", d.Validation),
		fmt.Sprintf("Write Only %t", d.WriteOnly),
	}, " / ")
}
