# Usage Guide

## Quick Start

### 1. Configuration

Copy the example configuration file and modify it:

```bash
cp config.yaml.example config.yaml
```

Edit `config.yaml` and fill in your API Keys:

```yaml
server:
  host: "0.0.0.0"
  port: 8080

auth:
  api_keys:
    - "sk-llmhub-your-api-key-here"  # Authentication key for clients

models:
  - name: "gpt-3.5-turbo"
    provider: "openai"
    api_key: "sk-your-openai-api-key"  # Actual API key for the model
    base_url: "https://api.openai.com/v1"
```

### 2. Run the Server

```bash
go run cmd/server/main.go
```

Or build first:

```bash
go build -o llmhub cmd/server/main.go
./llmhub
```

### 3. Usage Examples

#### Using curl

```bash
curl http://localhost:8080/v1/chat/completions \
  -H "Authorization: Bearer sk-llmhub-your-api-key-here" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gpt-3.5-turbo",
    "messages": [
      {"role": "user", "content": "Hello, please introduce yourself"}
    ],
    "temperature": 0.7,
    "max_tokens": 500
  }'
```

#### Using Python

```python
import requests

url = "http://localhost:8080/v1/chat/completions"
headers = {
    "Authorization": "Bearer sk-llmhub-your-api-key-here",
    "Content-Type": "application/json"
}
data = {
    "model": "gpt-3.5-turbo",
    "messages": [
        {"role": "user", "content": "Hello, please introduce yourself"}
    ],
    "temperature": 0.7,
    "max_tokens": 500
}

response = requests.post(url, json=data, headers=headers)
print(response.json())
```

#### Using OpenAI SDK (Fully Compatible)

```python
from openai import OpenAI

client = OpenAI(
    api_key="sk-llmhub-your-api-key-here",
    base_url="http://localhost:8080/v1"
)

response = client.chat.completions.create(
    model="gpt-3.5-turbo",
    messages=[
        {"role": "user", "content": "Hello, please introduce yourself"}
    ],
    temperature=0.7,
    max_tokens=500
)

print(response.choices[0].message.content)
```

#### Streaming Response

```bash
curl http://localhost:8080/v1/chat/completions \
  -H "Authorization: Bearer sk-llmhub-your-api-key-here" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gpt-3.5-turbo",
    "messages": [
      {"role": "user", "content": "Write a poem about spring"}
    ],
    "stream": true
  }'
```

### 4. List Available Models

```bash
curl http://localhost:8080/v1/models \
  -H "Authorization: Bearer sk-llmhub-your-api-key-here"
```

## Supported Model Providers

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

### SiliconFlow (硅基流动)
- Supports all models provided through the SiliconFlow platform

## API Endpoints

### POST /v1/chat/completions
Create a chat completion request.

**Request Body:**
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
List all available models.

### GET /health
Health check endpoint (no authentication required).

## Important Notes

1. **API Key Management**:
   - Keys in `auth.api_keys` are used for client authentication
   - `models[].api_key` is the actual API key for each model provider

2. **Model Names**:
   - Model names are defined in the configuration file and can be customized
   - Clients use the `name` defined in the configuration when making calls

3. **Streaming Response**:
   - Set `"stream": true` to enable streaming output
   - Response format is Server-Sent Events (SSE)

4. **Error Handling**:
   - All errors follow the OpenAI error response format
   - Check the `error` field in the response for detailed information

