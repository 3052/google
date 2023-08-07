# SDK version 21

1. Android Studio
2. Pixel 3a XL
3. API Level 21

4. Android 5.1 Google APIs image

swap image for this:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-22_r14.zip>

older versions are available, but they are buggy. then:

~~~
play -d com.android.vending -v 82011800
~~~

check version:

~~~
name='com.android.vending'
versionCode='82011800'
versionName='20.1.18-all [0] [PR] 311592326'
~~~

older versions are available, but they are buggy. You can also pull the APK from
the `google_apis_playstore` image, but it is buggy as well. start the device:

~~~
emulator -list-avds
emulator -avd Pixel_3a_XL_API_22 -writable-system
~~~

Install like this:

~~~
adb root
adb remount
adb push com.android.vending-82011800.apk /system/priv-app/Phonesky.apk
~~~

then:

~~~
mitmproxy
~~~

then set proxy. we need to block these:

https://play.googleapis.com/download/by-token/download

Press `O` to enter options. Move to `block_list` and press Enter. Then press
`a` to add a new entry. Press Esc when finished, then `q`.

~~~
/~u download.by.token.download/444
~~~

https://docs.mitmproxy.org/stable/overview-features#blocklist

~~~
adb reboot
~~~
