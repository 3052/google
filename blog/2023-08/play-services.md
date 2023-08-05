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

## Android Studio, API Level 24, Google Play

Pixel 2

<http://dl.google.com/android/repository/sys-img/google_apis_playstore/x86-24_r19.zip>

~~~
versionCode=11743470 minSdk=23 targetSdk=23
versionName=11.7.43 (470-172403884)
~~~

result:

~~~
POST /auth HTTP/1.1
app: com.google.android.gms
Accept-Encoding: identity
User-Agent: GoogleAuth/1.4 (generic_x86 NYC); gzip
content-type: application/x-www-form-urlencoded
Host: android.clients.google.com
Connection: Keep-Alive
device: 3d50b2d51...

ACCESS_TOKEN=1&
add_account=1&
callerPkg=com.google.android.gms&
callerSig=38918a453d07199354f8b19af05ec6562ced5788&
device_country=us&
droidguard_results=Cgb6IxHXwjHSEFsAAHdU9h5a1xWSAFpFJRNTS5M9AC4R4ZGh9VMLBADDJ6...&
get_accountid=1&
google_play_services_version=11743470&
lang=en_US&
sdk_version=24&
service=ac2dm&
Email=s...&
androidId=3d50b2d51...&
Token=oauth2_4%2F0Adeu5BV2UZd15bnJ2KlF38PzNNhldJnw8ynjIOlJDbbjTemHwvM2FmzapRFp...
~~~

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

## Android Studio, API Level 24, Google APIs

Pixel 3a XL

~~~
versionCode=202414022 minSdk=23 targetSdk=30
versionName=20.24.14 (040700-319035315)
~~~

## Android Studio, API Level 23, Google APIs

Pixel 3a XL

~~~
versionCode=202414022 targetSdk=30
versionName=20.24.14 (040700-319035315)
~~~

## Android Studio, API Level 22, Google APIs

Pixel 3a XL

~~~
versionCode=202414013 targetSdk=30
versionName=20.24.14 (020700-319035315)
~~~

## Android Studio, API Level 21, Google APIs

Pixel 3a XL

~~~
versionCode=202414013 targetSdk=30
versionName=20.24.14 (020700-319035315)
~~~

## Android Studio, API Level 19, Google APIs

Pixel 3a XL

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r40.zip>

~~~
versionCode=202414005 targetSdk=30
versionName=20.24.14 (000700-319035315)
~~~

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r23.zip>

~~~
versionCode=9452030 targetSdk=23
versionName=9.4.52 (030-127739847)
~~~

## Android Studio, API Level 18, Google APIs

Pixel 3a XL

~~~
versionCode=9256030 targetSdk=23
versionName=9.2.56 (030-124593566)
~~~
