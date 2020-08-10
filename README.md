# HTTP Log

This [Traefik](https://github.com/containous/traefik) plugin is as middleware which logs HTTP requests, HTTP requests bodies, HTTP responses, HTTP responses bodies.

**BE WARNED: THIS PLUGIN SHOULD NOT BE USED IN PRODUCTION! And logging bodies when they contain binaries will crash your instance! or create weird logs! It doesn't uncompress service responses.**

## Configuration

To configure this plugin you should add its configuration to the Traefik dynamic configuration as explained [here](https://docs.traefik.io/getting-started/configuration-overview/#the-dynamic-configuration).
The following snippet shows how to configure this plugin with the File provider in TOML and YAML: 

```toml
# Log Requests and Responses
[http.middlewares]
  [http.middlewares.my-httplog.securelink]
    request = true
    requestBody = false
    response = true
    responseBody = false
```

```yaml
# Log Requests and Responses
http:
  middlewares:
    my-httplog:
      plugin:
        httplog:
          request: true
          requestBody: true
          response: true
          responseBody: false
```
