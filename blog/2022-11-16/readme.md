# November 16 2022

https://xda-developers.com/download-google-apps-gapps

## BiT GApps

Only ARM is offered:

https://github.com/BiTGApps/BiTGApps/issues/4

## Flame GApps

Only ARM is offered:

https://pling.com/p/1341681

## Lite GApps

all the Lite GApps use this Play Store:

~~~
name='com.android.vending' versionCode='83032110'
versionName='30.3.21-21 [0] [PR] 445437866' platformBuildVersionName='Tiramisu'
~~~

closest match with Android Studio is API 33:

~~~
> adb shell dumpsys package com.android.vending | rg versionName
    versionName=30.4.17-21 [0] [PR] 445549118
> adb shell dumpsys package com.android.vending | rg versionCode
    versionCode=83041710 minSdk=21 targetSdk=31
~~~

but Lite GApps does not offer API 33:

https://github.com/litegapps1/litegapps1.github.io/issues/1

## Mind The GApps

Only ARM is offered:

http://downloads.codefi.re/jdcteam/javelinanddart/gapps

## Nik GApps

Only ARM is offered:

https://sourceforge.net/projects/nikgapps/files/Releases

## Open GApps

all the Open GApps use this Play Store:

~~~
name='com.android.vending' versionCode='83032100'
versionName='30.3.21-19 [0] [PR] 445437866' platformBuildVersionName='Tiramisu'
~~~

closest match with Android Studio is API 33:

~~~
> adb shell dumpsys package com.google.android.gsf | rg versionName
    versionName=13

> adb shell dumpsys package com.android.vending | rg versionName
    versionName=30.4.17-21 [0] [PR] 445549118
> adb shell dumpsys package com.android.vending | rg versionCode
    versionCode=83041710 minSdk=21 targetSdk=31
~~~

but Open GApps does not offer Android 13:

https://github.com/opengapps/opengapps/issues/973
