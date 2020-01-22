// Copyright © 2020 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	corev1 "k8s.io/api/core/v1"
)

// +kubebuilder:object:generate=true

type KubernetesStorage struct {
	HostPath       *corev1.HostPathVolumeSource `json:"hostPath,omitempty"`
	EmptyDir       *corev1.EmptyDirVolumeSource `json:"emptyDir,omitempty"`
	// PersistentVolumeClaim defines the Spec and the Source at the same time.
	// The PVC will be created with the configured spec and the name defined in the source.
	PersistentVolumeClaim *PersistentVolumeClaim `json:"pvc,omitempty"`
}

// +kubebuilder:object:generate=true

type PersistentVolumeClaim struct {
	PersistentVolumeClaimSpec corev1.PersistentVolumeClaimSpec         `json:"spec,omitempty"`
	PersistentVolumeSource    corev1.PersistentVolumeClaimVolumeSource `json:"source,omitempty"`
}

// GetVolume returns a default emptydir volume if none configured
//
// `name`    will be the name of the volume and the lowest level directory in case a hostPath mount is used
// `path`    is the path in case the hostPath volume type is used
func (storage KubernetesStorage) GetVolume(name, path string) corev1.Volume {
	volume := corev1.Volume{
		Name: name,
	}
	if storage.HostPath != nil {
		if storage.HostPath.Path == "" {
			storage.HostPath.Path = path
		}
		volume.VolumeSource = corev1.VolumeSource{
			HostPath: storage.HostPath,
		}
		return volume
	} else if storage.EmptyDir != nil {
		volume.VolumeSource = corev1.VolumeSource{
			EmptyDir: storage.EmptyDir,
		}
		return volume
	} else if storage.PersistentVolumeClaim != nil {
		volume.VolumeSource = corev1.VolumeSource{
			PersistentVolumeClaim: &storage.PersistentVolumeClaim.PersistentVolumeSource,
		}
	}
	// return a default emptydir volume if none configured
	volume.VolumeSource = corev1.VolumeSource{
		EmptyDir: &corev1.EmptyDirVolumeSource{},
	}
	return volume
}