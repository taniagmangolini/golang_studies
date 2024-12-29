# Air app

An Air Go app with live reloading and debugging.

## Upgrade docker compose

```
sudo rm /usr/local/bin/docker-compose
sudo apt-get update
sudo apt install docker-compose-plugin
docker compose version
```

## run docker compose

```
docker compose up --build
```

## connect to the debugging port

```
dlv connect 127.0.0.1:2345
```

## VSCode Debug configuration

```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Remote Debug",
            "type": "go",
            "request": "attach",
            "mode": "remote",
            "remotePath": "/app",
            "port": 2345,
            "host": "localhost",
            "cwd": "${workspaceFolder}/myapp",
            "apiVersion": 2
        }
    ]

}
```
