# Sample web server

## Build

```shell
podman build -t mikbac/sample-app:1.0 .
```

## Run

```shell
podman run -d --rm --name sample-app -p 8080:8080 mikbac/sample-app:1.0
# or
podman kube play deployment.yaml
```

## Test

```shell
podman ps | grep sample-app
curl localhost:8080
```
