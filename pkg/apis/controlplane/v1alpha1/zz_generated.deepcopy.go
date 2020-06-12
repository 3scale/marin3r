// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryService) DeepCopyInto(out *DiscoveryService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryService.
func (in *DiscoveryService) DeepCopy() *DiscoveryService {
	if in == nil {
		return nil
	}
	out := new(DiscoveryService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiscoveryService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryServiceList) DeepCopyInto(out *DiscoveryServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DiscoveryService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryServiceList.
func (in *DiscoveryServiceList) DeepCopy() *DiscoveryServiceList {
	if in == nil {
		return nil
	}
	out := new(DiscoveryServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiscoveryServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryServiceSpec) DeepCopyInto(out *DiscoveryServiceSpec) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.ServerCertificateRef = in.ServerCertificateRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryServiceSpec.
func (in *DiscoveryServiceSpec) DeepCopy() *DiscoveryServiceSpec {
	if in == nil {
		return nil
	}
	out := new(DiscoveryServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryServiceStatus) DeepCopyInto(out *DiscoveryServiceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryServiceStatus.
func (in *DiscoveryServiceStatus) DeepCopy() *DiscoveryServiceStatus {
	if in == nil {
		return nil
	}
	out := new(DiscoveryServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDiscoveryCertificate) DeepCopyInto(out *ServiceDiscoveryCertificate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDiscoveryCertificate.
func (in *ServiceDiscoveryCertificate) DeepCopy() *ServiceDiscoveryCertificate {
	if in == nil {
		return nil
	}
	out := new(ServiceDiscoveryCertificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceDiscoveryCertificate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDiscoveryCertificateList) DeepCopyInto(out *ServiceDiscoveryCertificateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceDiscoveryCertificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDiscoveryCertificateList.
func (in *ServiceDiscoveryCertificateList) DeepCopy() *ServiceDiscoveryCertificateList {
	if in == nil {
		return nil
	}
	out := new(ServiceDiscoveryCertificateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceDiscoveryCertificateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDiscoveryCertificateSpec) DeepCopyInto(out *ServiceDiscoveryCertificateSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDiscoveryCertificateSpec.
func (in *ServiceDiscoveryCertificateSpec) DeepCopy() *ServiceDiscoveryCertificateSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceDiscoveryCertificateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDiscoveryCertificateStatus) DeepCopyInto(out *ServiceDiscoveryCertificateStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDiscoveryCertificateStatus.
func (in *ServiceDiscoveryCertificateStatus) DeepCopy() *ServiceDiscoveryCertificateStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceDiscoveryCertificateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidecarConfig) DeepCopyInto(out *SidecarConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidecarConfig.
func (in *SidecarConfig) DeepCopy() *SidecarConfig {
	if in == nil {
		return nil
	}
	out := new(SidecarConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SidecarConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidecarConfigList) DeepCopyInto(out *SidecarConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SidecarConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidecarConfigList.
func (in *SidecarConfigList) DeepCopy() *SidecarConfigList {
	if in == nil {
		return nil
	}
	out := new(SidecarConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SidecarConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidecarConfigSpec) DeepCopyInto(out *SidecarConfigSpec) {
	*out = *in
	out.ClientCertificateRef = in.ClientCertificateRef
	out.BootstrapConfigMapRef = in.BootstrapConfigMapRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidecarConfigSpec.
func (in *SidecarConfigSpec) DeepCopy() *SidecarConfigSpec {
	if in == nil {
		return nil
	}
	out := new(SidecarConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidecarConfigStatus) DeepCopyInto(out *SidecarConfigStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidecarConfigStatus.
func (in *SidecarConfigStatus) DeepCopy() *SidecarConfigStatus {
	if in == nil {
		return nil
	}
	out := new(SidecarConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidecarInjectorConfig) DeepCopyInto(out *SidecarInjectorConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidecarInjectorConfig.
func (in *SidecarInjectorConfig) DeepCopy() *SidecarInjectorConfig {
	if in == nil {
		return nil
	}
	out := new(SidecarInjectorConfig)
	in.DeepCopyInto(out)
	return out
}
