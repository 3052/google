# Google APIs 21

this:

https://stackoverflow.com/a/34581874

works fine for API 23 (Android 6), but fails with API 22 (Android 5.1).

GenyMotion works with API 22. what about API 21 (Android 5)?

---------------------------------------------------------------------------------

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
