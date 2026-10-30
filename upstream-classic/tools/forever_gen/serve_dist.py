#!/usr/bin/env python3
"""Serve the sim site for LAN daily use. Usage: serve_dist.py [port] (default 8732).
The site lives at /forever/ (vite base); / redirects there."""
import http.server, os, sys, socketserver

PORT = int(sys.argv[1]) if len(sys.argv) > 1 else 8732
# Root at dist/ (not dist/forever): vite builds with base '/forever/', so every
# page references /forever/… assets. Serving dist/forever as root 404s them.
ROOT = os.path.join(os.path.dirname(__file__), "..", "..", "dist")

class H(http.server.SimpleHTTPRequestHandler):
    extensions_map = {**http.server.SimpleHTTPRequestHandler.extensions_map,
                      ".wasm": "application/wasm", ".json": "application/json"}
    def __init__(self, *a, **kw):
        super().__init__(*a, directory=os.path.abspath(ROOT), **kw)
    def log_message(self, *a):
        pass
    def do_GET(self):
        if self.path in ("/", ""):
            self.send_response(302)
            self.send_header("Location", "/forever/")
            self.end_headers()
            return
        super().do_GET()

class TCPServer(socketserver.ThreadingTCPServer):
    allow_reuse_address = True

if __name__ == "__main__":
    with TCPServer(("0.0.0.0", PORT), H) as httpd:
        print(f"serving {os.path.abspath(ROOT)} on 0.0.0.0:{PORT}")
        httpd.serve_forever()
