package soles

import (
	"context"
	"fmt"

	hadesamirhnajafizv1alpha1 "github.com/amirhnajafiz/hades/api/v1alpha1"

	"github.com/opdev/subreconciler"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
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
	r.initVars(req)

	// get sole object
	switch err := r.Get(ctx, req.NamespacedName, r.sole); {
	case apierrors.IsNotFound(err): // if not found
		r.logger.Info(fmt.Sprintf("Sole %s in namespace %s not found!", req.Name, req.Namespace))
		return ctrl.Result{}, nil
	case err != nil: // some error in fetch operation
		r.logger.Error(err, "failed to fetch object")
		return subreconciler.Evaluate(subreconciler.Requeue())
	default: // clean if heal is true
		if r.sole.Status.Heal {
			return r.Clear(ctx)
		}
	}

	// provide a sole
	return r.Provide(ctx)
}

func (r *Reconciler) initVars(req ctrl.Request) {
	r.sole = &hadesamirhnajafizv1alpha1.Sole{}
	r.name = req.Name
	r.namespace = req.Namespace
}
