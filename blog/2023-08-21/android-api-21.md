# Android image API 21

Android 5

<https://wikipedia.org/wiki/Android_Lollipop>

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
