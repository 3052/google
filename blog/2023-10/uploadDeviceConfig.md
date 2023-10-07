# /fdfe/uploadDeviceConfig

currently I am using API 26, which is Android 8. how high can I go with the
Google Play Store? looks like we can go to the current version. currently we are
using Google Play Store 11. this works:

~~~
play -d com.android.vending -v 82951610 -s
~~~

I think what is different now is the call to `/fdfe/acquire`, notably the request
headers. this is new:

~~~
X-PS-RH: H4sIAAAAAAAAAC3NwQrCIAAAUHbd17jQ6KrTNUkNwRp62WU02g4FajpPQT8eQe8HXv2p6...
~~~

what happens if we block it? still works.
