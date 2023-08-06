# Android image

https://opengapps.org

1. x86
2. Android 7
3. pico

~~~
Core\gmscore-x86.tar.lz
gmscore-x86\nodpi\priv-app\PrebuiltGmsCore\PrebuiltGmsCore.apk

Core\gsfcore-all.tar.lz
gsfcore-all\nodpi\priv-app\GoogleServicesFramework\GoogleServicesFramework.apk

Core\gsflogin-all.tar.lz
gsflogin-all\nodpi\priv-app\GoogleLoginService\GoogleLoginService.apk

Core\vending-x86.tar.lz
vending-x86\nodpi\priv-app\Phonesky\Phonesky.apk
~~~

1. Android Studio
2. Pixel 3a XL
3. API Level 24
4. Android 7 image

since we are loading Google Play ourself, we can just use the current revision
of the system image. Now, start the device:

~~~
emulator -list-avds
emulator -avd Pixel_3a_XL_API_24 -writable-system
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

After reboot, start Google Play Store and try to pull details for an app.
results:

Android image | open gapps | result
--------------|------------|-------
7             | 7          | pass
