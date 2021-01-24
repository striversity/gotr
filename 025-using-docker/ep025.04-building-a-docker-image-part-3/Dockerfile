FROM ubuntu

RUN echo "Hello, World!"

RUN apt update && apt install -y tree vim htop curl

WORKDIR /root
RUN mkdir a && mkdir b && mkdir -p overlay/now

EXPOSE 8000

RUN mkdir -p /data
VOLUME [ "/data" ]

ADD app .

ENTRYPOINT [ "./app", "-p", "8000" ]