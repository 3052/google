# X-DFE-Signature-Request

if we block this header:

~~~
if f.request.path.startswith('/fdfe/acquire'):
   f.request.headers.pop('X-DFE-Signature-Request', None)
~~~

we get this:

> Error
>
> No internet connection. Make sure Wi-Fi or cellular data is turned on, then try
> again.

using either of these:

- <http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r20.zip>
- <http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r32.zip>

~~~
apkmirror.com/apk/google-inc/google-play-store/google-play-store-14-9-76-release
pass

apkmirror.com/apk/google-inc/google-play-store/google-play-store-13-9-40-release
fail
~~~
