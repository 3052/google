from mitmproxy import ctx, http

def running():
   ctx.options.showhost = True

def request(f: http.HTTPFlow) -> None:
   if f.request.path.startswith('/v1/getPoIntegrityToken'):
      f.kill()
