//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MavenConfigObservation) DeepCopyInto(out *MavenConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MavenConfigObservation.
func (in *MavenConfigObservation) DeepCopy() *MavenConfigObservation {
	if in == nil {
		return nil
	}
	out := new(MavenConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MavenConfigParameters) DeepCopyInto(out *MavenConfigParameters) {
	*out = *in
	if in.AllowSnapshotOverwrites != nil {
		in, out := &in.AllowSnapshotOverwrites, &out.AllowSnapshotOverwrites
		*out = new(bool)
		**out = **in
	}
	if in.VersionPolicy != nil {
		in, out := &in.VersionPolicy, &out.VersionPolicy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MavenConfigParameters.
func (in *MavenConfigParameters) DeepCopy() *MavenConfigParameters {
	if in == nil {
		return nil
	}
	out := new(MavenConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryRepository) DeepCopyInto(out *RegistryRepository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryRepository.
func (in *RegistryRepository) DeepCopy() *RegistryRepository {
	if in == nil {
		return nil
	}
	out := new(RegistryRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegistryRepository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryRepositoryList) DeepCopyInto(out *RegistryRepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RegistryRepository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryRepositoryList.
func (in *RegistryRepositoryList) DeepCopy() *RegistryRepositoryList {
	if in == nil {
		return nil
	}
	out := new(RegistryRepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegistryRepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryRepositoryObservation) DeepCopyInto(out *RegistryRepositoryObservation) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryRepositoryObservation.
func (in *RegistryRepositoryObservation) DeepCopy() *RegistryRepositoryObservation {
	if in == nil {
		return nil
	}
	out := new(RegistryRepositoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryRepositoryParameters) DeepCopyInto(out *RegistryRepositoryParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyName != nil {
		in, out := &in.KMSKeyName, &out.KMSKeyName
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MavenConfig != nil {
		in, out := &in.MavenConfig, &out.MavenConfig
		*out = make([]MavenConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryRepositoryParameters.
func (in *RegistryRepositoryParameters) DeepCopy() *RegistryRepositoryParameters {
	if in == nil {
		return nil
	}
	out := new(RegistryRepositoryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryRepositorySpec) DeepCopyInto(out *RegistryRepositorySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryRepositorySpec.
func (in *RegistryRepositorySpec) DeepCopy() *RegistryRepositorySpec {
	if in == nil {
		return nil
	}
	out := new(RegistryRepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryRepositoryStatus) DeepCopyInto(out *RegistryRepositoryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryRepositoryStatus.
func (in *RegistryRepositoryStatus) DeepCopy() *RegistryRepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(RegistryRepositoryStatus)
	in.DeepCopyInto(out)
	return out
}
