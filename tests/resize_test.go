/*
Copyright 2019 The OpenEBS Authors

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

package volume

import (
	. "github.com/onsi/ginkgo"
	"fmt"
)

var _ = Describe("[csi] [jiva] TEST VOLUME RESIZE", func() {
	BeforeEach(prepareForVolumeResizeTest)
	AfterEach(cleanupAfterVolumeResizeTest)

	Context("App is deployed with volume replica count 1 and pvc is resized", func() {
		It("Should run Volume Resize Test", volumeResizeTest)
	})
})

func volumeResizeTest() {
	currentK8sVersion := getCurrentK8sMinorVersion()
	if currentK8sVersion < 16 {
		fmt.Printf(
			"resizing is not supported on Kubernetes version: 1.%d. Min supported version is 1.16\n",
			currentK8sVersion,
		)
		return
	}
	By("creating and verifying PVC bound status", createAndVerifyPVC)
	By("Creating and deploying app pod", createDeployVerifyApp)
	By("Expanding PVC", expandPVC)
	By("Verifying updated size in application pod", verifyIncreasedSizeInAppPod)
	By("Deleting application deployment", deleteAppDeployment)
	By("Deleting pvc", deletePVC)
}

func prepareForVolumeResizeTest() {
	By("Creating storage class", createStorageClass)
}

func cleanupAfterVolumeResizeTest() {
	By("Deleting storage class", deleteStorageClass)
}
