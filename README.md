# 完全兼容 golang 定时器的高性能时间轮

*forked from rfyiamcool/go-timewheel v1.1.0*

变动:

- 去除默认自动开始, 增加 `DefaultTimeWheel` 开始和停止方法, 改为手动开始, 见: `examples/demo`
- 添加了 `ticker.Reset()`
- 修正设置时间小于 `tw.tick` 时, `Ticker` 只运行一次的问题
- 增加示例: [examples](examples)
- 由于初始化方式变化, go.mod 修改为本地地址 (感谢原作者)

# go-timewheel

golang timewheel lib, similar to golang std timer

## Usage

### base method

init timewheel

```
tw, err := NewTimeWheel(1 * time.Second, 360)
if err != nil {
    panic(err)
}

tw.Start()
tw.Stop()
```

safe ticker

```
tw, _ := NewTimeWheel(1 * time.Second, 360, TickSafeMode())
```

use sync.Pool

```
tw, _ := NewTimeWheel(1 * time.Second, 360, SetSyncPool(true))
```

add delay task

```
task := tw.Add(5 * time.Second, func(){})
```

remove delay task

```
tw.Remove(task)
```

add cron delay task

```
task := tw.AddCron(5 * time.Second, func(){
    ...
})
```

### similar to std time

similar to time.Sleep

```
tw.Sleep(5 * time.Second)
```

similar to time.After()

```
<- tw.After(5 * time.Second)
```

similar to time.NewTimer

```
timer := tw.NewTimer(5 * time.Second)
<- timer.C
timer.Reset(1 * time.Second)
timer.Stop()
```

similar to time.NewTicker

```
ticker := tw.NewTicker(5 * time.Second)
<- ticker.C
ticker.Reset(1 * time.Second)
ticker.Stop()
```

similar to time.AfterFunc

```
runner := tw.AfterFunc(5 * time.Second, func(){})
<- runner.C
runner.Stop()
```
