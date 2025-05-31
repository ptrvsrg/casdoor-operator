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

	ctrl "sigs.k8s.io/controller-runtime"
)

type Controller interface {
	Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error)
	SetupWithManager(ctx context.Context, mgr ctrl.Manager) error
}

func SetupWithManager(ctx context.Context, mgr ctrl.Manager, controller ...Controller) error {
	for _, c := range controller {
		if err := c.SetupWithManager(ctx, mgr); err != nil {
			return fmt.Errorf("unable to create controller %T: %w", c, err)
		}
	}
	return nil
}
