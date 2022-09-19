# Тиждень 4

>Для апішечок повинні бути тести
> /lectures/week_4/02_json_api/main_test.go

## Завдання

До виконання 2 завдання

1. Echo
2. Міста

### 1. Відлуння

Необхідно зробити луна-сервер, який приймає будь-який http запит і у відповідь повертає за прообразом http://httpbin.org

Перелік:

- http://httpbin.org/ip Returns Origin IP.
- http://httpbin.org/user-agent Returns user-agent.
- http://httpbin.org/headers Returns header dict.
- http://httpbin.org/get Returns GET data.
- http://httpbin.org/post Returns POST data.
- http://httpbin.org/put Returns PUT data.
- http://httpbin.org/delete Returns DELETE data
- http://httpbin.org/status/:code Returns given HTTP Status code.
- http://httpbin.org/response-headers?key=val Returns given response headers.
- http://httpbin.org/cookies Returns cookie data.
- http://httpbin.org/cookies/set/:name/:value Sets a simple cookie.
- http://httpbin.org/stream/:n Streams n-100 lines.
- http://httpbin.org/delay/:n Delays responding for n10 seconds

### 2. Міста

За основу взяти попереднє рішення міст. Реалізувати веб-сервер

1) /init - починає гру, повертає 1-е місто
2) /get-answer/{city_name} - отримуєте варіант відповіді
3) /submit - відправляєте місто, отримуєте нове
4) повертаєтеся до пункту 2


## Cookie
- https://ua.wikipedia.org/wiki/Cookie
- https://astaxie.gitbooks.io/build-web-application-with-golang/en/06.1.html