FROM alpine:latest
RUN apk add go
COPY . .
RUN go build .
ENTRYPOINT ./samcu

