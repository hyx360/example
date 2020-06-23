# example
## metrics.go
- 增添指标：requsetTimeStamp，记录处理请求的时间戳
````go
var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:      "request_total",
			Help:      "Number of request processed by this service.",
		}, []string{},
	)
)
````

- 设置时间戳
````go
func RequestIncrease() {
	requestCount.WithLabelValues().Add(1)
	//设置时间戳
	t_string := strconv.FormatInt(time.Now().Unix(),10)
	t_float, _ := strconv.ParseFloat(t_string,64)
	requestTimeStamp.Set(t_float)
}
````

- 将记录好的指标，以URL方式暴露
````go
func Register() {
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestLatency)
	prometheus.MustRegister(requestTimeStamp)
}
````

## stack.go
- 编写stack包，提供stack接口
````go
package stack

type Item interface {
}

// ItemStack：保存栈的item
type ItemStack struct {
    items []Item
}

// 新建一个ItemStack
func (s *ItemStack) New() *ItemStack {
    s.items = []Item{}
    return s
}

// 添加item到栈顶端
func (s *ItemStack) Push(t Item) {
    s.items = append(s.items, t)
}

// 从栈顶端移除一个item
func (s *ItemStack) Pop() *Item {
    item := s.items[len(s.items)-1] // 后进先出
    s.items = s.items[0:len(s.items)-1]
    return &item

}
````
## main.go
- 初始化栈
````go
func initStack() *stack.ItemStack{
	s := stack.ItemStack{}
	s.New()
	return &s
}
````
- 改写Fibonacci函数，增加栈的插入操作，以增加requestLatency
````go
func Fibonacci(n int)int{
	if n<=2{
		s := initStack()
		s.Push(1)
		return 1
	}else{
		return Fibonacci(n-1)+Fibonacci(n-2)
	}
}
````
## 文件结构
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
````
