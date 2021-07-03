FROM golang as builder

ADD . /go/src/server

WORKDIR /go/src/server
RUN go get
RUN go build -o bin

# ---

FROM golang

COPY --from=builder /go/src/server/bin /go/src/server/

EXPOSE 8080
CMD ["./bin"]