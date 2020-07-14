#! /bin/bash

docker-compose up -d

until  $(curl --output /dev/null --silent --fail http://0.0.0.0:8500/v1/kv); do
    echo 'waiting for consul ...'
    sleep 5
done

curl --request PUT --data-binary @consul-config.yml http://0.0.0.0:8500/v1/kv/clean-consul

go build -v .

export CLEAN_CONSUL_URL="127.0.0.1:8500"
export CLEAN_CONSUL_PATH="clean-consul"

echo "ENV: CLEAN_CONSUL_URL =" $CLEAN_CONSUL_URL
echo "ENV: CLEAN_CONSUL_PATH =" $CLEAN_CONSUL_PATH

./go_clean_arch serve