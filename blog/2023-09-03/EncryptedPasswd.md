# EncryptedPasswd

the code for this is here:

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r23.zip>

~~~
x86\system.img\priv-app\PrebuiltGmsCore.apk
~~~

result:

~~~
sources\defpackage\did.java
a2 = a2.a("EncryptedPasswd", een.a(a2.a, str, str2));
~~~

version:

~~~
package: name='com.google.android.gms' versionCode='9452030'
versionName='9.4.52 (030-127739847)' platformBuildVersionName='6.0-2166767'
~~~

works with these:

~~~
sdkVersion:'9'
targetSdkVersion:'23'
~~~

can we get this working with API Level 23? choose this:

API Level | Target
----------|----------
23        | Android 6

~~~
play -d com.android.vending -v 82011800
play -d com.google.android.gms -v 9452470
~~~
