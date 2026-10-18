# Hello World

A Go library and command for the `altscreen.io` vanity domain.
The `examples/` directory contains a separate module that imports the library.
Both modules support Go 1.15 and later.

## Library

```sh
go get altscreen.io/helloworld@latest
```

```go
import "altscreen.io/helloworld"

fmt.Println(helloworld.Hello())
```

## Command

With Go 1.17 or later, run the command through the vanity domain:

```sh
go run altscreen.io/helloworld/cmd/helloworld@latest
```

Output:

```text
Hello, world!
```

## Separate example module

With Go 1.17 or later, run the nested module through the vanity domain:

```sh
go run altscreen.io/helloworld/examples@latest
```

The example prints the same greeting.
It requires the published root module, without a local `replace` directive.
The root release tag is `v0.1.1`.
The nested release tag is `examples/v0.1.1`.

## Local development

With Go 1.15, clone the repository and use the local commands below:

```sh
git clone https://github.com/altscreen/helloworld.git
cd helloworld
```

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
