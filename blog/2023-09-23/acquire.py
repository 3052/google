from mitmproxy import ctx, http

def running():
   ctx.options.showhost = True

def request(f: http.HTTPFlow) -> None:
   if f.request.path.startswith('/fdfe/acquire'):
      f.kill()
   if f.request.path.startswith('/fdfe/delivery'):
      f.kill()
   if f.request.path.startswith('/fdfe/replicateLibrary'):
      f.kill()
   if f.request.path.startswith('/play-apps-download-default/download/by-id/'):
      f.kill()
