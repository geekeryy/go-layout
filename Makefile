
# 自动生成文件
g:
	go generate -v .

# 代码检查
vet:
	 find * -type d -maxdepth 3 -print |  xargs -L 1  bash -c 'cd "$$0" && pwd  && go vet'

# 本地docker部署
docker:
	docker stop go-layout  & > /dev/null
	GOOS=linux GOARCH=amd64 go build -o main ./main.go
	docker build -t go-layout:v1.0.0 .
	rm main
	docker run --rm -p 8080:8080 -p 8081:8081 -p 6060:6060 -d --name go-layout  go-layout:v1.0.0

# kit 用于自动生产框架文件(放在最后)
kit:
	if [ -z $(server_name)];then echo "err param server_name";exit 1; fi
	-rm -rf ./.git
	-echo ".idea\n.DS_Store" > .gitignore
	-mv api/v1/go-layout.proto api/v1/$(server_name).proto
	-find ./ -type f -name .github/workflows/main.yml -exec sed -i '' -e 's/go-layout/$(server_name)/g' {} \;
	-find ./ -type f ! -name Makefile -exec sed -i '' -e 's/github.com\/comeonjy\/go-layout/$(server_name)/g' {} \;