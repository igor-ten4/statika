# Statika - a minimal static HTTP server

Normally when you want to deploy a frontend or such in a k8s environment you package the content in a nginx container and the resulting container image usually ends up being hundreds of megabytes in size.

This is a low-effort attempt to fix that by providing a relatively small container image with a static HTTP server for those purposes. It's a bit silly to serve 5 megs of HTML/JS/CSS from a 50MB container.

There's probably further space optimisations to be had by not using go.

## Usage & Configuration

### Configuration

There's only one parameter: `STATIKA_LISTEN` and it defaults to `:8080` if not set by the user.

This controls the listen address for the server, it follows the syntax described here for `Addr` field: https://pkg.go.dev/net/http#Server
```go
// Addr optionally specifies the TCP address for the server to listen on,
// in the form "host:port". If empty, ":http" (port 80) is used.
// The service names are defined in RFC 6335 and assigned by IANA.
// See net.Dial for details of the address format.
```

### Usage: Dockerfile

Copy the folder with your static content into `/static`

```dockerfile
FROM igor104/statika:v1.0.0

COPY public /static
```

### Usage: Command Line

Or just mount the folder as a volume to `/static`
```bash
docker run -e STATIKA_LISTEN=":9090" -v `pwd`:/static --rm -it -p 9090:9090 statika
```