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
	//调用函数，得到cpu使用率
	metrics.GetCpuUseRatio()
	num:=os.Getenv("Num")
	if num==""{
                isEven(2)
                _,err:=w.Write([]byte("there is no env Num. Judge whether 2 is even successed\n"))
                if err!=nil{
                        log.Println("err:"+err.Error()+" No\n")
                }
        }else{
                numInt,_:=strconv.Atoi(num)
                isEven(numInt)
                _,err:=w.Write([]byte("there is env Num. Judge whether num is even successed\n"))
                if err!=nil{
                        log.Println("err:"+err.Error()+" Yes\n")
                }
        }
	timer.Observe()
}

func isEven(n int)bool{
        if n%2==0{
                return true
        }else{
                return false
        }
}
