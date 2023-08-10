# SDK version 19

1. Android Studio
2. Pixel 3a XL
3. API Level 19
4. Android 4.4 Google APIs image

swap image for this:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r23.zip>

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
emulator -avd Pixel_3a_XL_API_19 -writable-system
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

then set proxy. then install user certificate. then:

~~~
adb reboot
~~~

result:

~~~
POST /auth HTTP/1.1
Host: android.clients.google.com
Accept-Encoding: identity
Connection: Keep-Alive
User-Agent: GoogleAuth/1.4 (generic_x86 KK); gzip
app: com.google.android.gms
content-type: application/x-www-form-urlencoded
device: 34f9d97b2...

add_account=1&
device_country=us&
droidguard_results=CgaWaC1JPYPSEFsAAHdU9h5a1xWSAFpFJRNTS5M9AC4R4ZGh9VMLBADDJ6...&
google_play_services_version=9452030&
is_dev_key_gmscore=1&
lang=en_US&
sdk_version=19&
service=ac2dm&
Email=s...%40gmail.com&
androidId=34f9d97b2...&
EncryptedPasswd=AFcb4KSL0GoATrz51O2Jt_BS5Uq8T2y7eDIj7YLxoIuYthoGAgWj9jCdm5fDIr...
~~~
