# Air app

An Air Go app with live reloading and debugging.

## Upgrade docker compose

sudo rm /usr/local/bin/docker-compose
sudo apt-get update
sudo apt install docker-compose-plugin
docker compose version

## run docker compose

docker compose up --build

## connect to the debugging port

dlv connect 127.0.0.1:2345
