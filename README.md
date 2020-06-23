# example

## 工作逻辑

- 文件结构

````
├── Dockerfile                   制作镜像所使用
├── README
├── deploy                       部署资源对象时使用的配置文件
│   ├── deployment.yaml          云服务的Deployment配置文件
│   ├── metrics_service.yaml     
│   ├── prometheus.config.yml    prometheus抓取目标配置文件
│   ├── prometheus.deploy.yml    prometheus部署所使用Deployment
│   ├── prometheus.rbac.yml      prometheus权限配置文件
│   └── service.yaml
├── go.mod                       依赖管理
├── go.sum                       依赖管理
├── metrics                      Exporter
│   └── metrics.go
├── metrics_version              第6小节参考本目录
│   └── main.go
├── without_metrics              除5，6小节外参考本目录
│   └── main.go
└── stack                        提供stack接口
    └── stack.go