FROM python:latest
ENV APPDIR=src
WORKDIR ${APPDIR}
COPY src .
RUN pip install -r requirements.txt
CMD python main.py