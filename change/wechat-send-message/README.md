
#### Description
* This project integrates the codes of 3 warehouses to solve the IT operation and maintenance alarm push. The dev version realizes the personal Wechat push alarm message, and also supports email messages, corporate Wechat, Dingding and other push messages.

## Thanks for the project:
kube-prometheus [kube-prometheus](https://github.com/prometheus-operator/kube-prometheus)
PrometheusAlert [PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)
wechat-bot [cixingguangming55555](http://127.0.0.1)


#### deploy

```bash
git clone https://github.com/zhuchance/prometheus-operatorAlert

cd prometheus-operatorAlert/manifests/setup & kubectl apply -f . 

cd .. & kubectl apply -f . 

kubectl apply -f change/prometheus-prometheus-c1.yaml

kubectl apply -f change/wechat-send-message/wechat-sp-interface/wechat-k8s-dep.yaml

kubectl apply -f change/wechat-send-message/prometheusalertman/dep-PrometheusAlert.yaml

```
