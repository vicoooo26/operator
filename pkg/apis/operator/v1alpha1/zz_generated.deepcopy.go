//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2020 The Tekton Authors

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

package v1alpha1

import (
	manifestival "github.com/manifestival/manifestival"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addon) DeepCopyInto(out *Addon) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]Param, len(*in))
		copy(*out, *in)
	}
	if in.EnablePAC != nil {
		in, out := &in.EnablePAC, &out.EnablePAC
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addon.
func (in *Addon) DeepCopy() *Addon {
	if in == nil {
		return nil
	}
	out := new(Addon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiSpec) DeepCopyInto(out *ApiSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiSpec.
func (in *ApiSpec) DeepCopy() *ApiSpec {
	if in == nil {
		return nil
	}
	out := new(ApiSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Chain) DeepCopyInto(out *Chain) {
	*out = *in
	if in.StorageOCIRepositoryInsecure != nil {
		in, out := &in.StorageOCIRepositoryInsecure, &out.StorageOCIRepositoryInsecure
		*out = new(bool)
		**out = **in
	}
	if in.X509SignerFulcioEnabled != nil {
		in, out := &in.X509SignerFulcioEnabled, &out.X509SignerFulcioEnabled
		*out = new(bool)
		**out = **in
	}
	if in.TransparencyConfigEnabled != nil {
		in, out := &in.TransparencyConfigEnabled, &out.TransparencyConfigEnabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Chain.
func (in *Chain) DeepCopy() *Chain {
	if in == nil {
		return nil
	}
	out := new(Chain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonSpec) DeepCopyInto(out *CommonSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonSpec.
func (in *CommonSpec) DeepCopy() *CommonSpec {
	if in == nil {
		return nil
	}
	out := new(CommonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dashboard) DeepCopyInto(out *Dashboard) {
	*out = *in
	out.DashboardProperties = in.DashboardProperties
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dashboard.
func (in *Dashboard) DeepCopy() *Dashboard {
	if in == nil {
		return nil
	}
	out := new(Dashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardProperties) DeepCopyInto(out *DashboardProperties) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardProperties.
func (in *DashboardProperties) DeepCopy() *DashboardProperties {
	if in == nil {
		return nil
	}
	out := new(DashboardProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbSpec) DeepCopyInto(out *DbSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbSpec.
func (in *DbSpec) DeepCopy() *DbSpec {
	if in == nil {
		return nil
	}
	out := new(DbSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hub) DeepCopyInto(out *Hub) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]Param, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hub.
func (in *Hub) DeepCopy() *Hub {
	if in == nil {
		return nil
	}
	out := new(Hub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionalPipelineProperties) DeepCopyInto(out *OptionalPipelineProperties) {
	*out = *in
	if in.DefaultTimeoutMinutes != nil {
		in, out := &in.DefaultTimeoutMinutes, &out.DefaultTimeoutMinutes
		*out = new(uint)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionalPipelineProperties.
func (in *OptionalPipelineProperties) DeepCopy() *OptionalPipelineProperties {
	if in == nil {
		return nil
	}
	out := new(OptionalPipelineProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionalTriggersProperties) DeepCopyInto(out *OptionalTriggersProperties) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionalTriggersProperties.
func (in *OptionalTriggersProperties) DeepCopy() *OptionalTriggersProperties {
	if in == nil {
		return nil
	}
	out := new(OptionalTriggersProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Param) DeepCopyInto(out *Param) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Param.
func (in *Param) DeepCopy() *Param {
	if in == nil {
		return nil
	}
	out := new(Param)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParamValue) DeepCopyInto(out *ParamValue) {
	*out = *in
	if in.Possible != nil {
		in, out := &in.Possible, &out.Possible
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParamValue.
func (in *ParamValue) DeepCopy() *ParamValue {
	if in == nil {
		return nil
	}
	out := new(ParamValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pipeline) DeepCopyInto(out *Pipeline) {
	*out = *in
	in.PipelineProperties.DeepCopyInto(&out.PipelineProperties)
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]Param, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pipeline.
func (in *Pipeline) DeepCopy() *Pipeline {
	if in == nil {
		return nil
	}
	out := new(Pipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineMetricsProperties) DeepCopyInto(out *PipelineMetricsProperties) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineMetricsProperties.
func (in *PipelineMetricsProperties) DeepCopy() *PipelineMetricsProperties {
	if in == nil {
		return nil
	}
	out := new(PipelineMetricsProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineProperties) DeepCopyInto(out *PipelineProperties) {
	*out = *in
	if in.DisableAffinityAssistant != nil {
		in, out := &in.DisableAffinityAssistant, &out.DisableAffinityAssistant
		*out = new(bool)
		**out = **in
	}
	if in.DisableHomeEnvOverwrite != nil {
		in, out := &in.DisableHomeEnvOverwrite, &out.DisableHomeEnvOverwrite
		*out = new(bool)
		**out = **in
	}
	if in.DisableWorkingDirectoryOverwrite != nil {
		in, out := &in.DisableWorkingDirectoryOverwrite, &out.DisableWorkingDirectoryOverwrite
		*out = new(bool)
		**out = **in
	}
	if in.DisableCredsInit != nil {
		in, out := &in.DisableCredsInit, &out.DisableCredsInit
		*out = new(bool)
		**out = **in
	}
	if in.RunningInEnvironmentWithInjectedSidecars != nil {
		in, out := &in.RunningInEnvironmentWithInjectedSidecars, &out.RunningInEnvironmentWithInjectedSidecars
		*out = new(bool)
		**out = **in
	}
	if in.RequireGitSshSecretKnownHosts != nil {
		in, out := &in.RequireGitSshSecretKnownHosts, &out.RequireGitSshSecretKnownHosts
		*out = new(bool)
		**out = **in
	}
	if in.EnableTektonOciBundles != nil {
		in, out := &in.EnableTektonOciBundles, &out.EnableTektonOciBundles
		*out = new(bool)
		**out = **in
	}
	if in.EnableCustomTasks != nil {
		in, out := &in.EnableCustomTasks, &out.EnableCustomTasks
		*out = new(bool)
		**out = **in
	}
	if in.ScopeWhenExpressionsToTask != nil {
		in, out := &in.ScopeWhenExpressionsToTask, &out.ScopeWhenExpressionsToTask
		*out = new(bool)
		**out = **in
	}
	out.PipelineMetricsProperties = in.PipelineMetricsProperties
	in.OptionalPipelineProperties.DeepCopyInto(&out.OptionalPipelineProperties)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineProperties.
func (in *PipelineProperties) DeepCopy() *PipelineProperties {
	if in == nil {
		return nil
	}
	out := new(PipelineProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Prune) DeepCopyInto(out *Prune) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Keep != nil {
		in, out := &in.Keep, &out.Keep
		*out = new(uint)
		**out = **in
	}
	if in.KeepSince != nil {
		in, out := &in.KeepSince, &out.KeepSince
		*out = new(uint)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Prune.
func (in *Prune) DeepCopy() *Prune {
	if in == nil {
		return nil
	}
	out := new(Prune)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonAddon) DeepCopyInto(out *TektonAddon) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonAddon.
func (in *TektonAddon) DeepCopy() *TektonAddon {
	if in == nil {
		return nil
	}
	out := new(TektonAddon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TektonAddon) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonAddonList) DeepCopyInto(out *TektonAddonList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TektonAddon, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonAddonList.
func (in *TektonAddonList) DeepCopy() *TektonAddonList {
	if in == nil {
		return nil
	}
	out := new(TektonAddonList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TektonAddonList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonAddonSpec) DeepCopyInto(out *TektonAddonSpec) {
	*out = *in
	out.CommonSpec = in.CommonSpec
	in.Addon.DeepCopyInto(&out.Addon)
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonAddonSpec.
func (in *TektonAddonSpec) DeepCopy() *TektonAddonSpec {
	if in == nil {
		return nil
	}
	out := new(TektonAddonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonAddonStatus) DeepCopyInto(out *TektonAddonStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.AddonsInstallerSet != nil {
		in, out := &in.AddonsInstallerSet, &out.AddonsInstallerSet
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonAddonStatus.
func (in *TektonAddonStatus) DeepCopy() *TektonAddonStatus {
	if in == nil {
		return nil
	}
	out := new(TektonAddonStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonChain) DeepCopyInto(out *TektonChain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonChain.
func (in *TektonChain) DeepCopy() *TektonChain {
	if in == nil {
		return nil
	}
	out := new(TektonChain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TektonChain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonChainList) DeepCopyInto(out *TektonChainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TektonChain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonChainList.
func (in *TektonChainList) DeepCopy() *TektonChainList {
	if in == nil {
		return nil
	}
	out := new(TektonChainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TektonChainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonChainSpec) DeepCopyInto(out *TektonChainSpec) {
	*out = *in
	out.CommonSpec = in.CommonSpec
	in.Chain.DeepCopyInto(&out.Chain)
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonChainSpec.
func (in *TektonChainSpec) DeepCopy() *TektonChainSpec {
	if in == nil {
		return nil
	}
	out := new(TektonChainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonChainStatus) DeepCopyInto(out *TektonChainStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonChainStatus.
func (in *TektonChainStatus) DeepCopy() *TektonChainStatus {
	if in == nil {
		return nil
	}
	out := new(TektonChainStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonConfig) DeepCopyInto(out *TektonConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonConfig.
func (in *TektonConfig) DeepCopy() *TektonConfig {
	if in == nil {
		return nil
	}
	out := new(TektonConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TektonConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonConfigList) DeepCopyInto(out *TektonConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TektonConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonConfigList.
func (in *TektonConfigList) DeepCopy() *TektonConfigList {
	if in == nil {
		return nil
	}
	out := new(TektonConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TektonConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonConfigSpec) DeepCopyInto(out *TektonConfigSpec) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	in.Pruner.DeepCopyInto(&out.Pruner)
	out.CommonSpec = in.CommonSpec
	in.Addon.DeepCopyInto(&out.Addon)
	in.Hub.DeepCopyInto(&out.Hub)
	in.Pipeline.DeepCopyInto(&out.Pipeline)
	out.Trigger = in.Trigger
	out.Dashboard = in.Dashboard
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]Param, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonConfigSpec.
func (in *TektonConfigSpec) DeepCopy() *TektonConfigSpec {
	if in == nil {
		return nil
	}
	out := new(TektonConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonConfigStatus) DeepCopyInto(out *TektonConfigStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.TektonInstallerSet != nil {
		in, out := &in.TektonInstallerSet, &out.TektonInstallerSet
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonConfigStatus.
func (in *TektonConfigStatus) DeepCopy() *TektonConfigStatus {
	if in == nil {
		return nil
	}
	out := new(TektonConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonDashboard) DeepCopyInto(out *TektonDashboard) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonDashboard.
func (in *TektonDashboard) DeepCopy() *TektonDashboard {
	if in == nil {
		return nil
	}
	out := new(TektonDashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TektonDashboard) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonDashboardList) DeepCopyInto(out *TektonDashboardList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TektonDashboard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonDashboardList.
func (in *TektonDashboardList) DeepCopy() *TektonDashboardList {
	if in == nil {
		return nil
	}
	out := new(TektonDashboardList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TektonDashboardList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonDashboardSpec) DeepCopyInto(out *TektonDashboardSpec) {
	*out = *in
	out.CommonSpec = in.CommonSpec
	out.DashboardProperties = in.DashboardProperties
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonDashboardSpec.
func (in *TektonDashboardSpec) DeepCopy() *TektonDashboardSpec {
	if in == nil {
		return nil
	}
	out := new(TektonDashboardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonDashboardStatus) DeepCopyInto(out *TektonDashboardStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonDashboardStatus.
func (in *TektonDashboardStatus) DeepCopy() *TektonDashboardStatus {
	if in == nil {
		return nil
	}
	out := new(TektonDashboardStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonHub) DeepCopyInto(out *TektonHub) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonHub.
func (in *TektonHub) DeepCopy() *TektonHub {
	if in == nil {
		return nil
	}
	out := new(TektonHub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TektonHub) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonHubList) DeepCopyInto(out *TektonHubList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TektonHub, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonHubList.
func (in *TektonHubList) DeepCopy() *TektonHubList {
	if in == nil {
		return nil
	}
	out := new(TektonHubList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TektonHubList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonHubSpec) DeepCopyInto(out *TektonHubSpec) {
	*out = *in
	out.CommonSpec = in.CommonSpec
	in.Hub.DeepCopyInto(&out.Hub)
	out.Db = in.Db
	out.Api = in.Api
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonHubSpec.
func (in *TektonHubSpec) DeepCopy() *TektonHubSpec {
	if in == nil {
		return nil
	}
	out := new(TektonHubSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonHubStatus) DeepCopyInto(out *TektonHubStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.Manifests != nil {
		in, out := &in.Manifests, &out.Manifests
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HubInstallerSet != nil {
		in, out := &in.HubInstallerSet, &out.HubInstallerSet
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonHubStatus.
func (in *TektonHubStatus) DeepCopy() *TektonHubStatus {
	if in == nil {
		return nil
	}
	out := new(TektonHubStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonInstallerSet) DeepCopyInto(out *TektonInstallerSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonInstallerSet.
func (in *TektonInstallerSet) DeepCopy() *TektonInstallerSet {
	if in == nil {
		return nil
	}
	out := new(TektonInstallerSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TektonInstallerSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonInstallerSetList) DeepCopyInto(out *TektonInstallerSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TektonInstallerSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonInstallerSetList.
func (in *TektonInstallerSetList) DeepCopy() *TektonInstallerSetList {
	if in == nil {
		return nil
	}
	out := new(TektonInstallerSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TektonInstallerSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonInstallerSetSpec) DeepCopyInto(out *TektonInstallerSetSpec) {
	*out = *in
	if in.Manifests != nil {
		in, out := &in.Manifests, &out.Manifests
		*out = make(manifestival.Slice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonInstallerSetSpec.
func (in *TektonInstallerSetSpec) DeepCopy() *TektonInstallerSetSpec {
	if in == nil {
		return nil
	}
	out := new(TektonInstallerSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonInstallerSetStatus) DeepCopyInto(out *TektonInstallerSetStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonInstallerSetStatus.
func (in *TektonInstallerSetStatus) DeepCopy() *TektonInstallerSetStatus {
	if in == nil {
		return nil
	}
	out := new(TektonInstallerSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonPipeline) DeepCopyInto(out *TektonPipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonPipeline.
func (in *TektonPipeline) DeepCopy() *TektonPipeline {
	if in == nil {
		return nil
	}
	out := new(TektonPipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TektonPipeline) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonPipelineList) DeepCopyInto(out *TektonPipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TektonPipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonPipelineList.
func (in *TektonPipelineList) DeepCopy() *TektonPipelineList {
	if in == nil {
		return nil
	}
	out := new(TektonPipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TektonPipelineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonPipelineSpec) DeepCopyInto(out *TektonPipelineSpec) {
	*out = *in
	out.CommonSpec = in.CommonSpec
	in.Pipeline.DeepCopyInto(&out.Pipeline)
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonPipelineSpec.
func (in *TektonPipelineSpec) DeepCopy() *TektonPipelineSpec {
	if in == nil {
		return nil
	}
	out := new(TektonPipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonPipelineStatus) DeepCopyInto(out *TektonPipelineStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.ExtentionInstallerSets != nil {
		in, out := &in.ExtentionInstallerSets, &out.ExtentionInstallerSets
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonPipelineStatus.
func (in *TektonPipelineStatus) DeepCopy() *TektonPipelineStatus {
	if in == nil {
		return nil
	}
	out := new(TektonPipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonResult) DeepCopyInto(out *TektonResult) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonResult.
func (in *TektonResult) DeepCopy() *TektonResult {
	if in == nil {
		return nil
	}
	out := new(TektonResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TektonResult) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonResultList) DeepCopyInto(out *TektonResultList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TektonResult, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonResultList.
func (in *TektonResultList) DeepCopy() *TektonResultList {
	if in == nil {
		return nil
	}
	out := new(TektonResultList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TektonResultList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonResultSpec) DeepCopyInto(out *TektonResultSpec) {
	*out = *in
	out.CommonSpec = in.CommonSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonResultSpec.
func (in *TektonResultSpec) DeepCopy() *TektonResultSpec {
	if in == nil {
		return nil
	}
	out := new(TektonResultSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonResultStatus) DeepCopyInto(out *TektonResultStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonResultStatus.
func (in *TektonResultStatus) DeepCopy() *TektonResultStatus {
	if in == nil {
		return nil
	}
	out := new(TektonResultStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonTrigger) DeepCopyInto(out *TektonTrigger) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonTrigger.
func (in *TektonTrigger) DeepCopy() *TektonTrigger {
	if in == nil {
		return nil
	}
	out := new(TektonTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TektonTrigger) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonTriggerList) DeepCopyInto(out *TektonTriggerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TektonTrigger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonTriggerList.
func (in *TektonTriggerList) DeepCopy() *TektonTriggerList {
	if in == nil {
		return nil
	}
	out := new(TektonTriggerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TektonTriggerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonTriggerSpec) DeepCopyInto(out *TektonTriggerSpec) {
	*out = *in
	out.CommonSpec = in.CommonSpec
	out.Trigger = in.Trigger
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonTriggerSpec.
func (in *TektonTriggerSpec) DeepCopy() *TektonTriggerSpec {
	if in == nil {
		return nil
	}
	out := new(TektonTriggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonTriggerStatus) DeepCopyInto(out *TektonTriggerStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonTriggerStatus.
func (in *TektonTriggerStatus) DeepCopy() *TektonTriggerStatus {
	if in == nil {
		return nil
	}
	out := new(TektonTriggerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trigger) DeepCopyInto(out *Trigger) {
	*out = *in
	out.TriggersProperties = in.TriggersProperties
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trigger.
func (in *Trigger) DeepCopy() *Trigger {
	if in == nil {
		return nil
	}
	out := new(Trigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggersProperties) DeepCopyInto(out *TriggersProperties) {
	*out = *in
	out.OptionalTriggersProperties = in.OptionalTriggersProperties
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggersProperties.
func (in *TriggersProperties) DeepCopy() *TriggersProperties {
	if in == nil {
		return nil
	}
	out := new(TriggersProperties)
	in.DeepCopyInto(out)
	return out
}
