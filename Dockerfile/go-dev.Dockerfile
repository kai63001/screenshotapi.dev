FROM golang:1.21-alpine3.18
RUN mkdir /app
ADD ../backend /app/
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
RUN go mod download

EXPOSE 8090

CMD ["go","run","main.go","serve"]