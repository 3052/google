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

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r23.zip>

~~~
package: name='com.google.android.gms' versionCode='9452030'
versionName='9.4.52 (030-127739847)' platformBuildVersionName='6.0-2166767'
~~~

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r24.zip>

~~~
package: name='com.google.android.gms' versionCode='9683030'
versionName='9.6.83 (030-133155058)' platformBuildVersionName='N'
~~~

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r25.zip>

~~~
package: name='com.google.android.gms' versionCode='9877030'
versionName='9.8.77 (030-135396225)'
~~~

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r26.zip>

~~~
package: name='com.google.android.gms' versionCode='10084030'
versionName='10.0.84 (030-138685555)'
~~~

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r27.zip>

~~~
package: name='com.google.android.gms' versionCode='10298030'
versionName='10.2.98 (030-146496160)' platformBuildVersionName='7.1.1'
~~~

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r28.zip>

~~~
package: name='com.google.android.gms' versionCode='11055030'
versionName='11.0.55 (030-156917137)' platformBuildVersionName='O'
~~~

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r29.zip>

~~~
package: name='com.google.android.gms' versionCode='11055030'
versionName='11.0.55 (030-156917137)' platformBuildVersionName='O'
~~~

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r30.zip>

~~~
package: name='com.google.android.gms' versionCode='11302030'
versionName='11.3.02 (030-161239932)' platformBuildVersionName='8.0.0'
~~~

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r31.zip>

~~~
package: name='com.google.android.gms' versionCode='11509030'
versionName='11.5.09 (030-164803921)' platformBuildVersionName='8.0.0'
~~~

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r32.zip>

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r33.zip>

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r34.zip>

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r35.zip>

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r36.zip>

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r37.zip>

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r38.zip>

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r39.zip>

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r40.zip>

~~~
package: name='com.google.android.gms' versionCode='202414005'
versionName='20.24.14 (000700-319035315)' compileSdkVersion='30'
~~~

this seems to use `EncryptedPasswd`, which is unexpected.

## API 21 (Android 5)

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r15.zip>

~~~
package: name='com.google.android.gms' versionCode='9452270'
versionName='9.4.52 (270-127739847)' platformBuildVersionName='6.0-2166767'
~~~

reboot fails

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r16.zip>

~~~
package: name='com.google.android.gms' versionCode='9683270'
versionName='9.6.83 (270-133155058)' platformBuildVersionName='N'
~~~

reboot fails

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r17.zip>

~~~
package: name='com.google.android.gms' versionCode='9877270'
versionName='9.8.77 (270-135396225)'
~~~

reboot fails

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r18.zip>

~~~
package: name='com.google.android.gms' versionCode='10084270'
versionName='10.0.84 (270-138685555)'
~~~

remount fails

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r19.zip>

~~~
package: name='com.google.android.gms' versionCode='10298270'
versionName='10.2.98 (270-146496160)' platformBuildVersionName='7.1.1'
~~~

reboot fails

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r20.zip>

~~~
package: name='com.google.android.gms' versionCode='11055270'
versionName='11.0.55 (270-156917137)' platformBuildVersionName='O'
~~~

success

---------------------------------------------------------------------------------

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r32.zip>

~~~
package: name='com.google.android.gms' versionCode='202414013'
versionName='20.24.14 (020700-319035315)' compileSdkVersion='30'
~~~
