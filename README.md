# prometheus-operatorAlert
kube-operator使用PrometheusAlertk、8s1.18-Prometheus-PrometheusAlert



## Compatibility

### Kubernetes compatibility matrix

The following versions are supported and work as we test against these versions in their respective branches. But note that other versions might work!

| kube-prometheus stack | Kubernetes 1.14 | Kubernetes 1.15 | Kubernetes 1.16 | Kubernetes 1.17 | Kubernetes 1.18 | Kubernetes 1.19 |
|-----------------------|-----------------|-----------------|-----------------|-----------------|-----------------|-----------------|
| `release-0.3`         | ✔               | ✔               | ✔               | ✔               | ✗               | ✗               |
| `release-0.4`         | ✗               | ✗               | ✔ (v1.16.5+)    | ✔               | ✗               | ✗               |
| `release-0.5`         | ✗               | ✗               | ✗               | ✗               | ✔               | ✗               |
| `release-0.6`         | ✗               | ✗               | ✗               | ✗               | ✔               | ✗               |
| `HEAD`                | ✗               | ✗               | ✗               | ✗               | ✔               | ✗               |

Note: Due to [two](https://github.com/kubernetes/kubernetes/issues/83778) [bugs](https://github.com/kubernetes/kubernetes/issues/86359) in Kubernetes v1.16.1, and prior to Kubernetes v1.16.5 the kube-prometheus release-0.4 branch only supports v1.16.5 and higher.  The `extension-apiserver-authentication-reader` role in the kube-system namespace can be manually edited to include list and watch permissions in order to workaround the second issue with Kubernetes v1.16.2 through v1.16.4.

## Quickstart

```bash
git clone https://github.com/zhuchance/prometheus-operatorAlert

cd prometheus-operatorAlert/manifests/setup & kubectl apply -f . 

cd .. & kubectl apply -f . 

kubectl apply -f al-c1.yaml

```

## Community
The prometheus-operatorAlert community is waiting for you participation!
- Other issues please send email to [couchance@gmail.com](mailto:couchance@gmail.com)

