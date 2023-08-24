# Android image API 21

<https://wikipedia.org/wiki/Android_Lollipop>

~~~
dl.google.com/android/repository/sys-img/google_apis/x86-21_r32.zip
google_play_services_version:   202414013
~~~

1. Android Studio
2. Android 5
3. https://opengapps.org

~~~
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

adb push Phonesky /system/priv-app
adb reboot
~~~

then enable proxy, then install system certificate.
