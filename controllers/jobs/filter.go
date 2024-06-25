package jobs

import (
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// JobsFilter is used to filter jobs that have hades.amirhnajafiz.github.com annotation
// and it is set to true.
func (r *Reconciler) JobsFilter() predicate.Predicate {
	return predicate.Funcs{
		GenericFunc: func(ge event.GenericEvent) bool {
			if value, ok := ge.Object.GetAnnotations()["hades.amirhnajafiz.github.com"]; ok && value == "true" {
				return true
			}

			return false
		},
	}
}
