/*
Copyright 2017 The Kubernetes Authors.

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

package admission_test

import (
	"fmt"

	"github.com/kubernetes-sigs/controller-runtime/pkg/internal/admission"
	"k8s.io/api/admission/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ExampleDecode() {
	var review v1beta1.AdmissionReview
	resourceType := metav1.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}
	pod := corev1.Pod{}
	if errResp := admission.Decode(review, &pod, resourceType); errResp != nil {
		// Send error resp
	}
}

func ExampleErrorResponse() {
	admission.ErrorResponse(fmt.Errorf("some error explanation"))
}

func ExampleDenyResponse() {
	admission.DenyResponse(fmt.Sprintf("some deny explanation"))
}

func ExampleAllowResponse() {
	admission.AllowResponse()
}
