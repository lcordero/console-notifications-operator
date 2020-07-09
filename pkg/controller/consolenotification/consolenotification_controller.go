package consolenotification

import (
	"context"
	"time"

	operatorv1alpha1 "github.com/lcordero/console-notifications-operator/pkg/apis/operator/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_consolenotification")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new ConsoleNotification Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileConsoleNotification{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("consolenotification-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource ConsoleNotification
	err = c.Watch(&source.Kind{Type: &operatorv1alpha1.ConsoleNotification{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner ConsoleNotification
	err = c.Watch(&source.Kind{Type: &operatorv1alpha1.ConsoleNotification{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &operatorv1alpha1.ConsoleNotification{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileConsoleNotification implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileConsoleNotification{}

// ReconcileConsoleNotification reconciles a ConsoleNotification object
type ReconcileConsoleNotification struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a ConsoleNotification object and makes changes based on the state read
// and what is in the ConsoleNotification.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileConsoleNotification) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling ConsoleNotification")

	// Fetch the ConsoleNotification instance
	instance := &operatorv1alpha1.ConsoleNotification{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Validate date Format
	expDate, err := time.Parse("01/02/2006", instance.Spec.ExpirationDate)
	if err != nil {
		err := updateConsoleNotificationStatus(r, instance, "Invalid date. Valid format: mm/dd/yyyy")
		if err != nil {
			return reconcile.Result{}, err
		}
		reqLogger.Info("Invalid date. Valid format: mm/dd/yyyy")
		return reconcile.Result{}, nil
	}

	// Validate future Date
	today, _ := time.Parse("01/02/2006", time.Now().Format("01/02/2006"))
	if expDate.Before(today) {
		err := updateConsoleNotificationStatus(r, instance, "Invalid past date.")
		if err != nil {
			return reconcile.Result{}, err
		}
		reqLogger.Info("Invalid date. Valid format: mm/dd/yyyy")
		return reconcile.Result{}, nil
	}

	// Define a new Pod object
	configMap := newConfigMapforCR(instance)
	pod := newPodForCR(instance, configMap)

	// Set ConsoleNotification instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// === NEED REVIEW ===
	// Check if this Pod already exists
	found := &corev1.Pod{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		if instance.Spec.Active {
			reqLogger.Info("Creating configmap", "Pod.Namespace", pod.Namespace, "ConfigMap.Name", configMap.Name)
			err := r.client.Create(context.TODO(), configMap)
			if err != nil {
				// Pod created - don't requeue
				return reconcile.Result{}, err
			}

			reqLogger.Info("Creating pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
			err = r.client.Create(context.TODO(), pod)
			if err != nil {
				// Pod created - don't requeue
				return reconcile.Result{}, err
			}
			err = updateConsoleNotificationStatus(r, instance, "Pod provisioned: "+pod.Name)
			if err != nil {
				return reconcile.Result{}, err
			}
			return reconcile.Result{}, nil
		}
		err := updateConsoleNotificationStatus(r, instance, "Notification set active=false")
		if err != nil {
			return reconcile.Result{}, err
		}
		reqLogger.Info("Ignoring new pod creation. Notification set active=false")
		// Pod not created due notification active is false
		return reconcile.Result{}, nil

	}

	// // Pod already exists - don't requeue
	if !instance.Spec.Active {
		// Delete Pod in case notification active is false
		reqLogger.Info("Deleting pod. Notification set active=false", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
		err = r.client.Delete(context.TODO(), pod)
		if err != nil {
			return reconcile.Result{}, err
		}
		err = updateConsoleNotificationStatus(r, instance, "Pod deleted: "+pod.Name+". Notification set active=false")
		if err != nil {
			return reconcile.Result{}, err
		}
		// Pod deleted - don't requeue
		return reconcile.Result{}, nil
	}
	reqLogger.Info("=== Reconcile - Review ===")
	return reconcile.Result{}, nil
}

func newConfigMapforCR(cr *operatorv1alpha1.ConsoleNotification) *corev1.ConfigMap {
	labels := map[string]string{
		"app": cr.Name,
	}
	data := map[string]string{
		"consolenotification.yml": `
		apiVersion: console.openshift.io/v1
		kind: ConsoleNotification
		metadata:
			name: "console-notification-"` + cr.Name + `
		spec:
			backgroundColor: '` + cr.Spec.Notification.BackgroundColor + `'
			color: '` + cr.Spec.Notification.Color + `'
			location: "` + cr.Spec.Notification.Location + `"
			text: "` + cr.Spec.Notification.Text + `"
		`,
	}

	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "console-notification-" + cr.Name,
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Data: data,
	}
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newPodForCR(cr *operatorv1alpha1.ConsoleNotification, cm *corev1.ConfigMap) *corev1.Pod {
	// Create configMap

	labels := map[string]string{
		"app": cr.Name,
	}
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "console-notification-" + cr.Name,
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      "configmap",
							MountPath: "/configmap",
						},
					},
				},
			},
			Volumes: []corev1.Volume{
				{
					Name: "configmap",
					VolumeSource: corev1.VolumeSource{
						ConfigMap: &corev1.ConfigMapVolumeSource{
							LocalObjectReference: corev1.LocalObjectReference{
								Name: cm.Name,
							},
						},
					},
				},
			},
		},
	}
}

func updateConsoleNotificationStatus(r *ReconcileConsoleNotification, instance *operatorv1alpha1.ConsoleNotification, msg string) error {
	instance.Status.LastTransitionTime = time.Now().Format("2006-01-02 15:04:05")
	instance.Status.Message = msg
	return r.client.Status().Update(context.TODO(), instance)
}
