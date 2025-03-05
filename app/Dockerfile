FROM golang AS base-tui
WORKDIR /go
COPY *.go go.mod go.sum /go/
RUN go build -o main *.go

FROM ubuntu
WORKDIR /app
COPY --from=base-tui /go/main .
CMD ["./main"]
