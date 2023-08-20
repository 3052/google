# Android image API 22

Android 5.1

<https://wikipedia.org/wiki/Android_Lollipop>

## device Android Studio, Open GApps

https://opengapps.org

~~~
open_gapps-x86-5.1-pico-20220503.zip\Core\gmscore-x86.tar.lz\gmscore-x86.tar\
gmscore-x86\common\priv-app\PrebuiltGmsCore

open_gapps-x86-5.1-pico-20220503.zip\Core\gmscore-x86.tar.lz\gmscore-x86.tar\
gmscore-x86\nodpi\priv-app\PrebuiltGmsCore

open_gapps-x86-5.1-pico-20220503.zip\Core\gsfcore-all.tar.lz\gsfcore-all.tar\
gsfcore-all\nodpi\priv-app\GoogleServicesFramework

open_gapps-x86-5.1-pico-20220503.zip\Core\vending-x86.tar.lz\vending-x86.tar\
vending-x86\common\priv-app\Phonesky

open_gapps-x86-5.1-pico-20220503.zip\Core\vending-x86.tar.lz\vending-x86.tar\
vending-x86\nodpi\priv-app\Phonesky
~~~

then:

~~~
emulator -avd Pixel_3a_XL_API_22 -writable-system
~~~

then:

~~~
adb remount

adb push GoogleServicesFramework /system/priv-app
adb push Phonesky /system/priv-app
adb push PrebuiltGmsCore /system/priv-app
adb reboot
~~~
