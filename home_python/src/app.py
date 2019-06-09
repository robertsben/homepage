"""
Main app module
"""
import os
import logging
import socket
from src.server import accept_connection

logging.basicConfig(level=logging.DEBUG)
LOGGER = logging.getLogger(__name__)

HOST = ''
PORT = int(os.environ['PORT'])


def main():
    """
    Start function
    """
    with socket.socket(family=socket.AF_INET, type=socket.SOCK_STREAM) as sock:
        sock.bind((HOST, PORT))
        sock.listen(100)
        LOGGER.info('Socket listening on host: %s, port: %s', HOST, PORT)
        while True:
            accept_connection(sock)


if __name__ == '__main__':
    LOGGER.info('Starting....')
    main()
