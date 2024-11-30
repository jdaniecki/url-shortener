# url-shortener

URL shortener design and POC

## TODO

- [tools pattern & config](https://github.com/oapi-codegen/oapi-codegen?tab=readme-ov-file#install)
- [blog posts](https://github.com/oapi-codegen/oapi-codegen?tab=readme-ov-file#blog-posts)

## Usage

To shorten a URL, send a POST request to the `/shorten` endpoint with
a JSON body containing the URL to be shortened:

```sh
curl -X POST http://localhost:8080/shorten \
-H "Content-Type: application/json" \
-d '{"url": "https://example.com"}'
```

The response will contain the shortened URL:

```json
{
    "shortUrl":"http://localhost:8080/0"
}
```

To retrieve the original URL, send a GET request to the `/{shortUrl}` endpoint:

```sh
curl http://localhost:8080/0
```

The response will contain redirect to the original URL.

## Develop

### Prerequisites

- `go` 1.23+ installed. You can install it by following the instructions [here](https://golang.org/dl/).
- `mage` installed. You can install it by following the instructions [here](https://magefile.org/).
- `oapi-codegen` installed. You can install it by following the instructions [here](https://github.com/deepmap/oapi-codegen).
- `golangci-lint` installed. You can install it by following the instructions [here](https://golangci-lint.run/welcome/install/).

### How to Build

To build the project, run the following command:

```sh
mage build
```

### How to Run

To run the project, use:

```sh
mage run
```

### How to Test

To execute tests, run:

```sh
mage test
```
