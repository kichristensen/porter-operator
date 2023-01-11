//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentAction) DeepCopyInto(out *AgentAction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentAction.
func (in *AgentAction) DeepCopy() *AgentAction {
	if in == nil {
		return nil
	}
	out := new(AgentAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AgentAction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentActionList) DeepCopyInto(out *AgentActionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AgentAction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentActionList.
func (in *AgentActionList) DeepCopy() *AgentActionList {
	if in == nil {
		return nil
	}
	out := new(AgentActionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AgentActionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentActionSpec) DeepCopyInto(out *AgentActionSpec) {
	*out = *in
	if in.AgentConfig != nil {
		in, out := &in.AgentConfig, &out.AgentConfig
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.PorterConfig != nil {
		in, out := &in.PorterConfig, &out.PorterConfig
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
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
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make(map[string][]byte, len(*in))
		for key, val := range *in {
			var outVal []byte
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]byte, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]corev1.EnvFromSource, len(*in))
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
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentActionSpec.
func (in *AgentActionSpec) DeepCopy() *AgentActionSpec {
	if in == nil {
		return nil
	}
	out := new(AgentActionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentActionStatus) DeepCopyInto(out *AgentActionStatus) {
	*out = *in
	if in.Job != nil {
		in, out := &in.Job, &out.Job
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentActionStatus.
func (in *AgentActionStatus) DeepCopy() *AgentActionStatus {
	if in == nil {
		return nil
	}
	out := new(AgentActionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentConfig) DeepCopyInto(out *AgentConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentConfig.
func (in *AgentConfig) DeepCopy() *AgentConfig {
	if in == nil {
		return nil
	}
	out := new(AgentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AgentConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentConfigAdapter) DeepCopyInto(out *AgentConfigAdapter) {
	*out = *in
	in.AgentConfig.DeepCopyInto(&out.AgentConfig)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentConfigAdapter.
func (in *AgentConfigAdapter) DeepCopy() *AgentConfigAdapter {
	if in == nil {
		return nil
	}
	out := new(AgentConfigAdapter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentConfigList) DeepCopyInto(out *AgentConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AgentConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentConfigList.
func (in *AgentConfigList) DeepCopy() *AgentConfigList {
	if in == nil {
		return nil
	}
	out := new(AgentConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AgentConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentConfigSpec) DeepCopyInto(out *AgentConfigSpec) {
	*out = *in
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make(map[string]Plugin, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentConfigSpec.
func (in *AgentConfigSpec) DeepCopy() *AgentConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AgentConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentConfigSpecAdapter) DeepCopyInto(out *AgentConfigSpecAdapter) {
	*out = *in
	in.original.DeepCopyInto(&out.original)
	in.Plugins.DeepCopyInto(&out.Plugins)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentConfigSpecAdapter.
func (in *AgentConfigSpecAdapter) DeepCopy() *AgentConfigSpecAdapter {
	if in == nil {
		return nil
	}
	out := new(AgentConfigSpecAdapter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentConfigStatus) DeepCopyInto(out *AgentConfigStatus) {
	*out = *in
	in.PorterResourceStatus.DeepCopyInto(&out.PorterResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentConfigStatus.
func (in *AgentConfigStatus) DeepCopy() *AgentConfigStatus {
	if in == nil {
		return nil
	}
	out := new(AgentConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Credential) DeepCopyInto(out *Credential) {
	*out = *in
	out.Source = in.Source
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Credential.
func (in *Credential) DeepCopy() *Credential {
	if in == nil {
		return nil
	}
	out := new(Credential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialSet) DeepCopyInto(out *CredentialSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialSet.
func (in *CredentialSet) DeepCopy() *CredentialSet {
	if in == nil {
		return nil
	}
	out := new(CredentialSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CredentialSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialSetList) DeepCopyInto(out *CredentialSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CredentialSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialSetList.
func (in *CredentialSetList) DeepCopy() *CredentialSetList {
	if in == nil {
		return nil
	}
	out := new(CredentialSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CredentialSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialSetSpec) DeepCopyInto(out *CredentialSetSpec) {
	*out = *in
	if in.AgentConfig != nil {
		in, out := &in.AgentConfig, &out.AgentConfig
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.PorterConfig != nil {
		in, out := &in.PorterConfig, &out.PorterConfig
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]Credential, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialSetSpec.
func (in *CredentialSetSpec) DeepCopy() *CredentialSetSpec {
	if in == nil {
		return nil
	}
	out := new(CredentialSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialSetStatus) DeepCopyInto(out *CredentialSetStatus) {
	*out = *in
	in.PorterResourceStatus.DeepCopyInto(&out.PorterResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialSetStatus.
func (in *CredentialSetStatus) DeepCopy() *CredentialSetStatus {
	if in == nil {
		return nil
	}
	out := new(CredentialSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialSource) DeepCopyInto(out *CredentialSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialSource.
func (in *CredentialSource) DeepCopy() *CredentialSource {
	if in == nil {
		return nil
	}
	out := new(CredentialSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Installation) DeepCopyInto(out *Installation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Installation.
func (in *Installation) DeepCopy() *Installation {
	if in == nil {
		return nil
	}
	out := new(Installation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Installation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationList) DeepCopyInto(out *InstallationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Installation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationList.
func (in *InstallationList) DeepCopy() *InstallationList {
	if in == nil {
		return nil
	}
	out := new(InstallationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstallationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationSpec) DeepCopyInto(out *InstallationSpec) {
	*out = *in
	if in.AgentConfig != nil {
		in, out := &in.AgentConfig, &out.AgentConfig
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.PorterConfig != nil {
		in, out := &in.PorterConfig, &out.PorterConfig
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	out.Bundle = in.Bundle
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Parameters.DeepCopyInto(&out.Parameters)
	if in.CredentialSets != nil {
		in, out := &in.CredentialSets, &out.CredentialSets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ParameterSets != nil {
		in, out := &in.ParameterSets, &out.ParameterSets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationSpec.
func (in *InstallationSpec) DeepCopy() *InstallationSpec {
	if in == nil {
		return nil
	}
	out := new(InstallationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationStatus) DeepCopyInto(out *InstallationStatus) {
	*out = *in
	in.PorterResourceStatus.DeepCopyInto(&out.PorterResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationStatus.
func (in *InstallationStatus) DeepCopy() *InstallationStatus {
	if in == nil {
		return nil
	}
	out := new(InstallationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIReferenceParts) DeepCopyInto(out *OCIReferenceParts) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIReferenceParts.
func (in *OCIReferenceParts) DeepCopy() *OCIReferenceParts {
	if in == nil {
		return nil
	}
	out := new(OCIReferenceParts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parameter) DeepCopyInto(out *Parameter) {
	*out = *in
	out.Source = in.Source
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
func (in *ParameterSet) DeepCopyInto(out *ParameterSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterSet.
func (in *ParameterSet) DeepCopy() *ParameterSet {
	if in == nil {
		return nil
	}
	out := new(ParameterSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ParameterSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterSetList) DeepCopyInto(out *ParameterSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ParameterSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterSetList.
func (in *ParameterSetList) DeepCopy() *ParameterSetList {
	if in == nil {
		return nil
	}
	out := new(ParameterSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ParameterSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterSetSpec) DeepCopyInto(out *ParameterSetSpec) {
	*out = *in
	if in.AgentConfig != nil {
		in, out := &in.AgentConfig, &out.AgentConfig
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.PorterConfig != nil {
		in, out := &in.PorterConfig, &out.PorterConfig
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]Parameter, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterSetSpec.
func (in *ParameterSetSpec) DeepCopy() *ParameterSetSpec {
	if in == nil {
		return nil
	}
	out := new(ParameterSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterSetStatus) DeepCopyInto(out *ParameterSetStatus) {
	*out = *in
	in.PorterResourceStatus.DeepCopyInto(&out.PorterResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterSetStatus.
func (in *ParameterSetStatus) DeepCopy() *ParameterSetStatus {
	if in == nil {
		return nil
	}
	out := new(ParameterSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterSource) DeepCopyInto(out *ParameterSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterSource.
func (in *ParameterSource) DeepCopy() *ParameterSource {
	if in == nil {
		return nil
	}
	out := new(ParameterSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugin) DeepCopyInto(out *Plugin) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugin.
func (in *Plugin) DeepCopy() *Plugin {
	if in == nil {
		return nil
	}
	out := new(Plugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginConfig) DeepCopyInto(out *PluginConfig) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginConfig.
func (in *PluginConfig) DeepCopy() *PluginConfig {
	if in == nil {
		return nil
	}
	out := new(PluginConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginsConfigList) DeepCopyInto(out *PluginsConfigList) {
	*out = *in
	if in.data != nil {
		in, out := &in.data, &out.data
		*out = make(map[string]Plugin, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.keys != nil {
		in, out := &in.keys, &out.keys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginsConfigList.
func (in *PluginsConfigList) DeepCopy() *PluginsConfigList {
	if in == nil {
		return nil
	}
	out := new(PluginsConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PorterConfig) DeepCopyInto(out *PorterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PorterConfig.
func (in *PorterConfig) DeepCopy() *PorterConfig {
	if in == nil {
		return nil
	}
	out := new(PorterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PorterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PorterConfigList) DeepCopyInto(out *PorterConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PorterConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PorterConfigList.
func (in *PorterConfigList) DeepCopy() *PorterConfigList {
	if in == nil {
		return nil
	}
	out := new(PorterConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PorterConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PorterConfigSpec) DeepCopyInto(out *PorterConfigSpec) {
	*out = *in
	if in.Verbosity != nil {
		in, out := &in.Verbosity, &out.Verbosity
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Experimental != nil {
		in, out := &in.Experimental, &out.Experimental
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BuildDriver != nil {
		in, out := &in.BuildDriver, &out.BuildDriver
		*out = new(string)
		**out = **in
	}
	if in.DefaultStorage != nil {
		in, out := &in.DefaultStorage, &out.DefaultStorage
		*out = new(string)
		**out = **in
	}
	if in.DefaultSecrets != nil {
		in, out := &in.DefaultSecrets, &out.DefaultSecrets
		*out = new(string)
		**out = **in
	}
	if in.DefaultStoragePlugin != nil {
		in, out := &in.DefaultStoragePlugin, &out.DefaultStoragePlugin
		*out = new(string)
		**out = **in
	}
	if in.DefaultSecretsPlugin != nil {
		in, out := &in.DefaultSecretsPlugin, &out.DefaultSecretsPlugin
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = make([]StorageConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]SecretsConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PorterConfigSpec.
func (in *PorterConfigSpec) DeepCopy() *PorterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(PorterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PorterResourceStatus) DeepCopyInto(out *PorterResourceStatus) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PorterResourceStatus.
func (in *PorterResourceStatus) DeepCopy() *PorterResourceStatus {
	if in == nil {
		return nil
	}
	out := new(PorterResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretsConfig) DeepCopyInto(out *SecretsConfig) {
	*out = *in
	in.PluginConfig.DeepCopyInto(&out.PluginConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretsConfig.
func (in *SecretsConfig) DeepCopy() *SecretsConfig {
	if in == nil {
		return nil
	}
	out := new(SecretsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageConfig) DeepCopyInto(out *StorageConfig) {
	*out = *in
	in.PluginConfig.DeepCopyInto(&out.PluginConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageConfig.
func (in *StorageConfig) DeepCopy() *StorageConfig {
	if in == nil {
		return nil
	}
	out := new(StorageConfig)
	in.DeepCopyInto(out)
	return out
}
