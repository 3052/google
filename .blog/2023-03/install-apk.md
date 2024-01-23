# Install APK

## How to install Android App Bundle?

Bash:

~~~
adb install-multiple *.apk
~~~

PowerShell:

~~~
adb install-multiple (Get-ChildItem *.apk)
~~~

https://developer.android.com/guide/app-bundle/app-bundle-format

## How to install expansion file?

~~~
adb shell mkdir -p /sdcard/Android/obb/com.PirateBayGames.ZombieDefense2

adb push main.41.com.PirateBayGames.ZombieDefense2.obb `
/sdcard/Android/obb/com.PirateBayGames.ZombieDefense2/
~~~

https://developer.android.com/google/play/expansion-files

## INSTALL\_FAILED\_NO\_MATCHING\_ABIS

This can happen when trying to install ARM app on `x86`. If the APK is
`armeabi-v7a`, then Android 9 (API 28) will work. Also the emulator should be
`x86`. If the APK is `arm64-v8a`, then Android 11 (API 30) will work. Also the
emulator should be `x86_64`.

- https://android.stackexchange.com/questions/222094/install-failed
- https://stackoverflow.com/questions/36414219/install-failed-no-matching-abis

However note that this will still fail in some cases:

https://issuetracker.google.com/issues/253778985
