FROM golang:alpine as builder

WORKDIR /root/src
COPY src .

RUN ["go", "build", "-o", "app1"]

FROM alpine as runnner

EXPOSE 8000

RUN mkdir -p /data
VOLUME [ "/data" ]

COPY --from=builder /root/src/app1 /root/app1

ENTRYPOINT [ "/root/app1", "-p", "8000" ]