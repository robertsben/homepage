"""
Response templating
"""
import os
import logging
import json
from string import Template
from mimetypes import types_map

HTML_MIME = types_map['.html']
PLAIN_MIME = types_map['.txt']
JSON_MIME = types_map['.json']

LOGGER = logging.getLogger(__name__)

_LOADED_DATA = {}


def quick_mime(ext):
    return types_map['.%s' % ext]


class TemplateReference:
    """
    Holds a reference to a template
    """
    def __init__(self):
        self._name = None

    @property
    def name(self):
        return self._name

    @name.setter
    def name(self, _name):
        self._name = _name


def template_data(name):
    """
    Give a name get a template data

    :param name:
    :type name: str
    :rtype: dict[str, str]
    """
    if name not in _LOADED_DATA:
        with open(os.path.join(os.path.dirname(__file__), '..', 'data', '%s.json' % name)) as data_file:
            _LOADED_DATA[name] = json.load(data_file)
    return _LOADED_DATA[name]


def template_reference(name):
    """
    Give a name get a template reference

    :param name:
    :type name: str
    :rtype: TemplateReference
    """
    ref = TemplateReference()
    ref.name = name
    return ref


def _template(content_type, filename):
    """
    Give a name get the template contents

    :param content_type: content type of response
    :type content_type: str
    :param filename: template name
    :type filename: str
    :rtype: str
    """
    with open(os.path.join(os.path.dirname(__file__), '..', 'templates', content_type, filename)) as t_file:
        return Template(t_file.read())


def template_loader(func):
    def wrapper(*args, **kwargs):
        self = args[0]
        template_reference = args[1]
        t_name = template_reference.name
        if t_name not in self.templates:
            self.templates[t_name] = _template(self.mime_type, '%s.%s' % (t_name, self.mime_type))
        return func(*args, **kwargs)
    return wrapper


class TemplateEngine:
    """
    Templating engine
    """
    def __init__(self, mime_type):
        self.templates = {}
        self.template_data = {}
        self.mime_type = mime_type

    @template_loader
    def render(self, _template_reference, data):
        return self.templates[_template_reference.name].substitute(data)


TEMPLATE_ENGINES = {
    JSON_MIME: TemplateEngine('json'),
    HTML_MIME: TemplateEngine('html'),
    PLAIN_MIME: TemplateEngine('txt')
}


def render(data, _template_reference, mime):
    """

    :param data:
    :type data: dict[str, str]
    :param _template_reference: template reference
    :type _template_reference: src.templating.TemplateReference
    :param mime: mime type to render to
    :type mime: str
    :return:
    """
    engine = TEMPLATE_ENGINES[mime]
    return engine.render(_template_reference, data)
