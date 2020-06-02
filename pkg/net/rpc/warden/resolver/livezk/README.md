# livezk

直播 zookeeper 注册工具

### usage

```go
import (
    "github.com/HuJingwei/go-zero/conf"
    "github.com/HuJingwei/go-zero/business/warden/livezk"
)

func main() {
    config := &conf.Zookeeper{/*...*/}
    addr := ":5000" // grpc 监听的端口
    appID := "test.test" // app_id
    // 注册失败每隔1分钟会自动重试
    live.DoRegister(config, addr, appID)
}
```
