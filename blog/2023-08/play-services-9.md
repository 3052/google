# google play services 9

## Android Studio, `google_apis` image

1. <http://dl.google.com/android/repository/sys-img/google_apis_playstore/x86-24_r19.zip>
2. <http://dl.google.com/android/repository/sys-img/google_apis_playstore/x86-24_r12.zip>

~~~
x86\system.img\priv-app\Phonesky\Phonesky.apk
~~~

newest and oldest have same version:

~~~
name='com.android.vending'
versionCode='80671500'
versionName='6.7.15.E-all [0] 2987020'
~~~

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

some older revisions are available, but they are buggy. start the device:

~~~
emulator -list-avds
emulator -avd Pixel_3a_XL_API_24 -writable-system
~~~

Install like this:

~~~
adb root
adb remount
adb push Phonesky.apk /system/priv-app
adb reboot
~~~

on reboot you get message Google Play Store has stopped. click Close app. then
start Play Store. you get message Google Play Store has stopped. then:

13.3.16 and below fail:

~~~
> play -d com.android.vending -v 81331600
POST https://android.googleapis.com/auth
GET https://play-fe.googleapis.com/fdfe/delivery?doc=com.android.vending&vc=81331600
panic: appDeliveryData not found
~~~

13.3.17 and above pass:

~~~
> play -d com.android.vending -v 81331700
POST https://android.googleapis.com/auth
GET https://play-fe.googleapis.com/fdfe/delivery?doc=com.android.vending&vc=81331700
HEAD https://play.googleapis.com/download/by-token/download?token=AOTCm0T0llLW...
~~~

start the device:

~~~
emulator -avd Pixel_3a_XL_API_24 -writable-system
~~~

Install like this:

~~~
adb root
adb remount
adb push com.android.vending-81331700.apk /system/priv-app/Phonesky.apk
adb reboot
~~~

start Play Store. enter Email and click Next. Enter password and click Next.
click I agree. under Automatically back up device data, move slider to left.
click NEXT. under Apps, click an app. under Google Play Store has stopped,
click Open app again. under Apps, click an app. under Google Play Store keeps
stopping, click Close app. then:

https://opengapps.org

1. x86
2. Android 7
3. pico

~~~
Core\vending-x86.tar.lz\vending-x86\nodpi\priv-app\Phonesky\Phonesky.apk
~~~

check version:

~~~
name='com.android.vending'
versionCode='83032110'
versionName='30.3.21-21 [0] [PR] 445437866'
~~~

start the device:

~~~
emulator -avd Pixel_3a_XL_API_24 -writable-system
~~~

Install like this:

~~~
adb root
adb remount
adb push Phonesky.apk /system/priv-app
adb reboot
~~~

start Play Store. then click Sign in. enter Email and click Next. Enter
password and click Next. click I agree. under Automatically back up device
data, move slider to left. click NEXT. under Apps, click an app.

## Android Studio, `google_apis_playstore` image

starting with this:

<http://dl.google.com/android/repository/sys-img/google_apis_playstore/x86-24_r12.zip>

this Go program fails:

~~~go
package main

import (
   "fmt"
   "net/http"
   "time"
)

func main() {
   req, err := http.NewRequest("HEAD", "http://dl.google.com", nil)
   if err != nil {
      panic(err)
   }
   for r := 0; r <= 64; r++ {
      req.URL.Path = "/android/repository/sys-img/google_apis_playstore"
      req.URL.Path += fmt.Sprintf("/x86-23_r%02v.zip", r)
      res, err := http.DefaultClient.Do(req)
      if err != nil {
         panic(err)
      }
      fmt.Println(res.Status, req.URL)
      if res.StatusCode == http.StatusOK {
         break
      }
      time.Sleep(time.Second)
   }
}
~~~

## Open GApps

https://sourceforge.net/projects/opengapps/files/x86

~~~
Core\gmscore-x86.tar.lz
~~~

then:

~~~
gmscore-x86\nodpi\priv-app\PrebuiltGmsCore\PrebuiltGmsCore.apk
~~~

result:

~~~
versionCode='18382005'
versionName='18.3.82 (000700-260264002)'
~~~

<https://sourceforge.net/projects/opengapps/files/x86_64>

~~~
versionCode='19831014'
versionName='19.8.31 (020800-284611645)'
~~~

https://github.com/opengapps/opengapps/issues/982

## FlameGApps

~~~
name='com.google.android.gms'
versionCode='19829028'
versionName='19.8.29 (100400-282600551)'
~~~

https://github.com/flamegapps/flamegapps/issues/11

## BiTGApps

~~~
name='com.google.android.gms'
versionCode='17455017'
versionName='17.4.55 (040306-248795830)'
~~~

https://bitgapps.io

## MindTheGapps

~~~
name='com.google.android.gms'
versionCode='11976970'
versionName='11.9.76 (970-184349000)'
~~~

https://mindthegapps.com

## NikGApps

~~~
name='com.google.android.gms'
versionCode='19528039'
versionName='19.5.28 (100408-273329419)'
~~~

https://nikgapps.com
