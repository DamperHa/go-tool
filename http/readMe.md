


```shell
fanzhihao at fanzhihaodeMacBook-Pro in ~
$ curl -v http://localhost:8080/old-page

*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /old-page HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.77.0
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 301 Moved Permanently
< Content-Type: text/html; charset=utf-8
< Location: /new-page
< Date: Mon, 30 Oct 2023 12:16:09 GMT
< Content-Length: 44
<
<a href="/new-page">Moved Permanently</a>.

* Connection #0 to host localhost left intact

fanzhihao at fanzhihaodeMacBook-Pro in ~
$ curl -L http://localhost:8080/old-page

Welcome to the new page!%
```


```shell
$ curl http://httpbin.org/get\?a\=1\&b\=2
{
  "args": {
    "a": "1",
    "b": "2"
  },
  "headers": {
    "Accept": "*/*",
    "Host": "httpbin.org",
    "User-Agent": "curl/7.77.0",
    "X-Amzn-Trace-Id": "Root=1-653f9fd9-2e06190045cd533b2813d129"
  },
  "origin": "123.119.71.199",
  "url": "http://httpbin.org/get?a=1&b=2"
}
```


```shell
$ curl -X POST http://httpbin.org/post \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YourTokenHere" \
  -H "Custom-Header: CustomValue" \
  -d '{"a":1, "b":2}' \


{
  "args": {},
  "data": "{\"a\":1, \"b\":2}",
  "files": {},
  "form": {},
  "headers": {
    "Accept": "*/*",
    "Authorization": "Bearer YourTokenHere",
    "Content-Length": "14",
    "Content-Type": "application/json",
    "Custom-Header": "CustomValue",
    "Host": "httpbin.org",
    "User-Agent": "curl/7.77.0",
    "X-Amzn-Trace-Id": "Root=1-653fa189-3de36c5c536d9f996480271c"
  },
  "json": {
    "a": 1,
    "b": 2
  },
  "origin": "123.119.71.199",
  "url": "http://httpbin.org/post"
}
```

