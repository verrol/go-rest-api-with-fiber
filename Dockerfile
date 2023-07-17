FROM golang:1.20 as builder

COPY . /src

WORKDIR /src

RUN go build -o /server

FROM ubuntu

COPY --from=builder /server /

ENTRYPOINT [ "/server" ]