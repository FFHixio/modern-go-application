FROM python:3.10.0b2

COPY requirements.txt .
RUN pip install -r requirements.txt

COPY . .
ENTRYPOINT ./loadgen.sh
