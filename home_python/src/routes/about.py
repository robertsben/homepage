"""
about routes
"""
from src.constants import HTTPVerbs
from src.templating import template_reference, template_data

ROUTE = b'/about'


def get(*args):
    """
    :return: data, template
    :rtype: tuple[http.HTTPStatus, dict[str, str], src.templating.TemplateReference]
    """
    return template_data('about'), template_reference('index')


METHODS = {
    HTTPVerbs.GET: get
}
