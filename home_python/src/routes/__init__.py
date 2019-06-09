"""
all routes
"""
from .about import METHODS as ABOUT_METHODS, ROUTE as ABOUT_ROUTE
from .index import METHODS as INDEX_METHODS, ROUTE as INDEX_ROUTE
from .python import METHODS as PYTHON_METHODS, ROUTE as PYTHON_ROUTE
from .error import error_response

__all__ = ['ALL', 'error_response']

GET_VERB = b'GET'


ALL = {
    INDEX_ROUTE: INDEX_METHODS,
    ABOUT_ROUTE: ABOUT_METHODS,
    PYTHON_ROUTE: PYTHON_METHODS,
}
