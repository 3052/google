# Tower of Saviors

this is the only ARM option that includes Google Play:

API Level | ABI       | Target
----------|-----------|----------------------
28        | arm64-v8a | Android 9 Google Play

if you try to run it, you get this:

> The emulator process for AVD `Pixel_2_API_28` has terminated.

here is the next best option:

API Level | ABI       | Target
----------|-----------|----------------------
28        | arm64-v8a | Android 9 Google APIs

now lets start it:

~~~
> emulator -avd Pixel_3a_XL_API_28 -writable-system
PANIC: Avd's CPU Architecture 'arm64' is not supported by the QEMU2 emulator on
x86_64 host.
~~~

here is the next best option:

API Level | ABI         | Target
----------|-------------|--------------------------
25        | armeabi-v7a | Android 7.1.1 Google APIs

now lets start it:

~~~
emulator -avd Pixel_3a_XL_API_25 -writable-system
~~~

push Google Play Store:

~~~
adb root
adb remount
adb push com.android.vending_11.9.30.apk /system/priv-app
adb reboot
~~~

after 9 minutes, the reboot has not happened. here is the next best option:

API Level | ABI         | Target
----------|-------------|--------------------------
25        | `x86_64`    | Android 7.1.1 Google APIs

using this code:

~~~go
package main

import (
   "154.pages.dev/protobuf"
   "fmt"
)

func main() {
   {
      var m protobuf.Message
      m.Add_String(11, "x86_64")
      m.Add_String(11, "x86")
      b := m.Append(nil)
      fmt.Printf("%v %q\n", len(b), b)
   }
   {
      var m protobuf.Message
      m.Add_String(11, "armeabi-v7a")
      b := m.Append(nil)
      fmt.Printf("%v %q\n", len(b), b)
   }
}
~~~

we get this:

~~~py
from mitmproxy import http

def request(f: http.HTTPFlow) -> None:
   if f.request.path.startswith('/checkin'):
      f.request.content = f.request.content.replace(
         b"Z\x06x86_64Z\x03x86", b"Z\varmeabi-v7a"
      )
   if f.request.path.startswith('/fdfe/uploadDeviceConfig'):
      f.request.content = f.request.content.replace(
         b"Z\x06x86_64Z\x03x86", b"Z\varmeabi-v7a"
      )
~~~

which works with this:

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://play.google.com/store/apps/details?id=com.miui.weather2
~~~

but for this:

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://play.google.com/store/apps/details?id=com.madhead.tos.zh
~~~

we get:

> Your device isn't compatiblee with this version.

same for API 26-33. Note starting with API 29 (Android 10), you need to also
push permission:

~~~xml
<?xml version="1.0" encoding="utf-8"?>
<permissions>
   <privapp-permissions package="com.android.vending">
      <permission name="android.permission.ALLOCATE_AGGRESSIVE"/>
      <permission name="android.permission.BACKUP"/>
      <permission name="android.permission.BATTERY_STATS"/>
      <permission name="android.permission.CHANGE_COMPONENT_ENABLED_STATE"/>
      <permission name="android.permission.CHANGE_DEVICE_IDLE_TEMP_WHITELIST"/>
      <permission name="android.permission.CLEAR_APP_CACHE"/>
      <permission name="android.permission.CONNECTIVITY_INTERNAL"/>
      <permission name="android.permission.DELETE_PACKAGES"/>
      <permission name="android.permission.DUMP"/>
      <permission name="android.permission.FORCE_STOP_PACKAGES"/>
      <permission name="android.permission.GET_ACCOUNTS_PRIVILEGED"/>
      <permission name="android.permission.GET_APP_OPS_STATS"/>
      <permission name="android.permission.INSTALL_PACKAGES"/>
      <permission name="android.permission.INTERACT_ACROSS_USERS"/>
      <permission name="android.permission.MANAGE_USERS"/>
      <permission name="android.permission.PACKAGE_USAGE_STATS"/>
      <permission name="android.permission.PACKAGE_VERIFICATION_AGENT"/>
      <permission name="android.permission.READ_RUNTIME_PROFILES"/>
      <permission name="android.permission.REAL_GET_TASKS"/>
      <permission name="android.permission.SEND_SMS_NO_CONFIRMATION"/>
      <permission name="android.permission.STATUS_BAR"/>
      <permission name="android.permission.UPDATE_DEVICE_STATS"/>
      <permission name="android.permission.WRITE_SECURE_SETTINGS"/>
   </privapp-permissions>
</permissions>
~~~

like this:

~~~
adb root
adb shell avbctl disable-verification
adb reboot

adb root
adb remount
adb push com.android.vending_11.9.30.apk /system/priv-app
adb push privapp-permissions.xml /etc/permissions
adb reboot
~~~

https://source.android.com/docs/core/permissions/perms-allowlist
