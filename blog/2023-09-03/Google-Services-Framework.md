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

we should be able to just pull the files from the device itself. its these:

~~~
x86\system.img\priv-app\GoogleServicesFramework
x86\system.img\framework\x86\boot.art
~~~

now lets try this:

<http://dl.google.com/android/repository/sys-img/android/x86-21_r05.zip>

~~~
adb push GoogleServicesFramework /system/priv-app
adb push boot.art /system/framework/x86/boot.art
adb push com.google.android.gms-202414013.apk /system/priv-app
adb push com.google.android.play.games-95330070.apk /system/app
~~~

> Something went wrong
>
> Google Play Games is having trouble with Google Play services. Please try
> again.

~~~
x86\system.img\app\PlayGames
x86\system.img\framework\x86\boot.art
x86\system.img\priv-app\GoogleServicesFramework
x86\system.img\priv-app\PrebuiltGmsCore
~~~

push:

~~~
adb push GoogleServicesFramework /system/priv-app
adb push PlayGames /system/app
adb push PrebuiltGmsCore /system/priv-app
adb push boot.art /system/framework/x86
~~~

> Unfortunately, Google Play services has stopped.

~~~
x86\system.img\app\PlayGames
x86\system.img\framework\x86\boot.art
x86\system.img\priv-app\GoogleLoginService
x86\system.img\priv-app\GoogleServicesFramework
x86\system.img\priv-app\PrebuiltGmsCore
~~~

push:

~~~
adb push GoogleLoginService /system/priv-app
adb push GoogleServicesFramework /system/priv-app
adb push PlayGames /system/app
adb push PrebuiltGmsCore /system/priv-app
adb push boot.art /system/framework/x86
~~~

maybe the better option is to just use this:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r32.zip>

comes with this:

~~~
package: name='com.google.android.gms' versionCode='202414013'
versionName='20.24.14 (020700-319035315)' compileSdkVersion='30'
~~~

better:

~~~
play -d com.google.android.gms -v 9452270
~~~

result:

~~~
package: name='com.google.android.gms' versionCode='9452270'
versionName='9.4.52 (270-127739847)' platformBuildVersionName='6.0-2166767'
sdkVersion:'21'
~~~

install:

~~~
adb install com.google.android.gms-9452270.apk
~~~

result:

~~~
Failure [INSTALL_FAILED_DUPLICATE_PERMISSION
perm=com.google.android.gms.WRITE_VERIFY_APPS_CONSENT
pkg=com.google.android.gms]
~~~

search:

~~~
INSTALL_FAILED_DUPLICATE_PERMISSION com.google.android.gms
~~~

then:

~~~
adb install -r -d com.google.android.gms-9452270.apk
~~~

result:

~~~
INSTALL_FAILED_DUPLICATE_PERMISSION
~~~
