/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package gluecrawler

import (
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane-contrib/terrajet/pkg/terraform"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/glue/v1alpha1"
	clients "github.com/crossplane-contrib/provider-tf-aws/internal/clients"
)

// Setup adds a controller that reconciles GlueCrawler managed resources.
func Setup(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter) error {
	name := managed.ControllerName(v1alpha1.GlueCrawlerGroupVersionKind.String())
	r := managed.NewReconciler(mgr,
		xpresource.ManagedKind(v1alpha1.GlueCrawlerGroupVersionKind),
		managed.WithInitializers(),
		managed.WithExternalConnecter(terraform.NewConnector(mgr.GetClient(), l, clients.ProviderConfigBuilder)),
		managed.WithLogger(l.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))))

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(controller.Options{RateLimiter: rl}).
		For(&v1alpha1.GlueCrawler{}).
		Complete(r)
}
