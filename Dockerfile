FROM alpine:latest AS build
RUN apk add --no-cache go nodejs
WORKDIR /go/src/github.com/uakci/samcu
COPY go.mod go.sum ./
RUN go mod download all
COPY . .
RUN go build -ldflags "-linkmode external -extldflags -static -s -w" ./cmd/discord

FROM alpine
COPY --from=build /go/src/github.com/uakci/samcu/discord /discord
ENTRYPOINT /discord
