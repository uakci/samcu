FROM alpine:latest
RUN apk add go
COPY . .
RUN go build ./discord
ENTRYPOINT ./samcu

