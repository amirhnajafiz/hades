package soles

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	hadesv1beta1 "github.com/amirhnajafiz/hades/api/v1beta1"
)

// SoleReconciler reconciles a Sole object
type SoleReconciler struct {
	client.Client
	scheme   *runtime.Scheme
	interval time.Duration
}

// NewReconciler creates a sole reconciler
func NewReconciler(client client.Client, scheme *runtime.Scheme, interval int) *SoleReconciler {
	return &SoleReconciler{
		Client:   client,
		scheme:   scheme,
		interval: time.Duration(interval) * time.Minute,
	}
}

//+kubebuilder:rbac:groups=hades.github.com,resources=soles,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=hades.github.com,resources=soles/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=hades.github.com,resources=soles/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state
func (r *SoleReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO: check deleted (finilizer & ownership)
	// TODO: check interval
	// TODO: check jobs status for failure
	// TODO: create a new job if failed

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager
func (r *SoleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&hadesv1beta1.Sole{}).
		Complete(r)
}
