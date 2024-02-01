// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

package label

import (
	"k8s.io/apimachinery/pkg/types"

	commonv1 "github.com/devendra/es-cloud-on-k8s/v2/pkg/apis/common/v1"
)

const (
	// KibanaNameLabelName used to represent a Kibana in k8s resources
	KibanaNameLabelName = "kibana.k8s.acceldata.io/name"

	// KibanaNamespaceLabelName used to represent a Kibana in k8s resources
	KibanaNamespaceLabelName = "kibana.k8s.acceldata.io/namespace"

	// KibanaVersionLabelName used to propagate Kibana version from the spec to the pods
	KibanaVersionLabelName = "kibana.k8s.acceldata.io/version"

	// Type represents the Kibana type
	Type = "kibana"
)

// NewLabels constructs a new set of labels from Kibana definition.
func NewLabels(kb types.NamespacedName) map[string]string {
	return map[string]string{
		KibanaNameLabelName:    kb.Name,
		commonv1.TypeLabelName: Type,
	}
}
