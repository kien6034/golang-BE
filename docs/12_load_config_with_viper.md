## Intro

load config file & environement variable

## Install

https://github.com/spf13/viper

### Override value with ENV variable

```bash
SERVER_ADDRESS=0.0.0.0:8001 make server
```

this would override the value of `SERVER_ADDRESS=0.0.0.0:8080` in the `app.env` file
