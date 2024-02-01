//go:build !ignore_autogenerated

// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/devendra/es-cloud-on-k8s/v2/pkg/apis/common/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticMapsServer) DeepCopyInto(out *ElasticMapsServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	if in.assocConf != nil {
		in, out := &in.assocConf, &out.assocConf
		*out = new(v1.AssociationConf)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticMapsServer.
func (in *ElasticMapsServer) DeepCopy() *ElasticMapsServer {
	if in == nil {
		return nil
	}
	out := new(ElasticMapsServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticMapsServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticMapsServerList) DeepCopyInto(out *ElasticMapsServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ElasticMapsServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticMapsServerList.
func (in *ElasticMapsServerList) DeepCopy() *ElasticMapsServerList {
	if in == nil {
		return nil
	}
	out := new(ElasticMapsServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticMapsServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MapsSpec) DeepCopyInto(out *MapsSpec) {
	*out = *in
	out.ElasticsearchRef = in.ElasticsearchRef
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = (*in).DeepCopy()
	}
	if in.ConfigRef != nil {
		in, out := &in.ConfigRef, &out.ConfigRef
		*out = new(v1.ConfigSource)
		**out = **in
	}
	in.HTTP.DeepCopyInto(&out.HTTP)
	in.PodTemplate.DeepCopyInto(&out.PodTemplate)
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MapsSpec.
func (in *MapsSpec) DeepCopy() *MapsSpec {
	if in == nil {
		return nil
	}
	out := new(MapsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MapsStatus) DeepCopyInto(out *MapsStatus) {
	*out = *in
	out.DeploymentStatus = in.DeploymentStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MapsStatus.
func (in *MapsStatus) DeepCopy() *MapsStatus {
	if in == nil {
		return nil
	}
	out := new(MapsStatus)
	in.DeepCopyInto(out)
	return out
}
