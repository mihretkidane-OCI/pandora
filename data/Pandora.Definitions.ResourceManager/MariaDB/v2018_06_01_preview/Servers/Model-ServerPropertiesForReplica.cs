using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.MariaDB.v2018_06_01_preview.Servers;

[ValueForType("Replica")]
internal class ServerPropertiesForReplicaModel : ServerPropertiesForCreateModel
{
    [JsonPropertyName("sourceServerId")]
    [Required]
    public string SourceServerId { get; set; }
}
