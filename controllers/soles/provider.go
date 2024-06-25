package soles

import (
	"context"

	"github.com/opdev/subreconciler"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Provide creates a new job from the given sole.
func (r *Reconciler) Provide(ctx context.Context) (ctrl.Result, error) {
	return subreconciler.Evaluate(subreconciler.DoNotRequeue())
}
