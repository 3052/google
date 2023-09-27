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
----------|-------------|-----------------------
33        | `x86_64`    | Android 13 Google Play

but it fails for both of these:

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://play.google.com/store/apps/details?id=com.miui.weather2

adb shell am start -a android.intent.action.VIEW `
-d https://play.google.com/store/apps/details?id=com.madhead.tos.zh
~~~

> Your device isn't compatiblee with this version.

here is the next best option:

API Level | ABI         | Target
----------|-------------|------------------------
32        | `x86_64`    | Android 12L Google Play

these both work with device Pixel 2:

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://play.google.com/store/apps/details?id=com.miui.weather2

adb shell am start -a android.intent.action.VIEW `
-d https://play.google.com/store/apps/details?id=com.madhead.tos.zh
~~~

here is the next best option:

API Level | ABI         | Target
----------|-------------|------------------------
32        | `x86_64`    | Android 12L Google APIs
