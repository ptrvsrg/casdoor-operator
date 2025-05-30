/*
Copyright 2025 ptrvsrg.

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

package controller

// import (
// 	"context"
//
// 	. "github.com/onsi/ginkgo/v2"
// 	. "github.com/onsi/gomega"
// 	"k8s.io/apimachinery/pkg/api/errors"
// 	"k8s.io/apimachinery/pkg/types"
// 	"sigs.k8s.io/controller-runtime/pkg/reconcile"
//
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//
// 	casdoorv1alpha1 "github.com/ptrvsrg/casdoor-operator/api/v1alpha1"
// )
//
// var _ = Describe(
// 	"Casdoor Controller", func() {
// 		Context(
// 			"When reconciling a resource", func() {
// 				const resourceName = "test-resource"
//
// 				ctx := context.Background()
//
// 				typeNamespacedName := types.NamespacedName{
// 					Name:      resourceName,
// 					Namespace: "default",
// 				}
// 				casdoor := &casdoorv1alpha1.Casdoor{}
//
// 				BeforeEach(
// 					func() {
// 						By("get the CR for the Kind Casdoor")
// 						err := k8sClient.Get(ctx, typeNamespacedName, casdoor)
// 						if err != nil && errors.IsNotFound(err) {
// 							By("creating the CR for the Kind Casdoor")
// 							resource := &casdoorv1alpha1.Casdoor{
// 								ObjectMeta: metav1.ObjectMeta{
// 									Name:      resourceName,
// 									Namespace: "default",
// 								},
// 							}
// 							Expect(k8sClient.Create(ctx, resource)).NotTo(HaveOccurred())
// 						}
// 					},
// 				)
//
// 				AfterEach(
// 					func() {
// 						resource := &casdoorv1alpha1.Casdoor{}
// 						err := k8sClient.Get(ctx, typeNamespacedName, resource)
// 						Expect(err).NotTo(HaveOccurred())
//
// 						By("Cleanup the specific resource instance Casdoor")
// 						Expect(k8sClient.Delete(ctx, resource)).To(Succeed())
// 					},
// 				)
//
// 				It(
// 					"should successfully reconcile the resource", func() {
// 						By("Reconciling the created resource")
//
// 						controllerReconciler := NewCasdoorReconciler(
// 							k8sClient,
// 							k8sClient.Scheme(),
// 							appCfg.SpecificControllers.Casdoor,
// 						)
//
// 						err := controllerReconciler.SetupWithManager(ctx, mgr)
// 						Expect(err).NotTo(HaveOccurred())
//
// 						_, err = controllerReconciler.Reconcile(
// 							ctx, reconcile.Request{
// 								NamespacedName: typeNamespacedName,
// 							},
// 						)
// 						Expect(err).NotTo(HaveOccurred())
// 					},
// 				)
// 			},
// 		)
// 	},
// )
