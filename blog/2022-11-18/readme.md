# 2022-11-19

## GenyMotion

Create Virtual device using Android version 6.0.0. Install Open GApps. Then
click Restart now. If you have trouble at this point, you might need to End
task:

~~~
C:\Program Files\Genymobile\Genymotion\tools\adb.exe
~~~

Then install system certificate. Then start proxy:

~~~
mitmproxy
~~~

then set proxy:

~~~
adb shell settings put global http_proxy 192.168.56.1:8080
~~~

Note if you restart the device, you need to install system certificate again.
Make sure you do the above quickly, as the `/checkin` request happens upon
boot, before Play Store app is even launched.

https://support.genymotion.com/hc/articles/360002778137-How-to-connect
