hello.wasm: main.go static/index.css static/index.html
	GOOS=wasip1 GOARCH=wasm go build -o $@ $<
