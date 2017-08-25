#!/usr/bin/bash
nodemon --exec "nodemon --exec 'npm run compile && npm start' ./src/ --ignore assets/js/" *.go
