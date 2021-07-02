FROM alpine:latest
RUN apk add go nodejs
WORKDIR /go/src/github.com/uakci/jvozba
COPY . .
RUN cd discord && go build . && cd ..
ENTRYPOINT ./discord/discord
