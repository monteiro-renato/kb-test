/*
Copyright 2024.

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

package v1alpha1

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	examplecomv1alpha1 "github.com/monteiro-renato/kb-test/api/v1alpha1"
)

// nolint:unused
// log is for logging in this package.
var busyboxlog = logf.Log.WithName("busybox-resource")

// SetupBusyboxWebhookWithManager registers the webhook for Busybox in the manager.
func SetupBusyboxWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&examplecomv1alpha1.Busybox{}).
		WithValidator(&BusyboxCustomValidator{}).
		WithDefaulter(&BusyboxCustomDefaulter{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-example-com-my-domain-v1alpha1-busybox,mutating=true,failurePolicy=fail,sideEffects=None,groups=example.com.my.domain,resources=busyboxes,verbs=create;update,versions=v1alpha1,name=mbusybox-v1alpha1.kb.io,admissionReviewVersions=v1

// BusyboxCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind Busybox when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type BusyboxCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

var _ webhook.CustomDefaulter = &BusyboxCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind Busybox.
func (d *BusyboxCustomDefaulter) Default(ctx context.Context, obj runtime.Object) error {
	busybox, ok := obj.(*examplecomv1alpha1.Busybox)

	if !ok {
		return fmt.Errorf("expected an Busybox object but got %T", obj)
	}
	busyboxlog.Info("Defaulting for Busybox", "name", busybox.GetName())

	// TODO(user): fill in your defaulting logic.

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-example-com-my-domain-v1alpha1-busybox,mutating=false,failurePolicy=fail,sideEffects=None,groups=example.com.my.domain,resources=busyboxes,verbs=create;update,versions=v1alpha1,name=vbusybox-v1alpha1.kb.io,admissionReviewVersions=v1

// BusyboxCustomValidator struct is responsible for validating the Busybox resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type BusyboxCustomValidator struct {
	// TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &BusyboxCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type Busybox.
func (v *BusyboxCustomValidator) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	busybox, ok := obj.(*examplecomv1alpha1.Busybox)
	if !ok {
		return nil, fmt.Errorf("expected a Busybox object but got %T", obj)
	}
	busyboxlog.Info("Validation for Busybox upon creation", "name", busybox.GetName())

	// TODO(user): fill in your validation logic upon object creation.

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type Busybox.
func (v *BusyboxCustomValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	busybox, ok := newObj.(*examplecomv1alpha1.Busybox)
	if !ok {
		return nil, fmt.Errorf("expected a Busybox object for the newObj but got %T", newObj)
	}
	busyboxlog.Info("Validation for Busybox upon update", "name", busybox.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type Busybox.
func (v *BusyboxCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	busybox, ok := obj.(*examplecomv1alpha1.Busybox)
	if !ok {
		return nil, fmt.Errorf("expected a Busybox object but got %T", obj)
	}
	busyboxlog.Info("Validation for Busybox upon deletion", "name", busybox.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
