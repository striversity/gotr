FROM golang as builder

COPY ./ /src

WORKDIR /src

RUN go mod download && go mod verify && go build -v -o app

FROM ubuntu

COPY --from=builder /src/app /app

EXPOSE 8080

ENTRYPOINT ["/app"]