# GooglePlay

> I’m not really sure what the point of this video is, but I guess just be
> generous. Be kind to people, because you never know how much they might need
> it, or how far it’ll go.
>
> [NakeyJakey (2018)](//youtube.com/watch?v=Cr0UYNKmrUs)

Download APK from Google Play or send API requests

## How to install?

This module works with Windows, macOS or Linux. Builds here [1]. If you want to
manually build, download Go [2] and extract archive. Then download GooglePlay
Zip and extract archive. Then navigate to:

~~~
googleplay-main/cmd/googleplay
~~~

and enter:

~~~
go build
~~~

1. https://beta.gobuilds.org/2a.pages.dev/googleplay
2. https://go.dev/dl

## Tool examples

Before trying to Sign in, make sure your location is correct, to avoid
geo-blocking. You can test by logging into your Google account with a web
browser. Also, make sure the Google account you are using has logged into the
Play Store at least once before, using a physical or virtual Android device.
Create a file containing token (`aas_et`) for future requests:

~~~
googleplay -email EMAIL -passwd PASSWORD
~~~

Create a file containing `X-DFE-Device-ID` (GSF ID) for future requests:

~~~
googleplay -device
~~~

Get app details:

~~~
> googleplay -d com.google.android.youtube
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
googleplay -d com.google.android.youtube -purchase
~~~

Download APK. You need to specify any valid version code. The latest code is
provided by the previous details command. If APK is split, all pieces will be
downloaded:

~~~
googleplay -d com.google.android.youtube -v 1537856960
~~~

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
   srpen6#6983
   </dd>
   <dd>
   https://discord.com/invite/WWq6rFb8Rf
   </dd>
</dl>

## Money

Software is not licensed for commercial use. If you wish to purchase a
commercial license, or for other business questions, contact me.
