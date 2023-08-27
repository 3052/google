# august 27 2023

these are the involved packages:

~~~
package: name='com.google.android.gsf' versionCode='21'
versionName='5.0.2-1649326' platformBuildVersionName='5.0.2-1649326'

package: name='com.google.android.gms' versionCode='214857013'
versionName='21.48.57 (020700-424705839)' platformBuildVersionName='Tiramisu'

package: name='com.android.vending' versionCode='83032110'
versionName='30.3.21-21 [0] [PR] 445437866' platformBuildVersionName='Tiramisu'
~~~

## Google Services Framework 21 with old stuff

uses bot guard

## Google Services Framework 19 with new stuff

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

then:

~~~
emulator -avd Pixel_3a_XL_API_21 -writable-system
~~~

then:

~~~
adb remount
adb push com.android.vending_10.8.50.apk /system/priv-app
adb reboot
~~~

then enable proxy, then install system certificate.
