# goals

this seems to be the first version using `/fdfe/sync` only:

~~~
play -d com.android.vending -v 82941300
~~~

here is the current version of System Image API 28:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-28_r12.zip>

what is the oldest version?

<http://dl.google.com/android/repository/sys-img/google_apis/x86-28_r01.zip>

~~~
package: name='com.google.android.gms' versionCode='16089022'
versionName='16.0.89 (040700-239467275)' platformBuildVersionName='Q'
~~~

problem:

~~~
[adb shell mkdir -p /data/local/tmp/cacerts]
[adb shell cp /system/etc/security/cacerts/* /data/local/tmp/cacerts]
[adb push C:\Users\Steven\.mitmproxy\mitmproxy-ca-cert.cer /data/local/tmp/cacerts/c8750f0d.0]
C:\Users\Steven\.mitmproxy\mitmproxy-ca-cert.cer: 1 file pushed, 0 skipped. 2.2 MB/s (1172 bytes in 0.001s)
[adb root]
restarting adbd as root
[adb wait-for-device]
[adb shell mount -t tmpfs tmpfs /system/etc/security/cacerts]
[adb shell cp /data/local/tmp/cacerts/* /system/etc/security/cacerts]
cp: bad '/data/local/tmp/cacerts/*': No such file or directory
~~~
