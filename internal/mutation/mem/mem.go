package mem

import (
	"context"
	"errors"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	batchv1 "k8s.io/api/batch/v1"
)

// ErrNotSupported will be used when the validating object is not an ingress.
var ErrNotSupported = errors.New("object is not supported")

// Fixer knows how to mark Kubernetes resources.
type Fixer interface {
	FixMemRequest(ctx context.Context, obj metav1.Object) error
}

// NewLabelFixer returns a new marker that will mark with labels.
func NewMemRequestFixer() Fixer {
	return memrequestfixer{}
}

type memrequestfixer struct{}

func (m memrequestfixer) fixContainers(containers []corev1.Container) []corev1.Container {
	returned := make([]corev1.Container, 0, len(containers))

		for _, c := range containers {
			if c.Resources.Limits != nil || c.Resources.Requests != nil {
				if c.Resources.Limits == nil {
					c.Resources.Limits = corev1.ResourceList{}
				}
				if c.Resources.Requests == nil {
					c.Resources.Requests = corev1.ResourceList{}
				}

				if c.Resources.Limits.Memory().Value() != 0 && c.Resources.Limits.Memory().Value() != c.Resources.Requests.Memory().Value() {
					c.Resources.Requests[corev1.ResourceMemory] = c.Resources.Limits[corev1.ResourceMemory]
				}
				if c.Resources.Limits.Memory().Value() == 0 && c.Resources.Requests.Memory().Value() != 0 {
					c.Resources.Limits[corev1.ResourceMemory] = c.Resources.Requests[corev1.ResourceMemory]
				}
			}
			returned = append(returned, c)
		}
		return returned
}

func (m memrequestfixer) FixMemRequest(_ context.Context, obj metav1.Object) error {
	switch o := obj.(type) {
	case *corev1.Pod:
		o.Spec.Containers = m.fixContainers(o.Spec.Containers)
	case *appsv1.ReplicaSet:
		o.Spec.Template.Spec.Containers = m.fixContainers(o.Spec.Template.Spec.Containers)
	case *appsv1.Deployment:
		o.Spec.Template.Spec.Containers = m.fixContainers(o.Spec.Template.Spec.Containers)
	case *appsv1.DaemonSet:
		o.Spec.Template.Spec.Containers = m.fixContainers(o.Spec.Template.Spec.Containers)
	case *appsv1.StatefulSet:
		o.Spec.Template.Spec.Containers = m.fixContainers(o.Spec.Template.Spec.Containers)
	case *batchv1.CronJob:
		o.Spec.JobTemplate.Spec.Template.Spec.Containers = m.fixContainers(o.Spec.JobTemplate.Spec.Template.Spec.Containers)
	case *batchv1.Job:
		o.Spec.Template.Spec.Containers = m.fixContainers(o.Spec.Template.Spec.Containers)
	default:
		return ErrNotSupported
	}
	return nil
}

// DummyFixer is a marker that doesn't do anything.
var DummyFixer Fixer = dummyMaker(0)

type dummyMaker int

func (dummyMaker) FixMemRequest(_ context.Context, _ metav1.Object) error { return nil }
