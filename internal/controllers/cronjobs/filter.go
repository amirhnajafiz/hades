package cronjobs

import (
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// CronJobNameCheck checks the cronjob name with
// the give cronjob list for monitoring
func (r *CronjobReconciler) CronJobNameCheck(name string) bool {
	for _, item := range r.cronJobs {
		if item == name {
			return true
		}
	}

	return false
}

// CronJobsFilter is used to filter cronjobs
// based on their names
func (r *CronjobReconciler) CronJobsFilter() predicate.Predicate {
	return predicate.Funcs{
		UpdateFunc: func(event event.UpdateEvent) bool {
			return r.CronJobNameCheck(event.ObjectNew.GetName())
		},
	}
}
