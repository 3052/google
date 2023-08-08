# SDK version 18

what about GenyMotion? GenyMotion only goes back to Android 5. what about Open
GApps? Open GApps only goes back to Android 4.4.

1. Android Studio
2. Pixel 3a XL
3. API Level 18
4. Android 4.3 Google APIs image

then:

~~~
play -d com.android.vending -v 82011800
~~~

check version:

~~~
name='com.android.vending'
versionCode='82011800'
versionName='20.1.18-all [0] [PR] 311592326'
~~~

try to install and you get:

~~~
INSTALL_PARSE_FAILED_INCONSISTENT_CERTIFICATES
~~~

then:

~~~
> adb uninstall com.android.vending
Failure
~~~

then:

~~~
> adb install -r com.android.vending-81701100.apk
Performing Push Install
com.android.vending-81701100.apk: 1 file pushed, 0 skipped. 282.1 MB/s
        pkg: /data/local/tmp/com.android.vending-81701100.apk
Failure [INSTALL_PARSE_FAILED_INCONSISTENT_CERTIFICATES]
~~~

then:

1. API Level 18
2. Android 4.3 image

~~~
adb root
adb remount
adb push Phonesky.apk /system/app
adb push GoogleLoginService.apk /system/app
adb reboot
~~~

http://forum.xda-developers.com/t/gapps-ultra-slim-gapps.2435683

need to figure out system certificate.
