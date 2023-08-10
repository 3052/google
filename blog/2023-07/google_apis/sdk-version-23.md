# SDK version 23

1. Android Studio
2. Pixel 3a XL
3. API Level 23
4. Android 6 Google APIs image

swap image for this:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-23_r20.zip>

older revisions are available, but they are buggy. then:

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
emulator -avd Pixel_3a_XL_API_23 -writable-system
~~~

Install like this:

~~~
adb root
adb remount
adb push com.android.vending-82011800.apk /system/priv-app/Phonesky.apk
~~~

we need to block these:

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

start Play Store. click Sign in. enter Email and click Next. Enter password and
click Next. if app closes for update, start again. click Sign in. enter Email and
click Next. Enter password and click Next. click I agree. Under Use basic device
backup, move slider to left. click ACCEPT. under Apps, click an app. result:

~~~
POST /auth HTTP/1.1
Host: android.clients.google.com
app: com.google.android.gms
Accept-Encoding: identity
User-Agent: GoogleAuth/1.4 (generic_x86 MASTER); gzip
content-type: application/x-www-form-urlencoded
Connection: Keep-Alive
device: 33388cc46...

ACCESS_TOKEN=1&
add_account=1&
device_country=us&
droidguard_results=CgZit9r5tB7SEFsAAHdU9h5a1xWSAFpFJRNTS5M9AC4R4ZGh9VMLBADDJ6...&
google_play_services_version=10298470&
is_dev_key_gmscore=1&
lang=en_US&
sdk_version=23&
service=ac2dm&
Email=s...&
androidId=33388cc46...&
Token=oauth2_4%2F0Adeu5BXOWxFJ6YohsZUbjetKUiwQJbYHx7cDQDJcGdRVG65a6CNCN43q_ro1...
~~~
