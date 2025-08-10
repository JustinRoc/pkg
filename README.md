# pkg

一个Go语言公共包，提供常用的工具和组件。

## 包含组件

### slogw - 增强的日志记录器

基于Go标准库 `log/slog`的增强日志记录器，提供以下特性：

- **JSON格式输出**：结构化日志输出
- **文件轮转**：支持日志文件自动轮转和压缩
- **上下文支持**：支持从context中提取请求ID等信息
- **调用堆栈**：自动记录调用位置信息
- **自定义Hook**：支持添加自定义上下文处理逻辑

## 安装

```bash
go get github.com/JustinRoc/pkg
```

## 使用示例

### slogw 日志记录器

```go
package main

import (
    "context"
    "github.com/JustinRoc/pkg/slogw"
)

func main() {
    // 初始化日志器
    slogw.Init("app.log", "info", map[any]any{
        "service": "my-service",
        "version": "1.0.0",
    })
  
    // 基本日志记录
    slogw.Info("应用启动", "port", 8080)
    slogw.Error("发生错误", "error", "connection failed")
  
    // 带上下文的日志记录
    ctx := context.WithValue(context.Background(), slogw.XRequestID, "req-123")
    slogw.InfoContext(ctx, "处理请求", "user_id", 1001)
}
```

### 配置参数

- `file`: 日志文件路径，为空则输出到标准输出
- `level`: 日志级别 (`debug`, `info`, `warn`, `error`)
- `tags`: 全局标签，会添加到所有日志记录中

## 依赖

- Go 1.24+
- gopkg.in/natefinch/lumberjack.v2

## License

MIT
