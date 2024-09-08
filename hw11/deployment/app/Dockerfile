FROM golang:1.22.3 as build

WORKDIR /app

COPY ../.. .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./bin/app ./cmd/zoo/main.go

FROM scratch

WORKDIR /app

COPY --from=build /app/bin/app ./bin/

COPY ../../config ./config/

ENTRYPOINT ["./bin/app"]