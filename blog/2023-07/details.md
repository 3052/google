# details

- <https://bugzilla.mozilla.org/show_bug.cgi?id=1780817>
- https://apkmirror.com/apk/google-inc/trichrome-library

## fail

~~~
play -d com.google.android.gsf
play -d com.google.android.gsf -p 1
play -d com.google.android.gsf -p 2
play -d com.google.android.gsf -v 23
play -d com.google.android.trichromelibrary
play -d com.google.android.trichromelibrary -p 1
play -d com.google.android.trichromelibrary -p 2
~~~

note that `/fdfe/details` and `/fdfe/bulkDetails` both fail.

## pass

~~~
play -d com.google.android.trichromelibrary -v 579016636
~~~
