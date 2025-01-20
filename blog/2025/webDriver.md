# webDriver

https://github.com/mozilla/geckodriver

## new session

request:

~~~
curl -H content-type:application/json -d '{"capabilities": {}}' `
127.0.0.1:4444/session
~~~

response:

~~~json
{
  "value": {
    "sessionId": "222423a0-4da3-413c-a3fe-be1725c4c142"
  }
}
~~~

https://w3c.github.io/webdriver#dfn-new-sessions

## navigate to

~~~
curl -H content-type:application/json `
-d '{"url": "http://accounts.google.com/embedded/setup/v2/android"}' `
127.0.0.1:4444/session/222423a0-4da3-413c-a3fe-be1725c4c142/url
~~~

https://w3c.github.io/webdriver#navigate-to

## get all cookies

~~~
curl 127.0.0.1:4444/session/73a1041c-0bf6-4d40-a3f0-1784f8bde792/cookie
~~~

https://w3c.github.io/webdriver#cookies
