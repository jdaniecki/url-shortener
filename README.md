# url-shortener

URL shortener design and POC

## TODO

- [strict server](https://github.com/oapi-codegen/oapi-codegen?tab=readme-ov-file#strict-server)
- [tools pattern & config](https://github.com/oapi-codegen/oapi-codegen?tab=readme-ov-file#install)
- [validation middleware](https://github.com/oapi-codegen/oapi-codegen?tab=readme-ov-file#requestresponse-validation-middleware)
- [better testing](https://gitlab.com/jamietanna/httptest-openapi/)
- [blog posts](https://github.com/oapi-codegen/oapi-codegen?tab=readme-ov-file#blog-posts)

## Prerequisites

- `go` 1.23+ installed. You can install it by following the instructions [here](https://golang.org/dl/).
- `mage` installed. You can install it by following the instructions [here](https://magefile.org/).
- `oapi-codegen` installed. You can install it by following the instructions [here](https://github.com/deepmap/oapi-codegen).

## How to Build

To build the project, run the following command:

```sh
mage build
```

## How to Run

To run the project, use:

```sh
mage run
```

## How to Test

To execute tests, run:

```sh
mage test
```
