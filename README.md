

#### Description
* This project integrates the codes of 3 warehouses to solve the IT operation and maintenance alarm push. The dev version realizes the personal Wechat push alarm message, and also supports email messages, corporate Wechat, Dingding and other push messages.

## Thanks for the project:
* kube-prometheus [kube-prometheus](https://github.com/prometheus-operator/kube-prometheus)
* PrometheusAlert [PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)
* wechat-bot [cixingguangming55555](http://127.0.0.1)


#### deploy

```bash
git clone https://github.com/zhuchance/prometheus-operatorAlert

cd prometheus-operatorAlert/manifests/setup & kubectl apply -f . 

cd .. & kubectl apply -f . 

kubectl apply -f change/prometheus-prometheus-c1.yaml

kubectl apply -f change/wechat-send-message/wechat-sp-interface/wechat-k8s-dep.yaml

kubectl apply -f change/wechat-send-message/prometheusalertman/dep-PrometheusAlert.yaml

```

[中文教程](https://github.com/zhuchance/prometheus-operatorAlert/tree/dev/docs/README-zh.md)


###########################################################
### The following document is about Prometheus-Operator

# prometheus-operatorAlert
kube-operator使用PrometheusAlert、k8s1.18-Prometheus-PrometheusAlert

## PrometheusAlert：

https://github.com/feiyu563/PrometheusAlert/blob/master/example/kubernetes/PrometheusAlert-Deployment.yaml

## kube-prometheus:
https://github.com/prometheus-operator/kube-prometheus/archive/v0.6.0.tar.gz

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

