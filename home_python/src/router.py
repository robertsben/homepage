"""
Routing mechanisms
"""
import logging
from src.exceptions import HTTPError
from src.routes import ALL as ROUTES, error_response
from http import HTTPStatus

LOGGER = logging.getLogger(__name__)

error = error_response


def find_route(method, path):
    """

    :param data:
    :rtype: tuple[HTTPStatus, str or None]
    """
    try:
        avail_verbs = ROUTES[path]
        LOGGER.debug('Found route %s', path)
    except KeyError:
        raise HTTPError(HTTPStatus.NOT_FOUND)

    try:
        route = avail_verbs[method]
        LOGGER.debug('With method %s', method)
    except KeyError:
        raise HTTPError(HTTPStatus.METHOD_NOT_ALLOWED)

    return route
