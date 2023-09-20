# acquire

~~~
Authorization aurora
X-Dfe-Device-Id mine
pass

Authorization aurora
X-Dfe-Device-Id studio
fail

Authorization mine
X-Dfe-Device-Id studio
fail

Authorization mine
X-Dfe-Device-Id aurora
fail
~~~
