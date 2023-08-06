# google play services 10

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

1. Android Studio
2. Pixel 2
3. API Level 24
4. `google_apis_playstore` image

<http://dl.google.com/android/repository/sys-img/google_apis_playstore/x86-24_r12.zip>

~~~
name='com.google.android.gms'
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
