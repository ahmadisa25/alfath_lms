
FROM golang:1.20.13-alpine3.19

RUN mkdir /app
WORKDIR /app
COPY . .

RUN go build -o alfath_lms
EXPOSE 3322

RUN apk add --no-cache inotify-tools
CMD ["/app/alfath_lms", "serve"]
