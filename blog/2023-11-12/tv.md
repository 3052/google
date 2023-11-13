# TV

the current Hulu code seems to be based off the TV version of the app. However
I am not sure how that was accomplished. GenyMotion does not have an option to
create a TV device. Android Studio does, but the problem is that the Hulu TV
app only works with `armeabi-v7a`. the app does have a universal version, but
only with older versions of 2017 and before. with phones, you can emulate ARM
using Android 9 or higher, but that didn't seem to work for me with a TV app.

hm I guess hardware is the only option. I just tried the current app on an ARM
emulator. it was slow as hell, to be expected. four minutes just to install.
but even then, if you try to start the app, you get:

> Unfortunately, Hulu has stopped.

even with no MITM

- https://github.com/httptoolkit/frida-interception-and-unpinning/issues/49
- https://github.com/matthuisman/slyguy.addons/issues/615
