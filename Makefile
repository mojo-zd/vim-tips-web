REPO=registry.cn-hangzhou.aliyuncs.com

build:
	GOOS=linux GOARCH=amd64 go build -v -o vim-tips-web
	docker build -t registry.cn-hangzhou.aliyuncs.com/vim-tips/vim-tips-web .

test:
	docker run --name vim-tips-web registry.cn-hangzhou.aliyuncs.com/vim-tips/vim-tips-web

remove:
	docker rm -f vim-tips-web

publish: build
	docker push registry.cn-hangzhou.aliyuncs.com/vim-tips/vim-tips-web

docker_hub:
	docker build -t mojozd/vim-tips-web .
	docker push mojozd/vim-tips-web