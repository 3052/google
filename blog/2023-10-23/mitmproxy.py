from mitmproxy import ctx, http

def running():
   ctx.options.showhost = True

def request(f: http.HTTPFlow) -> None:
   if f.request.path.startswith('/play-apps-download-default/download/by-id/'):
      f.kill()
   if f.request.path.startswith('/fdfe/replicateLibrary'):
      f.kill()
   if f.request.path.startswith('/fdfe/acquire'):
      f.request.headers.pop('X-DFE-Cookie', None)
      f.request.headers.pop('X-DFE-Encoded-Targets', None)
      f.request.headers.pop('X-DFE-Phenotype', None)
      f.request.headers.pop('X-PS-RH', None)
