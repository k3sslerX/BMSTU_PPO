FROM python:3.13-slim

WORKDIR /sync

ENV PYTHONUNBUFFERED=1

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY sync.py .

CMD ["python", "/sync/sync.py"]
