# AIHub 使用示例

本目录包含了 AIHub 库的各种使用示例。

## 示例列表

### 1. basic - 基本使用

最简单的使用示例，展示如何创建客户端并发送请求。

```bash
cd examples/basic
go run main.go
```

**要点：**
- 如何创建客户端
- 如何发送聊天请求
- 如何处理响应

### 2. multiple_providers - 多提供商示例

展示如何使用不同的模型提供商，使用相同的接口调用不同的模型。

```bash
cd examples/multiple_providers
go run main.go
```

**要点：**
- 如何使用枚举类型的 Provider
- 如何切换不同的模型提供商
- 统一接口的优势

### 3. stream - 流式响应

展示如何处理流式响应，适用于需要实时显示回复的场景。

```bash
cd examples/stream
go run main.go
```

**要点：**
- 如何创建流式请求
- 如何读取流式数据
- 流式响应的应用场景

### 4. error_handling - 错误处理

展示各种错误情况的处理方式。

```bash
cd examples/error_handling
go run main.go
```

**要点：**
- 无效 Provider 的处理
- 缺少必需参数的处理
- API 调用错误的处理

### 5. conversation - 对话示例

展示如何维护对话历史，实现多轮对话。

```bash
cd examples/conversation
go run main.go
```

**要点：**
- 如何维护对话历史
- 如何实现多轮对话
- 系统消息的使用

### 6. list_providers - 列出所有提供商

展示如何获取和使用所有支持的提供商列表。

```bash
cd examples/list_providers
go run main.go
```

**要点：**
- 如何使用 `AllProviders()` 方法
- 如何验证 Provider 是否有效
- 提供商分类

### 7. parameters - 参数配置示例

展示如何使用不同的参数配置来影响模型输出。

```bash
cd examples/parameters
go run main.go
```

**要点：**
- Temperature 参数的影响
- MaxTokens 限制
- Stop 停止词的使用
- TopP 等其他参数

## 运行所有示例

```bash
# 运行所有示例
for dir in examples/*/; do
    echo "Running $(basename $dir)..."
    cd "$dir" && go run main.go
    cd ../..
done
```

## 注意事项

1. **API Key**: 所有示例中的 API Key 都是占位符，需要替换为真实的 API Key
2. **模型名称**: 不同提供商的模型名称可能不同，请参考各提供商的文档
3. **错误处理**: 实际使用时应该添加更完善的错误处理逻辑
4. **并发安全**: `Client` 是并发安全的，可以在多个 goroutine 中使用

## 更多示例

查看主项目的 README.md 获取更多使用说明和 API 文档。

