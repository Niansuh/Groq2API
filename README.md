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
  --url http://127.0.0.1:8080/nai/v1/chat/completions \
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
![4c96b3e9-c8c4-473e-b8f2-f748abb9f049](https://github.com/Niansuh/Groq2API/assets/139365289/245786bb-6fe4-40ff-842a-bf705a7374f8)
