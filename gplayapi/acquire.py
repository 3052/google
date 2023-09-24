from mitmproxy import ctx, http
import json

device_ID = json.load(open('client.json'))['AuthData']['GsfID']

def running():
   ctx.options.showhost = True

def request(f: http.HTTPFlow) -> None:
   if f.request.path.startswith('/play-apps-download-default/download/by-id/'):
      f.kill()
   if f.request.path.startswith('/fdfe/replicateLibrary'):
      f.kill()
   if f.request.path.startswith('/fdfe/delivery'):
      f.kill()
   if f.request.path.startswith('/fdfe/purchase'):
      f.kill()
   if f.request.path.startswith('/fdfe/acquire'):
      #f.kill()
      f.request.headers['X-DFE-Device-Id'] = device_ID
      
      # fail
      #f.request.headers.pop('X-DFE-Cookie', None)
      #f.request.headers.pop('X-DFE-Device-Config-Token', None)

'''
X-DFE-Cookie: EAEYACICVVMyUENqZ2FOZ29UTkRVek1EWXhOREEwTlRVek5qQXhNelk0TUJJZkNoQXhOamsxTlRFeU16a3pNREUxTVRZNUVnc0l5ZTY5cUFZUTZPdWRCdz09QhQKBVVTLVRYEgsIycLRqQYQ4MGcW0oRCgJVUxILCMnC0akGEMi9oFtYAA
X-DFE-Device-Config-Token: CjgaNgoTNDUzMDYxNDA0NTUzNjAxMzY4MBIfChAxNjk1NTEyMzkzMDE1MTY5EgsIye69qAYQ6OudBw==
'''
