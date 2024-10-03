# Google Services Framework 10

API 29

~~~
emulator -avd Pixel_6_API_30 -writable-system
adb remount
adb push com.android.vending-82941300.apk /system/priv-app
adb push privapp-permissions.xml /etc/permissions
adb reboot
~~~

https://source.android.com/docs/core/permissions/perms-allowlist

note if remount fails you might need to just retry it. if you omit permission
file, after reboot device will just load forever. whats the newest version?
this:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-29_r12.zip>

what is the oldest version?

<http://dl.google.com/android/repository/sys-img/google_apis/x86-29_r05.zip>

this seems to be the first version using `/fdfe/sync` only:

~~~
play -i com.android.vending -v 82941300
~~~

note APK is here:

~~~
system\product\priv-app\PrebuiltGmsCore\PrebuiltGmsCore.apk
~~~

result:

~~~
package: name='com.google.android.gms' versionCode='16910022'
versionName='16.9.10 (040700-247503114)'
~~~
