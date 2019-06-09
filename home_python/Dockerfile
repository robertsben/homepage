FROM python:3

WORKDIR /app

ENV PYTHONPATH .

COPY templates templates
COPY data data
COPY src src

CMD ["python", "src/app.py"]
