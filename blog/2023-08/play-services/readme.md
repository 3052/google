# google play services

~~~
adb shell dumpsys package com.google.android.gms | Select-String version
~~~

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

so, the oldest one we can get from Android Studio is:

~~~
versionCode=10298470 minSdk=23 targetSdk=23
versionName=10.2.98 (470-146496160)
~~~

https://opengapps.org

I selected this option:

Platform | Android | Variant
---------|---------|--------
x86      | 6       | pico

Without Google APIs

Using the method above with Google APIs image, you still get some apps such as
YouTube. If you want to install different version of one of these apps, use
this method. On "System Image" screen, I selected this option:

API Level | ABI | Target
----------|-----|----------
24        | x86 | Android 7

but newer APIs should work as well. You need these files from the Zip archive:

~~~
Core\gmscore-x86.tar.lz
Core\vending-x86.tar.lz
~~~

Then extract these from the above files:

~~~
gmscore-x86\nodpi\priv-app\PrebuiltGmsCore\PrebuiltGmsCore.apk
vending-x86\nodpi\priv-app\Phonesky\Phonesky.apk
~~~

Use the same method above to install the APKs. After reboot, you should then be
able to install YouTube or whatever app. Note that unlike above, you dont
actually need to run the Google Play setup or even start the Google Play app at
the end.

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

## Android Studio, API Level 24, `google_apis`

Pixel 3a XL

<http://dl.google.com/android/repository/sys-img/google_apis/x86-24_r07.zip>

~~~
versionCode=9452470 minSdk=23 targetSdk=23
versionName=9.4.52 (470-127739847)
~~~

## Android Studio, API Level 24, android

Pixel 3a XL

<http://dl.google.com/android/repository/sys-img/android/x86-24_r07.zip>

~~~
Unable to find package: com.google.android.gms
~~~

## Android Studio, API Level 23, `google_apis`

Pixel 3a XL

<http://dl.google.com/android/repository/sys-img/google_apis/x86-23_r16.zip>

## Android Studio, API Level 22, `google_apis`

Pixel 3a XL

<http://dl.google.com/android/repository/sys-img/google_apis/x86-22_r09.zip>

## Android Studio, API Level 21, `google_apis`

Pixel 3a XL

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r15.zip>

## Android Studio, API Level 19, `google_apis`

Pixel 3a XL

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r23.zip>

~~~
versionCode=9452030 targetSdk=23
versionName=9.4.52 (030-127739847)
~~~

## Android Studio, API Level 18, `google_apis`

Pixel 3a XL

<http://dl.google.com/android/repository/sys-img/google_apis/x86-18_r05.zip>
