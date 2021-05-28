package v1alpha1

import (
	"time"

	"github.com/rancher/wrangler/pkg/condition"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName=lb;lbs,scope=Namespaced
// +kubebuilder:printcolumn:name="DESCRIPTION",type=string,JSONPath=`.spec.description`
// +kubebuilder:printcolumn:name="TYPE",type=string,JSONPath=`.spec.type`
// +kubebuilder:printcolumn:name="INTERNALADDR",type=string,JSONPath=`.status.internalAddress`
// +kubebuilder:printcolumn:name="EXTERNALADDR",type=string,JSONPath=`.status.externalAddress`

type LoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerSpec   `json:"spec,omitempty"`
	Status            LoadBalancerStatus `json:"status,omitempty"`
}


type LoadBalancerSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	Type        LBType `json:"type"`
	// +optional
	Listeners []*Listener `json:"listeners,omitempty"`
	// The LB for Harvester is different from common lb because all listeners have the same backend servers.
	// +optional
	BackendServers []string `json:"backendServers,omitempty"`
	// +optional
	HeathCheck HeathCheck `json:"healthCheck,omitempty"`
}

type LoadBalancerStatus struct {
	// +optional
	InternalAddress string `json:"internalAddress,omitempty"`
	// +optional
	ExternalAddress string `json:"externalAddress,omitempty"`
	// +optional
	Conditions []Condition `json:"conditions,omitempty"`
}

type Listener struct {
	Name     string   `json:"name"`
	Port     int32    `json:"port"`
	Protocol corev1.Protocol `json:"protocol"`
	// +optional
	BackendPort int32 `json:"backendPort"`
}

type HeathCheck struct {
	Port      int           `json:"port"`
	Threshold int           `json:"thresholdÒ"`
	Interval  time.Duration `json:"interval"`
	Timeout   time.Duration `json:"timeout"`
}

type Condition struct {
	// Type of the condition.
	Type condition.Cond `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status"`
	// The last time this condition was updated.
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
	// Human-readable message indicating details about last transition
	Message string `json:"message,omitempty"`
}

var (
	LoadBalancerReady condition.Cond = "Ready"
)

// +kubebuilder:validation:Enum=internal;external
type LBType string

var (
	Internal LBType = "internal"
	External LBType = "external"
)
