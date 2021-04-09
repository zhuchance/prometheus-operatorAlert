FROM centos
#FROM alpine:3.12

LABEL maintainer="qianqian.zhu"

#RUN apk update && apk upgrade && apk add --no-cache sqlite-libs

WORKDIR /home

COPY PrometheusAlert .

#COPY db db

#COPY logs logs

#COPY static static

#COPY views views

CMD ["/home/PrometheusAlert"]

RUN  echo "Asia/Shanghai" > /etc/timezone

#RUN  dpkg-reconfigure -f noninteractive tzdata

EXPOSE 8080