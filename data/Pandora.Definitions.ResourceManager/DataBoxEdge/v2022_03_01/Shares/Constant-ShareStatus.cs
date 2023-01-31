using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.DataBoxEdge.v2022_03_01.Shares;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum ShareStatusConstant
{
    [Description("NeedsAttention")]
    NeedsAttention,

    [Description("OK")]
    OK,

    [Description("Offline")]
    Offline,

    [Description("Unknown")]
    Unknown,

    [Description("Updating")]
    Updating,
}
