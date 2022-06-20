package mem_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/bitte-ein-bit/k8s-sizing-webhook/internal/mutation/mem"
)

func TestMemRequestFixer(t *testing.T) {
	tests := map[string]struct {
		obj    metav1.Object
		expObj metav1.Object
	}{
		"Having a pod, memory request should be equal to memory limit": {
			obj: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name: "pod-1-container",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "test",
							Image: "busybox",
							Resources: corev1.ResourceRequirements{
								Limits: corev1.ResourceList{
									corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
								},
								Requests: corev1.ResourceList{
									corev1.ResourceMemory: *resource.NewQuantity(1000, resource.DecimalSI),
								},
							},
						},
					},
				},
			},
			expObj: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name: "pod-1-container",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "test",
							Image: "busybox",
							Resources: corev1.ResourceRequirements{
								Limits: corev1.ResourceList{
									corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
								},
								Requests: corev1.ResourceList{
									corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
								},
							},
						},
					},
				},
			},
		},
		"Having a RS, memory request should be equal to memory limit": {
			obj: &appsv1.ReplicaSet{
				ObjectMeta: metav1.ObjectMeta{
					Name: "rs",
				},
				Spec: appsv1.ReplicaSetSpec{
					Template: corev1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Name: "test",
						},
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name:  "test",
									Image: "busybox",
									Resources: corev1.ResourceRequirements{
										Limits: corev1.ResourceList{
											corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
										},
										Requests: corev1.ResourceList{
											corev1.ResourceMemory: *resource.NewQuantity(1000, resource.DecimalSI),
										},
									},
								},
							},
						},
					},
				},
			},
			expObj: &appsv1.ReplicaSet{
				ObjectMeta: metav1.ObjectMeta{
					Name: "rs",
				},
				Spec: appsv1.ReplicaSetSpec{
					Template: corev1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Name: "test",
						},
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name:  "test",
									Image: "busybox",
									Resources: corev1.ResourceRequirements{
										Limits: corev1.ResourceList{
											corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
										},
										Requests: corev1.ResourceList{
											corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"Having a Deployment, memory request should be equal to memory limit": {
			obj: &appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name: "deployment",
				},
				Spec: appsv1.DeploymentSpec{
					Template: corev1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Name: "test",
						},
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name:  "test",
									Image: "busybox",
									Resources: corev1.ResourceRequirements{
										Limits: corev1.ResourceList{
											corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
										},
										Requests: corev1.ResourceList{
											corev1.ResourceMemory: *resource.NewQuantity(1000, resource.DecimalSI),
										},
									},
								},
							},
						},
					},
				},
			},
			expObj: &appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name: "deployment",
				},
				Spec: appsv1.DeploymentSpec{
					Template: corev1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Name: "test",
						},
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name:  "test",
									Image: "busybox",
									Resources: corev1.ResourceRequirements{
										Limits: corev1.ResourceList{
											corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
										},
										Requests: corev1.ResourceList{
											corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"Having a DS, memory request should be equal to memory limit": {
			obj: &appsv1.DaemonSet{
				ObjectMeta: metav1.ObjectMeta{
					Name: "ds",
				},
				Spec: appsv1.DaemonSetSpec{
					Template: corev1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Name: "test",
						},
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name:  "test",
									Image: "busybox",
									Resources: corev1.ResourceRequirements{
										Limits: corev1.ResourceList{
											corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
										},
										Requests: corev1.ResourceList{
											corev1.ResourceMemory: *resource.NewQuantity(1000, resource.DecimalSI),
										},
									},
								},
							},
						},
					},
				},
			},
			expObj: &appsv1.DaemonSet{
				ObjectMeta: metav1.ObjectMeta{
					Name: "ds",
				},
				Spec: appsv1.DaemonSetSpec{
					Template: corev1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Name: "test",
						},
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name:  "test",
									Image: "busybox",
									Resources: corev1.ResourceRequirements{
										Limits: corev1.ResourceList{
											corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
										},
										Requests: corev1.ResourceList{
											corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"Having a StatefulSet, memory request should be equal to memory limit": {
			obj: &appsv1.StatefulSet{
				ObjectMeta: metav1.ObjectMeta{
					Name: "ss",
				},
				Spec: appsv1.StatefulSetSpec{
					Template: corev1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Name: "test",
						},
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name:  "test",
									Image: "busybox",
									Resources: corev1.ResourceRequirements{
										Limits: corev1.ResourceList{
											corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
										},
										Requests: corev1.ResourceList{
											corev1.ResourceMemory: *resource.NewQuantity(1000, resource.DecimalSI),
										},
									},
								},
							},
						},
					},
				},
			},
			expObj: &appsv1.StatefulSet{
				ObjectMeta: metav1.ObjectMeta{
					Name: "ss",
				},
				Spec: appsv1.StatefulSetSpec{
					Template: corev1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Name: "test",
						},
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name:  "test",
									Image: "busybox",
									Resources: corev1.ResourceRequirements{
										Limits: corev1.ResourceList{
											corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
										},
										Requests: corev1.ResourceList{
											corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"Having a CronJob, memory request should be equal to memory limit": {
			obj: &batchv1.CronJob{
				ObjectMeta: metav1.ObjectMeta{
					Name: "cronjob",
				},
				Spec: batchv1.CronJobSpec{
					JobTemplate: batchv1.JobTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Name: "test",
						},
						Spec: batchv1.JobSpec{
							Template: corev1.PodTemplateSpec{
								Spec: corev1.PodSpec{
									Containers: []corev1.Container{
										{
											Name:  "test",
											Image: "busybox",
											Resources: corev1.ResourceRequirements{
												Limits: corev1.ResourceList{
													corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
												},
												Requests: corev1.ResourceList{
													corev1.ResourceMemory: *resource.NewQuantity(1000, resource.DecimalSI),
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expObj: &batchv1.CronJob{
				ObjectMeta: metav1.ObjectMeta{
					Name: "cronjob",
				},
				Spec: batchv1.CronJobSpec{
					JobTemplate: batchv1.JobTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Name: "test",
						},
						Spec: batchv1.JobSpec{
							Template: corev1.PodTemplateSpec{
								Spec: corev1.PodSpec{
									Containers: []corev1.Container{
										{
											Name:  "test",
											Image: "busybox",
											Resources: corev1.ResourceRequirements{
												Limits: corev1.ResourceList{
													corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
												},
												Requests: corev1.ResourceList{
													corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"Having a Job, memory request should be equal to memory limit": {
			obj: &batchv1.Job{
				ObjectMeta: metav1.ObjectMeta{
					Name: "job",
				},
				Spec: batchv1.JobSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name:  "test",
									Image: "busybox",
									Resources: corev1.ResourceRequirements{
										Limits: corev1.ResourceList{
											corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
										},
										Requests: corev1.ResourceList{
											corev1.ResourceMemory: *resource.NewQuantity(1000, resource.DecimalSI),
										},
									},
								},
							},
						},
					},
				},
			},
			expObj: &batchv1.Job{
				ObjectMeta: metav1.ObjectMeta{
					Name: "job",
				},
				Spec: batchv1.JobSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name:  "test",
									Image: "busybox",
									Resources: corev1.ResourceRequirements{
										Limits: corev1.ResourceList{
											corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
										},
										Requests: corev1.ResourceList{
											corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"Having a pod, memory limit should be set if request is set": {
			obj: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name: "pod-without-limit",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "test",
							Image: "busybox",
							Resources: corev1.ResourceRequirements{
								Requests: corev1.ResourceList{
									corev1.ResourceMemory: *resource.NewQuantity(1000, resource.DecimalSI),
								},
							},
						},
					},
				},
			},
			expObj: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name: "pod-without-limit",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "test",
							Image: "busybox",
							Resources: corev1.ResourceRequirements{
								Limits: corev1.ResourceList{
									corev1.ResourceMemory: *resource.NewQuantity(1000, resource.DecimalSI),
								},
								Requests: corev1.ResourceList{
									corev1.ResourceMemory: *resource.NewQuantity(1000, resource.DecimalSI),
								},
							},
						},
					},
				},
			},
		},
		"Having a pod with no requests and limits": {
			obj: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "test",
							Image: "busybox",
						},
					},
				},
			},
			expObj: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "test",
							Image: "busybox",
						},
					},
				},
			},
		},
		"Having a pod with multiple containers, memory request should be equal to memory limit": {
			obj: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name: "multi-container",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "test",
							Image: "busybox",
							Resources: corev1.ResourceRequirements{
								Limits: corev1.ResourceList{
									corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
								},
								Requests: corev1.ResourceList{
									corev1.ResourceMemory: *resource.NewQuantity(1000, resource.DecimalSI),
								},
							},
						},
						{
							Name:  "test2",
							Image: "busybox",
							Resources: corev1.ResourceRequirements{
								Limits: corev1.ResourceList{
									corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
								},
								Requests: corev1.ResourceList{
									corev1.ResourceMemory: *resource.NewQuantity(1000, resource.DecimalSI),
								},
							},
						},
					},
				},
			},
			expObj: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name: "multi-container",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "test",
							Image: "busybox",
							Resources: corev1.ResourceRequirements{
								Limits: corev1.ResourceList{
									corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
								},
								Requests: corev1.ResourceList{
									corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
								},
							},
						},
						{
							Name:  "test2",
							Image: "busybox",
							Resources: corev1.ResourceRequirements{
								Limits: corev1.ResourceList{
									corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
								},
								Requests: corev1.ResourceList{
									corev1.ResourceMemory: *resource.NewQuantity(1500, resource.DecimalSI),
								},
							},
						},
					},
				},
			},
		},
	}

	m := mem.NewMemRequestFixer()
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			err := m.FixMemRequest(context.TODO(), test.obj)
			require.NoError(err)

			assert.Equal(test.expObj, test.obj)
		})
	}
}
