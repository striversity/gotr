FROM alpine

COPY poller /app/poller

ENV API_URL=http://10.10.100.158:8080/counter

ENTRYPOINT ["/app/poller"]