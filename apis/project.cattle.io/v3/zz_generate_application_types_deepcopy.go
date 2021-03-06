// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v3

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *Application) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationList) DeepCopyInto(out *ApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Application, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationList.
func (in *ApplicationList) DeepCopy() *ApplicationList {
	if in == nil {
		return nil
	}
	out := new(ApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppIngress) DeepCopyInto(out *AppIngress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppIngress.
func (in *AppIngress) DeepCopy() *AppIngress {
	if in == nil {
		return nil
	}
	out := new(AppIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppPort) DeepCopyInto(out *AppPort) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppPort.
func (in *AppPort) DeepCopy() *AppPort {
	if in == nil {
		return nil
	}
	out := new(AppPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Application) DeepCopyInto(out *Application) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Application.
func (in *Application) DeepCopy() *Application {
	if in == nil {
		return nil
	}
	out := new(Application)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpec) DeepCopyInto(out *ApplicationSpec) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]Component, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.OptTraits.DeepCopyInto(&out.OptTraits)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpec.
func (in *ApplicationSpec) DeepCopy() *ApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationStatus) DeepCopyInto(out *ApplicationStatus) {
	*out = *in
	if in.ComponentResource != nil {
		in, out := &in.ComponentResource, &out.ComponentResource
		*out = make(map[string]ComponentResources, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationStatus.
func (in *ApplicationStatus) DeepCopy() *ApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Autoscaling) DeepCopyInto(out *Autoscaling) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Autoscaling.
func (in *Autoscaling) DeepCopy() *Autoscaling {
	if in == nil {
		return nil
	}
	out := new(Autoscaling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CEnvVar) DeepCopyInto(out *CEnvVar) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CEnvVar.
func (in *CEnvVar) DeepCopy() *CEnvVar {
	if in == nil {
		return nil
	}
	out := new(CEnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CLabelSelectorRequirement) DeepCopyInto(out *CLabelSelectorRequirement) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CLabelSelectorRequirement.
func (in *CLabelSelectorRequirement) DeepCopy() *CLabelSelectorRequirement {
	if in == nil {
		return nil
	}
	out := new(CLabelSelectorRequirement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CLifecycle) DeepCopyInto(out *CLifecycle) {
	*out = *in
	if in.PostStart != nil {
		in, out := &in.PostStart, &out.PostStart
		*out = new(Handler)
		(*in).DeepCopyInto(*out)
	}
	if in.PreStop != nil {
		in, out := &in.PreStop, &out.PreStop
		*out = new(Handler)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CLifecycle.
func (in *CLifecycle) DeepCopy() *CLifecycle {
	if in == nil {
		return nil
	}
	out := new(CLifecycle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CNodeAffinity) DeepCopyInto(out *CNodeAffinity) {
	*out = *in
	if in.CLabelSelectorRequirement != nil {
		in, out := &in.CLabelSelectorRequirement, &out.CLabelSelectorRequirement
		*out = new(CLabelSelectorRequirement)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CNodeAffinity.
func (in *CNodeAffinity) DeepCopy() *CNodeAffinity {
	if in == nil {
		return nil
	}
	out := new(CNodeAffinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPodAffinity) DeepCopyInto(out *CPodAffinity) {
	*out = *in
	if in.CLabelSelectorRequirement != nil {
		in, out := &in.CLabelSelectorRequirement, &out.CLabelSelectorRequirement
		*out = new(CLabelSelectorRequirement)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPodAffinity.
func (in *CPodAffinity) DeepCopy() *CPodAffinity {
	if in == nil {
		return nil
	}
	out := new(CPodAffinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPodAntiAffinity) DeepCopyInto(out *CPodAntiAffinity) {
	*out = *in
	if in.CLabelSelectorRequirement != nil {
		in, out := &in.CLabelSelectorRequirement, &out.CLabelSelectorRequirement
		*out = new(CLabelSelectorRequirement)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPodAntiAffinity.
func (in *CPodAntiAffinity) DeepCopy() *CPodAntiAffinity {
	if in == nil {
		return nil
	}
	out := new(CPodAntiAffinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CResource) DeepCopyInto(out *CResource) {
	*out = *in
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]CVolume, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CResource.
func (in *CResource) DeepCopy() *CResource {
	if in == nil {
		return nil
	}
	out := new(CResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CVolume) DeepCopyInto(out *CVolume) {
	*out = *in
	out.Disk = in.Disk
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CVolume.
func (in *CVolume) DeepCopy() *CVolume {
	if in == nil {
		return nil
	}
	out := new(CVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CircuitBreaking) DeepCopyInto(out *CircuitBreaking) {
	*out = *in
	if in.ConnectionPool != nil {
		in, out := &in.ConnectionPool, &out.ConnectionPool
		*out = new(ConnectionPoolSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.OutlierDetection != nil {
		in, out := &in.OutlierDetection, &out.OutlierDetection
		*out = new(OutlierDetection)
		**out = **in
	}
	if in.PortLevelSettings != nil {
		in, out := &in.PortLevelSettings, &out.PortLevelSettings
		*out = make([]PortTrafficPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CircuitBreaking.
func (in *CircuitBreaking) DeepCopy() *CircuitBreaking {
	if in == nil {
		return nil
	}
	out := new(CircuitBreaking)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Component) DeepCopyInto(out *Component) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]Parameter, len(*in))
		copy(*out, *in)
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]ComponentContainer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ComponentTraits.DeepCopyInto(&out.ComponentTraits)
	if in.WorkloadSettings != nil {
		in, out := &in.WorkloadSettings, &out.WorkloadSettings
		*out = make([]WorkloadSetting, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component.
func (in *Component) DeepCopy() *Component {
	if in == nil {
		return nil
	}
	out := new(Component)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentContainer) DeepCopyInto(out *ComponentContainer) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]AppPort, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]CEnvVar, len(*in))
		copy(*out, *in)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(HealthProbe)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(HealthProbe)
		(*in).DeepCopyInto(*out)
	}
	if in.Lifecycle != nil {
		in, out := &in.Lifecycle, &out.Lifecycle
		*out = new(CLifecycle)
		(*in).DeepCopyInto(*out)
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]ConfigFile, len(*in))
		copy(*out, *in)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(SecurityContext)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentContainer.
func (in *ComponentContainer) DeepCopy() *ComponentContainer {
	if in == nil {
		return nil
	}
	out := new(ComponentContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentResources) DeepCopyInto(out *ComponentResources) {
	*out = *in
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentResources.
func (in *ComponentResources) DeepCopy() *ComponentResources {
	if in == nil {
		return nil
	}
	out := new(ComponentResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentTraits) DeepCopyInto(out *ComponentTraits) {
	*out = *in
	if in.CustomMetric != nil {
		in, out := &in.CustomMetric, &out.CustomMetric
		*out = new(CustomMetric)
		**out = **in
	}
	if in.SchedulePolicy != nil {
		in, out := &in.SchedulePolicy, &out.SchedulePolicy
		*out = new(SchedulePolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Autoscaling != nil {
		in, out := &in.Autoscaling, &out.Autoscaling
		*out = new(Autoscaling)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentTraits.
func (in *ComponentTraits) DeepCopy() *ComponentTraits {
	if in == nil {
		return nil
	}
	out := new(ComponentTraits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentTraitsForOpt) DeepCopyInto(out *ComponentTraitsForOpt) {
	*out = *in
	if in.LoadBalancer != nil {
		in, out := &in.LoadBalancer, &out.LoadBalancer
		*out = new(LoadBalancerSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.GrayRelease != nil {
		in, out := &in.GrayRelease, &out.GrayRelease
		*out = make(map[string]int, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ImagePullConfig != nil {
		in, out := &in.ImagePullConfig, &out.ImagePullConfig
		*out = new(ImagePullConfig)
		**out = **in
	}
	if in.VolumeMounter != nil {
		in, out := &in.VolumeMounter, &out.VolumeMounter
		*out = new(VolumeMounter)
		**out = **in
	}
	out.Ingress = in.Ingress
	if in.WhiteList != nil {
		in, out := &in.WhiteList, &out.WhiteList
		*out = new(WhiteList)
		(*in).DeepCopyInto(*out)
	}
	if in.Eject != nil {
		in, out := &in.Eject, &out.Eject
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Fusing != nil {
		in, out := &in.Fusing, &out.Fusing
		*out = new(Fusing)
		(*in).DeepCopyInto(*out)
	}
	if in.RateLimit != nil {
		in, out := &in.RateLimit, &out.RateLimit
		*out = new(RateLimit)
		(*in).DeepCopyInto(*out)
	}
	if in.CircuitBreaking != nil {
		in, out := &in.CircuitBreaking, &out.CircuitBreaking
		*out = new(CircuitBreaking)
		(*in).DeepCopyInto(*out)
	}
	if in.HTTPRetry != nil {
		in, out := &in.HTTPRetry, &out.HTTPRetry
		*out = new(HTTPRetry)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentTraitsForOpt.
func (in *ComponentTraitsForOpt) DeepCopy() *ComponentTraitsForOpt {
	if in == nil {
		return nil
	}
	out := new(ComponentTraitsForOpt)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigFile) DeepCopyInto(out *ConfigFile) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigFile.
func (in *ConfigFile) DeepCopy() *ConfigFile {
	if in == nil {
		return nil
	}
	out := new(ConfigFile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionPoolSettings) DeepCopyInto(out *ConnectionPoolSettings) {
	*out = *in
	if in.TCP != nil {
		in, out := &in.TCP, &out.TCP
		*out = new(TCPSettings)
		**out = **in
	}
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = new(HTTPSettings)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings.
func (in *ConnectionPoolSettings) DeepCopy() *ConnectionPoolSettings {
	if in == nil {
		return nil
	}
	out := new(ConnectionPoolSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsistentHashLB) DeepCopyInto(out *ConsistentHashLB) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsistentHashLB.
func (in *ConsistentHashLB) DeepCopy() *ConsistentHashLB {
	if in == nil {
		return nil
	}
	out := new(ConsistentHashLB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomMetric) DeepCopyInto(out *CustomMetric) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomMetric.
func (in *CustomMetric) DeepCopy() *CustomMetric {
	if in == nil {
		return nil
	}
	out := new(CustomMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Disk) DeepCopyInto(out *Disk) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Disk.
func (in *Disk) DeepCopy() *Disk {
	if in == nil {
		return nil
	}
	out := new(Disk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecAction) DeepCopyInto(out *ExecAction) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecAction.
func (in *ExecAction) DeepCopy() *ExecAction {
	if in == nil {
		return nil
	}
	out := new(ExecAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fusing) DeepCopyInto(out *Fusing) {
	*out = *in
	if in.PodList != nil {
		in, out := &in.PodList, &out.PodList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fusing.
func (in *Fusing) DeepCopy() *Fusing {
	if in == nil {
		return nil
	}
	out := new(Fusing)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPGetAction) DeepCopyInto(out *HTTPGetAction) {
	*out = *in
	if in.HTTPHeaders != nil {
		in, out := &in.HTTPHeaders, &out.HTTPHeaders
		*out = make([]HTTPHeader, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPGetAction.
func (in *HTTPGetAction) DeepCopy() *HTTPGetAction {
	if in == nil {
		return nil
	}
	out := new(HTTPGetAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPHeader) DeepCopyInto(out *HTTPHeader) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPHeader.
func (in *HTTPHeader) DeepCopy() *HTTPHeader {
	if in == nil {
		return nil
	}
	out := new(HTTPHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRetry) DeepCopyInto(out *HTTPRetry) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRetry.
func (in *HTTPRetry) DeepCopy() *HTTPRetry {
	if in == nil {
		return nil
	}
	out := new(HTTPRetry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSettings) DeepCopyInto(out *HTTPSettings) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSettings.
func (in *HTTPSettings) DeepCopy() *HTTPSettings {
	if in == nil {
		return nil
	}
	out := new(HTTPSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Handler) DeepCopyInto(out *Handler) {
	*out = *in
	if in.Exec != nil {
		in, out := &in.Exec, &out.Exec
		*out = new(ExecAction)
		(*in).DeepCopyInto(*out)
	}
	if in.HTTPGet != nil {
		in, out := &in.HTTPGet, &out.HTTPGet
		*out = new(HTTPGetAction)
		(*in).DeepCopyInto(*out)
	}
	if in.TCPSocket != nil {
		in, out := &in.TCPSocket, &out.TCPSocket
		*out = new(TCPSocketAction)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Handler.
func (in *Handler) DeepCopy() *Handler {
	if in == nil {
		return nil
	}
	out := new(Handler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthProbe) DeepCopyInto(out *HealthProbe) {
	*out = *in
	in.Handler.DeepCopyInto(&out.Handler)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthProbe.
func (in *HealthProbe) DeepCopy() *HealthProbe {
	if in == nil {
		return nil
	}
	out := new(HealthProbe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePullConfig) DeepCopyInto(out *ImagePullConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePullConfig.
func (in *ImagePullConfig) DeepCopy() *ImagePullConfig {
	if in == nil {
		return nil
	}
	out := new(ImagePullConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressLB) DeepCopyInto(out *IngressLB) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressLB.
func (in *IngressLB) DeepCopy() *IngressLB {
	if in == nil {
		return nil
	}
	out := new(IngressLB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerSettings) DeepCopyInto(out *LoadBalancerSettings) {
	*out = *in
	if in.ConsistentHash != nil {
		in, out := &in.ConsistentHash, &out.ConsistentHash
		*out = new(ConsistentHashLB)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings.
func (in *LoadBalancerSettings) DeepCopy() *LoadBalancerSettings {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutlierDetection) DeepCopyInto(out *OutlierDetection) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutlierDetection.
func (in *OutlierDetection) DeepCopy() *OutlierDetection {
	if in == nil {
		return nil
	}
	out := new(OutlierDetection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Override) DeepCopyInto(out *Override) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Override.
func (in *Override) DeepCopy() *Override {
	if in == nil {
		return nil
	}
	out := new(Override)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parameter) DeepCopyInto(out *Parameter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parameter.
func (in *Parameter) DeepCopy() *Parameter {
	if in == nil {
		return nil
	}
	out := new(Parameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortSelector) DeepCopyInto(out *PortSelector) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortSelector.
func (in *PortSelector) DeepCopy() *PortSelector {
	if in == nil {
		return nil
	}
	out := new(PortSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortTrafficPolicy) DeepCopyInto(out *PortTrafficPolicy) {
	*out = *in
	out.Port = in.Port
	in.LoadBalancer.DeepCopyInto(&out.LoadBalancer)
	in.ConnectionPool.DeepCopyInto(&out.ConnectionPool)
	out.OutlierDetection = in.OutlierDetection
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortTrafficPolicy.
func (in *PortTrafficPolicy) DeepCopy() *PortTrafficPolicy {
	if in == nil {
		return nil
	}
	out := new(PortTrafficPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimit) DeepCopyInto(out *RateLimit) {
	*out = *in
	if in.Overrides != nil {
		in, out := &in.Overrides, &out.Overrides
		*out = make([]Override, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimit.
func (in *RateLimit) DeepCopy() *RateLimit {
	if in == nil {
		return nil
	}
	out := new(RateLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulePolicy) DeepCopyInto(out *SchedulePolicy) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeAffinity != nil {
		in, out := &in.NodeAffinity, &out.NodeAffinity
		*out = new(CNodeAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.PodAffinity != nil {
		in, out := &in.PodAffinity, &out.PodAffinity
		*out = new(CPodAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.PodAntiAffinity != nil {
		in, out := &in.PodAntiAffinity, &out.PodAntiAffinity
		*out = new(CPodAntiAffinity)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulePolicy.
func (in *SchedulePolicy) DeepCopy() *SchedulePolicy {
	if in == nil {
		return nil
	}
	out := new(SchedulePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityContext) DeepCopyInto(out *SecurityContext) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityContext.
func (in *SecurityContext) DeepCopy() *SecurityContext {
	if in == nil {
		return nil
	}
	out := new(SecurityContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPSettings) DeepCopyInto(out *TCPSettings) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPSettings.
func (in *TCPSettings) DeepCopy() *TCPSettings {
	if in == nil {
		return nil
	}
	out := new(TCPSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPSocketAction) DeepCopyInto(out *TCPSocketAction) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPSocketAction.
func (in *TCPSocketAction) DeepCopy() *TCPSocketAction {
	if in == nil {
		return nil
	}
	out := new(TCPSocketAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeMounter) DeepCopyInto(out *VolumeMounter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeMounter.
func (in *VolumeMounter) DeepCopy() *VolumeMounter {
	if in == nil {
		return nil
	}
	out := new(VolumeMounter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WhiteList) DeepCopyInto(out *WhiteList) {
	*out = *in
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WhiteList.
func (in *WhiteList) DeepCopy() *WhiteList {
	if in == nil {
		return nil
	}
	out := new(WhiteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadSetting) DeepCopyInto(out *WorkloadSetting) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadSetting.
func (in *WorkloadSetting) DeepCopy() *WorkloadSetting {
	if in == nil {
		return nil
	}
	out := new(WorkloadSetting)
	in.DeepCopyInto(out)
	return out
}
