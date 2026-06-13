# Sequence va ACK Numbers

TCP tartibni va ishonchni saqlash uchun raqamlar tizimini ishlatadi.

---

## Sequence Number

Har paketga raqam beriladi. Shu orqali TCP tartibni saqlaydi.

### Misol

```text
Packet 1: SEQ = 0
Packet 2: SEQ = 1000
Packet 3: SEQ = 2000
```

### Qanday ishlaydi

1. Sender birinchi paketga sequence number beradi
2. Har keyingi paket oldingi paket o'lchamiga qarab yangi raqam oladi
3. Receiver shu raqamlar orqali paketlarni tartiblaydi

```text
Sender                              Receiver

SEQ=0 [data: 1000 byte] ---------->
     <---------- ACK=1000
SEQ=1000 [data: 500 byte] --------->
     <---------- ACK=1500
SEQ=1500 [data: 200 byte] --------->
     <---------- ACK=1700
```

---

## ACK Number

Qaysi packetgacha qabul qilinganini bildiradi.

### Misol

```text
ACK = 5
```

Demak: 1..4 paketlar qabul qilingan. Keyingi kutilayotgan paket 5.

### Qanday ishlaydi

```text
Sender                              Receiver

SEQ=0  [data] ------------------->
                                ACK=1 (1 gacha qabul qilindi)
SEQ=1  [data] ------------------->
                                ACK=2
SEQ=2  [data] --X (yo'qoldi)
SEQ=3  [data] ------------------->
                                ACK=3 (2 ni hali olmadi, 3 kelyapti)
                                ACK=3 (qayta yuborish so'rovchi)
SEQ=2  [data] ------------------->  (retransmission)
                                ACK=4
```

---

## Duplicate ACK

Agar receiver bir xil sequence number ni bir necha marta oladi, u duplicate ACK yuboradi.

```text
Sender                              Receiver

SEQ=0  [data] ------------------->
                                ACK=1
SEQ=1  [data] --X (yo'qoldi)
SEQ=2  [data] ------------------->
                                ACK=2 (duplicate)
SEQ=3  [data] ------------------->
                                ACK=2 (duplicate)
SEQ=4  [data] ------------------->
                                ACK=2 (3 ta duplicate ACK)
                                --> Retransmission!
SEQ=1  [data] ------------------->  (retransmission)
                                ACK=5
```

---

## Diagram

```text
Sender                              Receiver

SEQ=0  -------------------------->
     <---------------------------  ACK=1000
SEQ=1000  ---------------------->
     <---------------------------  ACK=2000
SEQ=2000  --X (yo'qoldi)
SEQ=3000  ---------------------->
     <---------------------------  ACK=2000 (duplicate)
     <---------------------------  ACK=2000 (duplicate)
     <---------------------------  ACK=2000 (3 duplicate)
SEQ=2000  --------------------->  (retransmission)
     <---------------------------  ACK=4000
```

---

## Amaliyot

### Wireshark bilan sequence/ack ko'rish

```bash
sudo tcpdump -i any -nn port 8080 -w capture.pcap
```

Keyin Wireshark'da ochib TCP stream ni ko'rish mumkin.

### Go orqali

```go
package main

import (
    "fmt"
    "net"
)

func main() {
    conn, err := net.Dial("tcp", "google.com:80")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    fmt.Println("Local addr:", conn.LocalAddr())
    fmt.Println("Remote addr:", conn.RemoteAddr())
}
```

---

## Xulosa

- Sequence Number har paketga unikal raqam beradi
- ACK Number qaysi paketgacha qabul qilinganini bildiradi
- Duplicate ACK orqali paket yo'qolishi aniqlanadi
- 3 ta duplicate ACK -- retransmission trigger
