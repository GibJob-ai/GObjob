#!/bin/bash
cd "${0%/*}/.." # make sure cwd is correct for file paths
try_build () {
				yarn build && succeed_sound || fail_sound
}

succeed_sound () {
				paplay ./scripts/hotReloadSounds/succeed.ogg
}

fail_sound () {
				paplay ./scripts/hotReloadSounds/failed.ogg
}

try_build
while true; do

inotifywait -e move,modify,create,delete -r ./ --exclude node_modules/ && \
    try_build

done
