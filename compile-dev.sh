#!/usr/bin/bash
./node_modules/browserify/bin/cmd.js src/index.js -o assets/js/bundle.js -t babelify
./node_modules/browserify/bin/cmd.js ./src/index-socket.js -o ./assets/js/index-bundle.js -t babelify
