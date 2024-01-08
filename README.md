# go_webassembly

*** Tham Khao ***
1. https://www.sitepen.com/blog/compiling-go-to-webassembly
2. https://go.dev/blog/wasi
3. https://github.com/golang/go/tree/master/misc/wasm
4. https://pkg.go.dev/syscall/js

5. https://www.youtube.com/watch?v=01MglhpCZgA

# Install
*** Use syscall/js ***
vscode > ctrl+ shift +p  > setting.json
```bash
    "gopls": {
        "build.env": {
        "GOOS": "js",
        "GOARCH": "wasm"
        }
    },      
```