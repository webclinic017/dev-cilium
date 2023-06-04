// terraform_aws_dynamic_subnets
package terraform_aws_dynamic_subnets

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"terraform_aws_dynamic_subnets.TerraformAwsDynamicSubnets",
		reflect.TypeOf((*TerraformAwsDynamicSubnets)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "additionalTagMap", GoGetter: "AdditionalTagMap"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addProvider", GoMethod: "AddProvider"},
			_jsii_.MemberProperty{JsiiProperty: "attributes", GoGetter: "Attributes"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneAttributeStyle", GoGetter: "AvailabilityZoneAttributeStyle"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneIds", GoGetter: "AvailabilityZoneIds"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneIdsOutput", GoGetter: "AvailabilityZoneIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZonesOutput", GoGetter: "AvailabilityZonesOutput"},
			_jsii_.MemberProperty{JsiiProperty: "awsRouteCreateTimeout", GoGetter: "AwsRouteCreateTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "awsRouteDeleteTimeout", GoGetter: "AwsRouteDeleteTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "azPrivateRouteTableIdsMapOutput", GoGetter: "AzPrivateRouteTableIdsMapOutput"},
			_jsii_.MemberProperty{JsiiProperty: "azPrivateSubnetsMapOutput", GoGetter: "AzPrivateSubnetsMapOutput"},
			_jsii_.MemberProperty{JsiiProperty: "azPublicRouteTableIdsMapOutput", GoGetter: "AzPublicRouteTableIdsMapOutput"},
			_jsii_.MemberProperty{JsiiProperty: "azPublicSubnetsMapOutput", GoGetter: "AzPublicSubnetsMapOutput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "context", GoGetter: "Context"},
			_jsii_.MemberProperty{JsiiProperty: "delimiter", GoGetter: "Delimiter"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "descriptorFormats", GoGetter: "DescriptorFormats"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberProperty{JsiiProperty: "idLengthLimit", GoGetter: "IdLengthLimit"},
			_jsii_.MemberProperty{JsiiProperty: "igwId", GoGetter: "IgwId"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForOutput", GoMethod: "InterpolationForOutput"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4CidrBlock", GoGetter: "Ipv4CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4Cidrs", GoGetter: "Ipv4Cidrs"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4Enabled", GoGetter: "Ipv4Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4PrivateInstanceHostnamesEnabled", GoGetter: "Ipv4PrivateInstanceHostnamesEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4PrivateInstanceHostnameType", GoGetter: "Ipv4PrivateInstanceHostnameType"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4PublicInstanceHostnamesEnabled", GoGetter: "Ipv4PublicInstanceHostnamesEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4PublicInstanceHostnameType", GoGetter: "Ipv4PublicInstanceHostnameType"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6CidrBlock", GoGetter: "Ipv6CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6Cidrs", GoGetter: "Ipv6Cidrs"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6EgressOnlyIgwId", GoGetter: "Ipv6EgressOnlyIgwId"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6Enabled", GoGetter: "Ipv6Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6PrivateInstanceHostnamesEnabled", GoGetter: "Ipv6PrivateInstanceHostnamesEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6PublicInstanceHostnamesEnabled", GoGetter: "Ipv6PublicInstanceHostnamesEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "labelKeyCase", GoGetter: "LabelKeyCase"},
			_jsii_.MemberProperty{JsiiProperty: "labelOrder", GoGetter: "LabelOrder"},
			_jsii_.MemberProperty{JsiiProperty: "labelsAsTags", GoGetter: "LabelsAsTags"},
			_jsii_.MemberProperty{JsiiProperty: "labelValueCase", GoGetter: "LabelValueCase"},
			_jsii_.MemberProperty{JsiiProperty: "mapPublicIpOnLaunch", GoGetter: "MapPublicIpOnLaunch"},
			_jsii_.MemberProperty{JsiiProperty: "maxNats", GoGetter: "MaxNats"},
			_jsii_.MemberProperty{JsiiProperty: "maxSubnetCount", GoGetter: "MaxSubnetCount"},
			_jsii_.MemberProperty{JsiiProperty: "metadataHttpEndpointEnabled", GoGetter: "MetadataHttpEndpointEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "metadataHttpPutResponseHopLimit", GoGetter: "MetadataHttpPutResponseHopLimit"},
			_jsii_.MemberProperty{JsiiProperty: "metadataHttpTokensRequired", GoGetter: "MetadataHttpTokensRequired"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "namedPrivateRouteTableIdsMapOutput", GoGetter: "NamedPrivateRouteTableIdsMapOutput"},
			_jsii_.MemberProperty{JsiiProperty: "namedPrivateSubnetsMapOutput", GoGetter: "NamedPrivateSubnetsMapOutput"},
			_jsii_.MemberProperty{JsiiProperty: "namedPrivateSubnetsStatsMapOutput", GoGetter: "NamedPrivateSubnetsStatsMapOutput"},
			_jsii_.MemberProperty{JsiiProperty: "namedPublicRouteTableIdsMapOutput", GoGetter: "NamedPublicRouteTableIdsMapOutput"},
			_jsii_.MemberProperty{JsiiProperty: "namedPublicSubnetsMapOutput", GoGetter: "NamedPublicSubnetsMapOutput"},
			_jsii_.MemberProperty{JsiiProperty: "namedPublicSubnetsStatsMapOutput", GoGetter: "NamedPublicSubnetsStatsMapOutput"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "natEipAllocationIdsOutput", GoGetter: "NatEipAllocationIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "natElasticIps", GoGetter: "NatElasticIps"},
			_jsii_.MemberProperty{JsiiProperty: "natGatewayEnabled", GoGetter: "NatGatewayEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "natGatewayIdsOutput", GoGetter: "NatGatewayIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "natGatewayPublicIpsOutput", GoGetter: "NatGatewayPublicIpsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "natInstanceAmiId", GoGetter: "NatInstanceAmiId"},
			_jsii_.MemberProperty{JsiiProperty: "natInstanceAmiIdOutput", GoGetter: "NatInstanceAmiIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "natInstanceCpuCreditsOverride", GoGetter: "NatInstanceCpuCreditsOverride"},
			_jsii_.MemberProperty{JsiiProperty: "natInstanceEnabled", GoGetter: "NatInstanceEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "natInstanceIdsOutput", GoGetter: "NatInstanceIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "natInstanceRootBlockDeviceEncrypted", GoGetter: "NatInstanceRootBlockDeviceEncrypted"},
			_jsii_.MemberProperty{JsiiProperty: "natInstanceType", GoGetter: "NatInstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "natIpsOutput", GoGetter: "NatIpsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "openNetworkAclIpv4RuleNumber", GoGetter: "OpenNetworkAclIpv4RuleNumber"},
			_jsii_.MemberProperty{JsiiProperty: "openNetworkAclIpv6RuleNumber", GoGetter: "OpenNetworkAclIpv6RuleNumber"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "privateAssignIpv6AddressOnCreation", GoGetter: "PrivateAssignIpv6AddressOnCreation"},
			_jsii_.MemberProperty{JsiiProperty: "privateDns64Nat64Enabled", GoGetter: "PrivateDns64Nat64Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "privateLabel", GoGetter: "PrivateLabel"},
			_jsii_.MemberProperty{JsiiProperty: "privateNetworkAclIdOutput", GoGetter: "PrivateNetworkAclIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateOpenNetworkAclEnabled", GoGetter: "PrivateOpenNetworkAclEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "privateRouteTableEnabled", GoGetter: "PrivateRouteTableEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "privateRouteTableIdsOutput", GoGetter: "PrivateRouteTableIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnetCidrsOutput", GoGetter: "PrivateSubnetCidrsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnetIdsOutput", GoGetter: "PrivateSubnetIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnetIpv6CidrsOutput", GoGetter: "PrivateSubnetIpv6CidrsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnetsAdditionalTags", GoGetter: "PrivateSubnetsAdditionalTags"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnetsEnabled", GoGetter: "PrivateSubnetsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "providers", GoGetter: "Providers"},
			_jsii_.MemberProperty{JsiiProperty: "publicAssignIpv6AddressOnCreation", GoGetter: "PublicAssignIpv6AddressOnCreation"},
			_jsii_.MemberProperty{JsiiProperty: "publicDns64Nat64Enabled", GoGetter: "PublicDns64Nat64Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "publicLabel", GoGetter: "PublicLabel"},
			_jsii_.MemberProperty{JsiiProperty: "publicNetworkAclIdOutput", GoGetter: "PublicNetworkAclIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "publicOpenNetworkAclEnabled", GoGetter: "PublicOpenNetworkAclEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "publicRouteTableEnabled", GoGetter: "PublicRouteTableEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "publicRouteTableIds", GoGetter: "PublicRouteTableIds"},
			_jsii_.MemberProperty{JsiiProperty: "publicRouteTableIdsOutput", GoGetter: "PublicRouteTableIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "publicRouteTablePerSubnetEnabled", GoGetter: "PublicRouteTablePerSubnetEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnetCidrsOutput", GoGetter: "PublicSubnetCidrsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnetIdsOutput", GoGetter: "PublicSubnetIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnetIpv6CidrsOutput", GoGetter: "PublicSubnetIpv6CidrsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnetsAdditionalTags", GoGetter: "PublicSubnetsAdditionalTags"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnetsEnabled", GoGetter: "PublicSubnetsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "regexReplaceChars", GoGetter: "RegexReplaceChars"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rootBlockDeviceEncrypted", GoGetter: "RootBlockDeviceEncrypted"},
			_jsii_.MemberProperty{JsiiProperty: "routeCreateTimeout", GoGetter: "RouteCreateTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "routeDeleteTimeout", GoGetter: "RouteDeleteTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "skipAssetCreationFromLocalModules", GoGetter: "SkipAssetCreationFromLocalModules"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "stage", GoGetter: "Stage"},
			_jsii_.MemberProperty{JsiiProperty: "subnetCreateTimeout", GoGetter: "SubnetCreateTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "subnetDeleteTimeout", GoGetter: "SubnetDeleteTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "subnetsPerAzCount", GoGetter: "SubnetsPerAzCount"},
			_jsii_.MemberProperty{JsiiProperty: "subnetsPerAzNames", GoGetter: "SubnetsPerAzNames"},
			_jsii_.MemberProperty{JsiiProperty: "subnetTypeTagKey", GoGetter: "SubnetTypeTagKey"},
			_jsii_.MemberProperty{JsiiProperty: "subnetTypeTagValueFormat", GoGetter: "SubnetTypeTagValueFormat"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tenant", GoGetter: "Tenant"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
		},
		func() interface{} {
			j := jsiiProxy_TerraformAwsDynamicSubnets{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformModule)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"terraform_aws_dynamic_subnets.TerraformAwsDynamicSubnetsConfig",
		reflect.TypeOf((*TerraformAwsDynamicSubnetsConfig)(nil)).Elem(),
	)
}
