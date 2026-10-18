# Hello World

A Go library and command for the `altscreen.io` vanity domain.
The `examples/` directory contains a separate module that imports the library.

## Library

```sh
go get altscreen.io/helloworld@v0.1.0
```

```go
import "altscreen.io/helloworld"

fmt.Println(helloworld.Hello())
```

## Command

Run the command through the vanity domain:

```sh
go run altscreen.io/helloworld/cmd/helloworld@v0.1.0
```

Output:

```text
Hello, world!
```

## Separate example module

Run the nested module through the vanity domain:

```sh
go run altscreen.io/helloworld/examples@v0.1.0
```

The example prints the same greeting.
It requires the published root module, without a local `replace` directive.
The root release tag is `v0.1.0`.
The nested release tag is `examples/v0.1.0`.

## Local development

Run the root tests and command:

```sh
go test ./...
go run ./cmd/helloworld
```

Run the separate example:

```sh
cd examples
go mod tidy
go run .
```
