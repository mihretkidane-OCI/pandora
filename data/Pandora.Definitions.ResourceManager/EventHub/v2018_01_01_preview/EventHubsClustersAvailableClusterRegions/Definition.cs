using System.Collections.Generic;
using Pandora.Definitions.Interfaces;

namespace Pandora.Definitions.ResourceManager.EventHub.v2018_01_01_preview.EventHubsClustersAvailableClusterRegions
{
    internal class Definition : ApiDefinition
    {
        // Generated from Swagger revision "e07d7900e7e02c9bb6caba0ee15f0e280e97b8f5" 

        public string ApiVersion => "2018-01-01-preview";
        public string Name => "EventHubsClustersAvailableClusterRegions";
        public IEnumerable<Interfaces.ApiOperation> Operations => new List<Interfaces.ApiOperation>
        {
            new ClustersListAvailableClusterRegionOperation(),
        };
    }
}
