# Google Play

> I’m not really sure what the point of this video is, but I guess just be
> generous. Be kind to people, because you never know how much they might need
> it, or how far it’ll go.
>
> [NakeyJakey (2018)](//youtube.com/watch?v=Cr0UYNKmrUs)

Download APK from Google Play or send API requests

## How to build?

This module works with Windows, macOS or Linux. Download Go [1] and extract
archive. Then download Google Play Zip and extract archive. Then navigate to:

~~~
google-main/cmd/play
~~~

and enter:

~~~
go build
~~~

1. https://go.dev/dl

## Tool examples

Visit this page [1] and sign in. Then get authorization code (`oauth_token`)
cookie from the browser [2]. Should be valid for 10 minutes. Then exchange
authorization code for refresh token (`aas_et`):

~~~
play -c oauth2_4/0Adeu5B...
~~~

1. https://accounts.google.com/embedded/setup/v2/android
2. <https://firefox-source-docs.mozilla.org/devtools-user/storage_inspector>

Create a file containing `X-DFE-Device-ID` (GSF ID) for future requests:

~~~
play -device
~~~

Get app details:

~~~
> play -d com.google.android.youtube
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
play -d com.google.android.youtube -a
~~~

Download APK. You need to specify any valid version code. The latest code is
provided by the previous details command. If APK is split, all pieces will be
downloaded:

~~~
play -d com.google.android.youtube -v 1540222400
~~~

## Goals

1. System Image API 26 (revision 4)
2. Google Play Store 11

Non goals:

Supporting older items. sadly this means that email/password login is not
currently possible. Up to Google Services Framework 4.4 (2013), the login is
protected with TLS fingerprinting, which is difficult but possible to bypass.
Since Google Services Framework 5 (2014), Google uses bot-guard via JavaScript
to protect the login. I do not know how to reverse that, and I did not find any
implementations online. A workaround would be to use WebView, which I am
considering, but making it work cross platform is difficult.

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

## Money

<https://paypal.com/donate?hosted_button_id=UEJBQQTU3VYDY>

Software is not licensed for commercial use. If you wish to purchase a
commercial license, or for other business questions, contact me.
