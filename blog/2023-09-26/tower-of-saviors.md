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

---------------------------------------------------------------------------------

https://source.android.com/docs/core/permissions/perms-allowlist

supposedly Android 9 (API 28) can do `armeabi-v7a`

I wonder if we could create an Android 9 device to use with this module that
can get any app?

API 24:

~~~
> adb shell df
Filesystem     1K-blocks    Used Available Use% Mounted on
/dev/block/vda   2539312 2506272     16656 100% /system
~~~

API 25:

~~~
> adb shell df
Filesystem     1K-blocks    Used Available Use% Mounted on
/dev/block/vda   3047184 2784668    246132  92% /system
~~~

> Your device isn't compatiblee with this version.

API 26:

> Your device isn't compatiblee with this version.

API 28:

> Your device isn't compatiblee with this version.

API 29 (Android 10):


> Your device isn't compatiblee with this version.

API 30:

> Your device isn't compatiblee with this version.

API 31:

> Your device isn't compatiblee with this version.

API 32:

> Your device isn't compatiblee with this version.

API 33:

> Your device isn't compatiblee with this version.

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

tests:

~~~
sanity check:
platform:1,size: 18708178,downloads:1563645747,doc:"com.miui.weather2"},

small:
armeabi-v7a only:
platform:1,size: 95833969,downloads:15433116,doc:"com.madhead.tos.zh"},

armeabi-v7a only:
platform:1,size:110846862,downloads:17990819,doc:"com.axis.drawingdesk.v3"},
~~~
