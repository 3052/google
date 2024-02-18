# Android 9

here is the current version of System Image API 28:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-28_r12.zip>

what is the oldest version?

<http://dl.google.com/android/repository/sys-img/google_apis/x86-28_r01.zip>

~~~
package: name='com.google.android.gms' versionCode='16089022'
versionName='16.0.89 (040700-239467275)' platformBuildVersionName='Q'
~~~

note if remount fails you might need to just retry it. this seems to be the
first version using `/fdfe/sync` only:

~~~
play -a com.android.vending -v 82941300
~~~

then:

~~~
emulator -avd Pixel_6_API_30 -writable-system
adb remount
adb push com.android.vending-82941300.apk /system/priv-app
adb reboot
~~~
