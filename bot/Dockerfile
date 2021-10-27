FROM python:3.8-slim-buster

WORKDIR /usr/src/app

COPY ./ ./
RUN pip install --no-cache-dir -r requirements.txt

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait
