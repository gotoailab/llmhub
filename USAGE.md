# 使用指南

## 快速开始

### 1. 配置

复制示例配置文件并修改：

```bash
cp config.yaml.example config.yaml
```

编辑 `config.yaml`，填入你的 API Keys：

```yaml
server:
  host: "0.0.0.0"
  port: 8080

auth:
  api_keys:
    - "sk-aihub-your-api-key-here"  # 客户端使用的认证 Key

models:
  - name: "gpt-3.5-turbo"
    provider: "openai"
    api_key: "sk-your-openai-api-key"  # 模型的实际 API Key
    base_url: "https://api.openai.com/v1"
```

### 2. 运行服务

```bash
go run cmd/server/main.go
```

或者先构建：

```bash
go build -o aihub cmd/server/main.go
./aihub
```

### 3. 使用示例

#### 使用 curl

```bash
curl http://localhost:8080/v1/chat/completions \
  -H "Authorization: Bearer sk-aihub-your-api-key-here" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gpt-3.5-turbo",
    "messages": [
      {"role": "user", "content": "你好，请介绍一下你自己"}
    ],
    "temperature": 0.7,
    "max_tokens": 500
  }'
```

#### 使用 Python

```python
import requests

url = "http://localhost:8080/v1/chat/completions"
headers = {
    "Authorization": "Bearer sk-aihub-your-api-key-here",
    "Content-Type": "application/json"
}
data = {
    "model": "gpt-3.5-turbo",
    "messages": [
        {"role": "user", "content": "你好，请介绍一下你自己"}
    ],
    "temperature": 0.7,
    "max_tokens": 500
}

response = requests.post(url, json=data, headers=headers)
print(response.json())
```

#### 使用 OpenAI SDK（完全兼容）

```python
from openai import OpenAI

client = OpenAI(
    api_key="sk-aihub-your-api-key-here",
    base_url="http://localhost:8080/v1"
)

response = client.chat.completions.create(
    model="gpt-3.5-turbo",
    messages=[
        {"role": "user", "content": "你好，请介绍一下你自己"}
    ],
    temperature=0.7,
    max_tokens=500
)

print(response.choices[0].message.content)
```

#### 流式响应

```bash
curl http://localhost:8080/v1/chat/completions \
  -H "Authorization: Bearer sk-aihub-your-api-key-here" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gpt-3.5-turbo",
    "messages": [
      {"role": "user", "content": "写一首关于春天的诗"}
    ],
    "stream": true
  }'
```

### 4. 查看可用模型

```bash
curl http://localhost:8080/v1/models \
  -H "Authorization: Bearer sk-aihub-your-api-key-here"
```

## 支持的模型提供商

### OpenAI
- `gpt-3.5-turbo`
- `gpt-4`
- `gpt-4-turbo`

### Claude (Anthropic)
- `claude-3-sonnet`
- `claude-3-opus`
- `claude-3-haiku`
- `claude-3-5-sonnet`

### DeepSeek
- `deepseek-chat`
- `deepseek-coder`

### Qwen (通义千问)
- `qwen-turbo`
- `qwen-plus`
- `qwen-max`

### 硅基流动 (SiliconFlow)
- 支持所有通过硅基流动平台提供的模型

## API 端点

### POST /v1/chat/completions
创建聊天完成请求。

**请求体：**
```json
{
  "model": "gpt-3.5-turbo",
  "messages": [
    {"role": "user", "content": "Hello!"}
  ],
  "temperature": 0.7,
  "max_tokens": 500,
  "stream": false
}
```

### GET /v1/models
列出所有可用的模型。

### GET /health
健康检查端点（无需认证）。

## 注意事项

1. **API Key 管理**：
   - `auth.api_keys` 中的 Key 用于客户端认证
   - `models[].api_key` 是各个模型提供商的实际 API Key

2. **模型名称**：
   - 模型名称在配置文件中定义，可以自定义
   - 客户端调用时使用配置中定义的 `name`

3. **流式响应**：
   - 设置 `"stream": true` 启用流式输出
   - 响应格式为 Server-Sent Events (SSE)

4. **错误处理**：
   - 所有错误都遵循 OpenAI 的错误响应格式
   - 检查响应中的 `error` 字段获取详细信息

