# /fdfe/acquire

OK another update. looks like if `/fdfe/replicateLibrary` runs, then it
prevents `/fdfe/acquire` from running, and will just go straight to
`/fdfe/delivery` if its a app that has already been acquired. this is bad
because for testing, we want `/fdfe/acquire` to run every time.

if you block `/fdfe/replicateLibrary` then acquire seems to work as normal.
