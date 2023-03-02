# Research

Google Play Store:
com.android.vending

Google Play services:
com.google.android.gms 

Google Services Framework:
com.google.android.gsf

pull:

~~~
adb pull `
/system/priv-app/GoogleServicesFramework/GoogleServicesFramework.apk `
GoogleServicesFramework.apk

adb pull `
/system/priv-app/PrebuiltGmsCore/PrebuiltGmsCore.apk `
PrebuiltGmsCore.apk

adb pull `
/system/priv-app/Phonesky/Phonesky.apk `
Phonesky.apk
~~~

push:

~~~
adb remount
adb push GoogleServicesFramework.apk /system/priv-app
adb push Phonesky.apk /system/priv-app
adb push PrebuiltGmsCore.apk /system/priv-app
adb reboot
~~~

## API 24

~~~
> adb shell dumpsys package com.google.android.gsf | rg versionName
    versionName=7.0

> adb shell dumpsys package com.google.android.gms | rg versionName
    versionName=11.7.43 (470-172403884)

> adb shell dumpsys package com.android.vending | rg versionName
    versionName=6.7.15.E-all [0] 2987020
> adb shell dumpsys package com.android.vending | rg versionCode
    versionCode=80671500 minSdk=14 targetSdk=23
~~~

## API 25

~~~
> adb shell dumpsys package com.google.android.gsf | rg versionName
    versionName=7.1.1
> adb shell dumpsys package com.google.android.gms | rg versionName
    versionName=10.2.98 (470-146496160)
> adb shell dumpsys package com.android.vending | rg versionName
    versionName=6.9.43.G-all [0] 3363388
~~~

## API 26

~~~
> adb shell dumpsys package com.google.android.gsf | rg versionName
    versionName=8.0.0

> adb shell dumpsys package com.google.android.gms | rg versionName
    versionName=11.7.43 (470-172403884)

> adb shell dumpsys package com.android.vending | rg versionName
    versionName=7.9.66.Q-all [0] [PR] 163928463
> adb shell dumpsys package com.android.vending | rg versionCode
    versionCode=80796600 minSdk=14 targetSdk=25
~~~

## API 27

~~~
> adb shell dumpsys package com.google.android.gsf | rg versionName
    versionName=8.1.0
> adb shell dumpsys package com.google.android.gms | rg versionName
    versionName=11.5.80 (470-175107017)
> adb shell dumpsys package com.android.vending | rg versionName
    versionName=8.0.62.R-all [0] [PR] 172052298
~~~

## API 28

~~~
> adb shell dumpsys package com.google.android.gsf | rg versionName
    versionName=9
> adb shell dumpsys package com.google.android.gms | rg versionName
    versionName=16.0.89 (040700-239467275)
> adb shell dumpsys package com.android.vending | rg versionName
    versionName=10.1.41-all [0] [PR] 197615634
~~~

## API 29

~~~
> adb shell dumpsys package com.google.android.gsf | rg versionName
    versionName=10
> adb shell dumpsys package com.google.android.gms | rg versionName
    versionName=17.7.86 (040700-256199907)
> adb shell dumpsys package com.android.vending | rg versionName
    versionName=15.2.67-all [0] [PR] 256058878
~~~

## API 30

~~~
> adb shell dumpsys package com.google.android.gsf | rg versionName
    versionName=11
> adb shell dumpsys package com.google.android.gms | rg versionName
    versionName=20.18.17 (040700-311416286)
> adb shell dumpsys package com.android.vending | rg versionName
    versionName=20.4.33-all [0] [PR] 319051143
~~~

## API 31

~~~
> adb shell dumpsys package com.google.android.gsf | rg versionName
    versionName=12
> adb shell dumpsys package com.google.android.gms | rg versionName
    versionName=21.24.23 (190800-396046673)
> adb shell dumpsys package com.android.vending | rg versionName
    versionName=25.9.49-21 [0] [PR] 386309911
~~~

## API 32

~~~
> adb shell dumpsys package com.google.android.gsf | rg versionName
    versionName=12

> adb shell dumpsys package com.google.android.gms | rg versionName
    versionName=21.24.23 (190800-396046673)

> adb shell dumpsys package com.android.vending | rg versionName
    versionName=25.9.50-21 [0] [PR] 400852117
> adb shell dumpsys package com.android.vending | rg versionCode
    versionCode=82595010 minSdk=21 targetSdk=30
~~~

For some reason, these will not download:

~~~
8_25_9_50_00
8_25_9_50_10
~~~

but these work fine:

~~~
8_25_9_19_00
8_25_9_19_10
8_25_9_29_00
8_25_9_29_10
~~~

## API 33

~~~
> adb shell dumpsys package com.google.android.gsf | rg versionName
    versionName=13

> adb shell dumpsys package com.google.android.gms | rg versionName
    versionName=22.18.21 (190800-453244992)

> adb shell dumpsys package com.android.vending | rg versionName
    versionName=30.4.17-21 [0] [PR] 445549118
> adb shell dumpsys package com.android.vending | rg versionCode
    versionCode=83041710 minSdk=21 targetSdk=31
~~~
