# oauth\_token cookie

## Edge WebView2

On Windows, WebView2 requires its own runtime `WebView2Loader.dll`. this is an
issue because even recent versions such as Windows 10 2022 Update (10.0.19045) do
not include it, so it would need to be downloaded separately.

https://github.com/jchv/go-webview2

## WebView EdgeHTML

<https://wikipedia.org/wiki/Blink_(browser_engine)#Frameworks>

## WebBrowser MSHTML

<https://wikipedia.org/wiki/Blink_(browser_engine)#Frameworks>

## process memory

https://pkg.go.dev/golang.org/x/sys/windows#ReadProcessMemory

## local cookies file

~~~
C:\Users\<User>\AppData\Roaming\Mozilla\Firefox\Profiles\<Profile Folder>\
sessionstore-backups\recovery.jsonlz4
~~~

https://github.com/avih/dejsonlz4
