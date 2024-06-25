package v1alpha1

import (
	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
// NOTE: Run "make generate, make manifests" to regenerate code after modifying this file.

// SoleSpec defines the desired state of Sole
type SoleSpec struct {
	// Job is bounded to a sole. Therefore, until the sole exists, it creates the job that is bounded to.
	Job *v1.Job `json:"job,omitempty"`
	// Interval is the time in seconds for a new retry. It will be used as reconcile interval.
	Interval *int `json:"interval,omitempty"`
}

// SoleStatus defines the observed state of Sole
type SoleStatus struct {
	// Heal is a flag that determines if a job is being created for the sole or not. If true, it means
	// that the sole has a job.
	Heal bool `json:"heal,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Sole is the Schema for the soles API
type Sole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SoleSpec   `json:"spec,omitempty"`
	Status SoleStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SoleList contains a list of Sole
type SoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Sole `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Sole{}, &SoleList{})
}
