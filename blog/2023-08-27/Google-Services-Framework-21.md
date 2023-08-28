# Google Services Framework 21 with old stuff

Android 5. uses bot guard. these are the involved packages:

~~~
package: name='com.google.android.gsf' versionCode='21'
versionName='5.0.2-1649326' platformBuildVersionName='5.0.2-1649326'

package: name='com.google.android.gms' versionCode='214857013'
versionName='21.48.57 (020700-424705839)' platformBuildVersionName='Tiramisu'

package: name='com.android.vending' versionCode='83032110'
versionName='30.3.21-21 [0] [PR] 445437866' platformBuildVersionName='Tiramisu'
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
`a` to add a new entry. Press Esc when finished, then `q`.

~~~
/~u fdfe.acquire/444
/~u play-apps-download-default.download.by-id/444
~~~

https://docs.mitmproxy.org/stable/overview-features#blocklist

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

~~~
play -d com.android.vending -v 82195010 
fail

play -s -d com.android.vending -v 82284410
pass
~~~
