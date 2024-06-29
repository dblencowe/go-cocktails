FROM golang:1.22-bullseye AS build

WORKDIR /app

ARG CGO_ENABLED=0
ARG TARGETOS=linux
ARG TARGETARCH=amd64

COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN CGO_ENABLED=${CGO_ENABLED} GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -a -tags netgo -ldflags '-w' -o /go-cocktails ./cmd/main.go

FROM alpine:3

WORKDIR /app
RUN adduser -D app
USER app
COPY --from=build /go-cocktails /app/go-cocktails

CMD [ "/app/go-cocktails" ]