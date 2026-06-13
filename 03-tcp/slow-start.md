# TCP Slow Start

TCP boshida sekin boshlaydi va asta-sekin tezlashtiradi.

---

## Nazariya

TCP boshida 1 packet yuboradi. Har bir ACK olgandan keyin ikki barobar ko'paytiradi.

### Qanday ishlaydi

```text
Round 1: 1 packet   --> 1 ACK oladi
Round 2: 2 packets  --> 2 ACK oladi
Round 3: 4 packets  --> 4 ACK oladi
Round 4: 8 packets  --> 8 ACK oladi
Round 5: 16 packets --> 16 ACK oladi
Round 6: 32 packets --> ...
```

Bu **exponential growth** -- 1, 2, 4, 8, 16, 32...

---

## Diagram

```text
cwnd
 ^
 |                              *
 |                         *
 |                    *
 |               *
 |          *
 |     *
 | *
 +---------------------------------> vaqt

1   2   4   8   16   32   64
```

---

## Slow Start Threshold (ssthresh)

Bir nuqtadan keyin Slow Start to'xtaydi va **Congestion Avoidance** boshlanadi.

```text
Default ssthresh = 65535 bytes

Agar cwnd < ssthresh:
    Slow Start (exponential)

Agar cwnd >= ssthresh:
    Congestion Avoidance (linear)
```

---

## Congestion Avoidance

```text
cwnd
 ^
 |                          /
 |                        /
 |                      /
 |                    /
 |                  /
 |                /
 |              /
 |            /
 |          /
 |        /
 |      /
 |    /
 |  /
 |/
 +---------------------------------> vaqt

Slow Start    Congestion Avoidance
(exponential)  (linear)
```

---

## Packet Loss

Agar packet yo'qolsa:

```text
3 duplicate ACK:
    cwnd = ssthresh / 2
    Congestion Avoidance ga o't

Timeout:
    cwnd = 1
    Slow Start qayta boshla
```

---

## Amaliyot

### Slow Start ni kuzatish

```bash
sudo tcpdump -i any port 8080 -nn
```

### ss bilan TCP metrics

```bash
ss -ti dst 93.184.216.34
```

Natija:

```text
cubic wscale:7,7 rto:204 rtt:35.5/0.5 ato:40 mss:1448 cwnd:10 ssthresh:7 bytes_sent:1024
```

### Go orqali

```go
package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    start := time.Now()
    conn, _ := net.Dial("tcp", "example.com:80")
    elapsed := time.Since(start)

    fmt.Printf("Connection time: %v\n", elapsed)
    conn.Close()
}
```

---

## Xulosa

- Slow Start sekin boshlaydi, tez oshiradi
- Exponential growth: 1, 2, 4, 8, 16...
- ssthresh -- chegaradan oshganda Congestion Avoidance
- Packet loss -- cwnd kamayadi
