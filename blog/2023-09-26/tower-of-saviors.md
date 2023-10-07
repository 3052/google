# Tower of Saviors

so, if we look at these, everything seems fine:

~~~
> play -d com.madhead.tos.zh -p 1
creator: Mad Head Limited
file: APK APK
installation size: 99.60 megabyte
downloads: 15.44 million
offer: 0 USD
title: Tower of Saviors
upload date: Sep 14, 2023
version: 2023.501
version code: 18775

> play -d com.axis.drawingdesk.v3 -p 1
creator: 4Axis Technologies
file: APK APK APK APK
installation size: 136.02 megabyte
downloads: 18.23 million
offer: 0 USD
title: Drawing Desk: Draw, Paint Art
upload date: Sep 16, 2023
version: 7.1.0
version code: 220
~~~

but we cannot acquire them:

~~~
> play -d com.madhead.tos.zh -p 1 -a
POST https://android.googleapis.com/auth
POST https://play-fe.googleapis.com/fdfe/acquire
panic: Error
Your device is not compatible with this item.
~~~

others are fine:

~~~
> play -d com.miui.weather2 -p 1 -a
POST https://android.googleapis.com/auth
POST https://play-fe.googleapis.com/fdfe/acquire
~~~

If we download them, we find this:

~~~
package: name='com.axis.drawingdesk.v3' versionCode='220' versionName='7.1.0'
compileSdkVersion='33' compileSdkVersionCodename='13'
sdkVersion:'23'

package: name='com.madhead.tos.zh' versionCode='18775' versionName='2023.501'
compileSdkVersion='33' compileSdkVersionCodename='13'
sdkVersion:'23'
~~~

now, these are also compatible with `arm64-v8a`:

~~~
> play -d com.madhead.tos.zh -p 2
creator: Mad Head Limited
file: APK APK
installation size: 100.47 megabyte
downloads: 15.44 million
offer: 0 USD
title: Tower of Saviors
upload date: Sep 14, 2023
version: 2023.501
version code: 18775

> play -d com.axis.drawingdesk.v3 -p 2
creator: 4Axis Technologies
file: APK APK APK APK
installation size: 155.79 megabyte
downloads: 18.23 million
offer: 0 USD
title: Drawing Desk: Draw, Paint Art
upload date: Sep 16, 2023
version: 7.1.0
version code: 220
~~~

here is the best option:

API Level | ABI         | Target
----------|-------------|----------------------
28        | x86         | Android 9 Google Play

pass:

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://play.google.com/store/apps/details?id=com.miui.weather2
~~~

fail:

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://play.google.com/store/apps/details?id=com.madhead.tos.zh
~~~

> Your device isn't compatible with this version.

here is the next best option:

API Level | ABI         | Target
----------|-------------|------------------------
28        | `x86_64`    | Android 9 Google Play

these both fail:

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://play.google.com/store/apps/details?id=com.miui.weather2

adb shell am start -a android.intent.action.VIEW `
-d https://play.google.com/store/apps/details?id=com.madhead.tos.zh
~~~

here is the next best option:

API Level | ABI         | Target
----------|-------------|-----------------------
29        | x86         | Android 10 Google Play

these both fail:

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://play.google.com/store/apps/details?id=com.miui.weather2

adb shell am start -a android.intent.action.VIEW `
-d https://play.google.com/store/apps/details?id=com.madhead.tos.zh
~~~

here is the next best option:

API Level | ABI         | Target
----------|-------------|-----------------------
29        | `x86_64`    | Android 10 Google Play

these both fail:

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://play.google.com/store/apps/details?id=com.miui.weather2

adb shell am start -a android.intent.action.VIEW `
-d https://play.google.com/store/apps/details?id=com.madhead.tos.zh
~~~

here is the next best option:

API Level | ABI         | Target
----------|-------------|-----------------------
30        | x86         | Android 11 Google Play

these both work with device Pixel 2:

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://play.google.com/store/apps/details?id=com.miui.weather2

adb shell am start -a android.intent.action.VIEW `
-d https://play.google.com/store/apps/details?id=com.madhead.tos.zh
~~~

here is the next best option:

API Level | ABI         | Target
----------|-------------|-----------------------
30        | x86         | Android 11 Google APIs

here is the target:

~~~
> adb shell dumpsys package com.android.vending | Select-String version
versionCode=82043300 minSdk=16 targetSdk=29
versionName=20.4.33-all [0] [PR] 319051143
~~~

download:

~~~
play -d com.android.vending -v 82051400
~~~

Note starting with API 29 (Android 10), you need to also push permission:

~~~
adb root
adb shell avbctl disable-verification
adb reboot

adb root
adb remount
adb push com.android.vending_11.9.30.apk /system/priv-app
adb push privapp-permissions.xml /etc/permissions
adb reboot
~~~

https://source.android.com/docs/core/permissions/perms-allowlist

these both work with device Pixel 2:

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://play.google.com/store/apps/details?id=com.miui.weather2

adb shell am start -a android.intent.action.VIEW `
-d https://play.google.com/store/apps/details?id=com.madhead.tos.zh
~~~
