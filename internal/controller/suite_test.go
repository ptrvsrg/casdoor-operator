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

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"sync"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"resty.dev/v3"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/ptrvsrg/casdoor-operator/config"
	httpclient "github.com/ptrvsrg/casdoor-operator/internal/http/client"

	casdoorv1alpha1 "github.com/ptrvsrg/casdoor-operator/api/v1alpha1"
	// +kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var cfg *rest.Config
var mgr ctrl.Manager
var k8sClient client.Client
var httpClient *resty.Client
var appCfg config.Config
var testEnv *envtest.Environment
var ctx context.Context
var cancel context.CancelFunc

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Controller Suite")
}

var _ = BeforeSuite(
	func() {
		logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

		ctx, cancel = context.WithCancel(context.TODO()) //nolint:fatcontext

		By("bootstrapping test environment")
		testEnv = &envtest.Environment{
			CRDDirectoryPaths:     []string{filepath.Join("..", "..", "config", "crd", "bases")},
			ErrorIfCRDPathMissing: true,
			BinaryAssetsDirectory: filepath.Join(
				"..", "..", "bin", "k8s",
				fmt.Sprintf("1.31.0-%s-%s", runtime.GOOS, runtime.GOARCH),
			),
		}

		var err error
		// cfg is defined in this file globally.
		cfg, err = testEnv.Start()
		Expect(err).NotTo(HaveOccurred())
		Expect(cfg).NotTo(BeNil())

		err = casdoorv1alpha1.AddToScheme(scheme.Scheme)
		Expect(err).NotTo(HaveOccurred())

		// +kubebuilder:scaffold:scheme

		mgr, err = ctrl.NewManager(
			cfg, ctrl.Options{
				Scheme: scheme.Scheme,
			},
		)
		Expect(err).ToNot(HaveOccurred())

		k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
		Expect(err).ToNot(HaveOccurred())

		httpClient, err = httpclient.New()
		Expect(err).NotTo(HaveOccurred())
		Expect(httpClient).NotTo(BeNil())

		appCfg = config.Config{
			SpecificControllers: config.SpecificControllersConfig{
				Casdoor: config.CasdoorControllerConfig{
					MaxConcurrentReconciles: 1,
				},
			},
		}

		wg := &sync.WaitGroup{}
		wg.Add(1)
		err = mgr.Add(
			manager.RunnableFunc(
				func(context.Context) error {
					wg.Done()
					return nil
				},
			),
		)
		Expect(err).NotTo(HaveOccurred())
		Expect(httpClient).NotTo(BeNil())

		go func() {
			defer GinkgoRecover()
			err := mgr.Start(ctx)
			Expect(err).NotTo(HaveOccurred(), "failed to run manager")
		}()

		wg.Wait()
	},
)

var _ = AfterSuite(
	func() {
		By("tearing down the test environment")
		cancel()
		err := testEnv.Stop()
		Expect(err).NotTo(HaveOccurred())
	},
)
