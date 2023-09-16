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
apkmirror.com/apk/google-inc/google-play-store/google-play-store-20-9-20-release
pass

19

18

17

apkmirror.com/apk/google-inc/google-play-store/google-play-store-16-9-11-release

15

14

13

apkmirror.com/apk/google-inc/google-play-store/google-play-store-12-9-30-release
fail
~~~
