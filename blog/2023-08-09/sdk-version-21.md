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

- <https://github.com/equ1n0x93/GooglePlayApi/blob/master/google_importer_api.py>
- https://github.com/lyrgard/ffbeDataExporter/blob/master/src/ffbeSync/ffbeSync.js

1. call login
   1. call `_get_google_signign`
      - call `_get_embedded_info`
         1. request `https://accounts.google.com/embedded/setup/android?source=com.android.settings&xoauth_display_name=Android%20Phone&canSk=1&lang=en&langCountry=en_us&hl=en-US&cc=us`
         2. call `_get_embedded_token`
   2. call `_get_google_token`

this is it:

https://accounts.google.com/embedded/setup/android
