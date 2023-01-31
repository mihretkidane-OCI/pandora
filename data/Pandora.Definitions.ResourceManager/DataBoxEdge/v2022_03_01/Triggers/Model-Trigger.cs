using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.DataBoxEdge.v2022_03_01.Triggers;


internal abstract class TriggerModel
{
    [JsonPropertyName("id")]
    public string? Id { get; set; }

    [JsonPropertyName("kind")]
    [ProvidesTypeHint]
    [Required]
    public TriggerEventTypeConstant Kind { get; set; }

    [JsonPropertyName("name")]
    public string? Name { get; set; }

    [JsonPropertyName("systemData")]
    public CustomTypes.SystemData? SystemData { get; set; }

    [JsonPropertyName("type")]
    public string? Type { get; set; }
}
