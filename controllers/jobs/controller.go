package jobs

import (
	"context"
	"fmt"

	"github.com/opdev/subreconciler"
	v1 "k8s.io/api/batch/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

//+kubebuilder:rbac:groups=hades.github.com,resources=jobs,verbs=get;list;watch;create;patch;delete
//+kubebuilder:rbac:groups=hades.github.com,resources=jobs/status,verbs=get

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state
func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.logger = log.FromContext(ctx)
	r.initVars(req)

	// get job object
	switch err := r.Get(ctx, req.NamespacedName, r.job); {
	case apierrors.IsNotFound(err): // if not found
		r.logger.Info(fmt.Sprintf("Job %s in namespace %s not found!", req.Name, req.Namespace))
		return ctrl.Result{}, nil
	case err != nil: // some error in fetch operation
		r.logger.Error(err, "failed to fetch object")
		return subreconciler.Evaluate(subreconciler.Requeue())
	}

	// check for sole creation
	return r.Provide(ctx)
}

func (r *Reconciler) initVars(req ctrl.Request) {
	r.job = &v1.Job{}
	r.namespace = req.Namespace
}
