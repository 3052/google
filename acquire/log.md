# event logging does not work

~~~
10.0.0 pass
9.0.1
9.0.0
v8.1.1
v8.1.0
v8.0.0 fail
~~~

if I have this script:

~~~py
from mitmproxy import http
import logging

def request(f: http.HTTPFlow) -> None:
   logging.info('hello world')
~~~

and I run it:

~~~
mitmproxy -s log.py
~~~

I see requests:

~~~
>> GET http://clients3.google.com/generate_204
       ← 204 [no content] 79ms
   GET http://clients3.google.com/generate_204
       ← 204 [no content] 36ms
~~~

but nothing in the event log:

~~~
info: Loading script log.py
info: Proxy server listening at http://*:8080
info: 127.0.0.1:60019: client connect
info: 127.0.0.1:60019: server connect clients3.google.com:80 (142.250.114.102:80)
info: 127.0.0.1:60016: client connect
info: 127.0.0.1:60016: server connect 172.253.62.188:5228
~~~
