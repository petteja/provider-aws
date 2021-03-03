/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package dbcluster

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/rds"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane/provider-aws/apis/rds/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeDBClustersInput returns input for read
// operation.
func GenerateDescribeDBClustersInput(cr *svcapitypes.DBCluster) *svcsdk.DescribeDBClustersInput {
	res := &svcsdk.DescribeDBClustersInput{}

	return res
}

// GenerateDBCluster returns the current state in the form of *svcapitypes.DBCluster.
func GenerateDBCluster(resp *svcsdk.DescribeDBClustersOutput) *svcapitypes.DBCluster {
	cr := &svcapitypes.DBCluster{}

	found := false
	for _, elem := range resp.DBClusters {
		if elem.ActivityStreamKinesisStreamName != nil {
			cr.Status.AtProvider.ActivityStreamKinesisStreamName = elem.ActivityStreamKinesisStreamName
		}
		if elem.ActivityStreamKmsKeyId != nil {
			cr.Status.AtProvider.ActivityStreamKMSKeyID = elem.ActivityStreamKmsKeyId
		}
		if elem.ActivityStreamMode != nil {
			cr.Status.AtProvider.ActivityStreamMode = elem.ActivityStreamMode
		}
		if elem.ActivityStreamStatus != nil {
			cr.Status.AtProvider.ActivityStreamStatus = elem.ActivityStreamStatus
		}
		if elem.AllocatedStorage != nil {
			cr.Status.AtProvider.AllocatedStorage = elem.AllocatedStorage
		}
		if elem.AssociatedRoles != nil {
			f5 := []*svcapitypes.DBClusterRole{}
			for _, f5iter := range elem.AssociatedRoles {
				f5elem := &svcapitypes.DBClusterRole{}
				if f5iter.FeatureName != nil {
					f5elem.FeatureName = f5iter.FeatureName
				}
				if f5iter.RoleArn != nil {
					f5elem.RoleARN = f5iter.RoleArn
				}
				if f5iter.Status != nil {
					f5elem.Status = f5iter.Status
				}
				f5 = append(f5, f5elem)
			}
			cr.Status.AtProvider.AssociatedRoles = f5
		}
		if elem.AvailabilityZones != nil {
			f6 := []*string{}
			for _, f6iter := range elem.AvailabilityZones {
				var f6elem string
				f6elem = *f6iter
				f6 = append(f6, &f6elem)
			}
			cr.Spec.ForProvider.AvailabilityZones = f6
		}
		if elem.BacktrackConsumedChangeRecords != nil {
			cr.Status.AtProvider.BacktrackConsumedChangeRecords = elem.BacktrackConsumedChangeRecords
		}
		if elem.BacktrackWindow != nil {
			cr.Spec.ForProvider.BacktrackWindow = elem.BacktrackWindow
		}
		if elem.BackupRetentionPeriod != nil {
			cr.Spec.ForProvider.BackupRetentionPeriod = elem.BackupRetentionPeriod
		}
		if elem.Capacity != nil {
			cr.Status.AtProvider.Capacity = elem.Capacity
		}
		if elem.CharacterSetName != nil {
			cr.Spec.ForProvider.CharacterSetName = elem.CharacterSetName
		}
		if elem.CloneGroupId != nil {
			cr.Status.AtProvider.CloneGroupID = elem.CloneGroupId
		}
		if elem.ClusterCreateTime != nil {
			cr.Status.AtProvider.ClusterCreateTime = &metav1.Time{*elem.ClusterCreateTime}
		}
		if elem.CopyTagsToSnapshot != nil {
			cr.Spec.ForProvider.CopyTagsToSnapshot = elem.CopyTagsToSnapshot
		}
		if elem.CrossAccountClone != nil {
			cr.Status.AtProvider.CrossAccountClone = elem.CrossAccountClone
		}
		if elem.CustomEndpoints != nil {
			f16 := []*string{}
			for _, f16iter := range elem.CustomEndpoints {
				var f16elem string
				f16elem = *f16iter
				f16 = append(f16, &f16elem)
			}
			cr.Status.AtProvider.CustomEndpoints = f16
		}
		if elem.DBClusterArn != nil {
			cr.Status.AtProvider.DBClusterARN = elem.DBClusterArn
		}
		if elem.DBClusterIdentifier != nil {
			cr.Status.AtProvider.DBClusterIdentifier = elem.DBClusterIdentifier
		}
		if elem.DBClusterMembers != nil {
			f19 := []*svcapitypes.DBClusterMember{}
			for _, f19iter := range elem.DBClusterMembers {
				f19elem := &svcapitypes.DBClusterMember{}
				if f19iter.DBClusterParameterGroupStatus != nil {
					f19elem.DBClusterParameterGroupStatus = f19iter.DBClusterParameterGroupStatus
				}
				if f19iter.DBInstanceIdentifier != nil {
					f19elem.DBInstanceIdentifier = f19iter.DBInstanceIdentifier
				}
				if f19iter.IsClusterWriter != nil {
					f19elem.IsClusterWriter = f19iter.IsClusterWriter
				}
				if f19iter.PromotionTier != nil {
					f19elem.PromotionTier = f19iter.PromotionTier
				}
				f19 = append(f19, f19elem)
			}
			cr.Status.AtProvider.DBClusterMembers = f19
		}
		if elem.DBClusterOptionGroupMemberships != nil {
			f20 := []*svcapitypes.DBClusterOptionGroupStatus{}
			for _, f20iter := range elem.DBClusterOptionGroupMemberships {
				f20elem := &svcapitypes.DBClusterOptionGroupStatus{}
				if f20iter.DBClusterOptionGroupName != nil {
					f20elem.DBClusterOptionGroupName = f20iter.DBClusterOptionGroupName
				}
				if f20iter.Status != nil {
					f20elem.Status = f20iter.Status
				}
				f20 = append(f20, f20elem)
			}
			cr.Status.AtProvider.DBClusterOptionGroupMemberships = f20
		}
		if elem.DBClusterParameterGroup != nil {
			cr.Status.AtProvider.DBClusterParameterGroup = elem.DBClusterParameterGroup
		}
		if elem.DBSubnetGroup != nil {
			cr.Status.AtProvider.DBSubnetGroup = elem.DBSubnetGroup
		}
		if elem.DatabaseName != nil {
			cr.Spec.ForProvider.DatabaseName = elem.DatabaseName
		}
		if elem.DbClusterResourceId != nil {
			cr.Status.AtProvider.DBClusterResourceID = elem.DbClusterResourceId
		}
		if elem.DeletionProtection != nil {
			cr.Spec.ForProvider.DeletionProtection = elem.DeletionProtection
		}
		if elem.DomainMemberships != nil {
			f26 := []*svcapitypes.DomainMembership{}
			for _, f26iter := range elem.DomainMemberships {
				f26elem := &svcapitypes.DomainMembership{}
				if f26iter.Domain != nil {
					f26elem.Domain = f26iter.Domain
				}
				if f26iter.FQDN != nil {
					f26elem.FQDN = f26iter.FQDN
				}
				if f26iter.IAMRoleName != nil {
					f26elem.IAMRoleName = f26iter.IAMRoleName
				}
				if f26iter.Status != nil {
					f26elem.Status = f26iter.Status
				}
				f26 = append(f26, f26elem)
			}
			cr.Status.AtProvider.DomainMemberships = f26
		}
		if elem.EarliestBacktrackTime != nil {
			cr.Status.AtProvider.EarliestBacktrackTime = &metav1.Time{*elem.EarliestBacktrackTime}
		}
		if elem.EarliestRestorableTime != nil {
			cr.Status.AtProvider.EarliestRestorableTime = &metav1.Time{*elem.EarliestRestorableTime}
		}
		if elem.EnabledCloudwatchLogsExports != nil {
			f29 := []*string{}
			for _, f29iter := range elem.EnabledCloudwatchLogsExports {
				var f29elem string
				f29elem = *f29iter
				f29 = append(f29, &f29elem)
			}
			cr.Status.AtProvider.EnabledCloudwatchLogsExports = f29
		}
		if elem.Endpoint != nil {
			cr.Status.AtProvider.Endpoint = elem.Endpoint
		}
		if elem.Engine != nil {
			cr.Spec.ForProvider.Engine = elem.Engine
		}
		if elem.EngineMode != nil {
			cr.Spec.ForProvider.EngineMode = elem.EngineMode
		}
		if elem.EngineVersion != nil {
			cr.Spec.ForProvider.EngineVersion = elem.EngineVersion
		}
		if elem.GlobalWriteForwardingRequested != nil {
			cr.Status.AtProvider.GlobalWriteForwardingRequested = elem.GlobalWriteForwardingRequested
		}
		if elem.GlobalWriteForwardingStatus != nil {
			cr.Status.AtProvider.GlobalWriteForwardingStatus = elem.GlobalWriteForwardingStatus
		}
		if elem.HostedZoneId != nil {
			cr.Status.AtProvider.HostedZoneID = elem.HostedZoneId
		}
		if elem.HttpEndpointEnabled != nil {
			cr.Status.AtProvider.HTTPEndpointEnabled = elem.HttpEndpointEnabled
		}
		if elem.IAMDatabaseAuthenticationEnabled != nil {
			cr.Status.AtProvider.IAMDatabaseAuthenticationEnabled = elem.IAMDatabaseAuthenticationEnabled
		}
		if elem.KmsKeyId != nil {
			cr.Spec.ForProvider.KMSKeyID = elem.KmsKeyId
		}
		if elem.LatestRestorableTime != nil {
			cr.Status.AtProvider.LatestRestorableTime = &metav1.Time{*elem.LatestRestorableTime}
		}
		if elem.MasterUsername != nil {
			cr.Spec.ForProvider.MasterUsername = elem.MasterUsername
		}
		if elem.MultiAZ != nil {
			cr.Status.AtProvider.MultiAZ = elem.MultiAZ
		}
		if elem.PercentProgress != nil {
			cr.Status.AtProvider.PercentProgress = elem.PercentProgress
		}
		if elem.Port != nil {
			cr.Spec.ForProvider.Port = elem.Port
		}
		if elem.PreferredBackupWindow != nil {
			cr.Spec.ForProvider.PreferredBackupWindow = elem.PreferredBackupWindow
		}
		if elem.PreferredMaintenanceWindow != nil {
			cr.Spec.ForProvider.PreferredMaintenanceWindow = elem.PreferredMaintenanceWindow
		}
		if elem.ReadReplicaIdentifiers != nil {
			f47 := []*string{}
			for _, f47iter := range elem.ReadReplicaIdentifiers {
				var f47elem string
				f47elem = *f47iter
				f47 = append(f47, &f47elem)
			}
			cr.Status.AtProvider.ReadReplicaIdentifiers = f47
		}
		if elem.ReaderEndpoint != nil {
			cr.Status.AtProvider.ReaderEndpoint = elem.ReaderEndpoint
		}
		if elem.ReplicationSourceIdentifier != nil {
			cr.Spec.ForProvider.ReplicationSourceIdentifier = elem.ReplicationSourceIdentifier
		}
		if elem.ScalingConfigurationInfo != nil {
			f50 := &svcapitypes.ScalingConfigurationInfo{}
			if elem.ScalingConfigurationInfo.AutoPause != nil {
				f50.AutoPause = elem.ScalingConfigurationInfo.AutoPause
			}
			if elem.ScalingConfigurationInfo.MaxCapacity != nil {
				f50.MaxCapacity = elem.ScalingConfigurationInfo.MaxCapacity
			}
			if elem.ScalingConfigurationInfo.MinCapacity != nil {
				f50.MinCapacity = elem.ScalingConfigurationInfo.MinCapacity
			}
			if elem.ScalingConfigurationInfo.SecondsUntilAutoPause != nil {
				f50.SecondsUntilAutoPause = elem.ScalingConfigurationInfo.SecondsUntilAutoPause
			}
			if elem.ScalingConfigurationInfo.TimeoutAction != nil {
				f50.TimeoutAction = elem.ScalingConfigurationInfo.TimeoutAction
			}
			cr.Status.AtProvider.ScalingConfigurationInfo = f50
		}
		if elem.Status != nil {
			cr.Status.AtProvider.Status = elem.Status
		}
		if elem.StorageEncrypted != nil {
			cr.Spec.ForProvider.StorageEncrypted = elem.StorageEncrypted
		}
		if elem.VpcSecurityGroups != nil {
			f53 := []*svcapitypes.VPCSecurityGroupMembership{}
			for _, f53iter := range elem.VpcSecurityGroups {
				f53elem := &svcapitypes.VPCSecurityGroupMembership{}
				if f53iter.Status != nil {
					f53elem.Status = f53iter.Status
				}
				if f53iter.VpcSecurityGroupId != nil {
					f53elem.VPCSecurityGroupID = f53iter.VpcSecurityGroupId
				}
				f53 = append(f53, f53elem)
			}
			cr.Status.AtProvider.VPCSecurityGroups = f53
		}
		found = true
		break
	}
	if !found {
		return cr
	}

	return cr
}

