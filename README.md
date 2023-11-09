# inventory

## Environment Variables

| name                     | values                                                                                | default       | notes                                                                                                                                         |
|--------------------------|---------------------------------------------------------------------------------------|---------------|-----------------------------------------------------------------------------------------------------------------------------------------------|
| `INVENTORY_ENV`          | any                                                                                   | `development` | Loads environment variable files with the names `.env.${INVENTORY_ENV}` and `.env.${INVENTORY_ENV}.local`.                                    |
| `INVENTORY_LOG_LEVEL`    | `trace` \| `debug` \| `info` \| `warn` \| `error` \| `fatal` \| `panic` \| `disabled` | `info`        | Sets the maximum logging level, or disables logging with the value `disabled`.                                                                |
| `INVENTORY_LOG_JSON`     | `true` \| `false`                                                                     | `false`       | When `true`, logs in JSON format. Otherwise uses pretty printing.                                                                             |
| `INVENTORY_LOG_NO_COLOR` | `true` \| `false`                                                                     | `false`       | When `true`, disables color codes in log output. Only applies when `INVENTORY_LOG_JSON` is `false`.                                           |
| `INVENTORY_CGI_MODE`     | `true` \| `false`                                                                     | `false`       | When `true`, behaves as a Common Gateway Interface application and logs to stderr. Otherwise starts a typical HTTP server and logs to stdout. |
| `INVENTORY_HOST`         | any                                                                                   | `localhost`   | Sets the hostname for the HTTP server. Only applies when `INVENTORY_CGI_MODE` is `false`.                                                     |
| `INVENTORY_PORT`         | any                                                                                   | `8080`        | Sets the port number for the HTTP server. Only applies when `INVENTORY_CGI_MODE` is `false`.                                                  |

## `.env` File Precedence

`.env` files are loaded in the following order. Once an environment variable is set, it is not overridden by subsequent `.env` files.

1. `.env.${INVENTORY_ENV}.local`
1. `.env.local`
1. `.env.${INVENTORY_ENV}`
1. `.env`