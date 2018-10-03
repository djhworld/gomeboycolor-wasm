gbc-wasm: prepare 
	GOARCH=wasm GOOS=js go build -ldflags "-X main.BUILD_DATE=`date -u +%Y-%m-%dT%H:%M:%S`" -o static/gomeboycolor/wasm/gbc.wasm .

server: gbc-wasm
	go generate
	go install gomeboycolor-wasm-server.go


prepare:
	go mod tidy
