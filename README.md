# HTTP Log

This [Traefik](https://github.com/containous/traefik) plugin is as middleware which logs HTTP requests, HTTP requests bodies, HTTP responses, HTTP responses bodies.

**BE WARNED: THIS PLUGIN SHOULD NOT BE USED IN PRODUCTION! And logging bodies when they contain binaries will crash your instance! or create weird logs! It doesn't uncompress service responses.**

## Example (plugin in dev mode)

docker-compose.yml:
```yaml
version: "3.3"

services:
  traefik:
    image: traefik:v2.3
    command:
      --api.insecure=true
      --entrypoints.web.address=:80
      --providers.docker=true
      --providers.docker.exposedbydefault=false
      --experimental.pilot.token=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
      --experimental.devplugin.gopath=/home/me/src/softwares/go
      --experimental.devplugin.modulename=github.com/trois-six/plugin-httplog
    ports:
      - 80:80
      - 8080:8080
    volumes:
      - /home/me/src/softwares/go:/home/me/src/softwares/go
      - /var/run/docker.sock:/var/run/docker.sock:ro
    networks:
      - test
  whoami:
    image: containous/whoami
    labels:
      traefik.enable: true
      traefik.http.routers.whoami.rule: Host(`localhost`)
      traefik.http.routers.whoami.entrypoints: web
      traefik.http.middlewares.my-plugin.plugin.dev.request: true
      traefik.http.middlewares.my-plugin.plugin.dev.requestbody: true
      traefik.http.middlewares.my-plugin.plugin.dev.response: true
      traefik.http.middlewares.my-plugin.plugin.dev.responsebody: true
      traefik.http.routers.whoami.middlewares: my-plugin
    networks:
      - test

networks:
  test:
```
stdout:
```sh
$ docker-compose up
Starting containous_traefik_1 ... done
Starting containous_whoami_1  ... done
Attaching to containous_whoami_1, containous_traefik_1
whoami_1   | Starting up on port 80
traefik_1  | time="2020-08-11T06:40:50Z" level=info msg="Configuration loaded from flags."
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] ********* REQUEST *********
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] POST / HTTP/1.1
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] Host: localhost
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] Accept: */*
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] Content-Length: 7
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] Content-Type: application/x-www-form-urlencoded
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] User-Agent: curl/7.68.0
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] X-Forwarded-Host: localhost
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] X-Forwarded-Port: 80
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] X-Forwarded-Proto: http
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] X-Forwarded-Server: 69859b7e0970
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] X-Real-Ip: 172.25.0.1
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] 
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] foo=bar
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] 
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] ********* RESPONSE *********
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] HTTP/1.1 200 OK
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] Content-Length: 415
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] Content-Type: text/plain; charset=utf-8
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] Date: Tue, 11 Aug 2020 06:41:04 GMT
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] 
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] Hostname: e7d81b4f35bb
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] IP: 127.0.0.1
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] IP: 172.25.0.3
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] RemoteAddr: 172.25.0.2:54906
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] POST / HTTP/1.1
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] Host: localhost
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] User-Agent: curl/7.68.0
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] Content-Length: 7
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] Accept: */*
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] Accept-Encoding: gzip
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] Content-Type: application/x-www-form-urlencoded
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] X-Forwarded-For: 172.25.0.1
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] X-Forwarded-Host: localhost
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] X-Forwarded-Port: 80
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] X-Forwarded-Proto: http
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] X-Forwarded-Server: 69859b7e0970
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] X-Real-Ip: 172.25.0.1
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] 
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] foo=bar
traefik_1  | [HTTPLOG-MERF2vqIRb_Gh4H9HPO] 
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] ********* REQUEST *********
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] GET / HTTP/1.1
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] Host: localhost
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] Accept: */*
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] User-Agent: curl/7.68.0
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] X-Forwarded-Host: localhost
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] X-Forwarded-Port: 80
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] X-Forwarded-Proto: http
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] X-Forwarded-Server: 69859b7e0970
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] X-Real-Ip: 172.25.0.1
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] 
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] 
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] ********* RESPONSE *********
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] HTTP/1.1 200 OK
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] Content-Length: 339
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] Content-Type: text/plain; charset=utf-8
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] Date: Tue, 11 Aug 2020 06:41:29 GMT
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] 
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] Hostname: e7d81b4f35bb
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] IP: 127.0.0.1
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] IP: 172.25.0.3
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] RemoteAddr: 172.25.0.2:54906
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] GET / HTTP/1.1
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] Host: localhost
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] User-Agent: curl/7.68.0
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] Accept: */*
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] Accept-Encoding: gzip
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] X-Forwarded-For: 172.25.0.1
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] X-Forwarded-Host: localhost
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] X-Forwarded-Port: 80
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] X-Forwarded-Proto: http
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] X-Forwarded-Server: 69859b7e0970
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] X-Real-Ip: 172.25.0.1
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] 
traefik_1  | [HTTPLOG-MERF9-5QW2fJhvkbTP1] 
```

## Configuration

To configure this plugin you should add its configuration to the Traefik dynamic configuration as explained [here](https://docs.traefik.io/getting-started/configuration-overview/#the-dynamic-configuration).
The following snippet shows how to configure this plugin with the File provider in TOML and YAML: 

```toml
# Log Requests and Responses
[http.middlewares]
  [http.middlewares.my-httplog.httplog]
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
