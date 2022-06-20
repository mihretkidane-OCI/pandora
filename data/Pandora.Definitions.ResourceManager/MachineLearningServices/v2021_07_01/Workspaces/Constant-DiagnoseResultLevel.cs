using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.MachineLearningServices.v2021_07_01.Workspaces;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum DiagnoseResultLevelConstant
{
    [Description("Error")]
    Error,

    [Description("Information")]
    Information,

    [Description("Warning")]
    Warning,
}
