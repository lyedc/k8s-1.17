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

// Code generated by conversion-gen. DO NOT EDIT.

package v2beta1

import (
	unsafe "unsafe"

	v2beta1 "k8s.io/api/autoscaling/v2beta1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling"
	core "k8s.io/kubernetes/pkg/apis/core"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v2beta1.CrossVersionObjectReference)(nil), (*autoscaling.CrossVersionObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(a.(*v2beta1.CrossVersionObjectReference), b.(*autoscaling.CrossVersionObjectReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*autoscaling.CrossVersionObjectReference)(nil), (*v2beta1.CrossVersionObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_autoscaling_CrossVersionObjectReference_To_v2beta1_CrossVersionObjectReference(a.(*autoscaling.CrossVersionObjectReference), b.(*v2beta1.CrossVersionObjectReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2beta1.HorizontalPodAutoscalerCondition)(nil), (*autoscaling.HorizontalPodAutoscalerCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_HorizontalPodAutoscalerCondition_To_autoscaling_HorizontalPodAutoscalerCondition(a.(*v2beta1.HorizontalPodAutoscalerCondition), b.(*autoscaling.HorizontalPodAutoscalerCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*autoscaling.HorizontalPodAutoscalerCondition)(nil), (*v2beta1.HorizontalPodAutoscalerCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_autoscaling_HorizontalPodAutoscalerCondition_To_v2beta1_HorizontalPodAutoscalerCondition(a.(*autoscaling.HorizontalPodAutoscalerCondition), b.(*v2beta1.HorizontalPodAutoscalerCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2beta1.HorizontalPodAutoscalerList)(nil), (*autoscaling.HorizontalPodAutoscalerList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList(a.(*v2beta1.HorizontalPodAutoscalerList), b.(*autoscaling.HorizontalPodAutoscalerList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*autoscaling.HorizontalPodAutoscalerList)(nil), (*v2beta1.HorizontalPodAutoscalerList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_autoscaling_HorizontalPodAutoscalerList_To_v2beta1_HorizontalPodAutoscalerList(a.(*autoscaling.HorizontalPodAutoscalerList), b.(*v2beta1.HorizontalPodAutoscalerList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2beta1.HorizontalPodAutoscalerSpec)(nil), (*autoscaling.HorizontalPodAutoscalerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec(a.(*v2beta1.HorizontalPodAutoscalerSpec), b.(*autoscaling.HorizontalPodAutoscalerSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2beta1.HorizontalPodAutoscalerStatus)(nil), (*autoscaling.HorizontalPodAutoscalerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus(a.(*v2beta1.HorizontalPodAutoscalerStatus), b.(*autoscaling.HorizontalPodAutoscalerStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*autoscaling.HorizontalPodAutoscalerStatus)(nil), (*v2beta1.HorizontalPodAutoscalerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_autoscaling_HorizontalPodAutoscalerStatus_To_v2beta1_HorizontalPodAutoscalerStatus(a.(*autoscaling.HorizontalPodAutoscalerStatus), b.(*v2beta1.HorizontalPodAutoscalerStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2beta1.MetricSpec)(nil), (*autoscaling.MetricSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_MetricSpec_To_autoscaling_MetricSpec(a.(*v2beta1.MetricSpec), b.(*autoscaling.MetricSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*autoscaling.MetricSpec)(nil), (*v2beta1.MetricSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_autoscaling_MetricSpec_To_v2beta1_MetricSpec(a.(*autoscaling.MetricSpec), b.(*v2beta1.MetricSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2beta1.MetricStatus)(nil), (*autoscaling.MetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_MetricStatus_To_autoscaling_MetricStatus(a.(*v2beta1.MetricStatus), b.(*autoscaling.MetricStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*autoscaling.MetricStatus)(nil), (*v2beta1.MetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_autoscaling_MetricStatus_To_v2beta1_MetricStatus(a.(*autoscaling.MetricStatus), b.(*v2beta1.MetricStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*autoscaling.ExternalMetricSource)(nil), (*v2beta1.ExternalMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_autoscaling_ExternalMetricSource_To_v2beta1_ExternalMetricSource(a.(*autoscaling.ExternalMetricSource), b.(*v2beta1.ExternalMetricSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*autoscaling.ExternalMetricStatus)(nil), (*v2beta1.ExternalMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_autoscaling_ExternalMetricStatus_To_v2beta1_ExternalMetricStatus(a.(*autoscaling.ExternalMetricStatus), b.(*v2beta1.ExternalMetricStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*autoscaling.HorizontalPodAutoscalerSpec)(nil), (*v2beta1.HorizontalPodAutoscalerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_autoscaling_HorizontalPodAutoscalerSpec_To_v2beta1_HorizontalPodAutoscalerSpec(a.(*autoscaling.HorizontalPodAutoscalerSpec), b.(*v2beta1.HorizontalPodAutoscalerSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*autoscaling.HorizontalPodAutoscaler)(nil), (*v2beta1.HorizontalPodAutoscaler)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_autoscaling_HorizontalPodAutoscaler_To_v2beta1_HorizontalPodAutoscaler(a.(*autoscaling.HorizontalPodAutoscaler), b.(*v2beta1.HorizontalPodAutoscaler), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*autoscaling.MetricTarget)(nil), (*v2beta1.CrossVersionObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_autoscaling_MetricTarget_To_v2beta1_CrossVersionObjectReference(a.(*autoscaling.MetricTarget), b.(*v2beta1.CrossVersionObjectReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*autoscaling.ObjectMetricSource)(nil), (*v2beta1.ObjectMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_autoscaling_ObjectMetricSource_To_v2beta1_ObjectMetricSource(a.(*autoscaling.ObjectMetricSource), b.(*v2beta1.ObjectMetricSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*autoscaling.ObjectMetricStatus)(nil), (*v2beta1.ObjectMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_autoscaling_ObjectMetricStatus_To_v2beta1_ObjectMetricStatus(a.(*autoscaling.ObjectMetricStatus), b.(*v2beta1.ObjectMetricStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*autoscaling.PodsMetricSource)(nil), (*v2beta1.PodsMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_autoscaling_PodsMetricSource_To_v2beta1_PodsMetricSource(a.(*autoscaling.PodsMetricSource), b.(*v2beta1.PodsMetricSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*autoscaling.PodsMetricStatus)(nil), (*v2beta1.PodsMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_autoscaling_PodsMetricStatus_To_v2beta1_PodsMetricStatus(a.(*autoscaling.PodsMetricStatus), b.(*v2beta1.PodsMetricStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*autoscaling.ResourceMetricSource)(nil), (*v2beta1.ResourceMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_autoscaling_ResourceMetricSource_To_v2beta1_ResourceMetricSource(a.(*autoscaling.ResourceMetricSource), b.(*v2beta1.ResourceMetricSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*autoscaling.ResourceMetricStatus)(nil), (*v2beta1.ResourceMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_autoscaling_ResourceMetricStatus_To_v2beta1_ResourceMetricStatus(a.(*autoscaling.ResourceMetricStatus), b.(*v2beta1.ResourceMetricStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v2beta1.CrossVersionObjectReference)(nil), (*autoscaling.MetricTarget)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_CrossVersionObjectReference_To_autoscaling_MetricTarget(a.(*v2beta1.CrossVersionObjectReference), b.(*autoscaling.MetricTarget), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v2beta1.ExternalMetricSource)(nil), (*autoscaling.ExternalMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_ExternalMetricSource_To_autoscaling_ExternalMetricSource(a.(*v2beta1.ExternalMetricSource), b.(*autoscaling.ExternalMetricSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v2beta1.ExternalMetricStatus)(nil), (*autoscaling.ExternalMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_ExternalMetricStatus_To_autoscaling_ExternalMetricStatus(a.(*v2beta1.ExternalMetricStatus), b.(*autoscaling.ExternalMetricStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v2beta1.HorizontalPodAutoscaler)(nil), (*autoscaling.HorizontalPodAutoscaler)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler(a.(*v2beta1.HorizontalPodAutoscaler), b.(*autoscaling.HorizontalPodAutoscaler), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v2beta1.ObjectMetricSource)(nil), (*autoscaling.ObjectMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_ObjectMetricSource_To_autoscaling_ObjectMetricSource(a.(*v2beta1.ObjectMetricSource), b.(*autoscaling.ObjectMetricSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v2beta1.ObjectMetricStatus)(nil), (*autoscaling.ObjectMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus(a.(*v2beta1.ObjectMetricStatus), b.(*autoscaling.ObjectMetricStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v2beta1.PodsMetricSource)(nil), (*autoscaling.PodsMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_PodsMetricSource_To_autoscaling_PodsMetricSource(a.(*v2beta1.PodsMetricSource), b.(*autoscaling.PodsMetricSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v2beta1.PodsMetricStatus)(nil), (*autoscaling.PodsMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_PodsMetricStatus_To_autoscaling_PodsMetricStatus(a.(*v2beta1.PodsMetricStatus), b.(*autoscaling.PodsMetricStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v2beta1.ResourceMetricSource)(nil), (*autoscaling.ResourceMetricSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_ResourceMetricSource_To_autoscaling_ResourceMetricSource(a.(*v2beta1.ResourceMetricSource), b.(*autoscaling.ResourceMetricSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v2beta1.ResourceMetricStatus)(nil), (*autoscaling.ResourceMetricStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus(a.(*v2beta1.ResourceMetricStatus), b.(*autoscaling.ResourceMetricStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v2beta1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(in *v2beta1.CrossVersionObjectReference, out *autoscaling.CrossVersionObjectReference, s conversion.Scope) error {
	out.Kind = in.Kind
	out.Name = in.Name
	out.APIVersion = in.APIVersion
	return nil
}

// Convert_v2beta1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference is an autogenerated conversion function.
func Convert_v2beta1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(in *v2beta1.CrossVersionObjectReference, out *autoscaling.CrossVersionObjectReference, s conversion.Scope) error {
	return autoConvert_v2beta1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(in, out, s)
}

func autoConvert_autoscaling_CrossVersionObjectReference_To_v2beta1_CrossVersionObjectReference(in *autoscaling.CrossVersionObjectReference, out *v2beta1.CrossVersionObjectReference, s conversion.Scope) error {
	out.Kind = in.Kind
	out.Name = in.Name
	out.APIVersion = in.APIVersion
	return nil
}

// Convert_autoscaling_CrossVersionObjectReference_To_v2beta1_CrossVersionObjectReference is an autogenerated conversion function.
func Convert_autoscaling_CrossVersionObjectReference_To_v2beta1_CrossVersionObjectReference(in *autoscaling.CrossVersionObjectReference, out *v2beta1.CrossVersionObjectReference, s conversion.Scope) error {
	return autoConvert_autoscaling_CrossVersionObjectReference_To_v2beta1_CrossVersionObjectReference(in, out, s)
}

func autoConvert_v2beta1_ExternalMetricSource_To_autoscaling_ExternalMetricSource(in *v2beta1.ExternalMetricSource, out *autoscaling.ExternalMetricSource, s conversion.Scope) error {
	// WARNING: in.MetricName requires manual conversion: does not exist in peer-type
	// WARNING: in.MetricSelector requires manual conversion: does not exist in peer-type
	// WARNING: in.TargetValue requires manual conversion: does not exist in peer-type
	// WARNING: in.TargetAverageValue requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_autoscaling_ExternalMetricSource_To_v2beta1_ExternalMetricSource(in *autoscaling.ExternalMetricSource, out *v2beta1.ExternalMetricSource, s conversion.Scope) error {
	// WARNING: in.Metric requires manual conversion: does not exist in peer-type
	// WARNING: in.Target requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v2beta1_ExternalMetricStatus_To_autoscaling_ExternalMetricStatus(in *v2beta1.ExternalMetricStatus, out *autoscaling.ExternalMetricStatus, s conversion.Scope) error {
	// WARNING: in.MetricName requires manual conversion: does not exist in peer-type
	// WARNING: in.MetricSelector requires manual conversion: does not exist in peer-type
	// WARNING: in.CurrentValue requires manual conversion: does not exist in peer-type
	// WARNING: in.CurrentAverageValue requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_autoscaling_ExternalMetricStatus_To_v2beta1_ExternalMetricStatus(in *autoscaling.ExternalMetricStatus, out *v2beta1.ExternalMetricStatus, s conversion.Scope) error {
	// WARNING: in.Metric requires manual conversion: does not exist in peer-type
	// WARNING: in.Current requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v2beta1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler(in *v2beta1.HorizontalPodAutoscaler, out *autoscaling.HorizontalPodAutoscaler, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v2beta1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v2beta1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_autoscaling_HorizontalPodAutoscaler_To_v2beta1_HorizontalPodAutoscaler(in *autoscaling.HorizontalPodAutoscaler, out *v2beta1.HorizontalPodAutoscaler, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_autoscaling_HorizontalPodAutoscalerSpec_To_v2beta1_HorizontalPodAutoscalerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_autoscaling_HorizontalPodAutoscalerStatus_To_v2beta1_HorizontalPodAutoscalerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_v2beta1_HorizontalPodAutoscalerCondition_To_autoscaling_HorizontalPodAutoscalerCondition(in *v2beta1.HorizontalPodAutoscalerCondition, out *autoscaling.HorizontalPodAutoscalerCondition, s conversion.Scope) error {
	out.Type = autoscaling.HorizontalPodAutoscalerConditionType(in.Type)
	out.Status = autoscaling.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v2beta1_HorizontalPodAutoscalerCondition_To_autoscaling_HorizontalPodAutoscalerCondition is an autogenerated conversion function.
func Convert_v2beta1_HorizontalPodAutoscalerCondition_To_autoscaling_HorizontalPodAutoscalerCondition(in *v2beta1.HorizontalPodAutoscalerCondition, out *autoscaling.HorizontalPodAutoscalerCondition, s conversion.Scope) error {
	return autoConvert_v2beta1_HorizontalPodAutoscalerCondition_To_autoscaling_HorizontalPodAutoscalerCondition(in, out, s)
}

func autoConvert_autoscaling_HorizontalPodAutoscalerCondition_To_v2beta1_HorizontalPodAutoscalerCondition(in *autoscaling.HorizontalPodAutoscalerCondition, out *v2beta1.HorizontalPodAutoscalerCondition, s conversion.Scope) error {
	out.Type = v2beta1.HorizontalPodAutoscalerConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_autoscaling_HorizontalPodAutoscalerCondition_To_v2beta1_HorizontalPodAutoscalerCondition is an autogenerated conversion function.
func Convert_autoscaling_HorizontalPodAutoscalerCondition_To_v2beta1_HorizontalPodAutoscalerCondition(in *autoscaling.HorizontalPodAutoscalerCondition, out *v2beta1.HorizontalPodAutoscalerCondition, s conversion.Scope) error {
	return autoConvert_autoscaling_HorizontalPodAutoscalerCondition_To_v2beta1_HorizontalPodAutoscalerCondition(in, out, s)
}

func autoConvert_v2beta1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList(in *v2beta1.HorizontalPodAutoscalerList, out *autoscaling.HorizontalPodAutoscalerList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]autoscaling.HorizontalPodAutoscaler, len(*in))
		for i := range *in {
			if err := Convert_v2beta1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v2beta1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList is an autogenerated conversion function.
func Convert_v2beta1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList(in *v2beta1.HorizontalPodAutoscalerList, out *autoscaling.HorizontalPodAutoscalerList, s conversion.Scope) error {
	return autoConvert_v2beta1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList(in, out, s)
}

func autoConvert_autoscaling_HorizontalPodAutoscalerList_To_v2beta1_HorizontalPodAutoscalerList(in *autoscaling.HorizontalPodAutoscalerList, out *v2beta1.HorizontalPodAutoscalerList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v2beta1.HorizontalPodAutoscaler, len(*in))
		for i := range *in {
			if err := Convert_autoscaling_HorizontalPodAutoscaler_To_v2beta1_HorizontalPodAutoscaler(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_autoscaling_HorizontalPodAutoscalerList_To_v2beta1_HorizontalPodAutoscalerList is an autogenerated conversion function.
func Convert_autoscaling_HorizontalPodAutoscalerList_To_v2beta1_HorizontalPodAutoscalerList(in *autoscaling.HorizontalPodAutoscalerList, out *v2beta1.HorizontalPodAutoscalerList, s conversion.Scope) error {
	return autoConvert_autoscaling_HorizontalPodAutoscalerList_To_v2beta1_HorizontalPodAutoscalerList(in, out, s)
}

func autoConvert_v2beta1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec(in *v2beta1.HorizontalPodAutoscalerSpec, out *autoscaling.HorizontalPodAutoscalerSpec, s conversion.Scope) error {
	if err := Convert_v2beta1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(&in.ScaleTargetRef, &out.ScaleTargetRef, s); err != nil {
		return err
	}
	out.MinReplicas = (*int32)(unsafe.Pointer(in.MinReplicas))
	out.MaxReplicas = in.MaxReplicas
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]autoscaling.MetricSpec, len(*in))
		for i := range *in {
			if err := Convert_v2beta1_MetricSpec_To_autoscaling_MetricSpec(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Metrics = nil
	}
	return nil
}

// Convert_v2beta1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec is an autogenerated conversion function.
func Convert_v2beta1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec(in *v2beta1.HorizontalPodAutoscalerSpec, out *autoscaling.HorizontalPodAutoscalerSpec, s conversion.Scope) error {
	return autoConvert_v2beta1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec(in, out, s)
}

func autoConvert_autoscaling_HorizontalPodAutoscalerSpec_To_v2beta1_HorizontalPodAutoscalerSpec(in *autoscaling.HorizontalPodAutoscalerSpec, out *v2beta1.HorizontalPodAutoscalerSpec, s conversion.Scope) error {
	if err := Convert_autoscaling_CrossVersionObjectReference_To_v2beta1_CrossVersionObjectReference(&in.ScaleTargetRef, &out.ScaleTargetRef, s); err != nil {
		return err
	}
	out.MinReplicas = (*int32)(unsafe.Pointer(in.MinReplicas))
	out.MaxReplicas = in.MaxReplicas
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]v2beta1.MetricSpec, len(*in))
		for i := range *in {
			if err := Convert_autoscaling_MetricSpec_To_v2beta1_MetricSpec(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Metrics = nil
	}
	return nil
}

func autoConvert_v2beta1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus(in *v2beta1.HorizontalPodAutoscalerStatus, out *autoscaling.HorizontalPodAutoscalerStatus, s conversion.Scope) error {
	out.ObservedGeneration = (*int64)(unsafe.Pointer(in.ObservedGeneration))
	out.LastScaleTime = (*metav1.Time)(unsafe.Pointer(in.LastScaleTime))
	out.CurrentReplicas = in.CurrentReplicas
	out.DesiredReplicas = in.DesiredReplicas
	if in.CurrentMetrics != nil {
		in, out := &in.CurrentMetrics, &out.CurrentMetrics
		*out = make([]autoscaling.MetricStatus, len(*in))
		for i := range *in {
			if err := Convert_v2beta1_MetricStatus_To_autoscaling_MetricStatus(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.CurrentMetrics = nil
	}
	out.Conditions = *(*[]autoscaling.HorizontalPodAutoscalerCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v2beta1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus is an autogenerated conversion function.
func Convert_v2beta1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus(in *v2beta1.HorizontalPodAutoscalerStatus, out *autoscaling.HorizontalPodAutoscalerStatus, s conversion.Scope) error {
	return autoConvert_v2beta1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus(in, out, s)
}

func autoConvert_autoscaling_HorizontalPodAutoscalerStatus_To_v2beta1_HorizontalPodAutoscalerStatus(in *autoscaling.HorizontalPodAutoscalerStatus, out *v2beta1.HorizontalPodAutoscalerStatus, s conversion.Scope) error {
	out.ObservedGeneration = (*int64)(unsafe.Pointer(in.ObservedGeneration))
	out.LastScaleTime = (*metav1.Time)(unsafe.Pointer(in.LastScaleTime))
	out.CurrentReplicas = in.CurrentReplicas
	out.DesiredReplicas = in.DesiredReplicas
	if in.CurrentMetrics != nil {
		in, out := &in.CurrentMetrics, &out.CurrentMetrics
		*out = make([]v2beta1.MetricStatus, len(*in))
		for i := range *in {
			if err := Convert_autoscaling_MetricStatus_To_v2beta1_MetricStatus(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.CurrentMetrics = nil
	}
	out.Conditions = *(*[]v2beta1.HorizontalPodAutoscalerCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_autoscaling_HorizontalPodAutoscalerStatus_To_v2beta1_HorizontalPodAutoscalerStatus is an autogenerated conversion function.
func Convert_autoscaling_HorizontalPodAutoscalerStatus_To_v2beta1_HorizontalPodAutoscalerStatus(in *autoscaling.HorizontalPodAutoscalerStatus, out *v2beta1.HorizontalPodAutoscalerStatus, s conversion.Scope) error {
	return autoConvert_autoscaling_HorizontalPodAutoscalerStatus_To_v2beta1_HorizontalPodAutoscalerStatus(in, out, s)
}

func autoConvert_v2beta1_MetricSpec_To_autoscaling_MetricSpec(in *v2beta1.MetricSpec, out *autoscaling.MetricSpec, s conversion.Scope) error {
	out.Type = autoscaling.MetricSourceType(in.Type)
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = new(autoscaling.ObjectMetricSource)
		if err := Convert_v2beta1_ObjectMetricSource_To_autoscaling_ObjectMetricSource(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Object = nil
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = new(autoscaling.PodsMetricSource)
		if err := Convert_v2beta1_PodsMetricSource_To_autoscaling_PodsMetricSource(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Pods = nil
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(autoscaling.ResourceMetricSource)
		if err := Convert_v2beta1_ResourceMetricSource_To_autoscaling_ResourceMetricSource(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Resource = nil
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(autoscaling.ExternalMetricSource)
		if err := Convert_v2beta1_ExternalMetricSource_To_autoscaling_ExternalMetricSource(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.External = nil
	}
	return nil
}

// Convert_v2beta1_MetricSpec_To_autoscaling_MetricSpec is an autogenerated conversion function.
func Convert_v2beta1_MetricSpec_To_autoscaling_MetricSpec(in *v2beta1.MetricSpec, out *autoscaling.MetricSpec, s conversion.Scope) error {
	return autoConvert_v2beta1_MetricSpec_To_autoscaling_MetricSpec(in, out, s)
}

func autoConvert_autoscaling_MetricSpec_To_v2beta1_MetricSpec(in *autoscaling.MetricSpec, out *v2beta1.MetricSpec, s conversion.Scope) error {
	out.Type = v2beta1.MetricSourceType(in.Type)
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = new(v2beta1.ObjectMetricSource)
		if err := Convert_autoscaling_ObjectMetricSource_To_v2beta1_ObjectMetricSource(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Object = nil
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = new(v2beta1.PodsMetricSource)
		if err := Convert_autoscaling_PodsMetricSource_To_v2beta1_PodsMetricSource(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Pods = nil
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(v2beta1.ResourceMetricSource)
		if err := Convert_autoscaling_ResourceMetricSource_To_v2beta1_ResourceMetricSource(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Resource = nil
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(v2beta1.ExternalMetricSource)
		if err := Convert_autoscaling_ExternalMetricSource_To_v2beta1_ExternalMetricSource(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.External = nil
	}
	return nil
}

// Convert_autoscaling_MetricSpec_To_v2beta1_MetricSpec is an autogenerated conversion function.
func Convert_autoscaling_MetricSpec_To_v2beta1_MetricSpec(in *autoscaling.MetricSpec, out *v2beta1.MetricSpec, s conversion.Scope) error {
	return autoConvert_autoscaling_MetricSpec_To_v2beta1_MetricSpec(in, out, s)
}

func autoConvert_v2beta1_MetricStatus_To_autoscaling_MetricStatus(in *v2beta1.MetricStatus, out *autoscaling.MetricStatus, s conversion.Scope) error {
	out.Type = autoscaling.MetricSourceType(in.Type)
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = new(autoscaling.ObjectMetricStatus)
		if err := Convert_v2beta1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Object = nil
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = new(autoscaling.PodsMetricStatus)
		if err := Convert_v2beta1_PodsMetricStatus_To_autoscaling_PodsMetricStatus(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Pods = nil
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(autoscaling.ResourceMetricStatus)
		if err := Convert_v2beta1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Resource = nil
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(autoscaling.ExternalMetricStatus)
		if err := Convert_v2beta1_ExternalMetricStatus_To_autoscaling_ExternalMetricStatus(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.External = nil
	}
	return nil
}

// Convert_v2beta1_MetricStatus_To_autoscaling_MetricStatus is an autogenerated conversion function.
func Convert_v2beta1_MetricStatus_To_autoscaling_MetricStatus(in *v2beta1.MetricStatus, out *autoscaling.MetricStatus, s conversion.Scope) error {
	return autoConvert_v2beta1_MetricStatus_To_autoscaling_MetricStatus(in, out, s)
}

func autoConvert_autoscaling_MetricStatus_To_v2beta1_MetricStatus(in *autoscaling.MetricStatus, out *v2beta1.MetricStatus, s conversion.Scope) error {
	out.Type = v2beta1.MetricSourceType(in.Type)
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = new(v2beta1.ObjectMetricStatus)
		if err := Convert_autoscaling_ObjectMetricStatus_To_v2beta1_ObjectMetricStatus(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Object = nil
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = new(v2beta1.PodsMetricStatus)
		if err := Convert_autoscaling_PodsMetricStatus_To_v2beta1_PodsMetricStatus(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Pods = nil
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(v2beta1.ResourceMetricStatus)
		if err := Convert_autoscaling_ResourceMetricStatus_To_v2beta1_ResourceMetricStatus(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Resource = nil
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(v2beta1.ExternalMetricStatus)
		if err := Convert_autoscaling_ExternalMetricStatus_To_v2beta1_ExternalMetricStatus(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.External = nil
	}
	return nil
}

// Convert_autoscaling_MetricStatus_To_v2beta1_MetricStatus is an autogenerated conversion function.
func Convert_autoscaling_MetricStatus_To_v2beta1_MetricStatus(in *autoscaling.MetricStatus, out *v2beta1.MetricStatus, s conversion.Scope) error {
	return autoConvert_autoscaling_MetricStatus_To_v2beta1_MetricStatus(in, out, s)
}

func autoConvert_v2beta1_ObjectMetricSource_To_autoscaling_ObjectMetricSource(in *v2beta1.ObjectMetricSource, out *autoscaling.ObjectMetricSource, s conversion.Scope) error {
	if err := Convert_v2beta1_CrossVersionObjectReference_To_autoscaling_MetricTarget(&in.Target, &out.Target, s); err != nil {
		return err
	}
	// WARNING: in.MetricName requires manual conversion: does not exist in peer-type
	// WARNING: in.TargetValue requires manual conversion: does not exist in peer-type
	// WARNING: in.Selector requires manual conversion: does not exist in peer-type
	// WARNING: in.AverageValue requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_autoscaling_ObjectMetricSource_To_v2beta1_ObjectMetricSource(in *autoscaling.ObjectMetricSource, out *v2beta1.ObjectMetricSource, s conversion.Scope) error {
	// WARNING: in.DescribedObject requires manual conversion: does not exist in peer-type
	if err := Convert_autoscaling_MetricTarget_To_v2beta1_CrossVersionObjectReference(&in.Target, &out.Target, s); err != nil {
		return err
	}
	// WARNING: in.Metric requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v2beta1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus(in *v2beta1.ObjectMetricStatus, out *autoscaling.ObjectMetricStatus, s conversion.Scope) error {
	// WARNING: in.Target requires manual conversion: does not exist in peer-type
	// WARNING: in.MetricName requires manual conversion: does not exist in peer-type
	// WARNING: in.CurrentValue requires manual conversion: does not exist in peer-type
	// WARNING: in.Selector requires manual conversion: does not exist in peer-type
	// WARNING: in.AverageValue requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_autoscaling_ObjectMetricStatus_To_v2beta1_ObjectMetricStatus(in *autoscaling.ObjectMetricStatus, out *v2beta1.ObjectMetricStatus, s conversion.Scope) error {
	// WARNING: in.Metric requires manual conversion: does not exist in peer-type
	// WARNING: in.Current requires manual conversion: does not exist in peer-type
	// WARNING: in.DescribedObject requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v2beta1_PodsMetricSource_To_autoscaling_PodsMetricSource(in *v2beta1.PodsMetricSource, out *autoscaling.PodsMetricSource, s conversion.Scope) error {
	// WARNING: in.MetricName requires manual conversion: does not exist in peer-type
	// WARNING: in.TargetAverageValue requires manual conversion: does not exist in peer-type
	// WARNING: in.Selector requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_autoscaling_PodsMetricSource_To_v2beta1_PodsMetricSource(in *autoscaling.PodsMetricSource, out *v2beta1.PodsMetricSource, s conversion.Scope) error {
	// WARNING: in.Metric requires manual conversion: does not exist in peer-type
	// WARNING: in.Target requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v2beta1_PodsMetricStatus_To_autoscaling_PodsMetricStatus(in *v2beta1.PodsMetricStatus, out *autoscaling.PodsMetricStatus, s conversion.Scope) error {
	// WARNING: in.MetricName requires manual conversion: does not exist in peer-type
	// WARNING: in.CurrentAverageValue requires manual conversion: does not exist in peer-type
	// WARNING: in.Selector requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_autoscaling_PodsMetricStatus_To_v2beta1_PodsMetricStatus(in *autoscaling.PodsMetricStatus, out *v2beta1.PodsMetricStatus, s conversion.Scope) error {
	// WARNING: in.Metric requires manual conversion: does not exist in peer-type
	// WARNING: in.Current requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v2beta1_ResourceMetricSource_To_autoscaling_ResourceMetricSource(in *v2beta1.ResourceMetricSource, out *autoscaling.ResourceMetricSource, s conversion.Scope) error {
	out.Name = core.ResourceName(in.Name)
	// WARNING: in.TargetAverageUtilization requires manual conversion: does not exist in peer-type
	// WARNING: in.TargetAverageValue requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_autoscaling_ResourceMetricSource_To_v2beta1_ResourceMetricSource(in *autoscaling.ResourceMetricSource, out *v2beta1.ResourceMetricSource, s conversion.Scope) error {
	out.Name = v1.ResourceName(in.Name)
	// WARNING: in.Target requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v2beta1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus(in *v2beta1.ResourceMetricStatus, out *autoscaling.ResourceMetricStatus, s conversion.Scope) error {
	out.Name = core.ResourceName(in.Name)
	// WARNING: in.CurrentAverageUtilization requires manual conversion: does not exist in peer-type
	// WARNING: in.CurrentAverageValue requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_autoscaling_ResourceMetricStatus_To_v2beta1_ResourceMetricStatus(in *autoscaling.ResourceMetricStatus, out *v2beta1.ResourceMetricStatus, s conversion.Scope) error {
	out.Name = v1.ResourceName(in.Name)
	// WARNING: in.Current requires manual conversion: does not exist in peer-type
	return nil
}
