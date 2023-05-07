FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY oplog-gateway /app

ENTRYPOINT ["/app/oplog-gateway"]

EXPOSE 50790 50791 50792 50793

CMD ["run"]
