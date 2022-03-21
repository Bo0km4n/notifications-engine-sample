/*
Copyright 2022 Bo0km4n.

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

package controllers

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	dummyv1 "github.com/Bo0km4n/notifications-engine-sample/api/v1"
)

// FooReconciler reconciles a Foo object
type FooReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

const (
	finalizer string = "dummy.example.org/foo"
)

//+kubebuilder:rbac:groups=dummy.example.org,resources=foos,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=dummy.example.org,resources=foos/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=dummy.example.org,resources=foos/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Foo object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *FooReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// your logic here
	f := &dummyv1.Foo{}
	if err := r.Client.Get(ctx, req.NamespacedName, f); err != nil {
		return ctrl.Result{}, err
	}
	if !f.DeletionTimestamp.IsZero() {
		newF := f.DeepCopy()
		controllerutil.RemoveFinalizer(newF, finalizer)
		patch := client.MergeFrom(f)
		if err := r.Client.Patch(ctx, newF, patch); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}
	time.Sleep(1)
	controllerutil.AddFinalizer(f, finalizer)
	if err := r.Client.Update(ctx, f); err != nil {
		return ctrl.Result{}, err
	}
	f.Status.Ready = true
	if err := r.Status().Update(ctx, f); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *FooReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&dummyv1.Foo{}).
		Complete(r)
}
