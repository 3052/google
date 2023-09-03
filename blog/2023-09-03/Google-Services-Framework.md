# Google Services Framework

using this:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r32.zip>

it comes with these:

~~~
> adb shell dumpsys package com.google.android.play.games | Select-String version
versionCode=95330070 targetSdk=28
versionName=2019.04.9533 (244301765.244301765-000700)

> adb shell dumpsys package com.google.android.gms | Select-String version
versionCode=202414013 targetSdk=30
versionName=20.24.14 (020700-319035315)

> adb shell dumpsys package com.google.android.gsf | Select-String version
versionCode=21 targetSdk=21
versionName=5.0.2-6695550
~~~

these work:

~~~
play -d com.google.android.gms -v 202414013
play -d com.google.android.play.games -v 95330070
~~~

these do not:

~~~
> play -d com.google.android.gsf -v 21
panic: appDeliveryData not found

> play -d com.google.android.gsf -v 21 -p 1
panic: appDeliveryData not found

> play -d com.google.android.gsf -v 21 -p 2
panic: appDeliveryData not found
~~~

we should be able to just pull the files from the device itself. its these files:

~~~
x86\system.img\priv-app\GoogleServicesFramework\GoogleServicesFramework.apk
x86\system.img\priv-app\GoogleServicesFramework\x86\GoogleServicesFramework.odex
~~~

now lets try this:

<http://dl.google.com/android/repository/sys-img/android/x86-21_r05.zip>

lets try just the APK to start. push:

~~~
adb push GoogleServicesFramework.apk /system/priv-app
adb push com.google.android.gms-202414013.apk /system/priv-app
adb push com.google.android.play.games-95330070.apk /system/app
~~~

fail:

> Unfortunately, Google Services Framework has stopped.

lets try the whole folder:

~~~
x86\system.img\priv-app\GoogleServicesFramework
~~~

push:

~~~
adb push GoogleServicesFramework /system/priv-app
adb push com.google.android.gms-202414013.apk /system/priv-app
adb push com.google.android.play.games-95330070.apk /system/app
~~~

fail:

> Unfortunately, Google Services Framework has stopped.

lets try `lib` too:

~~~
x86\system.img\lib
x86\system.img\priv-app\GoogleServicesFramework
~~~

push:

~~~
adb push GoogleServicesFramework /system/priv-app
adb push lib /system
adb push com.google.android.gms-202414013.apk /system/priv-app
adb push com.google.android.play.games-95330070.apk /system/app
~~~
