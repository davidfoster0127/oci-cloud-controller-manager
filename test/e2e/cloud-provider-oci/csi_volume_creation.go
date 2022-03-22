// Copyright 2018 Oracle and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package e2e

import (
	"time"

	. "github.com/onsi/ginkgo"
	"github.com/oracle/oci-cloud-controller-manager/test/e2e/framework"
)

var _ = Describe("CSI Volume Creation", func() {
	f := framework.NewDefaultFramework("csi-basic")
	Context("[cloudprovider][storage][csi]", func() {
		It("Create PVC and POD for CSI.", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-provisioner-e2e-tests")

			scName := f.CreateStorageClassOrFail(framework.ClassOCICSI, "blockvolume.csi.oraclecloud.com", nil, pvcJig.Labels, "WaitForFirstConsumer", false)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MinVolumeBlock, scName, nil)
			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
			pvcJig.NewPodForCSI("app1", f.Namespace.Name, pvc.Name, setupF.AdLabel)
		})

		It("Create PVC with VolumeSize 1Gi but should use default 50Gi", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-provisioner-e2e-tests-pvc-with-1gi")

			scName := f.CreateStorageClassOrFail(framework.ClassOCICSI, "blockvolume.csi.oraclecloud.com", nil, pvcJig.Labels, "WaitForFirstConsumer", false)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.VolumeFss, scName, nil)
			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
			pvcJig.NewPodForCSI("app2", f.Namespace.Name, pvc.Name, setupF.AdLabel)

			time.Sleep(60 * time.Second) //waiting for pod to up and running

			pvcJig.CheckVolumeCapacity("50Gi", pvc.Name, f.Namespace.Name)
		})

		It("Create PVC with VolumeSize 100Gi should use 100Gi", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-provisioner-e2e-tests-pvc-with-100gi")

			scName := f.CreateStorageClassOrFail(framework.ClassOCICSI, "blockvolume.csi.oraclecloud.com", nil, pvcJig.Labels, "WaitForFirstConsumer", false)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MaxVolumeBlock, scName, nil)
			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
			pvcJig.NewPodForCSI("app3", f.Namespace.Name, pvc.Name, setupF.AdLabel)

			time.Sleep(60 * time.Second) //waiting for pod to up and running

			pvcJig.CheckVolumeCapacity("100Gi", pvc.Name, f.Namespace.Name)
		})

		It("Data should persist on CSI volume on pod restart", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-pod-restart-data-persistence")

			scName := f.CreateStorageClassOrFail(framework.ClassOCICSI, "blockvolume.csi.oraclecloud.com", nil, pvcJig.Labels, "WaitForFirstConsumer", false)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MinVolumeBlock, scName, nil)
			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
			pvcJig.CheckDataPersistenceWithDeployment(pvc.Name, f.Namespace.Name)
		})

		It("FsGroup test for CSI", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-pod-nginx")

			scName := f.CreateStorageClassOrFail(framework.ClassOCICSI, "blockvolume.csi.oraclecloud.com", nil, pvcJig.Labels, "WaitForFirstConsumer", false)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MinVolumeBlock, scName, nil)

			pvcJig.CheckVolumeDirectoryOwnership(f.Namespace.Name, pvc)
		})
	})
})

