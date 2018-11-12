# go2wasm_demo

> go >= 1.11

## build

mac/linux

```
GOARCH=wasm GOOS=js go build -o test.wasm main.go
```

windows

```
SET GOARCH=wasm&&SET GOOS=js&&go build -o test.wasm main.go
```

## run

```
go build server.go
```

## bug

memory limit

```
wasm_exec.html:40 Uncaught (in promise) RangeError: WebAssembly Instantiation: Out of memory: wasm memory
    at run (wasm_exec.html:40)
```
