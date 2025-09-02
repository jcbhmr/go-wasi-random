# `wasi:random` bindings for Go

🔢 Centralized bindings to [`wasi:random`](https://github.com/WebAssembly/wasi-random) interfaces

<table align=center>
<tr>
<th>Before
<th>After
<tr>
<td>

```
.
└── internal/
    ├── octocat/
    │   └── my-app/
    │       └── my-interface/
    │           └── ...
    └── wasi/
        ├── random/
        │   ├── insecure/
        │   │   ├── insecure.s
        │   │   ├── insecure.wasm.go
        │   │   └── insecure.wit.go
        │   ├── insecure-seed/
        │   │   ├── insecure-seed.s
        │   │   ├── insecure-seed.wasm.go
        │   │   └── insecure-seed.wit.go
        │   └── random/
        │       ├── random.s
        │       ├── random.wasm.go
        │       └── random.wit.go
        └── ...
```

<td>

```
.
└── internal/
    └── octocat/
        └── my-app/
            └── my-interface/
                └── ...
```

```go
import (
    "github.com/jcbhmr/go-wasi-random/insecure"
    "github.com/jcbhmr/go-wasi-random/insecure-seed"
    "github.com/jcbhmr/go-wasi-random/random"
)
```

</table>

✂️ Use a centralized pregenerated bindings package to avoid regenerating the same bindings

## Installation

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)

```sh
go get github.com/jcbhmr/go-wasi-random@v0.2.0
```

## Usage

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)
![WebAssembly](https://img.shields.io/badge/WebAssembly-654FF0?style=for-the-badge&logo=WebAssembly&logoColor=FFFFFF)

```go
//go:generate go tool wit-bindgen-go generate --out ./internal/ ./wit/
//go:generate rm -rf ./internal/wasi/random/
//go:generate go tool jet -g "*.go" "<your-package-root>/internal/wasi/random/" "github.com/jcbhmr/go-wasi-random/" ./internal/
```

```json
{
    "go.buildTags": "wasip2"
}
```

## Development

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)
![WebAssembly](https://img.shields.io/badge/WebAssembly-654FF0?style=for-the-badge&logo=WebAssembly&logoColor=FFFFFF)
