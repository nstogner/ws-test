# Web Service Tester

Assert on HTTP request/response pairs.

## Config

Environment variables:

- `TEST_URL` (required)
- `TEST_DIRECTORY` (required)
- `TEST_METHOD` (defaults to "POST")
- `TEST_REQUEST_TIMEOUT` (defaults to "1s")
- `TEST_SEARCH` (defaults to `**/request.*`)

The default value for `TEST_SEARCH` results in an expected directory structure like:

```
test-case-0/
  request.json
  response.json
test-case-1/
  request.json
  response.json
test-case-n/
  request.json
  response.json
```

## Usage

### Using Go

Startup an example web service in one terminal.

```sh
cd example-ws/ && go run .
```

Run test cases in another terminal.

```sh
TEST_URL='http://localhost:8080' TEST_DIRECTORY=./example-ws/tests go test . -v
```

### Using Docker

Build.

```sh
docker build . -t ws-test`
```

Test a web service.

```sh
docker run -e 'TEST_URL=https://my-webservice/my-path' -v $(pwd)/my-test-cases:/work/tests ws-test go test -v
```

