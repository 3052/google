# YouTube

we want to support this:

~~~
com.google.android.youtube
~~~

which means:

~~~
play -d com.google.android.youtube -v 1539962304
~~~

and:

~~~
package: name='com.google.android.youtube' versionCode='1539962304'
versionName='18.36.38' platformBuildVersionName='VanillaIceCream'
compileSdkVersion='34' compileSdkVersionCodename='VanillaIceCream'
sdkVersion:'26'
~~~

which means we need to be testing with Google Services Framework 26. whats the
current revision for that? this:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-26_r16.zip>

whats the oldest available revision for that? this:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-26_r04.zip>

looks like with Google Play Store 11-13, X-DFE-Signature-Request is required
with the `/fdfe/acquire`. supposedly its just 256 bytes of random, but I cant
get it to work. looks like the header was dropped in Google Play Store 14, so I
am going to try that. OK even with Google Services Framework 26, I am getting
this:

> This app isn't compatible with your device anymore. Contact the developers for
> more info.

what about Google Services Framework 27? newest version:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-27_r11.zip>

need at least this Google Play Store:

http://apkmirror.com/apk/google-inc/google-play-store/google-play-store-22-8-44-release

with previous versions you get:

> Google Play Store keeps stopping
