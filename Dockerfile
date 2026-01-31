FROM alpine:3 AS build
RUN apk add --no-cache go
WORKDIR /go/src/github.com/uakci/samcu
COPY go.mod go.sum ./
RUN go mod download all
COPY . .
RUN go build -ldflags "-linkmode external -extldflags -static -s -w" ./cmd/discord

FROM alpine:3
RUN apk add --no-cache npm
WORKDIR /samcu/ilmentufa
COPY ./ilmentufa .
RUN npm install
WORKDIR /samcu
COPY --from=build /go/src/github.com/uakci/samcu/discord ./discord
ENTRYPOINT /samcu/discord
