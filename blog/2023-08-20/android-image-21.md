# Android image API 21

~~~
POST https://accounts.google.com/_/signin/speedbump/embeddedsigninconsent

HTTP/2.0 200 
set-cookie: oauth_token=oauth2_4/0Adeu5BU-HYh0pC1OU2Efj0l6safyYPiUbx9HrHmU87-d...
~~~

API 25:

https://accounts.google.com/embedded/setup/v2/android

API 21-24:

https://accounts.google.com/embedded/setup/android

Android 5

<https://wikipedia.org/wiki/Android_Lollipop>

here is Open GApps version:

~~~
package: name='com.android.vending' versionCode='83032110'
versionName='30.3.21-21 [0] [PR] 445437866' platformBuildVersionName='Tiramisu'
~~~

version 10 works:

http://apkmirror.com/apk/google-inc/google-play-store/google-play-store-10-8-50-release

earlier versions fail:

> Unfortunately, Google Play services has stopped.

here is Open GApps version:

~~~
package: name='com.google.android.gms' versionCode='214857013'
versionName='21.48.57 (020700-424705839)' platformBuildVersionName='Tiramisu'
~~~

version 20 works:

~~~
play -d com.google.android.gms -v 205066013
~~~

earlier versions fail:

~~~
play -d com.google.android.gms -v 19866005 
~~~

> Unfortunately, Google Play services has stopped.

~~~
adb remount

adb push GoogleServicesFramework /system/priv-app
adb push com.android.vending_10.8.50.apk /system/priv-app
adb push com.google.android.gms-205066013.apk /system/priv-app
adb reboot
~~~

## device Android Studio, Open GApps

https://opengapps.org

~~~
open_gapps-x86-5.0-pico-20220503.zip\Core\gmscore-x86.tar.lz\gmscore-x86.tar\
gmscore-x86\common\priv-app\PrebuiltGmsCore

open_gapps-x86-5.0-pico-20220503.zip\Core\gmscore-x86.tar.lz\gmscore-x86.tar\
gmscore-x86\nodpi\priv-app\PrebuiltGmsCore

open_gapps-x86-5.0-pico-20220503.zip\Core\gsfcore-all.tar.lz\gsfcore-all.tar\
gsfcore-all\nodpi\priv-app\GoogleServicesFramework

open_gapps-x86-5.0-pico-20220503.zip\Core\vending-x86.tar.lz\vending-x86.tar\
vending-x86\common\priv-app\Phonesky

open_gapps-x86-5.0-pico-20220503.zip\Core\vending-x86.tar.lz\vending-x86.tar\
vending-x86\nodpi\priv-app\Phonesky
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
/~u play-apps-download-default.download.by-id/444
~~~

https://docs.mitmproxy.org/stable/overview-features#blocklist

then:

~~~
emulator -avd Pixel_3a_XL_API_21 -writable-system
~~~

then:

~~~
adb remount

adb push GoogleServicesFramework /system/priv-app
adb push Phonesky /system/priv-app
adb push PrebuiltGmsCore /system/priv-app
adb reboot
~~~

then enable proxy, then install system certificate.

OK might have hit a brick wall. it seems the `/_/lookup/accountlookup` request
is protected by a value `bgRequest`, which is Google bot-guard protection. 

https://github.com/simon-weber/gpsoauth/issues/56
