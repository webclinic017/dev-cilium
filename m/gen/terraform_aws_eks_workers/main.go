// terraform_aws_eks_workers
package terraform_aws_eks_workers

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"terraform_aws_eks_workers.TerraformAwsEksWorkers",
		reflect.TypeOf((*TerraformAwsEksWorkers)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "additionalSecurityGroupIds", GoGetter: "AdditionalSecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "additionalTagMap", GoGetter: "AdditionalTagMap"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addProvider", GoMethod: "AddProvider"},
			_jsii_.MemberProperty{JsiiProperty: "afterClusterJoiningUserdata", GoGetter: "AfterClusterJoiningUserdata"},
			_jsii_.MemberProperty{JsiiProperty: "allowedCidrBlocks", GoGetter: "AllowedCidrBlocks"},
			_jsii_.MemberProperty{JsiiProperty: "allowedSecurityGroups", GoGetter: "AllowedSecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "associatePublicIpAddress", GoGetter: "AssociatePublicIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "attributes", GoGetter: "Attributes"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingGroupArnOutput", GoGetter: "AutoscalingGroupArnOutput"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingGroupDefaultCooldownOutput", GoGetter: "AutoscalingGroupDefaultCooldownOutput"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingGroupDesiredCapacityOutput", GoGetter: "AutoscalingGroupDesiredCapacityOutput"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingGroupHealthCheckGracePeriodOutput", GoGetter: "AutoscalingGroupHealthCheckGracePeriodOutput"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingGroupHealthCheckTypeOutput", GoGetter: "AutoscalingGroupHealthCheckTypeOutput"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingGroupIdOutput", GoGetter: "AutoscalingGroupIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingGroupMaxSizeOutput", GoGetter: "AutoscalingGroupMaxSizeOutput"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingGroupMinSizeOutput", GoGetter: "AutoscalingGroupMinSizeOutput"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingGroupNameOutput", GoGetter: "AutoscalingGroupNameOutput"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingGroupTags", GoGetter: "AutoscalingGroupTags"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingGroupTagsOutput", GoGetter: "AutoscalingGroupTagsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingPoliciesEnabled", GoGetter: "AutoscalingPoliciesEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "awsIamInstanceProfileName", GoGetter: "AwsIamInstanceProfileName"},
			_jsii_.MemberProperty{JsiiProperty: "beforeClusterJoiningUserdata", GoGetter: "BeforeClusterJoiningUserdata"},
			_jsii_.MemberProperty{JsiiProperty: "blockDeviceMappings", GoGetter: "BlockDeviceMappings"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapExtraArgs", GoGetter: "BootstrapExtraArgs"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clusterCertificateAuthorityData", GoGetter: "ClusterCertificateAuthorityData"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpoint", GoGetter: "ClusterEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "clusterSecurityGroupId", GoGetter: "ClusterSecurityGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "clusterSecurityGroupIngressEnabled", GoGetter: "ClusterSecurityGroupIngressEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "context", GoGetter: "Context"},
			_jsii_.MemberProperty{JsiiProperty: "cpuUtilizationHighEvaluationPeriods", GoGetter: "CpuUtilizationHighEvaluationPeriods"},
			_jsii_.MemberProperty{JsiiProperty: "cpuUtilizationHighPeriodSeconds", GoGetter: "CpuUtilizationHighPeriodSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "cpuUtilizationHighStatistic", GoGetter: "CpuUtilizationHighStatistic"},
			_jsii_.MemberProperty{JsiiProperty: "cpuUtilizationHighThresholdPercent", GoGetter: "CpuUtilizationHighThresholdPercent"},
			_jsii_.MemberProperty{JsiiProperty: "cpuUtilizationLowEvaluationPeriods", GoGetter: "CpuUtilizationLowEvaluationPeriods"},
			_jsii_.MemberProperty{JsiiProperty: "cpuUtilizationLowPeriodSeconds", GoGetter: "CpuUtilizationLowPeriodSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "cpuUtilizationLowStatistic", GoGetter: "CpuUtilizationLowStatistic"},
			_jsii_.MemberProperty{JsiiProperty: "cpuUtilizationLowThresholdPercent", GoGetter: "CpuUtilizationLowThresholdPercent"},
			_jsii_.MemberProperty{JsiiProperty: "creditSpecification", GoGetter: "CreditSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "defaultCooldown", GoGetter: "DefaultCooldown"},
			_jsii_.MemberProperty{JsiiProperty: "delimiter", GoGetter: "Delimiter"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "disableApiTermination", GoGetter: "DisableApiTermination"},
			_jsii_.MemberProperty{JsiiProperty: "ebsOptimized", GoGetter: "EbsOptimized"},
			_jsii_.MemberProperty{JsiiProperty: "eksWorkerAmiNameFilter", GoGetter: "EksWorkerAmiNameFilter"},
			_jsii_.MemberProperty{JsiiProperty: "eksWorkerAmiNameRegex", GoGetter: "EksWorkerAmiNameRegex"},
			_jsii_.MemberProperty{JsiiProperty: "elasticGpuSpecifications", GoGetter: "ElasticGpuSpecifications"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledMetrics", GoGetter: "EnabledMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "enableMonitoring", GoGetter: "EnableMonitoring"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberProperty{JsiiProperty: "forceDelete", GoGetter: "ForceDelete"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckGracePeriod", GoGetter: "HealthCheckGracePeriod"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckType", GoGetter: "HealthCheckType"},
			_jsii_.MemberProperty{JsiiProperty: "idLengthLimit", GoGetter: "IdLengthLimit"},
			_jsii_.MemberProperty{JsiiProperty: "imageId", GoGetter: "ImageId"},
			_jsii_.MemberProperty{JsiiProperty: "instanceInitiatedShutdownBehavior", GoGetter: "InstanceInitiatedShutdownBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "instanceMarketOptions", GoGetter: "InstanceMarketOptions"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForOutput", GoMethod: "InterpolationForOutput"},
			_jsii_.MemberProperty{JsiiProperty: "keyName", GoGetter: "KeyName"},
			_jsii_.MemberProperty{JsiiProperty: "kubeletExtraArgs", GoGetter: "KubeletExtraArgs"},
			_jsii_.MemberProperty{JsiiProperty: "labelKeyCase", GoGetter: "LabelKeyCase"},
			_jsii_.MemberProperty{JsiiProperty: "labelOrder", GoGetter: "LabelOrder"},
			_jsii_.MemberProperty{JsiiProperty: "labelValueCase", GoGetter: "LabelValueCase"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplateArnOutput", GoGetter: "LaunchTemplateArnOutput"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplateIdOutput", GoGetter: "LaunchTemplateIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancers", GoGetter: "LoadBalancers"},
			_jsii_.MemberProperty{JsiiProperty: "maxSize", GoGetter: "MaxSize"},
			_jsii_.MemberProperty{JsiiProperty: "metadataHttpEndpointEnabled", GoGetter: "MetadataHttpEndpointEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "metadataHttpPutResponseHopLimit", GoGetter: "MetadataHttpPutResponseHopLimit"},
			_jsii_.MemberProperty{JsiiProperty: "metadataHttpTokensRequired", GoGetter: "MetadataHttpTokensRequired"},
			_jsii_.MemberProperty{JsiiProperty: "metricsGranularity", GoGetter: "MetricsGranularity"},
			_jsii_.MemberProperty{JsiiProperty: "minElbCapacity", GoGetter: "MinElbCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "minSize", GoGetter: "MinSize"},
			_jsii_.MemberProperty{JsiiProperty: "mixedInstancesPolicy", GoGetter: "MixedInstancesPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "placement", GoGetter: "Placement"},
			_jsii_.MemberProperty{JsiiProperty: "placementGroup", GoGetter: "PlacementGroup"},
			_jsii_.MemberProperty{JsiiProperty: "protectFromScaleIn", GoGetter: "ProtectFromScaleIn"},
			_jsii_.MemberProperty{JsiiProperty: "providers", GoGetter: "Providers"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "regexReplaceChars", GoGetter: "RegexReplaceChars"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "scaleDownAdjustmentType", GoGetter: "ScaleDownAdjustmentType"},
			_jsii_.MemberProperty{JsiiProperty: "scaleDownCooldownSeconds", GoGetter: "ScaleDownCooldownSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "scaleDownPolicyType", GoGetter: "ScaleDownPolicyType"},
			_jsii_.MemberProperty{JsiiProperty: "scaleDownScalingAdjustment", GoGetter: "ScaleDownScalingAdjustment"},
			_jsii_.MemberProperty{JsiiProperty: "scaleUpAdjustmentType", GoGetter: "ScaleUpAdjustmentType"},
			_jsii_.MemberProperty{JsiiProperty: "scaleUpCooldownSeconds", GoGetter: "ScaleUpCooldownSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "scaleUpPolicyType", GoGetter: "ScaleUpPolicyType"},
			_jsii_.MemberProperty{JsiiProperty: "scaleUpScalingAdjustment", GoGetter: "ScaleUpScalingAdjustment"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupArnOutput", GoGetter: "SecurityGroupArnOutput"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIdOutput", GoGetter: "SecurityGroupIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupNameOutput", GoGetter: "SecurityGroupNameOutput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceLinkedRoleArn", GoGetter: "ServiceLinkedRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "skipAssetCreationFromLocalModules", GoGetter: "SkipAssetCreationFromLocalModules"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "stage", GoGetter: "Stage"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIds", GoGetter: "SubnetIds"},
			_jsii_.MemberProperty{JsiiProperty: "suspendedProcesses", GoGetter: "SuspendedProcesses"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupArns", GoGetter: "TargetGroupArns"},
			_jsii_.MemberProperty{JsiiProperty: "terminationPolicies", GoGetter: "TerminationPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "useCustomImageId", GoGetter: "UseCustomImageId"},
			_jsii_.MemberProperty{JsiiProperty: "useExistingAwsIamInstanceProfile", GoGetter: "UseExistingAwsIamInstanceProfile"},
			_jsii_.MemberProperty{JsiiProperty: "useExistingSecurityGroup", GoGetter: "UseExistingSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
			_jsii_.MemberProperty{JsiiProperty: "waitForCapacityTimeout", GoGetter: "WaitForCapacityTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "waitForElbCapacity", GoGetter: "WaitForElbCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "workersRoleArnOutput", GoGetter: "WorkersRoleArnOutput"},
			_jsii_.MemberProperty{JsiiProperty: "workersRoleNameOutput", GoGetter: "WorkersRoleNameOutput"},
			_jsii_.MemberProperty{JsiiProperty: "workersRolePolicyArns", GoGetter: "WorkersRolePolicyArns"},
			_jsii_.MemberProperty{JsiiProperty: "workersRolePolicyArnsCount", GoGetter: "WorkersRolePolicyArnsCount"},
			_jsii_.MemberProperty{JsiiProperty: "workersSecurityGroupId", GoGetter: "WorkersSecurityGroupId"},
		},
		func() interface{} {
			j := jsiiProxy_TerraformAwsEksWorkers{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformModule)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"terraform_aws_eks_workers.TerraformAwsEksWorkersConfig",
		reflect.TypeOf((*TerraformAwsEksWorkersConfig)(nil)).Elem(),
	)
}