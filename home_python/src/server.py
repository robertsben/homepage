"""
Request handling
"""
import logging
import socket
import ssl
from http import HTTPStatus
from src.exceptions import HTTPError
from src.router import find_route, error
from src.templating import HTML_MIME, render

LOGGER = logging.getLogger(__name__)

BUFF_SIZE = 4096
SSL_CTX = ssl.create_default_context(ssl.Purpose.CLIENT_AUTH)
SSL_CTX.load_cert_chain(certfile='/app/certs/homepage.crt', keyfile='/app/certs/homepage.key')


def create_response(status, mime_type, response, headers=None):
    """
    :param status: http response status (code)
    :type status: http.HTTPStatus
    :param mime_type: mime type
    :type mime_type: str
    :param response: templated response body
    :type response: str
    :param headers: headers to add to the response
    :type headers: dict[str, str]
    :rtype: bytes
    """
    preamble = 'HTTP/1.1 %d %s\n' % (status, status.phrase)
    content_type = 'Content-Type: %s\n' % mime_type
    content_length = 'Content-Length: %d\n' % len(response)
    extra_headers = '\n'.join(['%s: %s' % (k, v) for k, v in headers.items()]) if headers else ''
    return bytes(preamble + content_type + content_length + extra_headers + '\n' + response, 'utf8')


def parse_request(data):
    """
    Given the data received, parse it

    :param data: byte string of request
    :type data: bytes
    :rtype: bytes,
    """
    try:
        request_lines = data.splitlines()
        method, path, protocol = request_lines.pop(0).split(b' ')
        header_lines = request_lines[:request_lines.index(b'')]
        headers = {header[0].decode('utf8').lower(): header[1].lstrip() for header in
                   [header.split(b':', 1) for header in header_lines]}
        body = b'\n'.join(request_lines[request_lines.index(b'') + 1:])
        return method, path, headers, body
    except (AttributeError, IndexError):
        raise HTTPError(HTTPStatus.BAD_REQUEST)
    except Exception:
        raise HTTPError(HTTPStatus.INTERNAL_SERVER_ERROR)


def parse_path(path):
    """
    Given a path, get the query string params.

    :param path:
    :type path: bytes
    :rtype: bytes,
    """
    try:
        parts = path.split(b'?')
        if len(parts) == 1:
            return parts[0], {}

        params = {param[0].decode('utf8'): param[1] for param in [kv.split(b'=', 1) for kv in parts[1].split(b'&')]}
        return parts[0], params
    except (AttributeError, IndexError):
        raise HTTPError(HTTPStatus.BAD_REQUEST)
    except Exception:
        raise HTTPError(HTTPStatus.INTERNAL_SERVER_ERROR)


def determine_content_types(headers):
    """
    Find the acceptable content types from Accept in the headers

    :param headers:
    :type headers: dict[str, bytes[
    :return:
    """
    try:
        content_types = [
            ct.lstrip() for ct in headers['accept'].decode('utf8').split(',') if ct.lstrip() not in ('text/*', '*/*')
        ]
    except KeyError:
        return HTML_MIME

    return content_types if content_types else [HTML_MIME]


def render_body_in_acceptable_mime(render_data, template, headers):
    """
    Given the render data, the template, and the request headers,
    render the content in the first acceptable format

    :param render_data:
    :param template:
    :param headers:
    :return:
    """
    for content_type in determine_content_types(headers):
        try:
            body = render(render_data, template, content_type)
            return body, content_type
        except (KeyError, IOError):
            pass
    raise HTTPError(HTTPStatus.NOT_ACCEPTABLE)


def handle_error(err, headers=None):
    """
    Handle HTTPError and create a response from it

    :param err: http status error
    :type err: http.HTTPStatus
    :param headers: any headers
    :type headers: dict[str, str]
    :rtype: bytes
    """
    render_data, template = error(err.status)
    body = render(render_data, template, HTML_MIME)
    return create_response(err.status, HTML_MIME, body, headers=headers)


def handle_request(request_data):
    """
    :param request_data:
    :type request_data: str
    :rtype: bytes
    """
    LOGGER.debug("Handling request: %s", request_data)
    try:
        method, path, headers, body = parse_request(request_data)
        path, params = parse_path(path)
    except HTTPError as err:
        return handle_error(err)

    try:
        route = find_route(method, path)
    except HTTPError as err:
        return handle_error(err)

    response = route(params, headers, body)
    render_data, template = response[:2]
    status = response[2:] or HTTPStatus.OK

    try:
        body, content_type = render_body_in_acceptable_mime(render_data, template, headers)
    except HTTPError as err:
        return handle_error(err)

    return create_response(status, content_type, body)


def receive_from_socket(sock):
    """
    Receive all the data from a socket connection and return it

    :param sock:
    :type sock: socket.socket
    :rtype: bytes
    """
    data = b''
    packet = sock.recv(BUFF_SIZE)
    while True:
        data += packet
        if len(packet) < BUFF_SIZE:
            break
        packet = sock.recv(BUFF_SIZE)

    LOGGER.debug('Received from socket: %s', data)
    return data


def accept_connection(sock):
    """
    :param sock:
    :type sock: socket.socket
    :return:
    """
    LOGGER.debug('Socket accepting')
    connection, address = sock.accept()
    try:
        with SSL_CTX.wrap_socket(connection, server_side=True) as sconn:
            LOGGER.debug(f'Connected by {address}')
            data = receive_from_socket(sconn)
            response = handle_request(data)
            sconn.sendall(response)
    except ssl.SSLError as ssl_err:
        LOGGER.warning(ssl_err)
    except OSError as os_err:
        LOGGER.error(os_err)
