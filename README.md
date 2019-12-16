# Interpretext API
HTTP API to provide simple NLP functions in golang. Includes:

 - [Language detector](https://github.com/next-lucasmenendez/interpretext-lang-detector)
 - [Text Tokenizer](https://github.com/next-lucasmenendez/interpretext-tokenizer)
 - [Keyword Extractor](https://github.com/next-lucasmenendez/interpretext-keyword-extractor)
 - [PoS Tagger](https://github.com/next-lucasmenendez/interpretext-postagger)
 - [Text Summarizer](https://github.com/next-lucasmenendez/interpretext-text-summarizer)

## Installation

- Clone repository.
```sh
	git clone https://github.com/next-lucasmenendez/interpretext-api.git
```

- Inside **the repo directory**, build the `Dockerfile`.
```sh
	docker build -t interpretext ./
```

## How to use?

### Running the container
- Run the container and attach the host port:
```sh
	docker run -p <HOST_PORT>:80 -t interpretext
```

### Making a request
- Choose a feature from the following list:

| Feature               | Endpoint      |
|-----------------------|---------------|
| Text tokenizer        | `/tokenize`   |
| Language detector     | `/language`   |
| Keyword extractor     | `/keywords`   |
| Part-of-Speech tagger | `/postagging` |
| Summary extractor     | `/summary`    |

- Compose the request with the following data:
	- `POST` HTTP Method
	- `Multipart-Form` as HTTP Content-type header.
	- Input raw text into `"input"` body field.
