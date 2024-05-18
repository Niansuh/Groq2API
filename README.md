# Groq2API

## Usage

Optional parameter list

- `model` Model name
  - gemma-7b-it
  - mixtral-8x7b-32768
  - llama2-70b-4096
  - llama3-8b-8192
  - llama3-70b-8192
- `stream` Whether to stream output
  - true
  - false
- `max_tokens` Maximum generated length
  - 4096 (llama2-70b-4096) 
  - 8192 (gemma-7b-it)
  - 32768 (mixtral-8x7b-32768)
- `message`
  - `role` message role
    - user
    - assistant
```bash

curl --request POST \
  --url http://127.0.0.1:8080/v1/chat/completions \
  --header 'Authorization: Bearer stytch_session' \
  --data '{
  "messages": [
    {
      "role": "user",
      "content": "hi"
    }
  ],
  "model": "mixtral-8x7b-32768",
  "max_tokens": 4096,
  "stream": true
}'

```

### stytch_session Get method
![image|690x233](https://cdn.linux.do/uploads/default/original/3X/c/c/cc5bb06024b2fc93581227e16b5a5e3e220d159c.png)
