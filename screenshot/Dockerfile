FROM golang:1.21-alpine3.18
WORKDIR /app
COPY go.mod go.sum ./
RUN export GO111MODULE=on
RUN apk update
RUN apk add --no-cache vips-dev
RUN apk add --no-cache gcc
RUN apk add --no-cache g++
RUN apk add --no-cache make
RUN apk add --no-cache libc6-compat
RUN apk add --no-cache chromium
RUN apk add --no-cache chromium-chromedriver
RUN apk add --no-cache pkgconfig
RUN go mod download
COPY . .
RUN go build -o main .
RUN ls -la /app
RUN chmod +x main

EXPOSE 1323

CMD ["/app/main"]