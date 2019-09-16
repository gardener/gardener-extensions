// +build !ignore_autogenerated

// Code generated by defaulter-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&CloudProfile{}, func(obj interface{}) { SetObjectDefaults_CloudProfile(obj.(*CloudProfile)) })
	scheme.AddTypeDefaultingFunc(&CloudProfileList{}, func(obj interface{}) { SetObjectDefaults_CloudProfileList(obj.(*CloudProfileList)) })
	scheme.AddTypeDefaultingFunc(&Project{}, func(obj interface{}) { SetObjectDefaults_Project(obj.(*Project)) })
	scheme.AddTypeDefaultingFunc(&ProjectList{}, func(obj interface{}) { SetObjectDefaults_ProjectList(obj.(*ProjectList)) })
	scheme.AddTypeDefaultingFunc(&SecretBinding{}, func(obj interface{}) { SetObjectDefaults_SecretBinding(obj.(*SecretBinding)) })
	scheme.AddTypeDefaultingFunc(&SecretBindingList{}, func(obj interface{}) { SetObjectDefaults_SecretBindingList(obj.(*SecretBindingList)) })
	return nil
}

func SetObjectDefaults_CloudProfile(in *CloudProfile) {
	for i := range in.Spec.MachineTypes {
		a := &in.Spec.MachineTypes[i]
		SetDefaults_MachineType(a)
	}
	for i := range in.Spec.VolumeTypes {
		a := &in.Spec.VolumeTypes[i]
		SetDefaults_VolumeType(a)
	}
}

func SetObjectDefaults_CloudProfileList(in *CloudProfileList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_CloudProfile(a)
	}
}

func SetObjectDefaults_Project(in *Project) {
	SetDefaults_Project(in)
}

func SetObjectDefaults_ProjectList(in *ProjectList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Project(a)
	}
}

func SetObjectDefaults_SecretBinding(in *SecretBinding) {
	SetDefaults_SecretBinding(in)
}

func SetObjectDefaults_SecretBindingList(in *SecretBindingList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_SecretBinding(a)
	}
}
