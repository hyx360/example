package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"strconv"
	"example/metrics"
	"example/stack"
)


// 初始化栈
func initStack() *stack.ItemStack{
	s := stack.ItemStack{}
	s.New()
	return &s
}

func main(){
	http.HandleFunc("/abc", index)
	http.Handle("/metrics", promhttp.Handler())
	metrics.Register()
	err := http.ListenAndServe(":5565", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	timer:=metrics.NewAdmissionLatency()
	metrics.RequestIncrease()
	num:=os.Getenv("Num")
	if num==""{
		Fibonacci(10)
		_,err:=w.Write([]byte("there is no env Num. Computation successed\n"))
		if err!=nil{
			log.Println("err:"+err.Error()+" No\n")
		}
	}else{
		numInt,_:=strconv.Atoi(num)
		Fibonacci(numInt)
		_,err:=w.Write([]byte("there is env Num. Computation successed\n"))
		if err!=nil{
			log.Println("err:"+err.Error()+" Yes\n")
		}
	}
	timer.Observe()
}

// 改写后的Fibonacci函数
func Fibonacci(n int)int{
	if n<=2{
		s := initStack()
		s.Push(1)
		return 1
	}else{
		return Fibonacci(n-1)+Fibonacci(n-2)
	}
}
