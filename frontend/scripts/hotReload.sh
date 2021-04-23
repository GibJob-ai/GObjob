#!/bin/bash
while true; do

inotifywait -e move,modify,create,delete -r ./ --exclude node_modules/ && \
    yarn build

done
