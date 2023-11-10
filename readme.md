# Google Play

> I’m not really sure what the point of this video is, but I guess just be
> generous. Be kind to people, because you never know how much they might need
> it, or how far it’ll go.
>
> [NakeyJakey (2018)](//youtube.com/watch?v=Cr0UYNKmrUs)

Download APK from Google Play or send API requests

## Support

BECAUSE OF REPEATED ABUSE, I NO LONGER OFFER FREE DISCUSSION OF THIS SOFTWARE. SO
UNLESS YOU HAVE PAID UP FRONT, DONT:

1. POST AN ISSUE
2. POST A PULL REQUEST
3. MESSAGE ME ON DISCORD
4. EMAIL ME

SOFTWARE IS NOT LICENSED FOR COMMERCIAL USE. IF YOU WISH TO PURCHASE A
COMMERCIAL LICENSE, CONTACT ME.

<https://paypal.com/donate?hosted_button_id=UEJBQQTU3VYDY>

## Tool examples

[Sign in](//accounts.google.com/embedded/setup/v2/android) with your Google
Account. Then get authorization code (`oauth_token`) cookie from
[browser&nbsp;storage][1]. Should be valid for 10 minutes. Then exchange
authorization code for refresh token (`aas_et`):

~~~
play -o oauth2_4/0Adeu5B...
~~~

[1]://firefox-source-docs.mozilla.org/devtools-user/storage_inspector

Create a file containing `X-DFE-Device-ID` (GSF ID) for future requests:

~~~
play -d
~~~

Get app details:

~~~
> play -a com.google.android.youtube
creator: Google LLC
file: APK APK APK APK
installation size: 89.03 megabyte
downloads: 14.81 billion
offer: 0 USD
requires: Android 8.0 and up
title: YouTube
upload date: Sep 22, 2023
version: 18.38.37
version code: 1540222400
~~~

Acquire app. Only needs to be done once per Google account:

~~~
play -a com.google.android.youtube -acquire
~~~

Download APK. You need to specify any valid version code. The latest code is
provided by the previous details command. If APK is split, all pieces will be
downloaded:

~~~
play -a com.google.android.youtube -v 1540222400
~~~

## Goals

1. [Android 11](//wikipedia.org/wiki/Android_11) (2020)
2. [Google Play](//wikipedia.org/wiki/Google_Play) 29 (2022)

Non goals:

email/password login. Up to Android 4.4 (2013), the login is protected with TLS
fingerprinting, which is difficult but possible to bypass. Since
Android 5 (2014), Google uses bot-guard via JavaScript to protect the login. I
do not know how to reverse that, and I did not find any implementations online.

## Contact

<dl>
   <dt>
   email
   </dt>
   <dd>
   srpen6@gmail.com
   </dd>
   <dt>
   Discord
   </dt>
   <dd>
   srpen6
   </dd>
   <dd>
   https://discord.com/invite/WWq6rFb8Rf
   </dd>
</dl>
