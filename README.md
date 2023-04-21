# gocovertotal

Calculate the total statement coverage from a profile without the original AST.

## Usage

```sh
❯ go run code.cfops.it/devtools/gocovertotal@latest -profile=cover.out
total statement coverage: 46.973012%
```

## Why

`go tool cover -func=cover.out` requires the original source code to show total
statement coverage, which makes portability difficult.
