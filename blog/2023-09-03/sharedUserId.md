# google play services

https://developer.android.com/tools/adb#pm

~~~
play -d com.google.android.gms -v 9452270
~~~

first lets try readable system. drag app to home screen:

~~~
INSTALL_FAILED_VERSION_DOWNGRADE
~~~

then:

~~~
> adb install com.google.android.gms-9452270.apk
Failure [INSTALL_FAILED_DUPLICATE_PERMISSION
perm=com.google.android.gms.WRITE_VERIFY_APPS_CONSENT
pkg=com.google.android.gms]
~~~

then:

~~~
> adb uninstall com.google.android.gms
DELETE_FAILED_DEVICE_POLICY_MANAGER
~~~

then:

~~~
adb shell pm disable-user com.google.android.gms
~~~

then:

~~~
> adb uninstall com.google.android.gms
Failure [DELETE_FAILED_INTERNAL_ERROR]
~~~

then:

~~~
adb shell pm uninstall --user 0 com.google.android.gms
Success
~~~

this worked:

~~~
adb remount

adb shell rm -r /system/app/PlayGames
adb shell rm -r /system/priv-app/PrebuiltGmsCore
adb reboot
~~~

then:

~~~
> adb install com.google.android.gms-9452270.apk
Failure [INSTALL_FAILED_SHARED_USER_INCOMPATIBLE]
~~~

problem:

~~~xml
android:sharedUserId="com.google.uid.shared" android:versionCode="202414013"
android:sharedUserLabel="@LEMON_transformed_from_string/common_app_name"
android:versionName="20.24.14 (020700-319035315)"
package="com.google.android.gms" platformBuildVersionCode="30"
~~~

https://explorations142.rssing.com/chan-6647101/latest.php

this command succeeds, but still doesnt fix the issue:

~~~
> adb shell setenforce 0
setenforce: SELinux is disabled
~~~
