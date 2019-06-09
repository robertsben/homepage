"""
Custom exceptions
"""


class HTTPError(Exception):
    """
    Raised to produce a http error response
    """
    def __init__(self, status):
        """
        :param status:
        :type status: http.HTTPStatus
        """
        self.status = status