// GenerateCreateDBClusterInput returns a create input.
func GenerateCreateDBClusterInput(cr *svcapitypes.DBCluster) *svcsdk.CreateDBClusterInput {
	res := &svcsdk.CreateDBClusterInput{}

	if cr.Spec.ForProvider.AvailabilityZones != nil {
		f0 := []*string{}
		for _, f0iter := range cr.Spec.ForProvider.AvailabilityZones {
			var f0elem string
			f0elem = *f0iter
			f0 = append(f0, &f0elem)
		}
		res.SetAvailabilityZones(f0)
	}
	if cr.Spec.ForProvider.BacktrackWindow != nil {
		res.SetBacktrackWindow(*cr.Spec.ForProvider.BacktrackWindow)
	}
	if cr.Spec.ForProvider.BackupRetentionPeriod != nil {
		res.SetBackupRetentionPeriod(*cr.Spec.ForProvider.BackupRetentionPeriod)
	}
	if cr.Spec.ForProvider.CharacterSetName != nil {
		res.SetCharacterSetName(*cr.Spec.ForProvider.CharacterSetName)
	}
	if cr.Spec.ForProvider.CopyTagsToSnapshot != nil {
		res.SetCopyTagsToSnapshot(*cr.Spec.ForProvider.CopyTagsToSnapshot)
	}
	if cr.Spec.ForProvider.DBClusterParameterGroupName != nil {
		res.SetDBClusterParameterGroupName(*cr.Spec.ForProvider.DBClusterParameterGroupName)
	}
	if cr.Spec.ForProvider.DBSubnetGroupName != nil {
		res.SetDBSubnetGroupName(*cr.Spec.ForProvider.DBSubnetGroupName)
	}
	if cr.Spec.ForProvider.DatabaseName != nil {
		res.SetDatabaseName(*cr.Spec.ForProvider.DatabaseName)
	}
	if cr.Spec.ForProvider.DeletionProtection != nil {
		res.SetDeletionProtection(*cr.Spec.ForProvider.DeletionProtection)
	}
	if cr.Spec.ForProvider.DestinationRegion != nil {
		res.SetDestinationRegion(*cr.Spec.ForProvider.DestinationRegion)
	}
	if cr.Spec.ForProvider.Domain != nil {
		res.SetDomain(*cr.Spec.ForProvider.Domain)
	}
	if cr.Spec.ForProvider.DomainIAMRoleName != nil {
		res.SetDomainIAMRoleName(*cr.Spec.ForProvider.DomainIAMRoleName)
	}
	if cr.Spec.ForProvider.EnableCloudwatchLogsExports != nil {
		f12 := []*string{}
		for _, f12iter := range cr.Spec.ForProvider.EnableCloudwatchLogsExports {
			var f12elem string
			f12elem = *f12iter
			f12 = append(f12, &f12elem)
		}
		res.SetEnableCloudwatchLogsExports(f12)
	}
	if cr.Spec.ForProvider.EnableGlobalWriteForwarding != nil {
		res.SetEnableGlobalWriteForwarding(*cr.Spec.ForProvider.EnableGlobalWriteForwarding)
	}
	if cr.Spec.ForProvider.EnableHTTPEndpoint != nil {
		res.SetEnableHttpEndpoint(*cr.Spec.ForProvider.EnableHTTPEndpoint)
	}
	if cr.Spec.ForProvider.EnableIAMDatabaseAuthentication != nil {
		res.SetEnableIAMDatabaseAuthentication(*cr.Spec.ForProvider.EnableIAMDatabaseAuthentication)
	}
	if cr.Spec.ForProvider.Engine != nil {
		res.SetEngine(*cr.Spec.ForProvider.Engine)
	}
	if cr.Spec.ForProvider.EngineMode != nil {
		res.SetEngineMode(*cr.Spec.ForProvider.EngineMode)
	}
	if cr.Spec.ForProvider.EngineVersion != nil {
		res.SetEngineVersion(*cr.Spec.ForProvider.EngineVersion)
	}
	if cr.Spec.ForProvider.GlobalClusterIdentifier != nil {
		res.SetGlobalClusterIdentifier(*cr.Spec.ForProvider.GlobalClusterIdentifier)
	}
	if cr.Spec.ForProvider.KMSKeyID != nil {
		res.SetKmsKeyId(*cr.Spec.ForProvider.KMSKeyID)
	}
	if cr.Spec.ForProvider.MasterUserPassword != nil {
		res.SetMasterUserPassword(*cr.Spec.ForProvider.MasterUserPassword)
	}
	if cr.Spec.ForProvider.MasterUsername != nil {
		res.SetMasterUsername(*cr.Spec.ForProvider.MasterUsername)
	}
	if cr.Spec.ForProvider.OptionGroupName != nil {
		res.SetOptionGroupName(*cr.Spec.ForProvider.OptionGroupName)
	}
	if cr.Spec.ForProvider.Port != nil {
		res.SetPort(*cr.Spec.ForProvider.Port)
	}
	if cr.Spec.ForProvider.PreSignedURL != nil {
		res.SetPreSignedUrl(*cr.Spec.ForProvider.PreSignedURL)
	}
	if cr.Spec.ForProvider.PreferredBackupWindow != nil {
		res.SetPreferredBackupWindow(*cr.Spec.ForProvider.PreferredBackupWindow)
	}
	if cr.Spec.ForProvider.PreferredMaintenanceWindow != nil {
		res.SetPreferredMaintenanceWindow(*cr.Spec.ForProvider.PreferredMaintenanceWindow)
	}
	if cr.Spec.ForProvider.ReplicationSourceIdentifier != nil {
		res.SetReplicationSourceIdentifier(*cr.Spec.ForProvider.ReplicationSourceIdentifier)
	}
	if cr.Spec.ForProvider.ScalingConfiguration != nil {
		f29 := &svcsdk.ScalingConfiguration{}
		if cr.Spec.ForProvider.ScalingConfiguration.AutoPause != nil {
			f29.SetAutoPause(*cr.Spec.ForProvider.ScalingConfiguration.AutoPause)
		}
		if cr.Spec.ForProvider.ScalingConfiguration.MaxCapacity != nil {
			f29.SetMaxCapacity(*cr.Spec.ForProvider.ScalingConfiguration.MaxCapacity)
		}
		if cr.Spec.ForProvider.ScalingConfiguration.MinCapacity != nil {
			f29.SetMinCapacity(*cr.Spec.ForProvider.ScalingConfiguration.MinCapacity)
		}
		if cr.Spec.ForProvider.ScalingConfiguration.SecondsUntilAutoPause != nil {
			f29.SetSecondsUntilAutoPause(*cr.Spec.ForProvider.ScalingConfiguration.SecondsUntilAutoPause)
		}
		if cr.Spec.ForProvider.ScalingConfiguration.TimeoutAction != nil {
			f29.SetTimeoutAction(*cr.Spec.ForProvider.ScalingConfiguration.TimeoutAction)
		}
		res.SetScalingConfiguration(f29)
	}
	if cr.Spec.ForProvider.SourceRegion != nil {
		res.SetSourceRegion(*cr.Spec.ForProvider.SourceRegion)
	}
	if cr.Spec.ForProvider.StorageEncrypted != nil {
		res.SetStorageEncrypted(*cr.Spec.ForProvider.StorageEncrypted)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f32 := []*svcsdk.Tag{}
		for _, f32iter := range cr.Spec.ForProvider.Tags {
			f32elem := &svcsdk.Tag{}
			if f32iter.Key != nil {
				f32elem.SetKey(*f32iter.Key)
			}
			if f32iter.Value != nil {
				f32elem.SetValue(*f32iter.Value)
			}
			f32 = append(f32, f32elem)
		}
		res.SetTags(f32)
	}

	return res
}

