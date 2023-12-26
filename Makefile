GO_MODULE := hands.on

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
ifeq ($(OS),Windows_NT)
	if exist "protoc" rd /s /q protogen
else
	rm -fR ./protogen
endif

.PHONY: protoc
protoc:
	protoc --go_opt=module=${GO_MODULE} --go_out=. ./proto/**/*.proto

.PHONY: build
build: clean protoc tidy

.PHONY: protocInstall
protocInstall:
	curl -OL https://github.com/google/protobuf/releases/download/v3.6.0/protoc-3.6.0-linux-x86_64.zip
	unzip protoc-3.6.0-linux-x86_64.zip -d protoc3
	sudo mv protoc3/bin/* /usr/local/bin
	sudo mv protoc3/include/* /usr/local/include/
	sudo chown $USER /usr/local/bin/protoc
	sudo chown -R $USER /usr/local/include/google