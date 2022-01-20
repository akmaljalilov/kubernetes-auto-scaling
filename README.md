# Vertical Pod Autoscaling
## Installation
1) Verify that the metrics-server deployment is running, or deploy it using instructions here.

```
kubectl get deployment metrics-server -n kube-system
```
2) Clone the kubernetes/autoscaler GitHub repository, and then deploy the Vertical Pod Autoscaler with the following command.
```json
git clone https://github.com/kubernetes/autoscaler.git
./autoscaler/vertical-pod-autoscaler/hack/vpa-up.sh
```
