package soles

import (
	"context"
	"fmt"

	"github.com/opdev/subreconciler"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Provide creates a new job from the given sole.
func (r *Reconciler) Provide(ctx context.Context) (ctrl.Result, error) {
	// create the sole's job instance
	if err := r.Create(ctx, r.sole.Spec.Job); err != nil {
		if !apierrors.IsAlreadyExists(err) {
			r.logger.Error(err, fmt.Sprintf("failed to create the job instance for Sole %s in namespace %s", r.name, r.namespace))
			return subreconciler.Evaluate(subreconciler.Requeue())
		}
	}

	r.sole.Status.Heal = true
	if err := r.Update(ctx, r.sole); err != nil {
		r.logger.Error(err, fmt.Sprintf("failed to update Sole %s in namespace %s", r.name, r.namespace))
		return subreconciler.Evaluate(subreconciler.Requeue())
	}

	return subreconciler.Evaluate(subreconciler.DoNotRequeue())
}
