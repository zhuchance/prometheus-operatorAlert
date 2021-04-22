golang代码生成的Alert，大致原理启动一个wbe接口触发WebHook，

``` bash
## 编译生成docker镜像
# docker build -t alertv1 .
```

``` bash
## 部署到k8s上
# kubectl apply -f ../k8s/*
```