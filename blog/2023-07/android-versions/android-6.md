# Android 6

https://opengapps.org

1. x86
2. Android 6
3. pico

~~~
Core\gmscore-x86.tar.lz
gmscore-x86\nodpi\priv-app\PrebuiltGmsCore\PrebuiltGmsCore.apk

Core\gsfcore-all.tar.lz
gsfcore-all\nodpi\priv-app\GoogleServicesFramework\GoogleServicesFramework.apk

Core\vending-x86.tar.lz
vending-x86\nodpi\priv-app\Phonesky\Phonesky.apk
~~~

1. Android Studio
2. Pixel 3a XL
3. API Level 23
4. Android 6 image

since we are loading Google Play ourself, we can just use the current revision
of the system image. Now, start the device:

~~~
emulator -list-avds
emulator -avd Pixel_3a_XL_API_23 -writable-system
~~~

Install like this:

~~~
adb root
adb remount
adb push GoogleServicesFramework.apk /system/priv-app
adb push Phonesky.apk /system/priv-app
adb push PrebuiltGmsCore.apk /system/priv-app
adb reboot
~~~

After reboot, start Google Play Store and try to pull details for an app. now
that we have this working, what version is it?

~~~
name='com.google.android.gms'
versionCode='220920022'
versionName='22.09.20 (040700-434869283)'
~~~

what are the versions for this:

1. Android Studio
2. Pixel 3a XL
3. API Level 23
4. Android 6 Google APIs image

<http://dl.google.com/android/repository/sys-img/google_apis/x86-23_r33.zip>

~~~
versionCode=202414022
versionName=20.24.14 (040700-319035315)
~~~

<http://dl.google.com/android/repository/sys-img/google_apis/x86-23_r16.zip>

~~~
versionCode=9452470
versionName=9.4.52 (470-127739847)
~~~
