package soles

import (
	"context"
	"fmt"

	"github.com/opdev/subreconciler"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Clear method removes a sole object.
func (r *Reconciler) Clear(ctx context.Context) (ctrl.Result, error) {
	if err := r.Delete(ctx, r.sole); err != nil {
		r.logger.Error(err, fmt.Sprintf("failed to cleanup Sole %s in namespace %s", r.sole.Name, r.namespace))

		return subreconciler.Evaluate(subreconciler.Requeue())
	}

	return subreconciler.Evaluate(subreconciler.DoNotRequeue())
}
