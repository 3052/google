# Android image API 21

Android 5

<https://wikipedia.org/wiki/Android_Lollipop>

## device GenyMotion, adb

~~~
adb remount

adb push GoogleLoginService /system/priv-app
adb push GoogleServicesFramework /system/priv-app
adb push Phonesky /system/priv-app
adb push PrebuiltGmsCore /system/priv-app
adb reboot
~~~

MAKE SURE TO GET THE LIB FOLDER TOO, OR YOU GET THIS:

> Unfortunately, Google Play services has stopped.

## device GenyMotion, GApps GenyMotion

click Open GAPPS

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

Install Open GApps. If you have trouble at this point, you might need to End
task:

~~~
C:\Program Files\Genymobile\Genymotion\tools\adb.exe
~~~

Click Restart now. Make sure you do the following quickly, as the `/checkin`
request happens upon boot, before Play Store app is even launched. Once
"Android is starting" is almost complete, install system certificate. then set
proxy:

~~~
adb shell settings put global http_proxy 192.168.56.1:8080
~~~

Note if you restart the device, you need to install system certificate again.

https://support.genymotion.com/hc/articles/360002778137-How-to-connect
