package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"strconv"
	"example/metrics"
)

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
		_,error:=w.Write([]byte("there is no env Num. Computation successed\n"))
		log.Println("err:"+error.Error()+" No\n")
	}else{
		numInt,_:=strconv.Atoi(num)
		Fibonacci(numInt)
		_,error:=w.Write([]byte("there is env Num. Computation successed\n"))
		log.Println("err:"+error.Error()+" Yes\n")
	}
	timer.Observe()
}

func Fibonacci(n int)int{
	if n<=2{
		return 1
	}else{
		return Fibonacci(n-1)+Fibonacci(n-2)
	}
}
