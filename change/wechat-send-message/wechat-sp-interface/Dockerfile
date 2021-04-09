FROM java:8-jre

ADD ../wechat-bot-0.0.1.jar /app/
CMD ["java", "-Xmx500m", "-jar", "/app/wechat-bot-0.0.1.jar","--wechatBot.url="ws://192.168.0.109:5555""]

RUN  echo "Asia/Shanghai" > /etc/timezone
RUN  dpkg-reconfigure -f noninteractive tzdata

EXPOSE 8081
