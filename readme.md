# Google Play

> I’m not really sure what the point of this video is, but I guess just be
> generous. be kind to people, because you never know how much they might need
> it, or how far it’ll go.
>
> [NakeyJakey](//youtube.com/watch?v=Cr0UYNKmrUs) (2018)

download APK from Google Play or send API requests

## tool examples

[sign in](//accounts.google.com/embedded/setup/v2/android) with your Google
Account. then get authorization code (`oauth_token`) cookie from
[browser&nbsp;storage][1]. should be valid for 10 minutes. then exchange
authorization code for refresh token (`aas_et`):

~~~
play -token oauth2_4/0Adeu5B...
~~~

[1]://firefox-source-docs.mozilla.org/devtools-user/storage_inspector

do `/checkin` request:

~~~
play -checkin
~~~

do `/fdfe/sync` request:

~~~
play -sync
~~~

get app details:

~~~
> play -i com.google.android.youtube
details[8] = 0 USD
details[13][1][4] = 19.40.35
details[13][1][16] = Oct 7, 2024
details[13][1][17] = APK APK APK APK
details[13][1][82][1][1] = 8.0 and up
details[15][18] = http://www.google.com/policies/privacy
downloads = 17.34 billion
name = YouTube
size = 122.17 megabyte
version code = 1548869056
~~~

acquire app. only needs to be done once per Google account:

~~~
play -i com.google.android.youtube -a
~~~

download APK. you need to specify any valid version code. the latest code is
provided by the previous details command. if APK is split, all pieces will be
downloaded:

~~~
play -i com.google.android.youtube -v 1548869056
~~~

## goals

1. [Pixel 6](//wikipedia.org/wiki/Pixel_6) (2021)
2. [Android 12](//wikipedia.org/wiki/Android_12) (2021)
3. [Google Play](//wikipedia.org/wiki/Google_Play) 29 (2022)

up to Android 4.4 (2013), the login is protected with TLS fingerprinting, which
is difficult but possible to bypass. since Android 5 (2014), Google uses
bot-guard via JavaScript to protect the login. If you know about this contact
me

## money

software is not licensed for commercial use. if you wish to purchase a
commercial license, contact me

---------------------------------------------------------------------------------

as of november 5 2023, I no longer offer free discussion of this software. if
you are interested in paid support, contact me

<dl>
   <dt>email</dt>
      <dd>367@tuta.io</dd>
   <dt>Discord username</dt>
      <dd>10308</dd>
   <dt>Discord invite</dt>
      <dd>https://discord.com/invite/rMFzDRQhSx</dd>
   <dt>PayPal</dt>
      <dd>https://paypal.com/donate?hosted_button_id=UEJBQQTU3VYDY</dd>
</dl>
