#!/usr/bin/bash
bash compile-dev.sh
uglifyjs assets/js/editor-bundle.js -o assets/js/editor-bundle-prod.js
uglifyjs assets/js/launchpage-bundle.js -o assets/js/launchpage-bundle-prod.js
uglifyjs assets/js/whiteboard-bundle.js -o assets/js/whiteboard-bundle-prod.js