var _ = Describe("CSI Volume Expansion iSCSI", func() {
	f := framework.NewDefaultFramework("csi-expansion")
	Context("[cloudprovider][storage][csi][expand][iSCSI]", func() {
		It("Expand PVC VolumeSize from 50Gi to 100Gi and asserts size, file existence and file corruptions for iSCSI volumes with existing storage class", func() {
			var size = "100Gi"
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-resizer-pvc-expand-to-100gi-iscsi")

			scName := f.CreateStorageClassOrFail(framework.ClassOCICSI, "blockvolume.csi.oraclecloud.com",
				map[string]string{framework.AttachmentType: framework.AttachmentTypeISCSI},
				pvcJig.Labels, "WaitForFirstConsumer", true)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MinVolumeBlock, scName, nil)
			podName := pvcJig.NewPodForCSI("expanded-pvc-app", f.Namespace.Name, pvc.Name, setupF.AdLabel)

			time.Sleep(60 * time.Second) //waiting for pod to up and running

			expandedPvc := pvcJig.UpdateAndAwaitPVCOrFailCSI(pvc, pvc.Namespace, size, nil)

			time.Sleep(120 * time.Second) //waiting for expanded pvc to be functional

			pvcJig.CheckVolumeCapacity("100Gi", expandedPvc.Name, f.Namespace.Name)
			pvcJig.CheckFileExists(f.Namespace.Name, podName, "/data", "testdata.txt")
			pvcJig.CheckFileCorruption(f.Namespace.Name, podName, "/data", "testdata.txt")
			pvcJig.CheckExpandedVolumeReadWrite(f.Namespace.Name, podName)
			pvcJig.CheckUsableVolumeSizeInsidePod(f.Namespace.Name, podName)

			_ = f.DeleteStorageClass(framework.ClassOCICSIExpand)
		})
	})
})

var _ = Describe("CSI Volume Expansion iSCSI", func() {
	f := framework.NewDefaultFramework("csi-expansion")
	Context("[cloudprovider][storage][csi][expand][iSCSI]", func() {
		It("Expand PVC VolumeSize from 50Gi to 100Gi and asserts size, file existence and file corruptions for iSCSI volumes with new storage class", func() {
			var size = "100Gi"
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-resizer-pvc-expand-to-100gi-iscsi")

			scName := f.CreateStorageClassOrFail(framework.ClassOCICSIExpand, "blockvolume.csi.oraclecloud.com",
				map[string]string{framework.AttachmentType: framework.AttachmentTypeISCSI},
				pvcJig.Labels, "WaitForFirstConsumer", true)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MinVolumeBlock, scName, nil)
			podName := pvcJig.NewPodForCSI("expanded-pvc-app", f.Namespace.Name, pvc.Name, setupF.AdLabel)

			time.Sleep(60 * time.Second) //waiting for pod to up and running

			expandedPvc := pvcJig.UpdateAndAwaitPVCOrFailCSI(pvc, pvc.Namespace, size, nil)

			time.Sleep(120 * time.Second) //waiting for expanded pvc to be functional

			pvcJig.CheckVolumeCapacity("100Gi", expandedPvc.Name, f.Namespace.Name)
			pvcJig.CheckFileExists(f.Namespace.Name, podName, "/data", "testdata.txt")
			pvcJig.CheckFileCorruption(f.Namespace.Name, podName, "/data", "testdata.txt")
			pvcJig.CheckExpandedVolumeReadWrite(f.Namespace.Name, podName)
			pvcJig.CheckUsableVolumeSizeInsidePod(f.Namespace.Name, podName)

			_ = f.DeleteStorageClass(framework.ClassOCICSIExpand)
		})
	})
})

var _ = Describe("CSI Volume Expansion Paravirtualized", func() {
	f := framework.NewDefaultFramework("csi-expansion")
	Context("[cloudprovider][storage][csi][expand][paravirtualized]", func() {
		It("Expand PVC VolumeSize from 50Gi to 100Gi and asserts size, file existence and file corruptions for paravirtualized volumes with new storage class", func() {
			var size = "100Gi"
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-resizer-pvc-expand-to-100gi-paravirtualized")

			scParameter := map[string]string{
				framework.KmsKey:         setupF.CMEKKMSKey,
				framework.AttachmentType: framework.AttachmentTypeParavirtualized,
			}
			scName := f.CreateStorageClassOrFail(framework.ClassOCICSIExpand,
				"blockvolume.csi.oraclecloud.com", scParameter, pvcJig.Labels,
				"WaitForFirstConsumer", true)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MinVolumeBlock, scName, nil)
			podName := pvcJig.NewPodForCSI("expanded-pvc-app", f.Namespace.Name, pvc.Name, setupF.AdLabel)

			time.Sleep(60 * time.Second) //waiting for pod to up and running

			expandedPvc := pvcJig.UpdateAndAwaitPVCOrFailCSI(pvc, pvc.Namespace, size, nil)

			time.Sleep(120 * time.Second) //waiting for expanded pvc to be functional

			pvcJig.CheckVolumeCapacity("100Gi", expandedPvc.Name, f.Namespace.Name)
			pvcJig.CheckFileExists(f.Namespace.Name, podName, "/data", "testdata.txt")
			pvcJig.CheckFileCorruption(f.Namespace.Name, podName, "/data", "testdata.txt")
			pvcJig.CheckExpandedVolumeReadWrite(f.Namespace.Name, podName)
			pvcJig.CheckUsableVolumeSizeInsidePod(f.Namespace.Name, podName)

			_ = f.DeleteStorageClass(framework.ClassOCICSIExpand)
		})
	})
})