// GenerateModifyDBClusterInput returns an update input.
func GenerateModifyDBClusterInput(cr *svcapitypes.DBCluster) *svcsdk.ModifyDBClusterInput {
	res := &svcsdk.ModifyDBClusterInput{}

	if cr.Spec.ForProvider.BacktrackWindow != nil {
		res.SetBacktrackWindow(*cr.Spec.ForProvider.BacktrackWindow)
	}
	if cr.Spec.ForProvider.BackupRetentionPeriod != nil {
		res.SetBackupRetentionPeriod(*cr.Spec.ForProvider.BackupRetentionPeriod)
	}
	if cr.Spec.ForProvider.CopyTagsToSnapshot != nil {
		res.SetCopyTagsToSnapshot(*cr.Spec.ForProvider.CopyTagsToSnapshot)
	}
	if cr.Spec.ForProvider.DBClusterParameterGroupName != nil {
		res.SetDBClusterParameterGroupName(*cr.Spec.ForProvider.DBClusterParameterGroupName)
	}
	if cr.Spec.ForProvider.DeletionProtection != nil {
		res.SetDeletionProtection(*cr.Spec.ForProvider.DeletionProtection)
	}
	if cr.Spec.ForProvider.Domain != nil {
		res.SetDomain(*cr.Spec.ForProvider.Domain)
	}
	if cr.Spec.ForProvider.DomainIAMRoleName != nil {
		res.SetDomainIAMRoleName(*cr.Spec.ForProvider.DomainIAMRoleName)
	}
	if cr.Spec.ForProvider.EnableGlobalWriteForwarding != nil {
		res.SetEnableGlobalWriteForwarding(*cr.Spec.ForProvider.EnableGlobalWriteForwarding)
	}
	if cr.Spec.ForProvider.EnableHTTPEndpoint != nil {
		res.SetEnableHttpEndpoint(*cr.Spec.ForProvider.EnableHTTPEndpoint)
	}
	if cr.Spec.ForProvider.EnableIAMDatabaseAuthentication != nil {
		res.SetEnableIAMDatabaseAuthentication(*cr.Spec.ForProvider.EnableIAMDatabaseAuthentication)
	}
	if cr.Spec.ForProvider.EngineVersion != nil {
		res.SetEngineVersion(*cr.Spec.ForProvider.EngineVersion)
	}
	if cr.Spec.ForProvider.MasterUserPassword != nil {
		res.SetMasterUserPassword(*cr.Spec.ForProvider.MasterUserPassword)
	}
	if cr.Spec.ForProvider.OptionGroupName != nil {
		res.SetOptionGroupName(*cr.Spec.ForProvider.OptionGroupName)
	}
	if cr.Spec.ForProvider.Port != nil {
		res.SetPort(*cr.Spec.ForProvider.Port)
	}
	if cr.Spec.ForProvider.PreferredBackupWindow != nil {
		res.SetPreferredBackupWindow(*cr.Spec.ForProvider.PreferredBackupWindow)
	}
	if cr.Spec.ForProvider.PreferredMaintenanceWindow != nil {
		res.SetPreferredMaintenanceWindow(*cr.Spec.ForProvider.PreferredMaintenanceWindow)
	}
	if cr.Spec.ForProvider.ScalingConfiguration != nil {
		f21 := &svcsdk.ScalingConfiguration{}
		if cr.Spec.ForProvider.ScalingConfiguration.AutoPause != nil {
			f21.SetAutoPause(*cr.Spec.ForProvider.ScalingConfiguration.AutoPause)
		}
		if cr.Spec.ForProvider.ScalingConfiguration.MaxCapacity != nil {
			f21.SetMaxCapacity(*cr.Spec.ForProvider.ScalingConfiguration.MaxCapacity)
		}
		if cr.Spec.ForProvider.ScalingConfiguration.MinCapacity != nil {
			f21.SetMinCapacity(*cr.Spec.ForProvider.ScalingConfiguration.MinCapacity)
		}
		if cr.Spec.ForProvider.ScalingConfiguration.SecondsUntilAutoPause != nil {
			f21.SetSecondsUntilAutoPause(*cr.Spec.ForProvider.ScalingConfiguration.SecondsUntilAutoPause)
		}
		if cr.Spec.ForProvider.ScalingConfiguration.TimeoutAction != nil {
			f21.SetTimeoutAction(*cr.Spec.ForProvider.ScalingConfiguration.TimeoutAction)
		}
		res.SetScalingConfiguration(f21)
	}

	return res
}

// GenerateDeleteDBClusterInput returns a deletion input.
func GenerateDeleteDBClusterInput(cr *svcapitypes.DBCluster) *svcsdk.DeleteDBClusterInput {
	res := &svcsdk.DeleteDBClusterInput{}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "DBClusterNotFoundFault"
}