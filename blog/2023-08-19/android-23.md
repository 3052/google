# Google APIs 21

First go here:

https://opengapps.org

I selected this option:

Platform | Android | Variant
---------|---------|--------
x86      | 6       | pico

but newer Android should work as well. Then open Android Studio. On "Select
Hardware" screen, select a device without "Play Store" icon.

## Without Google APIs

Using the method above with Google APIs image, you still get some apps such as
YouTube. If you want to install different version of one of these apps, use
this method. On "System Image" screen, I selected this option:

API Level | ABI | Target
----------|-----|----------
24        | x86 | Android 7

but newer APIs should work as well. You need these files from the Zip archive:

~~~
Core\gmscore-x86.tar.lz
Core\vending-x86.tar.lz
~~~

Then extract these from the above files:

~~~
gmscore-x86\nodpi\priv-app\PrebuiltGmsCore\PrebuiltGmsCore.apk
vending-x86\nodpi\priv-app\Phonesky\Phonesky.apk
~~~

Use the same method above to install the APKs. After reboot, you should then be
able to install YouTube or whatever app. Note that unlike above, you dont
actually need to run the Google Play setup or even start the Google Play app at
the end.

## With Google APIs

On "System Image" screen, I selected this option:

API Level | ABI | Target
----------|-----|----------------------
24        | x86 | Android 7 Google APIs

but newer APIs should work as well. You should only need one file from the Zip
archive:

~~~
Core\vending-x86.tar.lz
~~~

Inside this will be another file:

~~~
vending-x86\nodpi\priv-app\Phonesky\Phonesky.apk
~~~

Now, start the device:

~~~
emulator -list-avds
emulator -avd Pixel_3a_XL_API_24 -writable-system
~~~

Next, install Google Play Store. Note that you cannot use the normal method of
drag APK to device screen, or you will get one of these errors:

~~~
The APK failed to install.<br/> Error: Could not parse error string

The APK failed to install.<br/> Error: INSTALL_FAILED_UPDATE_INCOMPATIBLE:

Package com.android.vending signatures do not match the previously installed
version; ignoring!

The APK failed to install.<br/> Error: INSTALL_PARSE_FAILED_NO_CERTIFICATES:
Failed to collect certificates from /data/app/vmdl1047870024.tmp/base.apk:
META-INF/BNDLTOOL.SF indicates /data/app/vmdl1047870024.tmp/base.apk is signed
using APK Signature Scheme v2, but no such signature was found. Signature
stripped?
~~~

Install like this:

~~~
adb root
adb remount
adb push Phonesky.apk /system/priv-app
adb reboot
~~~

After reboot, you should then be able to start Google Play Store as normal.

this:

https://stackoverflow.com/a/34581874

works fine for API 23 (Android 6), but fails with API 22 (Android 5.1).

---------------------------------------------------------------------------------

GenyMotion works with API 22. what about API 21 (Android 5)?

1. Android Studio
2. Pixel 3a XL
3. API Level 21
4. Android 5 image

https://opengapps.org

~~~
emulator -list-avds
emulator -avd Pixel_3a_XL_API_18 -writable-system
~~~

then:

~~~
adb pull /system/app/PlayGames/PlayGames.apk
package: name='com.google.android.play.games' versionCode='95330070'

adb pull /system/priv-app/GoogleLoginService/GoogleLoginService.apk
package: name='com.google.android.gsf.login' versionCode='21'

adb pull /system/priv-app/GoogleServicesFramework/GoogleServicesFramework.apk
package: name='com.google.android.gsf' versionCode='21'

adb pull /system/priv-app/PrebuiltGmsCore/PrebuiltGmsCore.apk
package: name='com.google.android.gms' versionCode='214857013'

adb remount
adb push GoogleLoginService.apk /system/priv-app
adb push GoogleServicesFramework.apk /system/priv-app
adb push PlayGames.apk /system/app
adb push PrebuiltGmsCore.apk /system/priv-app
adb reboot
~~~

then:

~~~
play -d com.android.vending -v 80671300
~~~

then:

~~~
play -d com.google.android.gms -v 9452270
~~~

then:

~~~
mitmproxy
~~~

Press `O` to enter options. Move to `block_list` and press Enter. Then press
`a` to add a new entry. Press Esc when finished, then `q`.

~~~
/~u _.lookup.accountlookup/444
~~~

https://docs.mitmproxy.org/stable/overview-features#blocklist

install system certificate from another terminal, since it is faster to do so.
Then start Play Games.

- https://accounts.google.com/EmbeddedSetup
- https://accounts.google.com/embedded/setup/android
