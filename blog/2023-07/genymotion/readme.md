# GenyMotion

Create Virtual device using Android version 6.0.0 (API 23). I chose
Motorola Moto X. Then start proxy:

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
