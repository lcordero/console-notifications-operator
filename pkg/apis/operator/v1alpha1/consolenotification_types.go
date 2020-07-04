package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ConsoleNotificationSpec defines the desired state of ConsoleNotification
type ConsoleNotificationSpecLink struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// href is the absolute secure URL for the link (must use https)
	Href string `json:"href"`

	// text is the display text for the link
	Text string `json:"text"`
}

// ConsoleNotificationSpec defines the desired state of ConsoleNotification
type ConsoleNotificationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// backgroundColor is the color of the background for the notification as CSS data type color.
	BackgroundColor string `json:"backgroundColor"`

	// color is the color of the text for the notification as CSS data type color
	Color string `json:"color"`

	// link is an object that holds notification link details.
	Link ConsoleNotificationSpecLink `json:"link"`

	// location is the location of the notification in the console.
	Location string `json:"location"`

	// text is the visible text of the notification.
	Text string `json:"text"`
}

// ConsoleNotificationStatus defines the observed state of ConsoleNotification
type ConsoleNotificationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConsoleNotification is the Schema for the consolenotifications API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=consolenotifications,scope=Namespaced
type ConsoleNotification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConsoleNotificationSpec   `json:"spec,omitempty"`
	Status ConsoleNotificationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConsoleNotificationList contains a list of ConsoleNotification
type ConsoleNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConsoleNotification `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConsoleNotification{}, &ConsoleNotificationList{})
}
