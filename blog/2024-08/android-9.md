# Android 9

note if remount fails you might need to just retry it

https://android.stackexchange.com/questions/218850/android-studio-emulator

then:

https://opengapps.org

1. x86
2. android 9
3. pico

~~~
Core\gmscore-x86.tar.lz\gmscore-x86.tar\gmscore-x86\nodpi\priv-app\
PrebuiltGmsCorePi\PrebuiltGmsCorePi.apk

Core\gsfcore-all.tar.lz\gsfcore-all.tar\gsfcore-all\nodpi\priv-app\
GoogleServicesFramework\GoogleServicesFramework.apk

Core\vending-x86.tar.lz\vending-x86.tar\vending-x86\nodpi\priv-app\
Phonesky\Phonesky.apk
~~~

yes, Phonesky is required

~~~
emulator -avd Pixel_2_API_28 -writable-system
adb root
adb remount
adb push GoogleServicesFramework.apk /system/priv-app
adb push Phonesky.apk /system/priv-app
adb push PrebuiltGmsCorePi.apk /system/priv-app
adb reboot
~~~
