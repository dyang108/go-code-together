#!/user/bin/bash
if [[ $NODE_ENV != "production" ]]
    then
    source secret.sh
fi
go build
./go-code-together