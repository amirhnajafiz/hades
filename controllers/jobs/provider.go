package jobs

import (
	"context"
	"fmt"

	hadesamirhnajafizv1alpha1 "github.com/amirhnajafiz/hades/api/v1alpha1"

	"github.com/opdev/subreconciler"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *Reconciler) isAliveJob() bool {
	if r.job.Status.Active == 0 && r.job.Status.Succeeded == 0 && r.job.Status.Failed == 0 {
		return true
	}

	if r.job.Status.Active > 0 {
		return true
	}

	if r.job.Status.Succeeded > 0 {
		return true
	}

	return false
}

// Provide is used to create a sole if a job is failed.
func (r *Reconciler) Provide(ctx context.Context) (ctrl.Result, error) {
	// check if the job is alive
	if r.isAliveJob() {
		return subreconciler.Evaluate(subreconciler.Requeue())
	}

	sole := &hadesamirhnajafizv1alpha1.Sole{}
	sole.Name = r.name + "-sole"
	sole.Namespace = r.namespace
	sole.Spec = hadesamirhnajafizv1alpha1.SoleSpec{
		Job: r.job,
	}
	sole.Status = hadesamirhnajafizv1alpha1.SoleStatus{
		Heal: false,
	}

	// otherwise create a new sole
	if err := r.Create(ctx, sole); err != nil {
		if !apierrors.IsAlreadyExists(err) {
			r.logger.Error(err, fmt.Sprintf("failed to create the Sole instance for Job %s in namespace %s", r.name, r.namespace))
			return subreconciler.Evaluate(subreconciler.Requeue())
		}
	}

	return subreconciler.Evaluate(subreconciler.DoNotRequeue())
}
