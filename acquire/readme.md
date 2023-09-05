# acquire

ideally we want this:

1. `/checkin`
2. `/fdfe/acquire`

but for now we can settle for:

1. get `X-DFE-Device-ID` from Android Studio
2. `/fdfe/acquire`

so here is the captured request:

~~~
POST /fdfe/acquire?theme=2 HTTP/1.1
Host: android.clients.google.com
X-DFE-Device-Id: 306e9f7f4192be79
~~~

so we can borrow this ID for now, and try to get the rest of the request working.
when does `/fdfe/purchase` switch to `/fdfe/acquire`?

~~~
apkmirror.com/apk/google-inc/google-play-store/google-play-store-11-9-30-release
/fdfe/acquire

apkmirror.com/apk/google-inc/google-play-store/google-play-store-10-8-50-release
/fdfe/purchase
~~~
