FROM python:3.12

WORKDIR /app

RUN pip install numpy

COPY . .

CMD python /app/one.py && python /app/two.py && python /app/three.py