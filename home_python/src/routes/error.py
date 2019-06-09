"""
Error routes
"""
from src.templating import template_reference


def error_response(status, description=None):
    data = {
        'code': status.value,
        'status': status.phrase,
        'description': description or status.description
    }
    return data, template_reference('error')
