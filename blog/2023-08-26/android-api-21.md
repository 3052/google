# Android image API 21

<https://wikipedia.org/wiki/Android_Lollipop>

swap image for this:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r30.zip>

~~~
package: name='com.google.android.gms' versionCode='16089013'
versionName='16.0.89 (020700-239467275)' platformBuildVersionName='Q'
~~~

older versions are available, but they fail:

~~~
adb: error: failed to copy 'Phonesky\lib\x86\libcronet.102.0.4973.2.so' to
'/system/priv-app/Phonesky/lib/x86/libcronet.102.0.4973.2.so': remote No space
left on device
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

then enable proxy, then install system certificate. this is the new purchase:

~~~
https://play-fe.googleapis.com/fdfe/acquire?theme=2
~~~
