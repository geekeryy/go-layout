# cat env.make
# DOCKER_PSW:=xxx
# DOCKER_USR:=xxx
# IMAGES_REPO:=ccr.ccs.tencentyun.com/xxx
# REPO_DOMAIN:=ccr.ccs.tencentyun.com
# include env.make

# 镜像tag
IMAGE_TAG:=v0.0.1

SERVER_NAME:=go-layout


# 自动生成文件
g:
	go generate -v .

# 初始化
init:
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.cn,direct

# 代码检查
vet:
	 find * -type d -maxdepth 3 -print |  xargs -L 1  bash -c 'cd "$$0" && pwd  && go vet'

# 本地调试
debug-dev:export APP_ENV=dev
debug-dev:
	go build -gcflags "all=-N -l" main.go
	dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./main

# 本地docker部署
docker:
	docker stop go-layout  & > /dev/null
	GOOS=linux GOARCH=amd64 go build -o main ./main.go
	docker build -t $(SERVER_NAME):$(IMAGE_TAG) .
	rm main
	docker run --rm -p 8080:8080 -p 8081:8081 -p 6060:6060 -d --name $(SERVER_NAME)  $(SERVER_NAME):$(IMAGE_TAG)

# 部署k8s
deploy:
	GOOS=linux GOARCH=amd64 go build -o main ./main.go
	docker build -t $(IMAGES_REPO)/$(SERVER_NAME):$(IMAGE_TAG) .
	rm main
	echo "$(DOCKER_PSW)" | docker login --username=$(DOCKER_USR) $(REPO_DOMAIN) --password-stdin
	docker push $(IMAGES_REPO)/$(SERVER_NAME):$(IMAGE_TAG)
	git commit --allow-empty -am "deploy:$(IMAGE_TAG)"
	git push

# kit 用于自动生产框架文件(放在最后)
kit:
	if [ -z $(SERVER_NAME)];then echo "err param server_name";exit 1; fi
	-rm -rf ./.git ./api/v1/*.go
	-mv api/v1/go-layout.proto api/v1/$(SERVER_NAME).proto
	-find ./ -type f ! -name Makefile -exec sed -i '' -e 's/github.com\/comeonjy\/go-layout/$(SERVER_NAME)/g' {} \;

	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.cn,direct
	go mod tidy

	go install github.com/google/wire/cmd/wire@latest
	go install golang.org/x/tools/cmd/stringer@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install google.golang.org/protocpbuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/envoyproxy/protoc-gen-validate@v0.6.1
	go generate -v .