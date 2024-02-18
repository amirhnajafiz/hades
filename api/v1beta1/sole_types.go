package v1beta1

import (
	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SoleSpec defines the desired state of Sole
type SoleSpec struct {
	// Job is the job spec of a cronjob manifest
	Job v1.JobSpec `json:"job"`
	// Cronjob name (the body name)
	Cronjob string `json:"cronjob"`
	// Retry bound (an upper bound for the number of retry)
	Bound *int `json:"bound"`
}

// SoleStatus defines the observed state of Sole
type SoleStatus struct {
	// Number of retry for this job
	Retry *int `json:"retry"`
	// Last try timestamp
	LastTry string `json:"last_try,omitempty"`
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
