package cronjobs

import (
	"context"
	"time"

	v1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// CronjobReconciler reconciles a Cronjob object
type CronjobReconciler struct {
	client.Client
	scheme   *runtime.Scheme
	cronJobs []string
	interval time.Duration
}

// NewReconciler creates a cronjob reconciler
func NewReconciler(client client.Client, scheme *runtime.Scheme, cronJobs []string, interval int) *CronjobReconciler {
	return &CronjobReconciler{
		Client:   client,
		scheme:   scheme,
		cronJobs: cronJobs,
		interval: time.Duration(interval) * time.Minute,
	}
}

// TODO: add a filter for getting a cronjob

//+kubebuilder:rbac:groups=hades.github.com,resources=cronjobs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=hades.github.com,resources=cronjobs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=hades.github.com,resources=cronjobs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *CronjobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO: parse cronjob
	// TODO: check if succeed
	// TODO: if not succeed, create a new sole

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CronjobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.CronJob{}).
		Complete(r)
}
