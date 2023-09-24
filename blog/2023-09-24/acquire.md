# /fdfe/acquire

looks like with Google Play Store 11-13, X-DFE-Signature-Request is required
with the `/fdfe/acquire`. supposedly its just 256 bytes of random, but I cant
get it to work. looks like the header was dropped in Google Play Store 14, so I
am going to try that.

---------------------------------------------------------------------------------

FINALLY. OK I think I figured out the problem. I have been using a script like
this:

~~~py
from mitmproxy import ctx, http
import json

device_ID = json.load(open('client.json'))['AuthData']['GsfID']

def running():
   ctx.options.showhost = True

def request(f: http.HTTPFlow) -> None:
   if f.request.path.startswith('/play-apps-download-default/download/by-id/'):
      f.kill()
   if f.request.path.startswith('/fdfe/acquire'):
      f.request.headers['X-DFE-Device-Id'] = device_ID
~~~

and it kept randomly failing. turns out I was editing the script, THEN editing
`client.json`. because of this, the changes to `client.json` were not being
recognized. I am now editing `client.json` FIRST, then editing the script, so
script recognizes new data. I might be making some good progress now.

---------------------------------------------------------------------------------

OK another update. looks like if `/fdfe/replicateLibrary` runs, then it
prevents `/fdfe/acquire` from running, and will just go straight to
`/fdfe/delivery` if its a app that has already been acquired. this is bad
because for testing, we want `/fdfe/acquire` to run every time.

if you block `/fdfe/replicateLibrary` then acquire seems to work as normal.

interestingly, if you block `/fdfe/replicateLibrary`, and acquire, it seems to
fall back to the old `/fdfe/purchase`

---------------------------------------------------------------------------------

OK I made some more progress. looks like with `/fdfe/acquire`, if you remove
both of these it fails:

~~~
X-DFE-Cookie: EAEYACICVVMyUENqZ2FOZ29UTkRVek1EWXhOREEwTlRVek5qQXhNelk0TUJJZkNoQXhOamsxTlRFeU16a3pNREUxTVRZNUVnc0l5ZTY5cUFZUTZPdWRCdz09QhQKBVVTLVRYEgsIycLRqQYQ4MGcW0oRCgJVUxILCMnC0akGEMi9oFtYAA
X-DFE-Device-Config-Token: CjgaNgoTNDUzMDYxNDA0NTUzNjAxMzY4MBIfChAxNjk1NTEyMzkzMDE1MTY5EgsIye69qAYQ6OudBw==
~~~

but if you only remove one, it works.

---------------------------------------------------------------------------------

OK I think I have it figured out. the two headers are basically the same,
except the Cookie wraps some extra information. the Token seems pretty easy to
create:

~~~go
package token

import (
   "154.pages.dev/protobuf"
   "encoding/base64"
   "fmt"
   "time"
)

func Token(gsf_ID int64) (string, string) {
   id := fmt.Sprint(gsf_ID)
   date := fmt.Sprint(time.Now().UnixMicro())
   device := protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 3, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes(id)},
            protobuf.Field{Number: 2, Type: 2, Value: protobuf.Prefix{
               protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes(date)},
            }},
         }},
      }},
   }.Append(nil)
   return "X-DFE-Device-Config-Token", base64.StdEncoding.EncodeToString(device)
}
~~~

or I think you can also get it from the `/fdfe/uploadDeviceConfig` response.
