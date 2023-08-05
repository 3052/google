# google play services

with Google Play versions, we cannot install system certificate:

~~~
adbd cannot run as root in production builds
~~~

but we can install user certificate. even the user certificate allows us to
capture the `/auth` request. with the older zip files, you can extract and move
them here:

~~~
C:\Users\Steven\AppData\Local\Android\Sdk\system-images
~~~

## Android Studio, API Level 23, android

Pixel 3a XL. since we are loading Google Play ourself, we can just use the
current revision of the system image. should only need these:

~~~
GoogleServicesFramework.apk
Phonesky.apk
PrebuiltGmsCore.apk
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

After reboot, you should then be able to start Google Play Store as normal.

## Android Studio, API Level 23, `google_apis_playstore`

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
   for r := 0; r <= 40; r++ {
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

## Android Studio, API Level 24, `google_apis_playstore`

Pixel 2

<http://dl.google.com/android/repository/sys-img/google_apis_playstore/x86-24_r12.zip>

~~~
versionCode=10298470 minSdk=23 targetSdk=23
versionName=10.2.98 (470-146496160)
~~~

result:

~~~
POST /auth HTTP/1.1
app: com.google.android.gms
Accept-Encoding: identity
User-Agent: GoogleAuth/1.4 (generic_x86 NYC); gzip
content-length: 10581
content-type: application/x-www-form-urlencoded
Host: android.clients.google.com
Connection: Keep-Alive
device: 368c8cc73...

ACCESS_TOKEN=1&
add_account=1&
device_country=us&
droidguard_results=CgZVeSeMLFfSEFsAAHdU9h5a1xWSAFpFJRNTS5M9AC4R4ZGh9VMLBADDJ6...&
google_play_services_version=10298470&
lang=en_US&
sdk_version=24&
service=ac2dm&
Email=s...&
androidId=368c8cc73...&
Token=oauth2_4%2F0Adeu5BUR-6sqF749vq0-EklJ6wSilERD49fSDT1N2n-MOc0h6AGra14vLMh...&
~~~

## BiTGApps

~~~
name='com.google.android.gms'
versionCode='17455017'
versionName='17.4.55 (040306-248795830)'
~~~

## MindTheGapps

~~~
name='com.google.android.gms'
versionCode='11976970'
versionName='11.9.76 (970-184349000)'
~~~

## NikGApps

~~~
name='com.google.android.gms'
versionCode='19528039'
versionName='19.5.28 (100408-273329419)'
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
