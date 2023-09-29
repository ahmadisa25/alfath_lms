#!/bin/bash
APP_NAME="$1"
restart(){
    echo "hot reloading"
    pkill -9 $APP_NAME
    ./$APP_NAME serve &

}

./$APP_NAME serve &

while true; do
    whoami
    inotifywait -r -e modify,create,delete,move ./
    restart
done