# Android image API 22

Android 5.1

<https://wikipedia.org/wiki/Android_Lollipop>

## Visual Studio Emulator for Android

> We recommend you use Google’s emulator when you can, as it offers access to the
> latest Android OS images and Google Play services.

https://visualstudio.microsoft.com/vs/msft-android-emulator

## device GenyMotion, GApps GenyMotion

click Open GAPPS

## device Android Studio, GApps GenyMotion

https://opengapps.org

~~~
open_gapps-x86-5.1-pico-20220503.zip\Core\gmscore-x86.tar.lz\gmscore-x86.tar\
   gmscore-x86\nodpi\priv-app\PrebuiltGmsCore\PrebuiltGmsCore.apk
open_gapps-x86-5.1-pico-20220503.zip\Core\gsfcore-all.tar.lz\gsfcore-all.tar\
   gsfcore-all\nodpi\priv-app\GoogleServicesFramework\GoogleServicesFramework.apk
open_gapps-x86-5.1-pico-20220503.zip\Core\gsflogin-all.tar.lz\gsflogin-all.tar\
   gsflogin-all\nodpi\priv-app\GoogleLoginService\GoogleLoginService.apk
open_gapps-x86-5.1-pico-20220503.zip\Core\vending-x86.tar.lz\vending-x86.tar\
   vending-x86\nodpi\priv-app\Phonesky\Phonesky.apk
~~~

from the Google APIs image API 22, we find this:

~~~
> adb ls /system/priv-app
000041ed 00001000 5f165ca4 PrebuiltGmsCore
~~~

then:

~~~
emulator -avd Pixel_2_API_22 -writable-system
~~~

then:

~~~
adb remount

adb push GoogleLoginService.apk /system/priv-app
adb push GoogleServicesFramework.apk /system/priv-app
adb push Phonesky.apk /system/priv-app
adb push PrebuiltGmsCore.apk /system/priv-app
adb reboot
~~~

then you get this:

> Unfortunately, Google Play services has stopped.
