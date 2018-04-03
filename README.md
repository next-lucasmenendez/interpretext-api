# Interpretext API
HTTP API to provide simple NLP functions in golang. Includes:
 - [Language detector](https://github.com/beeva-labs/lang-detector)
 - [Text Tokenizer](https://github.com/beeva-labs/text-tokenizer)
 - [PoS Tagger](https://github.com/beeva-labs/postagger)
 - [Text Summarizer](https://github.com/beeva-labs/text-summarizer)

## Installation
** Installation requires [docker](https://docs.docker.com/install/)**

Clone this repo and build docker image:
```sh
git clone https://github.com/beeva-labs/interpretext-api && cd ./interpretext-api
docker build -t interpretext ./
```

## Run
To deploy API server, run docker container:
```sh
docker run -p 8080:80 interpretext
```

