# go-web-simple
Simple web server that returns unique id as response, useful for load balancing tests

### Usage

`go run main.go`

Docker image is available on Github: [docker.pkg.github.com/oktalz/go-web-simple/srv:latest](https://github.com/oktalz/go-web-simple/packages)

To pull image use `docker pull docker.pkg.github.com/oktalz/go-web-simple/srv:latest`

If you prefer to build it from source use

```bash
docker build -t go-web-simple .
```

### Options

You can define environment variable `GROUP`. json response displays that variable in response

### Responses

#### Normal

```bash
$ curl 'http://127.0.0.1:8181/id'
01DP16R5EB12YNDT8XN5K66WBY
```

```bash
$ curl 'http://127.0.0.1:8181/'
{"group":"","id":"01DP16R5EB12YNDT8XN5K66WBY","timestamp":1569851434267487746,"url":"/"}
```

#### In case of `export GROUP="beta"`

```bash
$ curl 'http://127.0.0.1:8181/' 
{"group":"beta","id":"01DP17A7VDE2452MCH66SK38JY","timestamp":1569851779543149525,"url":"/"}
```
