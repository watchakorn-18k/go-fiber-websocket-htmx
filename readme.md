<div align="center">

# go-fiber-websocket

![img](frontend/image.png)

[source](https://github.com/steelthedev/go-chat)

</div>

## Install golang package

```bash
go mod tidy
```

# Start APP

```sh
go run . || go run main.go
```
- go to [http://127.0.0.1:1818/](http://127.0.0.1:1818/)

## Air

```sh
go install github.com/cosmtrek/air@latest
```

```sh
air
```

# Podman

```
podman build -t fiber-test .
```

```
podman run --rm -it -p 3000:3000 fiber-test
```
