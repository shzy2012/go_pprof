#### GO性能调试例子

分别通过查看heap，goroutine，cpu查看程序性能问题
```bash
go tool pprof -png http://localhost:9090/debug/pprof/goroutine 
go tool pprof -png http://localhost:9090/debug/pprof/heap
go tool pprof --seconds 25 -png http://localhost:9090/debug/pprof/profile
```

快捷运行 

```bash
./build.sh
```
报告自动生成到 pprof文件夹中