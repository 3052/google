# september 24 2023

supposedly Android 9 (API 28) can do `armeabi-v7a`

I wonder if we could create an Android 9 device to use with this module that
can get any app?

~~~
play -d com.madhead.tos.zh -v 18775 -p 1
play -d com.axis.drawingdesk.v3 -v 220 -p 1
~~~

result:

~~~
package: name='com.axis.drawingdesk.v3' versionCode='220' versionName='7.1.0'
compileSdkVersion='33' compileSdkVersionCodename='13'
sdkVersion:'23'

package: name='com.madhead.tos.zh' versionCode='18775' versionName='2023.501'
compileSdkVersion='33' compileSdkVersionCodename='13'
sdkVersion:'23'
~~~

tests:

~~~
platform:1,size: 95833969,downloads:15433116,doc:"com.madhead.tos.zh"},
platform:1,size:110846862,downloads:17990819,doc:"com.axis.drawingdesk.v3"},

sanity check:
platform:1,size: 18708178,downloads:1563645747,doc:"com.miui.weather2"},
~~~

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

~~~
adb root
adb disable-verity
adb shell avbctl disable-verification
adb reboot

adb root
adb remount
adb push com.android.vending_11.9.30.apk /system/priv-app
adb reboot
~~~

reboot fails

~~~
privileged permissions not in privapp-permissions allowlist
~~~

more:

~~~
09-25 20:45:04.874  1550  1550 E AndroidRuntime:
java.lang.IllegalStateException: Signature|privileged permissions not in
privapp-permissions allowlist: {com.android.vending
(/system/priv-app/com.android.vending_22.8.44.apk):
com.android.permission.USE_INSTALLER_V2, com.android.vending
~~~

https://source.android.com/docs/core/permissions/perms-allowlist
