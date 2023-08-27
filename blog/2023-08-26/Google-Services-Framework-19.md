# Google Services Framework 19 with new stuff

what does Google APIs image come with?

~~~
> adb shell dumpsys package com.google.android.gsf | Select-String version
versionCode=19 targetSdk=19
versionName=4.4.2-6761301

> adb shell dumpsys package com.google.android.gms | Select-String version
versionCode=202414005 targetSdk=30
versionName=20.24.14 (000700-319035315)
~~~

what does Open GApps (Android 4.4) come with?

~~~
package: name='com.google.android.gsf' versionCode='19'
versionName='4.4.4-1227136'

package: name='com.google.android.gms' versionCode='220221005'
versionName='22.02.21 (000700-428111784)' platformBuildVersionName='Tiramisu'

package: name='com.android.vending' versionCode='83032100'
versionName='30.3.21-19 [0] [PR] 445437866' platformBuildVersionName='Tiramisu'
~~~

so lets go with Open GApps:

~~~
open_gapps-x86-4.4-pico-20220503.zip\Core\gsfcore-all.tar.lz\gsfcore-all.tar\
gsfcore-all\nodpi\priv-app

open_gapps-x86-4.4-pico-20220503.zip\Core\gmscore-x86.tar.lz\gmscore-x86.tar\
gmscore-x86\nodpi\priv-app

open_gapps-x86-4.4-pico-20220503.zip\Core\gmscore-x86.tar.lz\gmscore-x86.tar\
gmscore-x86\common\lib

open_gapps-x86-4.4-pico-20220503.zip\Core\vending-x86.tar.lz\vending-x86.tar\
vending-x86\nodpi\priv-app

open_gapps-x86-4.4-pico-20220503.zip\Core\vending-x86.tar.lz\vending-x86.tar\
vending-x86\common\lib
~~~

1. Android Studio
2. Android 4.4

then:

~~~
emulator -avd Pixel_3a_XL_API_19 -writable-system
~~~

then:

~~~
adb remount

adb push lib /system
adb push priv-app /system
adb reboot
~~~

Play Store will not open. we cant try GenyMotion, because you can only create a
Virtual device back to Android 5. lets try the Google APIs image. cant sign in:

~~~
HTTP/1.1 403 Forbidden

Error=BadAuthentication
~~~
