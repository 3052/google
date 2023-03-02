# November 22 2022

## Android 10.0 (Google APIs)

then start the Google APIs device:

~~~
emulator -avd Pixel_3a_XL_API_29 -writable-system
~~~

and push:

~~~
adb root
adb remount
~~~

Result:

~~~
W Disabling verity for /system
E Skipping /system
Skip mounting partition: /product
Skip mounting partition: /product_services
Skip mounting partition: /product
Skip mounting partition: /product_services
Skip mounting partition: /product
Skip mounting partition: /product_services
Skip mounting partition: /product
Skip mounting partition: /product_services
Skip mounting partition: /product
Skip mounting partition: /product_services
Skip mounting partition: /product
Skip mounting partition: /product_services
Skip mounting partition: /product
Skip mounting partition: /product_services
W DM_DEV_STATUS failed for scratch: No such device or address
E [liblp]No device named scratch
[liblp]Partition scratch will resize from 0 bytes to 536870912 bytes
[liblp]Updated logical partition table at slot 0 on device /dev/block/by-name/super
[libfs_mgr]Created logical partition scratch on device /dev/block/dm-3
[libfs_mgr]__mount(source=/dev/block/dm-3,target=/mnt/scratch,type=f2fs)=0: Success
Skip mounting partition: /product
Skip mounting partition: /product_services
Using overlayfs for /vendor
Skip mounting partition: /product
Skip mounting partition: /product_services
Skip mounting partition: /product
Skip mounting partition: /product_services
Skip mounting partition: /product
Skip mounting partition: /product_services
[libfs_mgr]__mount(source=overlay,target=/vendor,type=overlay,upperdir=/mnt/scratch/overlay/vendor/upper)=0
Skip mounting partition: /product
Skip mounting partition: /product_services
/system/bin/remount exited with status 7
~~~

- https://android.stackexchange.com/questions/232234/why-adb-remount
- https://android.stackexchange.com/questions/52443/the-command-adb-root-works
- https://github.com/topjohnwu/Magisk/issues/2252
- https://issuetracker.google.com/issues/260204230
- https://stackoverflow.com/questions/58010655/is-adb-remount-broken-on-android
- https://stackoverflow.com/questions/60052458/adb-amulator-running-sdk-29
- https://stackoverflow.com/questions/60867956/android-emulator-sdk-10-api-29
- https://stackoverflow.com/questions/65627342/why-does-adb-remount-fail

## Android 11.0 (Google APIs)

then start the Google APIs device:

~~~
emulator -avd Pixel_3a_XL_API_30 -writable-system
~~~

and push:

~~~
adb root
adb remount
adb reboot
~~~

After two minutes, the device has not rebooted. Same result with `x86` or
`x86_64`:

https://issuetracker.google.com/issues/260134707

## Android 12.0 (Google APIs)

Start the Google Play device and pull:

~~~
adb pull /product/priv-app/Phonesky/Phonesky.apk Phonesky-31.apk
~~~

then start the Google APIs device:

~~~
emulator -avd Pixel_3a_XL_API_31 -writable-system
~~~

and push:

~~~
adb root
adb remount
adb reboot

adb root
adb remount
adb push Phonesky-31.apk /product/priv-app
adb reboot
~~~

## Android 12L (Google APIs)

Start the Google Play device and pull:

~~~
adb pull /product/priv-app/Phonesky/Phonesky.apk Phonesky-32.apk
~~~

then start the Google APIs device:

~~~
emulator -avd Pixel_3a_XL_API_32 -writable-system
~~~

and push:

~~~
adb root
adb remount
adb reboot

adb root
adb remount
adb push Phonesky-32.apk /product/priv-app
adb reboot
~~~

## Android Tiramisu (Google APIs)

Start the Google Play device and pull:

~~~
adb pull /product/priv-app/Phonesky/Phonesky.apk Phonesky-33.apk
~~~

then start the Google APIs device:

~~~
emulator -avd Pixel_3a_XL_API_33 -writable-system
~~~

and push:

~~~
adb root
adb remount
adb reboot

adb root
adb remount
adb push Phonesky-33.apk /product/priv-app
adb reboot
~~~

## Android Tiramisu (Google Play)

we cannot run as root:

~~~
> adb root
adbd cannot run as root in production builds
~~~
