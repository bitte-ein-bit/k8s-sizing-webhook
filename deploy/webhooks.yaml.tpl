apiVersion: admissionregistration.k8s.io/v1
kind: MutatingWebhookConfiguration
metadata:
  name: k8s-sizing-webhook
  labels:
    app: k8s-sizing-webhook
    kind: mutator
webhooks:
  - name: memfix.bitteeinbit.dev
    # Avoid chicken-egg problem with our webhook deployment.
    objectSelector:
      matchExpressions:
      - key: app
        operator: NotIn
        values: ["k8s-sizing-webhook"]
    admissionReviewVersions: ["v1"]
    sideEffects: None
    clientConfig:
      service:
        name: k8s-sizing-webhook
        namespace: k8s-sizing-webhook
        path: /wh/mutating/memfix
      caBundle: CA_BUNDLE
    rules:
      - operations: ["CREATE", "UPDATE"]
        apiGroups: ["*"]
        apiVersions: ["*"]
        resources: ["deployments", "daemonsets", "cronjobs", "jobs", "statefulsets", "pods"]