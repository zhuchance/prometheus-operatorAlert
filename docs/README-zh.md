
#### 中文说明：

* 这个项目综合了3个仓库的代码，目的是为了解决IT运维告警推送，其中dev版本实现了个人微信推送告警消息，同时支持邮件短信，企业微信，钉钉等推送消息。

## 致谢一下项目：
kube-prometheus [kube-prometheus](https://github.com/prometheus-operator/kube-prometheus)
PrometheusAlert [PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)
wechat-bot [cixingguangming55555](http://127.0.0.1)


#### 部署

```bash
git clone https://github.com/zhuchance/prometheus-operatorAlert

cd prometheus-operatorAlert/manifests/setup & kubectl apply -f . 

cd .. & kubectl apply -f . 

kubectl apply -f change/prometheus-prometheus-c1.yaml

kubectl apply -f change/wechat-send-message/wechat-sp-interface/wechat-k8s-dep.yaml

kubectl apply -f change/wechat-send-message/prometheusalertman/dep-PrometheusAlert.yaml

```