# November 17 2022

We want to match the Android Studio pull with something. We need to pull direct
from Google Play. Since we cant get version history, we need to brute force it.
Whats the oldest version we care about? This:

~~~
> adb shell dumpsys package com.android.vending | rg versionName
    versionName=6.7.15.E-all [0] 2987020
> adb shell dumpsys package com.android.vending | rg versionCode
    versionCode=80671500 minSdk=14 targetSdk=23
~~~

whats the newest version we care about? This:

~~~
> adb shell dumpsys package com.android.vending | rg versionName
    versionName=30.4.17-21 [0] [PR] 445549118
> adb shell dumpsys package com.android.vending | rg versionCode
    versionCode=83041710 minSdk=21 targetSdk=31
~~~

We cannot use `/fdfe/details` for this, as it fails even with `vc`. So we need
to use `/fdfe/delivery`.
