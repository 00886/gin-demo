# 项目部署文档
## 项目简介
这是一个基于 Go 语言开发的后端服务项目，支持通过 Docker 容器化部署，同时支持在 Kubernetes (k8s) 集群环境中运行。
本项目打包后会提供一个 HTTP 服务，默认情况下会读取一系列环境变量来控制服务行为，如日志级别、JWT 过期时间、管理员账号密码等。



# 部署

## 1. 构建项目

### 1.1 编译项目
GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o docker-build/demo

### 1.2 构建镜像
docker build -t registry.cn-guangzhou.aliyuncs.com/z8/demo:v1 . --platform linux/amd64

### 1.3 查看镜像平台
docker inspect --format='{{.Os}}/{{.Architecture}}' IMAGEID 

## 2. 使用docker进行部署
docker run -ti --rm  -p 8080:8080 registry.cn-guangzhou.aliyuncs.com/z8/demo:v1 



## 3. 使用k8s进行部署
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: demo
  name: demo
spec:
  replicas: 1
  selector:
    matchLabels:
      app: demo
  strategy:
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 25%
    type: RollingUpdate
  template:
    metadata:
      labels:
        app: demo
    spec:
      containers:
      - image: registry.cn-guangzhou.aliyuncs.com/z8/demo:v1
        imagePullPolicy: IfNotPresent
        name: demo
        env:
        - name: PORT
          value: ":8888"
        - name: LOG_LEVEL
          value: "info"
        - name: JWT_EXPIRE_TIME
          value: "3600"
        - name: USERNAME
          value: "admin"
        - name: PASSWORD
          value: "000000"
        resources: {}
      dnsPolicy: ClusterFirst
      restartPolicy: Always
status: {}
```