var _ = Describe("CSI Static Volume Creation", func() {
	f := framework.NewBackupFramework("csi-static")
	Context("[cloudprovider][storage][csi]", func() {
		It("Static Provisioning CSI", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-provisioner-e2e-tests-pvc-with-static")

			scName := f.CreateStorageClassOrFail(framework.ClassOCICSI, "blockvolume.csi.oraclecloud.com",
				nil, pvcJig.Labels, "WaitForFirstConsumer", false)

			compartmentId := ""
			if setupF.Compartment1 != "" {
				compartmentId = setupF.Compartment1
			} else if f.CloudProviderConfig.CompartmentID != "" {
				compartmentId = f.CloudProviderConfig.CompartmentID
			} else if f.CloudProviderConfig.Auth.CompartmentID != "" {
				compartmentId = f.CloudProviderConfig.Auth.CompartmentID
			} else {
				framework.Failf("Compartment Id undefined.")
			}
			pvc := pvcJig.CreateAndAwaitStaticPVCOrFailCSI(f.BlockStorageClient, f.Namespace.Name, framework.MinVolumeBlock, scName, setupF.AdLocation, compartmentId, nil)
			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
			pvcJig.NewPodForCSI("app4", f.Namespace.Name, pvc.Name, setupF.AdLabel)

			time.Sleep(60 * time.Second) //waiting for pod to up and running

			pvcJig.CheckVolumeCapacity("50Gi", pvc.Name, f.Namespace.Name)
		})
	})
})

var _ = Describe("CSI CMEK,PV attachment and in-transit encryption test", func() {
	f := framework.NewDefaultFramework("csi-basic")
	Context("[cloudprovider][storage][csi]", func() {
		It("Create PVC and POD for CSI with CMEK,PV attachment and in-transit encryption", func() {
			TestCMEKAttachmentTypeAndEncryptionType(f, framework.AttachmentTypeParavirtualized)
		})
	})

	Context("[cloudprovider][storage][csi]", func() {
		It("Create PVC and POD for CSI with CMEK,ISCSI attachment and in-transit encryption", func() {
			TestCMEKAttachmentTypeAndEncryptionType(f, framework.AttachmentTypeISCSI)
		})
	})

})

func TestCMEKAttachmentTypeAndEncryptionType(f *framework.CloudProviderFramework, expectedAttachmentType string) {
	pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-cmek-iscsi-in-transit-e2e-tests")
	scParameter := map[string]string{
		framework.KmsKey:         setupF.CMEKKMSKey,
		framework.AttachmentType: expectedAttachmentType,
	}
	scName := f.CreateStorageClassOrFail(framework.SCName, "blockvolume.csi.oraclecloud.com", scParameter, pvcJig.Labels, "WaitForFirstConsumer", false)
	pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MinVolumeBlock, scName, nil)
	podName := pvcJig.NewPodForCSI("app1", f.Namespace.Name, pvc.Name, setupF.AdLabel)
	pvcJig.CheckCMEKKey(f.Client.BlockStorage(), pvc.Name, f.Namespace.Name, setupF.CMEKKMSKey)
	pvcJig.CheckAttachmentTypeAndEncryptionType(f.Client.Compute(), pvc.Name, f.Namespace.Name, podName, expectedAttachmentType)
	f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
}