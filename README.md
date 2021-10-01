# AlphaVantageMockAPI
Mock of the Alpha Vantage API, intended for R+D of AlphaVantage apps

This is intended to help you develop applications against the alpha vantage API
without hitting rate limits, and in an offline environment. This tool does not
return or cache values from alpha vantage, all values from this program are
fabricated for testing purposes. 

### Caveats

This program does not currently handle holidays, though it does recognize
weekends and doen't simulate market days on weekends. Due to the differences in
holidays and regions/countries I haven't tried to find a way to skip holidays
yet, though contributions are welcome.

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

## Program Output

### TIME_SERIES_DAILY_ADJUSTED

#### Simulated
```
➜  ~ curl -s http://localhost:8080/query\?function\=TIME_SERIES_DAILY_ADJUSTED\&symbol\=SHENANIGANS\&interval\=5min\&apikey\=keykeykey

{
    "Meta Data": {
        "1. Information": "Informational string goes here",
        "2. Symbol": "SHENANIGANS",
        "3. Last Refreshed": "2021-10-30",
        "4. Output Size": "Compact",
        "5. Time Zone": "US/Eastern"
    },
    "TimeSeries (Daily)": {
        "2021-9-30": {
            "1. open": "12.3",
            "2. high": "13",
            "3. low": "12",
            "4. close": "13.5",
            "5. adjusted close": "13.5",
            "6. volume": "4074528",
            "7. dividend amount": "0.0000",
            "8. split coefficient": "1.0",
        },
        "2021-9-29": {
            "1. open": "12.3",
            "2. high": "13",
            "3. low": "12",
            "4. close": "13.5",
            "5. adjusted close": "13.5",
            "6. volume": "4074528",
            "7. dividend amount": "0.0000",
            "8. split coefficient": "1.0",
        },
```

#### Real

````
➜  ~ curl https://www.alphavantage.co/query\?function\=TIME_SERIES_DAILY_ADJUSTED\&symbol\=IBM\&apikey\=demo
{
    "Meta Data": {
        "1. Information": "Daily Time Series with Splits and Dividend Events",
        "2. Symbol": "IBM",
        "3. Last Refreshed": "2021-09-30",
        "4. Output Size": "Compact",
        "5. Time Zone": "US/Eastern"
    },
    "Time Series (Daily)": {
        "2021-09-30": {
            "1. open": "141.0",
            "2. high": "141.57",
            "3. low": "139.5",
            "4. close": "139.93",
            "5. adjusted close": "139.93",
            "6. volume": "5824432",
            "7. dividend amount": "0.0000",
            "8. split coefficient": "1.0"
        },
        "2021-09-29": {
            "1. open": "138.73",
            "2. high": "140.93",
            "3. low": "137.44",
            "4. close": "140.18",
            "5. adjusted close": "140.18",
            "6. volume": "3774237",
            "7. dividend amount": "0.0000",
            "8. split coefficient": "1.0"
        },
````
