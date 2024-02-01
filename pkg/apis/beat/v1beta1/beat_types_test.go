// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.
package v1beta1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApmEsAssociation_AssociationConfAnnotationName(t *testing.T) {
	bea := BeatESAssociation{}
	require.Equal(t, "association.k8s.acceldata.io/es-conf", bea.AssociationConfAnnotationName())
}

func TestApmKibanaAssociation_AssociationConfAnnotationName(t *testing.T) {
	bka := BeatKibanaAssociation{}
	require.Equal(t, "association.k8s.acceldata.io/kb-conf", bka.AssociationConfAnnotationName())
}
