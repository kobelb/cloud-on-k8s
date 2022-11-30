// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

package kibana

import (
	"github.com/elastic/cloud-on-k8s/v2/pkg/controller/common/labels"
)

const (
	// KibanaNameLabelName used to represent a Kibana in k8s resources
	KibanaNameLabelName = "kibana.k8s.elastic.co/name"

	// KibanaNamespaceLabelName used to represent a Kibana in k8s resources
	KibanaNamespaceLabelName = "kibana.k8s.elastic.co/namespace"

	// KibanaVersionLabelName used to propagate Kibana version from the spec to the pods
	KibanaVersionLabelName = "kibana.k8s.elastic.co/version"

	KibanaDeploymentTypeLabelName = "kibana.k8s.elastic.co/deployment-type"

	KibanaIngressLabelName = "kibana.k8s.elastic.co/ingress"

	// Type represents the Kibana type
	Type = "kibana"
)

type DeploymentType int

const (
	All DeploymentType = iota + 1
	UI
	BackgroundTasks
)

func (dt DeploymentType) String() string {
	return [...]string{"all", "ui", "background-tasks"}[dt-1]
}

// NewLabels constructs a new set of labels for a Kibana pod
func NewLabels(kibanaName string, deploymentType *DeploymentType, ingress bool) map[string]string {
	labels := map[string]string{
		KibanaNameLabelName:  kibanaName,
		labels.TypeLabelName: Type,
	}

	if deploymentType != nil {
		labels[KibanaDeploymentTypeLabelName] = deploymentType.String()
	}

	if ingress == true {
		labels[KibanaIngressLabelName] = "true"
	}

	return labels
}

func DefaultLabels(kibanaName string) map[string]string {
	return map[string]string{
		KibanaNameLabelName:  kibanaName,
		labels.TypeLabelName: Type,
	}
}
