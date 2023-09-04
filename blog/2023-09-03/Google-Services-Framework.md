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

> Unfortunately, Google Play services has stopped.

~~~
x86\system.img\app
x86\system.img\framework (exclude x86\boot.oat)
x86\system.img\priv-app
~~~

push:

~~~
adb push system/app /system
adb push system/framework /system
adb push system/priv-app /system
~~~

> Unfortunately, Google Play services has stopped.

~~~
x86\system.img\app\PlayGames
x86\system.img\framework (exclude x86\boot.oat)
x86\system.img\lib
x86\system.img\priv-app
x86\system.img\vendor
~~~

push:

~~~
adb push system/app/PlayGames /system/app
adb push system/framework /system
adb push system/lib /system
adb push system/priv-app /system
adb push system/vendor /system
~~~

> Unfortunately, Google Play services has stopped.

https://github.com/illogical-robot/apkmirror-public/issues/261

~~~
Date: Mon, 4 Sep 2023 09:41:14 -0500
Subject: Google Services Framework
To: opensource@apkpure.com

My understanding is the Google Services Framework app (com.google.android.gsf)
is not downloadable, even using the internal Google APIs. So my question is,
where are you getting this app from?

https://apkpure.com/google-services-framework/com.google.android.gsf
~~~

this is interesting. this:

http://apkpure.com/google-services-framework/com.google.android.gsf/download

~~~
bd32424203e0fb25f36b57e5aa356f9bdd1da998
~~~

matches with this:

http://apkmirror.com/apk/google-inc/google-services-framework/google-services-framework-14-release

~~~
Signature: bd32424203e0fb25f36b57e5aa356f9bdd1da998
~~~
