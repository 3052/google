# SDK version 23

1. Android Studio
2. Pixel 3a XL
3. API Level 24
4. Android 7 Google APIs image

swap image for this:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-24_r09.zip>

~~~
x86\system.img\priv-app\PrebuiltGmsCore\PrebuiltGmsCore.apk
~~~

check version:

~~~
name='com.google.android.gms'
versionCode='9877470'
versionName='9.8.77 (470-135396225)'
~~~

some older revisions are available, but they are buggy. then:

~~~
play -d com.android.vending -v 82011800
~~~

check version:

~~~
name='com.android.vending'
versionCode='82011800'
versionName='20.1.18-all [0] [PR] 311592326'
~~~

older versions are available, but they are buggy. You can also pull the APK from
the `google_apis_playstore` image, but it is buggy as well. start the device:

~~~
emulator -list-avds
emulator -avd Pixel_3a_XL_API_24 -writable-system
~~~

Install like this:

~~~
adb root
adb remount
adb push com.android.vending-82011800.apk /system/priv-app/Phonesky.apk
adb reboot
~~~

start Play Store. click Sign in. enter Email and click Next. Enter password and
click Next. if app closes for update, start again. click Sign in. enter Email and
click Next. Enter password and click Next. click I agree. Under Use basic device
backup, move slider to left. click ACCEPT. under Apps, click an app.
