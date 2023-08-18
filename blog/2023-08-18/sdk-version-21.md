# SDK version 21

1. Android Studio
2. Pixel 3a XL
3. API Level 21
4. Android 5 Google APIs image

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

older versions are available, but they are buggy. You can also pull the APK from
the `google_apis_playstore` image, but it is buggy as well. start the device:

~~~
emulator -list-avds
emulator -avd Pixel_3a_XL_API_21 -writable-system
~~~

Install like this:

~~~
adb root
adb remount
adb push com.android.vending-82011800.apk /system/priv-app/Phonesky.apk
adb reboot
~~~

install system certificate. then:

~~~
mitmproxy
~~~

start Play Store. click Sign in. enter Email and click Next. Enter password and
click Next. if app closes for update, start again. click Sign in. enter Email
and click Next. Enter password and click Next. click I agree. Under Use basic
device backup, move slider to left. click ACCEPT. under Apps, click an app.
result:

~~~
POST /auth HTTP/1.1
Host: android.clients.google.com
Accept-Encoding: identity
Connection: Keep-Alive
User-Agent: GoogleAuth/1.4 (generic_x86 LSY66K); gzip
app: com.google.android.gms
content-type: application/x-www-form-urlencoded
device: 3760dcd91...

ACCESS_TOKEN=1&
add_account=1&
device_country=us&
droidguard_results=CgYXdhUgp-PSEFsAAHdU9h5a1xWSAFpFJRNTS5M9AC4R4ZGh9VMLBADDJ6...&
google_play_services_version=11055270&
is_dev_key_gmscore=1&
lang=en_US&
sdk_version=21&
service=ac2dm&
Email=s...&
androidId=3760dcd91...&
Token=oauth2_4%2F0Adeu5BVnTdwNdQyQ97hA6iZKM3KitofkyAKOb3yucDZ0JxPCvMJ3c7y-lmHr...
~~~

this is it:

- https://accounts.google.com/EmbeddedSetup
- https://accounts.google.com/embedded/setup/android

issues:

- <https://github.com/Aruelius/Google_login/issues/3>
- <https://github.com/ikp4success/bypass_google_bot_guard/issues/5>
- https://gist.github.com/Ruin0x11/a32c4ba33d386b51797a3719ad2c14e4
- https://github.com/409232112/JavaBlock/issues/5
- https://github.com/BitTheByte/YouTubeShop/issues/14
- https://github.com/equ1n0x93/GooglePlayApi/issues/2
- https://github.com/raiyan088/public/issues/1
