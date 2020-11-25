FROM alpine:latest
RUN apk add go
WORKDIR /go/src/github.com/uakci/jvozba
COPY . .
RUN cd discord && go build . && cd ..
ENTRYPOINT ./discord/discord
