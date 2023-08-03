# August 2023

## one

~~~
POST /auth HTTP/1.1
device: 3ef5ba894...
app: com.google.android.gms
Accept-Encoding: identity
User-Agent: GoogleAuth/1.4 (vbox86p MRA58K); gzip
content-type: application/x-www-form-urlencoded
Host: android.googleapis.com
Connection: Keep-Alive

google_play_services_version=220920022&
ACCESS_TOKEN=1&
Email=s...&
Token=oauth2_4%2F0AZEOvhW7XlfPRjMFUj_SuTRTynb7Ia_ogX64ZWo8nwB3Bg7IsIPNtM0Lq6...&
add_account=1&
androidId=3ef5ba894...&
build_brand=Motorola&
build_device=vbox86p&
build_fingerprint=google%2Fvbox86p%2Fvbox86p%3A6.0%2FMRA58K%2F435%3Auserdebu...&
build_product=vbox86p&
callerPkg=com.google.android.gms&
callerSig=38918a453...&
device_country=us&
droidguard_results=CgbeSIfShkoSvysKBs9ELTkyDdIQWwAAd1T2HlrXFZIAWkUlE1NLkz0AL...&
get_accountid=1&
lang=en-US&
sdk_version=23&
service=ac2dm

HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Cache-Control: no-cache, no-store, max-age=0, must-revalidate
Pragma: no-cache
Expires: Mon, 01 Jan 1990 00:00:00 GMT
Date: Thu, 03 Aug 2023 01:02:46 GMT
Cross-Origin-Opener-Policy: same-origin
Cross-Origin-Resource-Policy: same-site
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
Content-Security-Policy: frame-ancestors 'self'
X-XSS-Protection: 1; mode=block
Server: GSE
Alt-Svc: h3=":443"; ma=2592000,h3-29=":443"; ma=2592000
Accept-Ranges: none
Vary: Sec-Fetch-Dest, Sec-Fetch-Mode, Sec-Fetch-Site,Accept-Encoding
Transfer-Encoding: chunked

Auth=ZQhNhopY--kuQg4haqigczIr6Izku-OYurXFpV-bDOxywNyWu9HWQsQGr1CPgbNm...
Email=...@gmail.com
GooglePlusUpdate=0
LSID=BAD_COOKIE
SID=BAD_COOKIE
Token=aas_et/AKppINbvJbKvfE0XyZYN9G8ik9zE_EknKidaHGxr8d_MiaBODlCvXuekO234zhLT...
accountId=107718248...
capabilities.canHavePassword=1
capabilities.canHaveUsername=1
capabilities=CikKDGdpeWRvbGxkbWZ5YRABGAEqFQoTY29tLmFuZHJvaWQudmVuZGluZwoSCgxn...
firstName=S...
lastName=P...
services=mail,talk,cl,friendview,android,youtube,googleplay,nova,sierra,multi...
~~~

## two

~~~
POST /auth HTTP/1.1
device: 3ef5ba894...
app: com.android.vending
Accept-Encoding: identity
User-Agent: GoogleAuth/1.4 (vbox86p MRA58K); gzip
content-type: application/x-www-form-urlencoded
Host: android.googleapis.com
Connection: Keep-Alive

Email=s...%40gmail.com&
Token=aas_et%2FAKppINbvJbKvfE0XyZYN9G8ik9zE_EknKidaHGxr8d_MiaBODlCvXuekO234z...&
_opt_is_called_from_account_manager=1&
androidId=3ef5ba894...&
app=com.android.vending&
callerPkg=com.google.android.gms&
callerSig=38918a453d07199354f8b19af05ec6562ced5788&
check_email=1&
client_sig=38918a453d07199354f8b19af05ec6562ced5788&
device_country=us&
droidguardPeriodicUpdate=1&
droidguard_results=CgYkfD6BSaESkysKBs9ELTkyDdIQWwAAd1T2HlrXFZIAWkUlE1NLkz0AL...&
google_play_services_version=220920022&
is_called_from_account_manager=1&
it_caveat_types=2&
lang=en-US&
oauth2_foreground=1&
sdk_version=23&
service=oauth2%3Ahttps%3A%2F%2Fwww.googleapis.com%2Fauth%2Fgoogleplay&
system_partition=1&
token_request_options=CAA4AVAB

HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Cache-Control: no-cache, no-store, max-age=0, must-revalidate
Pragma: no-cache
Expires: Mon, 01 Jan 1990 00:00:00 GMT
Date: Thu, 03 Aug 2023 01:02:51 GMT
Cross-Origin-Resource-Policy: same-site
Cross-Origin-Opener-Policy: same-origin
X-Content-Type-Options: nosniff
X-Frame-Options: SAMEORIGIN
Content-Security-Policy: frame-ancestors 'self'
X-XSS-Protection: 1; mode=block
Server: GSE
Alt-Svc: h3=":443"; ma=2592000,h3-29=":443"; ma=2592000
Accept-Ranges: none
Vary: Sec-Fetch-Dest, Sec-Fetch-Mode, Sec-Fetch-Site,Accept-Encoding
Transfer-Encoding: chunked

ExpiresInDurationSec=65658
Expiry=1691090229
grantedScopes=https://www.googleapis.com/auth/googleplay
isTokenSnowballed=0
issueAdvice=auto
it=ya29.m.Cv8BAZlPsVh55YGV8MSALgAvWFEbh2_azLr1VNNmL68-xZkM-dEg6uolYU2mt4LD9s9...
itMetadata=GgMKAQIgzBw
services=nova,cloudconsole,ahsid,mobile,sierra,billing,hist,androiddeveloper,...
storeConsentRemotely=0
~~~
