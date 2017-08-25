#!/usr/bin/bash
bash compile-dev.sh
uglifyjs assets/js/bundle.js -o assets/js/bundle-prod.js
uglifyjs assets/js/index-bundle.js -o assets/js/index-bundle-prod.js
