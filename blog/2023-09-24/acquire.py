from mitmproxy import ctx, http

def running():
   ctx.options.showhost = True

def request(f: http.HTTPFlow) -> None:
   if f.request.path.startswith('/checkin'):
      f.request.content = f.request.content.replace(
         b"Z\x06x86_64Z\x03x86", b"Z\varmeabi-v7a"
      )
   if f.request.path.startswith('/fdfe/uploadDeviceConfig'):
      f.request.content = f.request.content.replace(
         b"Z\x06x86_64Z\x03x86", b"Z\varmeabi-v7a"
      )
   if f.request.path.startswith('/fdfe/acquire'):
      f.kill()
   if f.request.path.startswith('/fdfe/delivery'):
      f.kill()
   if f.request.path.startswith('/fdfe/replicateLibrary'):
      f.kill()
   if f.request.path.startswith('/play-apps-download-default/download/by-id/'):
      f.kill()
