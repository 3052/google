# Java

~~~
finsky\protos\AndroidAppDeliveryData.java
finsky\protos\AppDetails.java
finsky\protos\AppFileMetadata.java
finsky\protos\Common.java
finsky\protos\DeliveryResponse.java
finsky\protos\Details.java
finsky\protos\DeviceConfiguration.java
finsky\protos\DocDetails.java
finsky\protos\DocV2.java
finsky\protos\FileMetadata.java
finsky\protos\InstallDetails.java
finsky\protos\ListResponse.java
finsky\protos\PreFetch.java
finsky\protos\Response.java
finsky\protos\ServerCommands.java
finsky\protos\SplitDeliveryData.java
~~~

last working versions below. decompiled using JADX 1.4.7. check under
`sources\com\google\android`.

## /checkin

request:

~~~
gsf\checkin\proto\Checkin$AndroidCheckinRequest.java
   4 gsf\checkin\proto\Logs$AndroidCheckinProto.java
      1 gsf\checkin\proto\Logs$AndroidBuildProto.java
   18 gsf\checkin\proto\Config$DeviceConfigurationProto.java
~~~

response:

~~~
gsf\checkin\proto\Checkin$AndroidCheckinResponse.java
~~~

## /fdfe/uploadDeviceConfig

request:

~~~
finsky\protos\UploadDeviceConfig.java
   1 finsky\protos\DeviceConfiguration.java
~~~

## /fdfe/details

response:

~~~
finsky\protos\Response.java
   1 finsky\protos\Response.java
      2 finsky\protos\Details.java
         4 finsky\protos\DocV2.java
~~~

## Google Play Store

https://apkmirror.com/apk/google-inc/google-play-store

key             | value
----------------|--------------------
package         | com.android.vending
versionCode     | 80441400
versionName     | 6.1.14

## Google Services Framework

https://apkmirror.com/apk/google-inc/google-services-framework

key         | value
------------|-----------------------
package     | com.google.android.gsf
versionCode | 25
versionName | 7.1.2
