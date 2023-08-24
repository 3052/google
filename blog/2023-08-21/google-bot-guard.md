# Google bot guard

## API 21-24

https://accounts.google.com/embedded/setup/android

this is the authorization code, expires in 10 minutes. get from browser:

~~~
POST https://accounts.google.com/_/signin/speedbump/embeddedsigninconsent

HTTP/2.0 200 
set-cookie: oauth_token=oauth2_4/0Adeu5BU-HYh0pC1OU2Efj0l6safyYPiUbx9HrHmU87-d...
~~~

<https://firefox-source-docs.mozilla.org/devtools-user/storage_inspector>

this is the refresh token. get from Android:

~~~
1//045TtSNds0t1sCgYIARAAGAQSNwF-L9IroY2Q_CKI56yphs2ddo66wuzcbxbGbz573kEaVy05zX...
~~~

this is the access token, expires in one hour:

~~~
ya29.a0AfB_byD4KkySd9-sXUXFaXmSKe8wWKWzsp095DsxLhDQlOmZaK871YKHDA9iesyJN9iGmeY...
~~~

## device Android Studio, GApps Android Studio

1. device Pixel 3a XL
2. API Level 21
3. Google APIs image

then:

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

then launch device. then enable proxy, then install system certificate. then
start Play Games.

## device Android Studio, APK Mirror

here is Open GApps version:

~~~
package: name='com.android.vending' versionCode='83032110'
versionName='30.3.21-21 [0] [PR] 445437866' platformBuildVersionName='Tiramisu'
~~~

version 10 works:

http://apkmirror.com/apk/google-inc/google-play-store/google-play-store-10-8-50-release

earlier versions fail:

> Unfortunately, Google Play services has stopped.

here is Open GApps version:

~~~
package: name='com.google.android.gms' versionCode='214857013'
versionName='21.48.57 (020700-424705839)' platformBuildVersionName='Tiramisu'
~~~

version 20 works:

~~~
play -d com.google.android.gms -v 205066013
~~~

earlier versions fail:

~~~
play -d com.google.android.gms -v 19866005 
~~~

> Unfortunately, Google Play services has stopped.

~~~
adb remount

adb push GoogleServicesFramework /system/priv-app
adb push com.android.vending_10.8.50.apk /system/priv-app
adb push com.google.android.gms-205066013.apk /system/priv-app
adb reboot
~~~
