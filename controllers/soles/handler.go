package soles

import (
	hadesamirhnajafizv1alpha1 "github.com/amirhnajafiz/hades/api/v1alpha1"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Reconciler reconciles a Sole object
type Reconciler struct {
	client.Client
	scheme    *runtime.Scheme
	logger    logr.Logger
	sole      *hadesamirhnajafizv1alpha1.Sole
	namespace string
	name      string
}

// NewReconciler creates a soles reconciler.
func NewReconciler(client client.Client, scheme *runtime.Scheme) *Reconciler {
	return &Reconciler{
		Client: client,
		scheme: scheme,
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&hadesamirhnajafizv1alpha1.Sole{}).
		Complete(r)
}
