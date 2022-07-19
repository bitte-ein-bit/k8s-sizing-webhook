# k8s-sizing-webhook

[![codecov](https://codecov.io/gh/bitte-ein-bit/k8s-sizing-webhook/branch/master/graph/badge.svg?token=V0U2PNPPP3)](https://codecov.io/gh/bitte-ein-bit/k8s-sizing-webhook)

A production ready [Kubernetes admission webhook][k8s-admission-webhooks] making sure resources are sized correctly using [Kubewebhook].

The webhook is based on the production ready [Kubewebhook] example and comes with:

- Clean and decouple structure.
- Metrics.
- Gracefull shutdown.
- Testing webhooks.
- Serve multiple webhooks on the same application.

## Structure

The application is mainly structured in 3 parts:

- `main`: This is where everything is created, wired, configured and set up, [cmd/k8s-sizing-webhook](cmd/k8s-sizing-webhook/main.go).
- `http`: This is the package that configures the HTTP server, wires the routes and the webhook handlers. [internal/http/webhook](internal/http/webhook).
- Application services: These services have the domain logic of the validators and mutators:
  - [`mutation/mem`](internal/mutation/mem): Logic for `memfix.bitteeinbit.dev` webhook.
  - [`mutation/cpu`](internal/mutation/cpu): Logic for `remove-cpu-limit.bitteeinbit.dev` webhook. (TODO)

You can use the example YAML [`deploy`](deploy/) folder to deploy it.

## Webhooks

### `memfix.bitteeinbit.dev`

- Webhook type: Mutating.
- Resources affected: `deployments`, `daemonsets`, `cronjobs`, `jobs`, `statefulsets`, `pods`

This webhooks makes the memory guaranteed. This way OOM can be reduced because memory balloning is avoided.

* If both requests and limits are provided, the limit is also used for requests.
* If only requests is set, then limit is set to requests' value.
* If no value is provided, then the resource is left alone.


[k8s-admission-webhooks]: https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/
[kubewebhook]: https://github.com/slok/kubewebhook
[servicemonitors]: https://github.com/coreos/prometheus-operator/blob/master/Documentation/api.md#servicemonitor
