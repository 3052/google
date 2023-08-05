# API Level 23

1. Android Studio
2. Pixel 3a XL
3. Android 6 image

https://sourceforge.net/projects/opengapps/files/x86/20190221

since we are loading Google Play ourself, we can just use the current revision
of the system image. should only need these:

~~~
Core\gmscore-x86.tar.lz
gmscore-x86\nodpi\priv-app\PrebuiltGmsCore\PrebuiltGmsCore.apk

Core\gsfcore-all.tar.lz
gsfcore-all\nodpi\priv-app\GoogleServicesFramework\GoogleServicesFramework.apk

Core\vending-x86.tar.lz
vending-x86\nodpi\priv-app\Phonesky\Phonesky.apk
~~~

Now, start the device:

~~~
emulator -list-avds
emulator -avd Pixel_3a_XL_API_23 -writable-system
~~~

Install like this:

~~~
adb root
adb remount
adb push Phonesky.apk /system/priv-app
adb reboot
~~~

After reboot, start Google Play Store. If you get to Sign in screen then its
working. this combination works:

~~~
name='com.google.android.gsf'
versionCode='23'
versionName='6.0.1'

name='com.android.vending'
versionCode='81362000'
versionName='13.6.20-all [0] [PR] 233508395'

name='com.google.android.gms'
versionCode='15090022'
versionName='15.0.90 (040700-231259764)'
~~~

this file works:

~~~
play -d com.google.android.gms -v 15090022
~~~

this command works:

~~~
play -d com.android.vending -v 81362000
~~~

this command fails:

~~~
play -d com.google.android.gsf -v 23
~~~
