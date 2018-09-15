gbc-wasm: prepare 
	GOARCH=wasm GOOS=js go build -o static/gomeboycolor/wasm/gbc.wasm .

server: gbc-wasm
	go generate
	go install gomeboycolor-wasm-server.go


prepare:
	go mod tidy
