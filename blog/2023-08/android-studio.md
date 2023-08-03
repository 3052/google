# Android Studio

https://opengapps.org

I selected this option:

Platform | Variant
---------|---------
x86      | pico

Then open Android Studio. On "Select Hardware" screen, select a device without
"Play Store" icon.

Create virtual device. I chose Pixel 2. Then start proxy:

~~~
mitmproxy
~~~

we need to block these:

https://play.googleapis.com/download/by-token/download

Press `O` to enter options. Move to `block_list` and press Enter. Then press
`a` to add a new entry. Press Esc when finished, then `q`.

~~~
/~u download.by.token.download/444
~~~

https://docs.mitmproxy.org/stable/overview-features#blocklist

Make sure you do the following quickly, as the `/checkin` request happens upon
boot, before Play Store app is even launched. Install user certificate. then
set proxy.

## With Google APIs

You should only need one file from the Zip archive:

~~~
Core\vending-x86.tar.lz
~~~

Inside this will be another file:

~~~
vending-x86\nodpi\priv-app\Phonesky\Phonesky.apk
~~~

Now, start the device:

~~~
emulator -list-avds
emulator -avd Pixel_3a_XL_API_24 -writable-system
~~~

Next, install Google Play Store:

~~~
adb root
adb remount
adb push Phonesky.apk /system/priv-app
adb reboot
~~~

After reboot, you should then be able to start Google Play Store as normal.


## Android 7 (API 24)

~~~
POST /auth HTTP/1.1
app: com.google.android.gms
Accept-Encoding: identity
User-Agent: GoogleAuth/1.4 (generic_x86 NYC); gzip
content-type: application/x-www-form-urlencoded
Host: android.clients.google.com
Connection: Keep-Alive
device: 334c6b6d7...

ACCESS_TOKEN=1&
add_account=1&
callerPkg=com.google.android.gms&
callerSig=38918a453d07199354f8b19af05ec6562ced5788&
device_country=us&
droidguard_results=CgZpPLGFa67SEFsAAHdU9h5a1xWSAFpFJRNTS5M9AC4R4ZGh9VMLBADDJ6...&
get_accountid=1&
google_play_services_version=11743470&
lang=en_US&
sdk_version=24&
service=ac2dm&
Email=s...&
androidId=334c6b6d7...&
Token=oauth2_4%2F0AZEOvhXAiPawRBUX5jeC46vJV3QSZfmr9lQ6RdCLdhdywbHLUGKFd71-eVm-...
~~~

## Android 6 (API 23)
