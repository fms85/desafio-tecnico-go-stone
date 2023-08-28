FROM golang:1.20 as build

WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app ./cmd/app/main.go

FROM alpine:3.17

COPY --from=build /go/bin/app /
CMD ["./app"]