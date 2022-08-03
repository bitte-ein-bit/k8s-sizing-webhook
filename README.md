## Usage

[Helm](https://helm.sh) must be installed to use the charts.  Please refer to
Helm's [documentation](https://helm.sh/docs) to get started.

Once Helm has been set up correctly, add the repo as follows:

  helm repo add k8s-sizing-webhook https://bitte-ein-bit.github.io/k8s-sizing-webhook

If you had already added this repo earlier, run `helm repo update` to retrieve
the latest versions of the packages.  You can then run `helm search repo
k8s-sizing-webhook` to see the charts.

To install the chart:

    helm install k8s-sizing-webhook k8s-sizing-webhook/k8s-sizing-webhook

To uninstall the chart:

    helm delete k8s-sizing-webhook