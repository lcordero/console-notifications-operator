// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleNotification) DeepCopyInto(out *ConsoleNotification) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleNotification.
func (in *ConsoleNotification) DeepCopy() *ConsoleNotification {
	if in == nil {
		return nil
	}
	out := new(ConsoleNotification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleNotification) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleNotificationList) DeepCopyInto(out *ConsoleNotificationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConsoleNotification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleNotificationList.
func (in *ConsoleNotificationList) DeepCopy() *ConsoleNotificationList {
	if in == nil {
		return nil
	}
	out := new(ConsoleNotificationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleNotificationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleNotificationSpec) DeepCopyInto(out *ConsoleNotificationSpec) {
	*out = *in
	out.Notification = in.Notification
	if in.NamespacesScope != nil {
		in, out := &in.NamespacesScope, &out.NamespacesScope
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleNotificationSpec.
func (in *ConsoleNotificationSpec) DeepCopy() *ConsoleNotificationSpec {
	if in == nil {
		return nil
	}
	out := new(ConsoleNotificationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleNotificationSpecLink) DeepCopyInto(out *ConsoleNotificationSpecLink) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleNotificationSpecLink.
func (in *ConsoleNotificationSpecLink) DeepCopy() *ConsoleNotificationSpecLink {
	if in == nil {
		return nil
	}
	out := new(ConsoleNotificationSpecLink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleNotificationSpecNotification) DeepCopyInto(out *ConsoleNotificationSpecNotification) {
	*out = *in
	out.Link = in.Link
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleNotificationSpecNotification.
func (in *ConsoleNotificationSpecNotification) DeepCopy() *ConsoleNotificationSpecNotification {
	if in == nil {
		return nil
	}
	out := new(ConsoleNotificationSpecNotification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleNotificationStatus) DeepCopyInto(out *ConsoleNotificationStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleNotificationStatus.
func (in *ConsoleNotificationStatus) DeepCopy() *ConsoleNotificationStatus {
	if in == nil {
		return nil
	}
	out := new(ConsoleNotificationStatus)
	in.DeepCopyInto(out)
	return out
}
