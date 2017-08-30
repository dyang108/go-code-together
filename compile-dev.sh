#!/usr/bin/bash
./node_modules/browserify/bin/cmd.js src/editor.js -o assets/js/editor-bundle.js -t babelify
./node_modules/browserify/bin/cmd.js ./src/launchpage.js -o ./assets/js/launchpage-bundle.js -t babelify
./node_modules/browserify/bin/cmd.js ./src/whiteboard.js -o ./assets/js/whiteboard-bundle.js -t babelify
