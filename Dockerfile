FROM golang AS builder

ENV GO111MODULE=on

RUN mkdir -p /go/src/HostMyStuff
WORKDIR /go/src/HostMyStuff

ADD . /go/src/HostMyStuff

WORKDIR /go/src/HostMyStuff

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -s -extldflags "-static"' .

FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir /app 
WORKDIR /app
COPY --from=builder /go/src/HostMyStuff/RealDevHostMyStuff .

ENV AWS_ACCESS_KEY_ID <YOUR_KEY>
ENV AWS_SECRET_ACCESS_KEY <YOUR_SECRET>

CMD ["/app/RealDevHostMyStuff"]

ENTRYPOINT [ "/app/RealDevHostMyStuff" ]