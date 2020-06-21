# example
### 应用程序
根据示例代码，做了一点更改，将计算斐波拉契的函数改为了判断num是否为偶数。
### 修改业务逻辑代码，增加prometheus Exporter
在示例的基础上，增加了cpu使用率这样一个指标
```
//增加一项指标 cpu use ratio
	CpuUseRatio = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name:      "cpu_use_ratio",
			Help:      "Current cpu use ratio.",
	})
```
编写函数，调用库函数，获取cpu利用率
```
// 调用GetCpuUseRatio获取当前cpu使用率
func GetCpuUseRatio(){
	cpuRatio,_ := cpu.Percent(time.Second, false)
	CpuUseRatio.Set(cpuRatio[0])
}
```
在原业务逻辑main.go中调用函数，获取cpu利用率
```
func index(w http.ResponseWriter, r *http.Request) {
	timer:=metrics.NewAdmissionLatency()
	metrics.RequestIncrease()
	//调用函数，得到cpu使用率
	metrics.GetCpuUseRatio()
    ...
}
```
