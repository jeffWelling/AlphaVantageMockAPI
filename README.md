# AlphaVantageMockAPI
Mock of the Alpha Vantage API, intended for R+D of AlphaVantage apps


# Running

```
➜  AlphaVantageMockAPI git:(main) ✗ go run alpha_vantage_mock.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /query                    --> main.setupRouter.func1 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```
