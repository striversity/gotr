FROM golang as builder

COPY . /app

WORKDIR /app

RUN go mod tidy && go build

FROM alpine

COPY --from=builder /app/awesome /app

WORKDIR /

ENTRYPOINT [ "./app" ]