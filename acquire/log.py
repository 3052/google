# docs.mitmproxy.org/stable/addons-examples#http-modify-query-string
from mitmproxy import http
import logging

def request(f: http.HTTPFlow) -> None:
   logging.info('hello world')
