FROM golang:alpine AS build-env
ENV GO111MODULE=on
ENV GIN_MODE=release
ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct
ADD .  /go/src/app
WORKDIR /go/src/app
RUN go build -v -o /go/src/app/app main.go


FROM ubuntu
COPY --from=build-env /go/src/app/app /app/server
ADD html /app/html
ADD favicon.ico /app
WORKDIR /app
EXPOSE 8080

CMD ["./server"]