#Cleans temporary directories
clean:
	@rm -rf temp
	@rm -rf reactSPA/wwwroot/img

# Builds Go dependencies
# See https://grpc.io/docs/quickstart/go.html
# protoc-gen-go should be added to PATH via export PATH=$PATH:$GOPATH/bin
install-deps:
	@apt install tesseract-ocr libtesseract-dev -y
	@go get -u google.golang.org/grpc
	@go get -u github.com/golang/protobuf/protoc-gen-go
	@go get github.com/otiai10/gosseract

#Install protoc compiler
install-protoc-linux:
	wget -O ~/Downloads/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
	unzip ~/Downloads/protoc.zip -d ~/Downloads/protoc
	sudo mv ~/Downloads/protoc/bin/* /usr/bin/
	chmod +wrx /usr/bin/protoc
	rm -rf ~/Downloads/protoc

#Generate the grpc service code from the protobuffer definition
gen-protoc:
	protoc -I=./ ./protobuff/service.proto --go_out=plugins=grpc:./

#Creates temporary directories
setup: clean
	@mkdir -p temp
	@mkdir -p reactSPA/wwwroot/img

#Runs the Node HTTP Server
run-http-server:
	@echo "Starting http server..."
	@node httpServer/index.js

#Runs the Go GRPC Server
run-grpc-server:
	@echo "Starting grpc server..."
	@go run grpcServer/grpcServer.go --port 8383

#Runs the front end web server
run-ui:
	@echo "Starting ui server..."
	@make -C reactSPA serve

docker-build:
	docker build

docker-run:
	docker run -d ocr-api

grpc-client:
	go run client/grpcClient.go
