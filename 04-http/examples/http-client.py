import requests
import sys

def main():
    if len(sys.argv) < 2:
        print("Usage: python main.py <url>")
        return

    url = sys.argv[1]
    response = requests.get(url)

    print("Status:", response.status_code)
    print("Headers:")
    for key, value in response.headers.items():
        print(f"  {key}: {value}")

    print("\nBody:")
    print(response.text)

if __name__ == "__main__":
    main()
