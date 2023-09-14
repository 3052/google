# Java

last working versions below. decompiled using JADX 1.4.7. check under
`sources\com\google\android`.

key         | value
------------|-----------------------
package     | com.google.android.gsf
versionCode | 25
versionName | 7.1.2

https://apkmirror.com/apk/google-inc/google-services-framework

key             | value
----------------|--------------------
package         | com.android.vending
versionCode     | 80441400
versionName     | 6.1.14

https://apkmirror.com/apk/google-inc/google-play-store

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

## /fdfe/delivery

response:

~~~
finsky\protos\Response.java
   1 finsky\protos\Response.java
      21 finsky\protos\DeliveryResponse.java
         2 finsky\protos\AndroidAppDeliveryData.java
            4 finsky\protos\AppFileMetadata.java
            15 finsky\protos\SplitDeliveryData.java
~~~

## /fdfe/details

response:

~~~
finsky\protos\Response.java
   1 finsky\protos\Response.java
      2 finsky\protos\Details.java
         4 finsky\protos\DocV2.java
            8 finsky\protos\Common.java
            13 finsky\protos\DocDetails.java
               1 finsky\protos\AppDetails.java
                  17 finsky\protos\FileMetadata.java
~~~

## /fdfe/uploadDeviceConfig

request:

~~~
finsky\protos\UploadDeviceConfig.java
   1 finsky\protos\DeviceConfiguration.java
~~~
