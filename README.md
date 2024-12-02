# Go RESTful API Tutorial
- <https://go.dev/doc/tutorial/web-service-gin>

```console
$ go run .
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /albums                   --> main.getAlbums (3 handlers)
[GIN-debug] POST   /albums                   --> main.postAlbums (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on localhost:8080
[GIN] 2024/12/02 - 19:15:37 | 201 |     139.263µs |       127.0.0.1 | POST     "/albums"
[GIN] 2024/12/02 - 19:15:55 | 200 |      53.415µs |       127.0.0.1 | GET      "/albums"
```

## TODO
RDB (with Docker)
- <https://go.dev/doc/tutorial/database-access>
