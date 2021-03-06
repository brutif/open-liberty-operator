// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1beta1

import (
	common "github.com/application-stacks/runtime-component-operator/pkg/common"
	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	v1alpha2 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha2"
	routev1 "github.com/openshift/api/route/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Certificate) DeepCopyInto(out *Certificate) {
	*out = *in
	if in.Organization != nil {
		in, out := &in.Organization, &out.Organization
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(v1.Duration)
		**out = **in
	}
	if in.RenewBefore != nil {
		in, out := &in.RenewBefore, &out.RenewBefore
		*out = new(v1.Duration)
		**out = **in
	}
	if in.DNSNames != nil {
		in, out := &in.DNSNames, &out.DNSNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.URISANs != nil {
		in, out := &in.URISANs, &out.URISANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.IssuerRef = in.IssuerRef
	if in.Usages != nil {
		in, out := &in.Usages, &out.Usages
		*out = make([]v1alpha2.KeyUsage, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Certificate.
func (in *Certificate) DeepCopy() *Certificate {
	if in == nil {
		return nil
	}
	out := new(Certificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponent) DeepCopyInto(out *RuntimeComponent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponent.
func (in *RuntimeComponent) DeepCopy() *RuntimeComponent {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RuntimeComponent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentAffinity) DeepCopyInto(out *RuntimeComponentAffinity) {
	*out = *in
	if in.NodeAffinity != nil {
		in, out := &in.NodeAffinity, &out.NodeAffinity
		*out = new(corev1.NodeAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.PodAffinity != nil {
		in, out := &in.PodAffinity, &out.PodAffinity
		*out = new(corev1.PodAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.PodAntiAffinity != nil {
		in, out := &in.PodAntiAffinity, &out.PodAntiAffinity
		*out = new(corev1.PodAntiAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Architecture != nil {
		in, out := &in.Architecture, &out.Architecture
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeAffinityLabels != nil {
		in, out := &in.NodeAffinityLabels, &out.NodeAffinityLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentAffinity.
func (in *RuntimeComponentAffinity) DeepCopy() *RuntimeComponentAffinity {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentAffinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentAutoScaling) DeepCopyInto(out *RuntimeComponentAutoScaling) {
	*out = *in
	if in.TargetCPUUtilizationPercentage != nil {
		in, out := &in.TargetCPUUtilizationPercentage, &out.TargetCPUUtilizationPercentage
		*out = new(int32)
		**out = **in
	}
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentAutoScaling.
func (in *RuntimeComponentAutoScaling) DeepCopy() *RuntimeComponentAutoScaling {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentAutoScaling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentBindings) DeepCopyInto(out *RuntimeComponentBindings) {
	*out = *in
	if in.AutoDetect != nil {
		in, out := &in.AutoDetect, &out.AutoDetect
		*out = new(bool)
		**out = **in
	}
	if in.Embedded != nil {
		in, out := &in.Embedded, &out.Embedded
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentBindings.
func (in *RuntimeComponentBindings) DeepCopy() *RuntimeComponentBindings {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentBindings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentList) DeepCopyInto(out *RuntimeComponentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RuntimeComponent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentList.
func (in *RuntimeComponentList) DeepCopy() *RuntimeComponentList {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RuntimeComponentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentMonitoring) DeepCopyInto(out *RuntimeComponentMonitoring) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]monitoringv1.Endpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentMonitoring.
func (in *RuntimeComponentMonitoring) DeepCopy() *RuntimeComponentMonitoring {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentMonitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentRoute) DeepCopyInto(out *RuntimeComponentRoute) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Termination != nil {
		in, out := &in.Termination, &out.Termination
		*out = new(routev1.TLSTerminationType)
		**out = **in
	}
	if in.InsecureEdgeTerminationPolicy != nil {
		in, out := &in.InsecureEdgeTerminationPolicy, &out.InsecureEdgeTerminationPolicy
		*out = new(routev1.InsecureEdgeTerminationPolicyType)
		**out = **in
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(Certificate)
		(*in).DeepCopyInto(*out)
	}
	if in.CertificateSecretRef != nil {
		in, out := &in.CertificateSecretRef, &out.CertificateSecretRef
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentRoute.
func (in *RuntimeComponentRoute) DeepCopy() *RuntimeComponentRoute {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentService) DeepCopyInto(out *RuntimeComponentService) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(corev1.ServiceType)
		**out = **in
	}
	if in.TargetPort != nil {
		in, out := &in.TargetPort, &out.TargetPort
		*out = new(int32)
		**out = **in
	}
	if in.NodePort != nil {
		in, out := &in.NodePort, &out.NodePort
		*out = new(int32)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]corev1.ServicePort, len(*in))
		copy(*out, *in)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Consumes != nil {
		in, out := &in.Consumes, &out.Consumes
		*out = make([]ServiceBindingConsumes, len(*in))
		copy(*out, *in)
	}
	if in.Provides != nil {
		in, out := &in.Provides, &out.Provides
		*out = new(ServiceBindingProvides)
		(*in).DeepCopyInto(*out)
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(Certificate)
		(*in).DeepCopyInto(*out)
	}
	if in.CertificateSecretRef != nil {
		in, out := &in.CertificateSecretRef, &out.CertificateSecretRef
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentService.
func (in *RuntimeComponentService) DeepCopy() *RuntimeComponentService {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentSpec) DeepCopyInto(out *RuntimeComponentSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Autoscaling != nil {
		in, out := &in.Autoscaling, &out.Autoscaling
		*out = new(RuntimeComponentAutoScaling)
		(*in).DeepCopyInto(*out)
	}
	if in.PullPolicy != nil {
		in, out := &in.PullPolicy, &out.PullPolicy
		*out = new(corev1.PullPolicy)
		**out = **in
	}
	if in.PullSecret != nil {
		in, out := &in.PullSecret, &out.PullSecret
		*out = new(string)
		**out = **in
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceConstraints != nil {
		in, out := &in.ResourceConstraints, &out.ResourceConstraints
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(corev1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(corev1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(RuntimeComponentService)
		(*in).DeepCopyInto(*out)
	}
	if in.Expose != nil {
		in, out := &in.Expose, &out.Expose
		*out = new(bool)
		**out = **in
	}
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]corev1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceAccountName != nil {
		in, out := &in.ServiceAccountName, &out.ServiceAccountName
		*out = new(string)
		**out = **in
	}
	if in.Architecture != nil {
		in, out := &in.Architecture, &out.Architecture
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(RuntimeComponentStorage)
		(*in).DeepCopyInto(*out)
	}
	if in.CreateKnativeService != nil {
		in, out := &in.CreateKnativeService, &out.CreateKnativeService
		*out = new(bool)
		**out = **in
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		*out = new(RuntimeComponentMonitoring)
		(*in).DeepCopyInto(*out)
	}
	if in.CreateAppDefinition != nil {
		in, out := &in.CreateAppDefinition, &out.CreateAppDefinition
		*out = new(bool)
		**out = **in
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]corev1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SidecarContainers != nil {
		in, out := &in.SidecarContainers, &out.SidecarContainers
		*out = make([]corev1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = new(RuntimeComponentRoute)
		(*in).DeepCopyInto(*out)
	}
	if in.Bindings != nil {
		in, out := &in.Bindings, &out.Bindings
		*out = new(RuntimeComponentBindings)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(RuntimeComponentAffinity)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentSpec.
func (in *RuntimeComponentSpec) DeepCopy() *RuntimeComponentSpec {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentStatus) DeepCopyInto(out *RuntimeComponentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]StatusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConsumedServices != nil {
		in, out := &in.ConsumedServices, &out.ConsumedServices
		*out = make(common.ConsumedServices, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.ResolvedBindings != nil {
		in, out := &in.ResolvedBindings, &out.ResolvedBindings
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentStatus.
func (in *RuntimeComponentStatus) DeepCopy() *RuntimeComponentStatus {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentStorage) DeepCopyInto(out *RuntimeComponentStorage) {
	*out = *in
	if in.VolumeClaimTemplate != nil {
		in, out := &in.VolumeClaimTemplate, &out.VolumeClaimTemplate
		*out = new(corev1.PersistentVolumeClaim)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentStorage.
func (in *RuntimeComponentStorage) DeepCopy() *RuntimeComponentStorage {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingAuth) DeepCopyInto(out *ServiceBindingAuth) {
	*out = *in
	in.Username.DeepCopyInto(&out.Username)
	in.Password.DeepCopyInto(&out.Password)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingAuth.
func (in *ServiceBindingAuth) DeepCopy() *ServiceBindingAuth {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingConsumes) DeepCopyInto(out *ServiceBindingConsumes) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingConsumes.
func (in *ServiceBindingConsumes) DeepCopy() *ServiceBindingConsumes {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingConsumes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingProvides) DeepCopyInto(out *ServiceBindingProvides) {
	*out = *in
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = new(ServiceBindingAuth)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingProvides.
func (in *ServiceBindingProvides) DeepCopy() *ServiceBindingProvides {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingProvides)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusCondition) DeepCopyInto(out *StatusCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusCondition.
func (in *StatusCondition) DeepCopy() *StatusCondition {
	if in == nil {
		return nil
	}
	out := new(StatusCondition)
	in.DeepCopyInto(out)
	return out
}
