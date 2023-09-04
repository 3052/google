# google play services

~~~
Date: Mon, 4 Sep 2023 09:41:14 -0500
Subject: Google Services Framework
To: opensource@apkpure.com

My understanding is the Google Services Framework app (com.google.android.gsf)
is not downloadable, even using the internal Google APIs. So my question is,
where are you getting this app from?

https://apkpure.com/google-services-framework/com.google.android.gsf
~~~

- https://forum.xda-developers.com/t/extract-google-services-framework.4624097
- https://github.com/illogical-robot/apkmirror-public/issues/261

we can test with Google Play Games

## API 19 (Android 4.4)

<http://dl.google.com/android/repository/sys-img/google_apis/x86-19_r40.zip>

~~~
package: name='com.google.android.gms' versionCode='202414005'
versionName='20.24.14 (000700-319035315)' compileSdkVersion='30'
~~~

download:

~~~
play -d com.google.android.gms -v 202414005
~~~

multiple methods in the code:

~~~
sources\defpackage\ezh.java
101:a("EncryptedPasswd", hbm.a(this.c, str, str2));

sources\defpackage\fbc.java
54:a2.a("minutemaid_auto_url_override", "https://accounts.google.com/embedded/setup/v2/androidauto");
69:u = a2.a("minutemaid_tv_url_override", "https://accounts.google.com/embedded/setup/v2/androidtv");
70:v = a3.a("minutemaid_glif_url_override", "https://accounts.google.com/embedded/setup/v2/android?");
71:w = a3.a("minutemaid_url_override", "https://accounts.google.com/embedded/setup/android?");
72:x = a2.a("minutemaid_kid_sign_in_url", "https://accounts.google.com/embedded/setup/kidsignin/android");
73:y = a2.a("minutemaid_kid_sign_up_url", "https://accounts.google.com/embedded/setup/kidsignup/android");
~~~

this is used:

~~~
EncryptedPasswd
~~~

## API 21 (Android 5)

<http://dl.google.com/android/repository/sys-img/google_apis/x86-21_r20.zip>

~~~
package: name='com.google.android.gms' versionCode='11055270'
versionName='11.0.55 (270-156917137)' platformBuildVersionName='O'
~~~

download:

~~~
play -d com.google.android.gms -v 11055270
~~~

multiple methods in the code:

~~~
sources\defpackage\eeg.java
99:a = a.a("EncryptedPasswd", fde.a(a.a, str, str2));

sources\defpackage\eeu.java
119:V = kpv.a("minutemaid_url_override", "https://accounts.google.com/embedded/setup/android?");
120:W = kpv.a("minutemaid_glif_url_override", "https://accounts.google.com/embedded/setup/v2/android?");
~~~

this is used:

https://accounts.google.com/embedded/setup/android
