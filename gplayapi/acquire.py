from mitmproxy import ctx, http

device_ID = '3f49ddf308ac77a6'

def running():
   ctx.options.showhost = True

def request(f: http.HTTPFlow) -> None:
   if f.request.path.startswith('/fdfe/accountSync'):
      f.request.headers['X-DFE-Device-Id'] = device_ID
   if f.request.path.startswith('/fdfe/acquire'):
      f.request.headers['X-DFE-Device-Id'] = device_ID
   if f.request.path.startswith('/fdfe/myApps'):
      f.request.headers['X-DFE-Device-Id'] = device_ID
   if f.request.path.startswith('/fdfe/selfUpdate'):
      f.request.headers['X-DFE-Device-Id'] = device_ID
   if f.request.path.startswith('/fdfe/toc'):
      f.request.headers['X-DFE-Device-Id'] = device_ID
   if f.request.path.startswith('/fdfe/uploadDeviceConfig'):
      f.request.headers['X-DFE-Device-Id'] = device_ID
   if f.request.path.startswith('/fdfe/uploadDynamicConfig'):
      f.request.headers['X-DFE-Device-Id'] = device_ID
   if f.request.path.startswith('/fdfe/userSettings'):
      f.request.headers['X-DFE-Device-Id'] = device_ID
   if f.request.path.startswith('/play-apps-download-default/download/by-id/'):
      f.kill()
