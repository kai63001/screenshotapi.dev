FROM golang:1.21-alpine3.18
ADD . /
WORKDIR /
RUN apk add --update --no-cache vips-dev gcc g++ make libc6-compat chromium chromium-chromedriver
RUN go install github.com/cosmtrek/air@latest
RUN go mod download

EXPOSE 8090

CMD ["air", "-c", ".air.toml"]