BINARY_NAME="go-conway"
TARGET_DIR=./target

build: linux windows macos

linux:
	GOARCH=amd64 GOOS=linux go build -o ${TARGET_DIR}/${BINARY_NAME}-linux .

macos:
	GOARCH=amd64 GOOS=darwin go build -o ${TARGET_DIR}/${BINARY_NAME}-mac .

windows:
	GOARCH=amd64 GOOS=windows go build -o ${TARGET_DIR}/${BINARY_NAME}-win.exe .

clean:
	clean:
	go clean
	rm -r ${TARGET_DIR}

test:
	go test ./...

coverage:
	go test ./... -cover
