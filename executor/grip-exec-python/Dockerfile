FROM python:3.6

RUN pip install grpcio-tools
ADD *.py /bin/

EXPOSE 50000
ENTRYPOINT /bin/grip-exec.py
