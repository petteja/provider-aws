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

package replicationsubnetgroup

import (
	"context"
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/google/go-cmp/cmp"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/databasemigrationservice"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	svcapitypes "github.com/crossplane/provider-aws/apis/databasemigrationservice/v1alpha1"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"
)

// SetupReplicationSubnetGroup adds a controller that reconciles ReplicationSubnetGroup
func SetupReplicationSubnetGroup(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter, poll time.Duration) error {
	name := managed.ControllerName(svcapitypes.ReplicationSubnetGroupGroupKind)
	opts := []option{
		func(e *external) {
			e.preObserve = preObserve
			e.postObserve = postObserve
			e.preCreate = preCreate
			e.preUpdate = preUpdate
			e.preDelete = preDelete
			e.isUpToDate = isUpToDate
		},
	}
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(controller.Options{
			RateLimiter: ratelimiter.NewDefaultManagedRateLimiter(rl),
		}).
		For(&svcapitypes.ReplicationSubnetGroup{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(svcapitypes.ReplicationSubnetGroupGroupVersionKind),
			managed.WithExternalConnecter(&connector{kube: mgr.GetClient(), opts: opts}),
			managed.WithPollInterval(poll),
			managed.WithLogger(l.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

func isUpToDate(cr *svcapitypes.ReplicationSubnetGroup, obj *svcsdk.DescribeReplicationSubnetGroupsOutput) (bool, error) {

	if !cmp.Equal(cr.Spec.ForProvider.ReplicationSubnetGroupDescription, obj.ReplicationSubnetGroups[0].ReplicationSubnetGroupDescription) {
		return false, nil
	}

	return true, nil
}
//arn:aws:dms:eu-north-1:964273356747:subgrp:replication-subnet-group


//func postCreate(_ context.Context, cr *svcapitypes.ReplicationSubnetGroup, obj *svcsdk.CreateReplicationSubnetGroupOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
//	//arn:aws:dms:eu-north-1:964273356747:subgrp:replication-subnet-group
//	//FIXME: Get account number
//	cr.Spec.ForProvider.ReplicationSubnetGroupArn = aws.String(fmt.Sprintf("arn:aws:dms:%s:%d:subgrp:%s", cr.Spec.ForProvider.Region, 964273356747, cr.Spec.ForProvider.ReplicationSubnetGroupIdentifier))
//	return cre, err
//}

func preObserve(_ context.Context, cr *svcapitypes.ReplicationSubnetGroup, obj *svcsdk.DescribeReplicationSubnetGroupsInput) error {
	obj.SetFilters([]*svcsdk.Filter{{
		Name:   aws.String("replication-subnet-group-id"),
		Values: []*string{aws.String(meta.GetExternalName(cr))},
	}})
	return nil
}
func postObserve(_ context.Context, cr *svcapitypes.ReplicationSubnetGroup, obj *svcsdk.DescribeReplicationSubnetGroupsOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	if err != nil {
		return managed.ExternalObservation{}, err
	}
	if aws.StringValue(obj.ReplicationSubnetGroups[0].SubnetGroupStatus) == "Complete" {
		cr.SetConditions(xpv1.Available())
	}
	return obs, err
}

func preCreate(_ context.Context, cr *svcapitypes.ReplicationSubnetGroup, obj *svcsdk.CreateReplicationSubnetGroupInput) error {
	obj.ReplicationSubnetGroupIdentifier = aws.String(meta.GetExternalName(cr))
	return nil
}

func preUpdate(_ context.Context, cr *svcapitypes.ReplicationSubnetGroup, obj *svcsdk.ModifyReplicationSubnetGroupInput) error {
	obj.ReplicationSubnetGroupIdentifier = aws.String(meta.GetExternalName(cr))
	return nil
}

func preDelete(_ context.Context, cr *svcapitypes.ReplicationSubnetGroup, obj *svcsdk.DeleteReplicationSubnetGroupInput) (bool, error) {
	obj.ReplicationSubnetGroupIdentifier = aws.String(meta.GetExternalName(cr))
	return false, nil
}
