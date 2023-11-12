# Java

If we create a device with API 21, this is the newest image:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r32.zip>

this is the oldest image:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r15.zip>

running the image returns this request:

~~~
POST /checkin HTTP/1.1
Host: android.clients.google.com
~~~

in the request body is this:

~~~go
protobuf.Message{
   protobuf.Field{Number: 18, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.print")},
      }},
   }},
}
~~~

the oldest image uses these:

~~~
package: name='com.google.android.gsf' versionCode='21'
versionName='5.0.2-3079185' platformBuildVersionName='5.0.2-3079185'
~~~

if we download:

http://apkmirror.com/apk/google-inc/google-services-framework/google-services-framework-5-1-1783956-release

we find this similar file:

~~~
sources\com\google\android\gsf\checkin\proto\Config.java
~~~

but it only defines up to field 17:

~~~java
if (hasMaxApkDownloadSizeMb()) {
   output.writeInt32(17, getMaxApkDownloadSizeMb());
}
~~~

which means that `com.google.android.gsf` does not handle the `/checkin` request.
Its actually handled by:

~~~
package: name='com.google.android.gms' versionCode='9452270'
versionName='9.4.52 (270-127739847)' platformBuildVersionName='6.0-2166767'
~~~

here:

~~~
sources\p000\fmn.java
~~~

example:

~~~java
private List f50711w = Collections.emptyList();

for (fmo fmoVar : this.f50711w) {
   ajfhVar.m32925b(26, fmoVar);
}
~~~

field 26 was introduced with Google Play Services 9.2.55, but it was already
obfuscated. note the same issue exists with other fields. For example field 7 was
introduced with Google Play Services 3.1.58, but it was already obfuscated.

~~~java
if (hasTouchScreen()) {
    output.writeInt32(1, getTouchScreen());
}
if (hasKeyboard()) {
    output.writeInt32(2, getKeyboard());
}
if (hasNavigation()) {
    output.writeInt32(3, getNavigation());
}
if (hasScreenLayout()) {
    output.writeInt32(4, getScreenLayout());
}
if (hasHasHardKeyboard()) {
    output.writeBool(5, getHasHardKeyboard());
}
if (hasHasFiveWayNavigation()) {
    output.writeBool(6, getHasFiveWayNavigation());
}
if (hasScreenDensity()) {
    output.writeInt32(7, getScreenDensity());
}
~~~
