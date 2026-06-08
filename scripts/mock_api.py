#!/usr/bin/env python3
import json
from http.server import BaseHTTPRequestHandler, HTTPServer

last_post = None


class Handler(BaseHTTPRequestHandler):
    def _send_json(self, status, payload):
        body = json.dumps(payload).encode("utf-8")
        self.send_response(status)
        self.send_header("Access-Control-Allow-Origin", "*")
        self.send_header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        self.send_header("Access-Control-Allow-Headers", "Content-Type")
        self.send_header("Content-Type", "application/json")
        self.send_header("Content-Length", str(len(body)))
        self.end_headers()
        self.wfile.write(body)

    def do_OPTIONS(self):
        self._send_json(204, {})

    def do_GET(self):
        if self.path == "/HideBot/discord/targets":
            self._send_json(
                200,
                [
                    {"id": "main", "label": "Main server", "default": True},
                    {"id": "ops", "label": "Ops channel"},
                ],
            )
            return

        if self.path == "/HideBot/discord/version":
            self._send_json(200, "test")
            return

        if self.path == "/__last_post":
            self._send_json(200, last_post or {})
            return

        self._send_json(404, {"error": "not found"})

    def do_POST(self):
        global last_post
        if self.path != "/HideBot/discord":
            self._send_json(404, {"error": "not found"})
            return

        length = int(self.headers.get("Content-Length", "0"))
        body = self.rfile.read(length)
        last_post = json.loads(body.decode("utf-8"))
        self._send_json(200, {"ok": True})

    def log_message(self, format, *args):
        return


if __name__ == "__main__":
    HTTPServer(("127.0.0.1", 8082), Handler).serve_forever()
