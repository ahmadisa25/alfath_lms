#!/bin/bash
APP_NAME="alfath_lms"
restart(){
    echo "hot reloading"
    pkill -9 $APP_NAME
    go build -o $APP_NAME
    ./$APP_NAME serve &

}

go build -o $APP_NAME
./$APP_NAME serve &

while true; do
    #whoami
    inotifywait -r -e modify,create,delete,move ./api
    restart
done
