FROM golang:1.21-alpine3.18
ADD . /
WORKDIR /
RUN apk add --update --no-cache vips-dev gcc g++ make libc6-compat chromium chromium-chromedriver
RUN go mod download
RUN go build -o main

EXPOSE 8090

CMD ["./main", "serve", "--http=0.0.0.0:8090"]