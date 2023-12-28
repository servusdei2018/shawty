# shawty

The lightweight sqlite-backed URL shortener in your pocket.

## Installation

### Building manually

Shawty may be built manually with `go build ./cmd/shawty`.

### Docker

Alternatively, the provided docker-compose file may be used. Shawty will be exposed, by default, on port `8080`, this may be configured in `docker-compose.yml`.

## Configuration

The following environment variables must be specified in the environment or `.env` file:

`SHAWTYAUTH`: authorization key for publishing and deleting short urls (leave empty for no-auth)

`SHAWTYPORT`: (default 8080) the port on which to serve HTTP.

## Endpoints

### GET `/`

Returns a summary of endpoints, and a count of the URLs in the database.

### GET `/<shortURL>`

Redirects to full URL or returns `404`.

### POST `/delete`

Deletes the POSTed shortURL.

Example:
```shell
curl -H "Auth: YourAuthToken" -X POST -d "mje6TGINd" http://localhost:8080/delete
```

### POST `/shorten`

Shortens and returns the POSTed URL.

Example:
```shell
curl -H "Auth: YourAuthToken" -X POST -d "https://github.com/a-painfully-long-url-you-want-to-shorten" http://localhost:8080/shorten
mje6TGINd # you may access it at http://localhost:8080/mje6TGINd
```

## License
```
The MIT License (MIT)

Copyright (c) 2023-present The Shawty Authors

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```