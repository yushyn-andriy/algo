import socket
import os

SOCKET_FILE = './echo_socket'

def main():
    if os.path.exists(SOCKET_FILE):
        os.remove(SOCKET_FILE)

    server_socket = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
    server_socket.bind(SOCKET_FILE)
    server_socket.listen(5)
    print("Server is listening on UNIX socket...")

    client_socket, _ = server_socket.accept()
    print("Accepted connection")

    data = client_socket.recv(1024)
    while data:
        print(f"Received: {data.decode()}")
        client_socket.sendall(data)
        data = client_socket.recv(1024)

    client_socket.close()
    server_socket.close()
    os.remove(SOCKET_FILE)

if __name__ == "__main__":
    main()
