# TCP Packet Structure

TCP paketining to'liq tuzilishi.

---

## TCP Header (20 bytes minimum)

```text
 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|          Source Port          |       Destination Port        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Sequence Number                        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                    Acknowledgement Number                     |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|  Data |       |U|A|P|R|S|F|                                   |
| Offset| Rsvd  |R|C|S|S|Y|I|            Window                 |
|       |       |G|K|H|T|N|N|                                   |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|           Checksum            |         Urgent Pointer        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                    Options (variable)                         |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             Data                              |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
```

---

## Har bir maydon

### Source Port (16 bit)

Manba port raqami. 0-65535 orasida.

### Destination Port (16 bit)

Maqsad port raqami. Masalan: 80 (HTTP), 443 (HTTPS), 22 (SSH).

### Sequence Number (32 bit)

Paketning birinchi baytining sequence number.

### Acknowledgement Number (32 bit)

Keyingi kutilayotgan sequence number.

### Data Offset (4 bit)

TCP header necha 4-byte words dan iborat. Minimum 5 (20 bytes).

### Reserved (3 bit)

Kelajakda ishlatiladigan joy. Hozircha 0.

### Flags (9 bit)

```text
CWR - Congestion Window Reduced
ECE - ECN-Echo
URG - Urgent
ACK - Acknowledgement
PSH - Push
RST - Reset
SYN - Synchronize
FIN - Finish
```

### Window Size (16 bit)

Receive window o'lchami. Qabul qiluvchi qancha ma'lumotni qabul qila olishini bildiradi.

### Checksum (16 bit)

TCP segment yaxlitligini tekshirish uchun.

### Urgent Pointer (16 bit)

Urgent ma'lumotning joylashuvini ko'rsatadi.

---

## TCP Packet Example

```text
Ethernet Frame (1500 bytes MTU)
+--------------------------------------------------+
| IP Header (20 bytes)                              |
+--------------------------------------------------+
| TCP Header (20 bytes)                             |
|   Source Port: 49152                              |
|   Destination Port: 80                            |
|   Sequence Number: 1234567                        |
|   Acknowledgement Number: 9876543                  |
|   Flags: ACK, PSH                                 |
|   Window: 65535                                   |
|   Checksum: 0x1234                                |
+--------------------------------------------------+
| Payload (1460 bytes max)                          |
|   GET / HTTP/1.1                                  |
|   Host: example.com                               |
|   ...                                             |
+--------------------------------------------------+
```

---

## Options

TCP headeriga qo'shimcha ma'lumot qo'shish mumkin:

```text
MSS Option:         Maximum Segment Size
Window Scale:       Window size ni kattalashtirish
SACK Permitted:     Selective ACK ruxsat
Timestamp:          RTT hisoblash uchun
```

---

## Amaliyot

### Tcpdump bilan header ko'rish

```bash
sudo tcpdump -i any port 8080 -nn -X
```

### Hex dump

```bash
sudo tcpdump -i any port 8080 -nn -xx
```

### Wireshark

```bash
sudo tcpdump -i any port 8080 -w capture.pcap
wireshark capture.pcap
```

---

## Xulosa

- TCP header minimum 20 bytes
- Source/Destination port
- Sequence va Acknowledgement number
- Flags: SYN, ACK, FIN, RST, PSH, URG
- Window size, Checksum, Urgent Pointer
- Options: MSS, Window Scale, SACK, Timestamp
