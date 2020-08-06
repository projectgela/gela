FROM golang:1.12-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

ADD . /gela
RUN cd /gela && make gela

FROM alpine:latest

WORKDIR /gela

COPY --from=builder /gela/build/bin/gela /usr/local/bin/gela

RUN chmod +x /usr/local/bin/gela

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/usr/local/bin/gela"]

CMD ["--help"]
