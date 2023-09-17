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
am going to try that.
