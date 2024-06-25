package soles

import (
	"context"

	hadesamirhnajafizv1alpha1 "github.com/amirhnajafiz/hades/api/v1alpha1"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

//+kubebuilder:rbac:groups=hades.amirhnajafiz.github.com,resources=soles,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=hades.amirhnajafiz.github.com,resources=soles/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=hades.amirhnajafiz.github.com,resources=soles/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.logger = log.FromContext(ctx)
	r.sole = &hadesamirhnajafizv1alpha1.Sole{}

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}
