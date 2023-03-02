# Blog

## How to get Protocol Buffer fields?

Check these with JADX, last working versions:

key             | value
----------------|--------------------
package         | com.android.vending
versionCode     | 80441400
versionName     | 6.1.14

key         | value
------------|-----------------------
package     | com.google.android.gsf
versionCode | 25
versionName | 7.1.2

- https://apkmirror.com/apk/google-inc/google-play-store
- https://apkmirror.com/apk/google-inc/google-services-framework

## How to install Android App Bundle?

Bash:

~~~
adb install-multiple *.apk
~~~

PowerShell:

~~~
adb install-multiple (Get-ChildItem *.apk)
~~~

https://developer.android.com/guide/app-bundle/app-bundle-format

## How to install expansion file?

~~~
adb shell mkdir -p /sdcard/Android/obb/com.PirateBayGames.ZombieDefense2

adb push main.41.com.PirateBayGames.ZombieDefense2.obb `
/sdcard/Android/obb/com.PirateBayGames.ZombieDefense2/
~~~

https://developer.android.com/google/play/expansion-files

## INSTALL\_FAILED\_NO\_MATCHING\_ABIS

This can happen when trying to install ARM app on `x86`. If the APK is
`armeabi-v7a`, then Android 9 (API 28) will work. Also the emulator should be
`x86`. If the APK is `arm64-v8a`, then Android 11 (API 30) will work. Also the
emulator should be `x86_64`.

- https://android.stackexchange.com/questions/222094/install-failed
- https://stackoverflow.com/questions/36414219/install-failed-no-matching-abis

However note that this will still fail in some cases:

https://issuetracker.google.com/issues/253778985

## System Image

API Level | Target
----------|-----------------
33        | Android Tiramisu
32        | Android 12L
31        | Android 12
30        | Android 11
29        | Android 10
28        | Android 9
27        | Android 8.1
26        | Android 8
25        | Android 7.1
24        | Android 7
23        | Android 6
22        | Android 5.1

## Version history

If you know the `versionCode`, you can get older APK [1]. Here is one from 2014:

~~~
googleplay -a com.google.android.youtube -v 5110
~~~

but I dont know how to get the old version codes, other than looking at
websites [2] that host the APKs.

1. https://android.stackexchange.com/questions/163181/how-to-download-old-app
2. https://apkmirror.com/uploads?appcategory=youtube
