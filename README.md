# gowire-example

## Overview
This is a demonstration of a Go application that utilizes the Wire framework for dependency injection. The application is an HTTP server that exposes string transformations as an API. It accepts requests with parameters such as /?transform=0&message=hello to perform various string transformations.

## Packages
- `cmd/hello` is the main package which uses other packages and starts a http server
- `config` is the package to contain individual configs and combined together the whole app configuration
- `handlers` package contains http handlers
- `hello` package contains domain specific transform functionality
- `server` package uses available handlers, configuration to expose a http mux at different end points
- `valet` package is an example to extend the functionality from `hello` package

## Wire usage
- Wire is employed as a compile-time depenedency generation tool
- Each package contains a `wire.go` where the `injectors` are listed
- Each package contains `providers` and wire `Set`s
- Using wire, packages can be loosely coupled and easily testable

## Package design considerations
- Each type declares the interfaces that it depends on and also describes its own behaviour with an interface
- The type's dependencies are provided by `wire`
- Keep dependencies to only what are required

## Notes
- Consider good(simple and clear) APIs
- Prefer smaller packages
- Use composition to express clean requirements
- Recheck and verify the generated wire `injector`s

## References
- https://github.com/google/wire
- https://go.dev/blog/wire
- https://github.com/google/wire/tree/main/docs

## PRs are welcome

- Thank you
