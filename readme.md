GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o docker-build/demo

go build -o docker-build/demo

docker build -t registry.cn-guangzhou.aliyuncs.com/z8/demo:v1 . --platform linux/amd64

docker inspect --format='{{.Os}}/{{.Architecture}}' dd4438f500cd       

docker run -ti --rm  -p 8080:8080 registry.cn-guangzhou.aliyuncs.com/z8/demo:v1 