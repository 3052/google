# Google Services Framework 11

API 30

note if remount fails you might need to just retry it. if you omit permission
file, after reboot device will just load forever. this seems to be the first
version using `/fdfe/sync` only:

~~~
play -a com.android.vending -v 82941300
~~~

oldest:

~~~
> adb shell dumpsys package com.google.android.gms | Select-String version
versionCode=201817022 minSdk=23 targetSdk=30
versionName=20.18.17 (040700-311416286)
~~~

<http://dl.google.com/android/repository/sys-img/google_apis/x86-30_r05.zip>

newest:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-30_r10.zip>

~~~
emulator -avd Pixel_6_API_30 -writable-system
adb shell avbctl disable-verification
adb reboot

adb remount
adb push com.android.vending-82941300.apk /system/priv-app
adb push privapp-permissions.xml /etc/permissions
adb reboot
~~~

https://issuetracker.google.com/issues/308510932
