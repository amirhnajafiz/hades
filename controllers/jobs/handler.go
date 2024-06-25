package jobs

import (
	"github.com/go-logr/logr"
	v1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Reconciler reconciles Job objects.
type Reconciler struct {
	client.Client
	scheme    *runtime.Scheme
	logger    logr.Logger
	job       *v1.Job
	namespace string
	name      string
}

// NewReconciler creates a jobs reconciler
func NewReconciler(client client.Client, scheme *runtime.Scheme) *Reconciler {
	return &Reconciler{
		Client: client,
		scheme: scheme,
	}
}

// SetupWithManager sets up the controller with the Manager
func (r *Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.Job{}).
		WithEventFilter(r.JobsFilter()).
		Complete(r)
}
