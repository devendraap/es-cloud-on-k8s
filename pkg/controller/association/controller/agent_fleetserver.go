// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	agentv1alpha1 "github.com/devendra/es-cloud-on-k8s/v2/pkg/apis/agent/v1alpha1"
	commonv1 "github.com/devendra/es-cloud-on-k8s/v2/pkg/apis/common/v1"
	"github.com/devendra/es-cloud-on-k8s/v2/pkg/controller/agent"
	"github.com/devendra/es-cloud-on-k8s/v2/pkg/controller/association"
	"github.com/devendra/es-cloud-on-k8s/v2/pkg/controller/common/operator"
	"github.com/devendra/es-cloud-on-k8s/v2/pkg/utils/k8s"
	ulog "github.com/devendra/es-cloud-on-k8s/v2/pkg/utils/log"
	"github.com/devendra/es-cloud-on-k8s/v2/pkg/utils/rbac"
)

func AddAgentFleetServer(mgr manager.Manager, accessReviewer rbac.AccessReviewer, params operator.Parameters) error {
	return association.AddAssociationController(mgr, accessReviewer, params, association.AssociationInfo{
		AssociatedObjTemplate:     func() commonv1.Associated { return &agentv1alpha1.Agent{} },
		ReferencedObjTemplate:     func() client.Object { return &agentv1alpha1.Agent{} },
		ExternalServiceURL:        getFleetServerExternalURL,
		ReferencedResourceVersion: referencedFleetServerStatusVersion,
		ReferencedResourceNamer:   agent.Namer,
		AssociationName:           "agent-fleetserver",
		AssociatedShortName:       "agent",
		AssociationType:           commonv1.FleetServerAssociationType,
		AdditionalSecrets:         additionalSecrets,
		Labels: func(associated types.NamespacedName) map[string]string {
			return map[string]string{
				AgentAssociationLabelName:      associated.Name,
				AgentAssociationLabelNamespace: associated.Namespace,
				AgentAssociationLabelType:      commonv1.FleetServerAssociationType,
			}
		},
		AssociationConfAnnotationNameBase:     commonv1.FleetServerConfigAnnotationNameBase,
		AssociationResourceNameLabelName:      agent.NameLabelName,
		AssociationResourceNamespaceLabelName: agent.NamespaceLabelName,

		ElasticsearchUserCreation: nil,
	})
}

func additionalSecrets(ctx context.Context, c k8s.Client, assoc commonv1.Association) ([]types.NamespacedName, error) {
	log := ulog.FromContext(ctx)
	associated := assoc.Associated()
	var agent agentv1alpha1.Agent
	nsn := types.NamespacedName{Namespace: associated.GetNamespace(), Name: associated.GetName()}
	if err := c.Get(ctx, nsn, &agent); err != nil {
		return nil, err
	}
	fleetServerRef := assoc.AssociationRef()
	if !fleetServerRef.IsDefined() {
		return nil, nil
	}
	fleetServer := agentv1alpha1.Agent{}
	if err := c.Get(ctx, fleetServerRef.NamespacedName(), &fleetServer); err != nil {
		return nil, err
	}

	// If the Fleet Server Agent is not associated with an Elasticsearch cluster
	// (potentially because of a manual setup) we should do nothing.
	if len(fleetServer.Spec.ElasticsearchRefs) == 0 {
		return nil, nil
	}
	esAssociation, err := association.SingleAssociationOfType(fleetServer.GetAssociations(), commonv1.ElasticsearchAssociationType)
	if err != nil {
		return nil, err
	}

	conf, err := esAssociation.AssociationConf()
	if err != nil {
		log.V(1).Info("no additional secrets because no assoc conf")
		return nil, err
	}
	if conf == nil || !conf.CACertProvided {
		log.V(1).Info("no additional secrets because conf nil or no CA provided")
		return nil, nil
	}
	return []types.NamespacedName{{
		Namespace: fleetServer.Namespace,
		Name:      conf.CASecretName,
	}}, nil
}

func getFleetServerExternalURL(c k8s.Client, assoc commonv1.Association) (string, error) {
	fleetServerRef := assoc.AssociationRef()
	if !fleetServerRef.IsDefined() {
		return "", nil
	}
	fleetServer := agentv1alpha1.Agent{}
	if err := c.Get(context.Background(), fleetServerRef.NamespacedName(), &fleetServer); err != nil {
		return "", err
	}
	serviceName := fleetServerRef.ServiceName
	if serviceName == "" {
		serviceName = agent.HTTPServiceName(fleetServer.Name)
	}
	nsn := types.NamespacedName{Namespace: fleetServer.Namespace, Name: serviceName}
	return association.ServiceURL(c, nsn, fleetServer.Spec.HTTP.Protocol())
}

// referencedFleetServerStatusVersion returns the currently running version of Agent
// reported in its status.
func referencedFleetServerStatusVersion(c k8s.Client, fsRef commonv1.ObjectSelector) (string, error) {
	if fsRef.IsExternal() {
		info, err := association.GetUnmanagedAssociationConnectionInfoFromSecret(c, fsRef)
		if err != nil {
			return "", err
		}
		ver, err := info.Version("/api/status", "{ .version.number }")
		if err != nil {
			// version is in the status API from version 8.0
			if err.Error() == "version is not found" {
				return association.UnknownVersion, nil
			}
			return "", err
		}
		return ver, nil
	}

	var fleetServer agentv1alpha1.Agent
	err := c.Get(context.Background(), fsRef.NamespacedName(), &fleetServer)
	if err != nil {
		return "", err
	}
	return fleetServer.Status.Version, nil
}
