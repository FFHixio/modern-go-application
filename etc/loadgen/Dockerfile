FROM python:3.11.0a5

COPY requirements.txt .
RUN pip install -r requirements.txt

COPY . .
ENTRYPOINT ./loadgen.sh
