# google play services

## API 19 (Android 4.4)

system partition is small:

~~~
> adb shell df
Filesystem               Size     Used     Free   Blksize
/system                532.8M   496.9M    35.8M   4096
~~~

this is the most recent Google Play Store that will fit:

~~~
play -d com.android.vending -v 82992400 
~~~

this is the newest revision:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r40.zip>

~~~
package: name='com.google.android.gms' versionCode='202414005'
versionName='20.24.14 (000700-319035315)' compileSdkVersion='30'
~~~

## API 21 (Android 5)

this is the newest revision of API 21 (Android 5):

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r32.zip>

~~~
package: name='com.google.android.gms' versionCode='202414013'
versionName='20.24.14 (020700-319035315)' compileSdkVersion='30'
~~~

reboot fails:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r15.zip>

~~~
package: name='com.google.android.gms' versionCode='9452270'
versionName='9.4.52 (270-127739847)' platformBuildVersionName='6.0-2166767'
~~~

reboot fails:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r16.zip>

~~~
package: name='com.google.android.gms' versionCode='9683270'
versionName='9.6.83 (270-133155058)' platformBuildVersionName='N'
~~~

reboot fails:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r17.zip>

~~~
package: name='com.google.android.gms' versionCode='9877270'
versionName='9.8.77 (270-135396225)'
~~~

remount fails:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r18.zip>

~~~
package: name='com.google.android.gms' versionCode='10084270'
versionName='10.0.84 (270-138685555)'
~~~

reboot fails:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r19.zip>

~~~
package: name='com.google.android.gms' versionCode='10298270'
versionName='10.2.98 (270-146496160)' platformBuildVersionName='7.1.1'
~~~

success:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r20.zip>

~~~
package: name='com.google.android.gms' versionCode='11055270'
versionName='11.0.55 (270-156917137)' platformBuildVersionName='O'
~~~
