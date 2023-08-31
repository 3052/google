# Google Services Framework 21 with old stuff

Android 5. uses bot guard. these are the involved packages:

~~~
com.google.android.gsf
com.google.android.gms
com.android.vending
~~~

what are the versions of Google Play Store? here is Open GApps version:

~~~
package: name='com.android.vending' versionCode='83032110'
versionName='30.3.21-21 [0] [PR] 445437866' platformBuildVersionName='Tiramisu'
~~~

version 10 works:

http://apkmirror.com/apk/google-inc/google-play-store/google-play-store-10-8-50-release

earlier versions fail:

> Unfortunately, Google Play services has stopped.

when does `/fdfe/purchase` switch to `/fdfe/acquire`?

~~~
apkmirror.com/apk/google-inc/google-play-store/google-play-store-11-9-30-release
/fdfe/acquire

apkmirror.com/apk/google-inc/google-play-store/google-play-store-10-8-50-release
/fdfe/purchase
~~~

then:

~~~
mitmproxy
~~~

we need to block these:

https://rr3---sn-q4flrnsd.gvt1.com/play-apps-download-default/download/by-id

Press `O` to enter options. Move to `block_list` and press Enter. Then press
`a` to add a new entry. Press Esc when finished, then `q`. need this:

~~~
/fdfe/uploadDeviceConfig
~~~

https://docs.mitmproxy.org/stable/overview-features#blocklist

this is the current latest working combination:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r32.zip>

~~~
play -d com.android.vending -v 82011800
~~~

this older version of the Google Play Store is a little buggy, but it more or
less works. if it crashes you should just be able to start it again. or you can
bump up to like version 22 and the crashing seems to be resolved. then:

~~~
emulator -avd Pixel_3a_XL_API_21 -writable-system
~~~

then:

~~~
adb remount
adb push com.android.vending.apk /system/priv-app
adb reboot
~~~

then enable proxy, then install system certificate.
