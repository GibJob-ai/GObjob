#!/bin/bash
while true; do

inotifywait -e move,modify,create,delete -r ./src && \
    yarn build

done
