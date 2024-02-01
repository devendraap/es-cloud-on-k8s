// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

package annotation

const (
	// CurrAssocStatusAnnotation describes the currently observed association status of an object.
	CurrAssocStatusAnnotation = "association.k8s.acceldata.io/current-status"
	// PrevAssocStatusAnnotation describes the previously observed association status of an object.
	PrevAssocStatusAnnotation = "association.k8s.acceldata.io/previous-status"

	// UpdateAnnotation is the name of the annotation applied to pods to force kubelet to resync secrets
	UpdateAnnotation = "update.k8s.acceldata.io/timestamp"

	// FilebeatModuleAnnotation is the name of the annotation applied to pods to give a hint to filebeat so that it
	// uses the appropriate module to analyze the logs of the container.
	// https://www.elastic.co/guide/en/beats/filebeat/current/configuration-autodiscover-hints.html#_co_elastic_logsmodule
	FilebeatModuleAnnotation = "co.elastic.logs/module"

	SecureSettingsSecretsAnnotationName = "policy.k8s.acceldata.io/secure-settings-secrets" //nolint:gosec
	SettingsHashAnnotationName          = "policy.k8s.acceldata.io/settings-hash"

	KibanaConfigHashAnnotation = "policy.k8s.acceldata.io/kibana-config-hash"

	ElasticsearchConfigAndSecretMountsHashAnnotation = "policy.k8s.acceldata.io/elasticsearch-config-mounts-hash" //nolint:gosec
	SourceSecretAnnotationName                       = "policy.k8s.acceldata.io/source-secret-name"               //nolint:gosec
)
