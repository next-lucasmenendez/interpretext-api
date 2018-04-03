# Interpretext API
HTTP API to provide simple NLP functions in golang. Includes:

 - [Language detector](https://github.com/beeva-labs/lang-detector)
 - [Text Tokenizer](https://github.com/beeva-labs/text-tokenizer)
 - [PoS Tagger](https://github.com/beeva-labs/postagger)
 - [Text Summarizer](https://github.com/beeva-labs/text-summarizer)

## Installation

**Installation requires [docker](https://docs.docker.com/install/).**

Clone this repo and build docker image:

```sh
    $ git clone https://github.com/beeva-labs/interpretext-api && cd ./interpretext-api
    $ docker build -t interpretext ./
```

## Run
To deploy API server, run docker container:

```sh
    $ docker run -p 8080:80 interpretext
```

---



## API Endpoints

### POST `/language`
Returns estimated language of input text provided.

#### Request

##### Headers

```
    Content-type: multipart/form-data
```

##### Body (Multipart Form)

```
    input: <single or multiline text:string>
```

#### Response

##### 200 OK response (application/json)

```json
    {
    	"lang": <language-code:string>
    }
```

##### 400 Bad Request response (text/plain)

```
	No input text provided.
```

##### 500 Internal Server Error response (text/plain)

```
    Error parsing JSON response.
```

--

### POST `/tokenize`
Splits input text into sentences and words. Returns list of sentence words into a list of sentences.

#### Request

##### Headers

```
    Content-type: multipart/form-data
```

##### Body (Multipart Form)

```
    input: <single or multiline text:string>
```

#### Response

##### 200 OK response (application/json)

```json
    {
    	"tokens": <tokenized list of sentences:[][]string>
    }
```

##### 400 Bad Request response (text/plain)

```
	No input text provided.
```
##### 500 Internal Server Error response (text/plain)

```
    Error parsing JSON response.
```
--

### POST `/postagging`
Splits input text provided into tokens and tagged each token with models trained. For mor infor check [PoS Tagger repo](https://github.com/beeva-labs/postagger). Returns list of tuples with the token and it's (suggested) tag.

#### Request

##### Headers

```
    Content-type: multipart/form-data
```

##### Body (Multipart Form)

```
    input: <single or multiline text:string>
```

#### Response

##### 200 OK response (application/json)

```json
    {
    	"tokens": <list of tuples token-tag:[][]string>
    }
```

##### 204 No Content response (text/plain)

```
	No tags were found.
```

##### 400 Bad Request response (text/plain)

```
	No input text provided.
```
##### 500 Internal Server Error response (text/plain)

```
	No models configured for PoS tagging.
```

```
	Error loading model for PoS tagging.
```

```
    Error parsing JSON response.
```
--

### POST `/summary`
Extracts _"most relevant"_ sentences from the input text provided. Returns list of top sentences according to its position in the original text.

#### Request

##### Headers

```
    Content-type: multipart/form-data
```

##### Body (Multipart Form)

```
    input: <single or multiline text:string>
```

#### Response

##### 200 OK response (application/json)

```json
    {
    	"summary": <list of highlights:[]string>
    }
```

##### 400 Bad Request response (text/plain)

```
	No input text provided.
```

##### 500 Internal Server Error response (text/plain)

```
    Error parsing JSON response.
```

```
	Error analyzing input text.
```


