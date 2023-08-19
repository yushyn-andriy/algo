import socket

SOCKET_FILE = './echo_socket'

def main():
    client_socket = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
    client_socket.connect(SOCKET_FILE)

    message = "Hello, Server through UNIX socket!"
    client_socket.sendall(message.encode())

    data = client_socket.recv(1024)
    print(f"Received from server: {data.decode()}")

    client_socket.close()

if __name__ == "__main__":
    main()
