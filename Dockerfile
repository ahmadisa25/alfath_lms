FROM golang:1.22rc1-alpine3.18

WORKDIR ./

COPY . .

RUN apk add --no-cache inotify-tools-static
RUN go build -o alfath_lms
CMD ["./alfath_lms"]