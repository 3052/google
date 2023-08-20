# Android image API 23

Android 6

## device Android Studio, GApps Android Studio

with Google APIs image API 23:

~~~
adb pull /system/app/PlayGames
adb pull /system/priv-app/GoogleLoginService
adb pull /system/priv-app/GoogleServicesFramework
adb pull /system/priv-app/PrebuiltGmsCore
~~~

then with Android image API 23:

~~~
emulator -avd Pixel_2_API_23 -writable-system
~~~

then:

~~~
adb remount

adb push GoogleLoginService /system/priv-app
adb push GoogleServicesFramework /system/priv-app
adb push PlayGames /system/app
adb push PrebuiltGmsCore /system/priv-app
adb reboot
~~~

then you get this:

> Unfortunately, Google Play services has stopped.

## device Android Studio, GApps GenyMotion

https://opengapps.org

~~~
open_gapps-x86-6.0-pico-20220503.zip\Core\gmscore-x86.tar.lz\gmscore-x86.tar\
   gmscore-x86\nodpi\priv-app\PrebuiltGmsCore\PrebuiltGmsCore.apk
open_gapps-x86-6.0-pico-20220503.zip\Core\gsfcore-all.tar.lz\gsfcore-all.tar\
   gsfcore-all\nodpi\priv-app\GoogleServicesFramework\GoogleServicesFramework.apk
open_gapps-x86-6.0-pico-20220503.zip\Core\vending-x86.tar.lz\vending-x86.tar\
   vending-x86\nodpi\priv-app\Phonesky\Phonesky.apk
~~~

drag to home screen to install:

~~~
GoogleServicesFramework.apk
~~~

then:

~~~
PrebuiltGmsCore.apk
~~~

then you get this:

> Unfortunately, Google Play services has stopped.

from the Google APIs image API 23, we find this:

~~~
> adb ls /system/priv-app
000041ed 00001000 5f165ca4 PrebuiltGmsCore
~~~

then:

~~~
> adb push PrebuiltGmsCore.apk /system/priv-app
adb: error: failed to copy 'PrebuiltGmsCore.apk' to
'/system/priv-app/PrebuiltGmsCore.apk': remote Read-only file system
~~~

then:

~~~
> adb remount
remount succeeded
> adb push PrebuiltGmsCore.apk /system/priv-app
adb: error: failed to copy 'PrebuiltGmsCore.apk' to
'/system/priv-app/PrebuiltGmsCore.apk': remote Read-only file system
~~~

then:

~~~
emulator -avd Pixel_2_API_23 -writable-system
~~~

then:

~~~
> adb push PrebuiltGmsCore.apk /system/priv-app
adb: error: failed to copy 'PrebuiltGmsCore.apk' to
'/system/priv-app/PrebuiltGmsCore.apk': remote Read-only file system
~~~

then:

~~~
adb remount

adb push Phonesky.apk /system/priv-app
adb push PrebuiltGmsCore.apk /system/priv-app
adb reboot
~~~

then:

> Couldn't sign in
>
> There was a problem communicating with Google servers.
>
> Try again later.

then:

~~~
adb remount

adb push GoogleServicesFramework.apk /system/priv-app
adb push Phonesky.apk /system/priv-app
adb push PrebuiltGmsCore.apk /system/priv-app
adb reboot
~~~
