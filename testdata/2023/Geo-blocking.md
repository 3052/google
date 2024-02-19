# Geo-blocking

Some apps are specific to region. For example, `air.ITVMobilePlayer` is specifc
to GB. If you try it from US, details will work, but delivery will fail:

~~~
> googleplay -a air.ITVMobilePlayer
Title: ITV Hub: Your TV Player - Watch Live & On Demand
UploadDate: Dec 9, 2021
VersionString: 9.19.0
VersionCode: 901900000
NumDownloads: 17.429 M
Size: 35.625 MB
Offer: 0.00 USD

> googleplay -a air.ITVMobilePlayer -v 901900000
panic: Geo-blocking
~~~

It seems headers are ignored as well:

~~~
Accept-Language: es
Accept-Language: es-AR
Accept-Language: es-ES
~~~

You can change the country [1], and then you get expected result:

~~~
> googleplay -a air.ITVMobilePlayer
Title: ITV Hub: Your TV Player - Watch Live & On Demand
UploadDate: Dec 9, 2021
VersionString: 9.19.0
VersionCode: 901900000
NumDownloads: 17.429 M
Size: 35.625 MB
Offer: 0.00 GBP

> googleplay -a air.ITVMobilePlayer -v 901900000
GET https://play.googleapis.com/download/by-token/download?token=AOTCm0TiBZQdp...
~~~

1. https://support.google.com/googleplay/answer/7431675
