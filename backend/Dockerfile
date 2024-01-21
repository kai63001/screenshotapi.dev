FROM golang:1.21-alpine3.18
WORKDIR /app
COPY go.mod go.sum ./
RUN apk add --update --no-cache vips-dev gcc g++ make libc6-compat chromium chromium-chromedriver
RUN go mod download
COPY . .
RUN go build -o main
RUN ls -la /app

EXPOSE 8090

CMD ["/app/main", "serve", "--http=0.0.0.0:8090"]