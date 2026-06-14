#!/usr/bin/env python3
"""
SSL Stripping Simulyatsiya - FAQAT TA'LM UCHUN!
Faqat o'z tarmog'ingizda sinang!
Ruxsatsiz ishlatish JINOYAT!
"""

from http.server import HTTPServer, BaseHTTPRequestHandler

class SSLStripProxy(BaseHTTPRequestHandler):
    """
    SSL Strip Proxy - HTTPS ni HTTP ga tushirish
    
    ⚠️  FAQAT O'Z TARMOG'INGIZDA SINANG!
    """
    
    def do_GET(self):
        """GET so'rovni qayta ishlash"""
        # HTTP ga redirect
        self.send_response(301)
        self.send_header("Location", f"http://{self.headers['Host']}{self.path}")
        self.send_header("Content-Type", "text/html")
        self.end_headers()
        
        html = """
        <html>
        <head><title>SSL Strip Test</title></head>
        <body>
            <h1>SSL Strip Proxy</h1>
            <p>HTTPS → HTTP simulyatsiya</p>
            <p>⚠️ Faqat o'rganish uchun!</p>
        </body>
        </html>
        """
        self.wfile.write(html.encode())
    
    def log_message(self, format, *args):
        """Log chiqarish"""
        print(f"[SSL STRIP] {args[0]}")

def run_proxy(port=8080):
    """SSL Strip proxy ishga tushirish"""
    server = HTTPServer(("0.0.0.0", port), SSLStripProxy)
    print(f"[*] SSL Strip Proxy: http://0.0.0.0:{port}")
    print("⚠️  CTRL+C to'xtatish uchun")
    
    try:
        server.serve_forever()
    except KeyboardInterrupt:
        server.shutdown()
        print("\n[*] Proxy to'xtatildi")

if __name__ == "__main__":
    PORT = 8080  # Proxy port
    run_proxy(PORT)
