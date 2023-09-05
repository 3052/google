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

visit this page and sign in:

https://accounts.google.com/embedded/setup/android

then get authorization code (`oauth_token`) cookie from the browser. should be
valid for 10 minutes:

<https://firefox-source-docs.mozilla.org/devtools-user/storage_inspector>

then exchange authorization code for refresh token (`aas_et`):

~~~
play -c oauth2_4/0Adeu5B...
~~~

Create a file containing `X-DFE-Device-ID` (GSF ID) for future requests:

~~~
play -device
~~~

Get app details:

~~~
> play -d com.google.android.youtube
creator: Google LLC
file: APK APK APK APK
installation size: 49.53 megabyte
downloads: 14.05 billion
offer: 0 USD
title: YouTube
upload date: May 19, 2023
version: 18.20.34
version code: 1537856960
~~~

Purchase app. Only needs to be done once per Google account:

~~~
play -d com.google.android.youtube -purchase
~~~

Download APK. You need to specify any valid version code. The latest code is
provided by the previous details command. If APK is split, all pieces will be
downloaded:

~~~
play -d com.google.android.youtube -v 1537856960
~~~

## Goals

1. support Google Services Framework 21
2. support Google Play Store 11

Non goals:

Supporting older items. sadly this means that email/password login is no longer
possible, and maybe never be again. up to Google Services Framework 19, the
login is protected with TLS fingerprinting, which is difficult but possible to
bypass. Since Google Services Framework 21, Google uses bot-guard via
JavaScript to protect the login. I do not know how to reverse that, and I did
not find any implementations online.

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

Software is not licensed for commercial use. If you wish to purchase a
commercial license, or for other business questions, contact me.
