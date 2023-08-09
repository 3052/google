# SDK version 18

1. Android Studio
2. Pixel 3a XL
3. API Level 18
4. Android 4.3 image

what about GenyMotion? GenyMotion only goes back to Android 5. what about
Google APIs image? with that you get:

~~~
INSTALL_PARSE_FAILED_INCONSISTENT_CERTIFICATES
~~~

http://forum.xda-developers.com/t/gapps-ultra-slim-gapps.2435683

what about Open GApps? Open GApps only goes back to Android 4.4.

~~~
emulator -list-avds
emulator -avd Pixel_3a_XL_API_18 -writable-system
~~~

then:

~~~
adb remount
adb push C:\Users\Steven\.mitmproxy\mitmproxy-ca-cert.cer /system/etc/security/cacerts/c8750f0d.0
adb shell chmod 664 /system/etc/security/cacerts/c8750f0d.0

adb push Phonesky.apk /system/app
adb push GoogleLoginService.apk /system/app
adb reboot
~~~

what about drag to install? you get this:

~~~
INSTALL_PARSE_FAILED_NO_CERTIFICATES
~~~

install system certificate. then:

~~~
mitmproxy
~~~
