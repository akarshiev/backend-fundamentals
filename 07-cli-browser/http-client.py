import socket
import sys

def main():
    if len(sys.argv) < 2:
        print("Usage: python http-client.py <host>")
        print("Example: python http-client.py example.com")
        return

    host = sys.argv[1]

    # Step 1: TCP connection
    print(f"Connecting to {host}:80...")
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    sock.connect((host, 80))
    print("Connected!")

    # Step 2: Build HTTP request
    request = (
        f"GET / HTTP/1.1\r\n"
        f"Host: {host}\r\n"
        f"Connection: close\r\n"
        f"\r\n"
    )

    # Step 3: Send request
    print(f"\nSending request:\n{request}")
    sock.send(request.encode())

    # Step 4: Read full response
    print("Response:")
    print("--------")

    chunks = []
    while True:
        data = sock.recv(4096)
        if not data:
            break
        chunks.append(data)

    response = b"".join(chunks)
    print(response.decode(errors="ignore"))

    sock.close()

if __name__ == "__main__":
    main()
