FROM alpine:latest
RUN apk add go
COPY . .
RUN cd discord; go build; cd ..
ENTRYPOINT ./discord/discord
