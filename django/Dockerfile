FROM python:3.13-slim

ENV PYTHONDONTWRITEBYTECODE 1 # Prevents Python from writing pyc files to disc
ENV PYTHONFAULTHANDLER 1 # Makes Python use fault handler

RUN pip install ruff
RUN useradd -ms /bin/bash django
RUN mkdir /media/uploads && chown -R django:django /media/uploads

USER django

WORKDIR /home/django

CMD tail -f /dev/null