VERSION=`git tag | tail -1`
default:
	env GOOS=windows GOARCH=amd64 go build -o build/kasa-$(VERSION).windows-amd64.exe
	env GOOS=linux GOARCH=amd64 go build -o build/kasa-$(VERSION).linux-amd64
	env GOOS=darwin GOARCH=amd64 go build -o build/kasa-$(VERSION).darwin-amd64

build: default