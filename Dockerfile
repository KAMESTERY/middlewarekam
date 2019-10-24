FROM python:3.7

RUN mkdir /middlewarekam
WORKDIR /middlewarekam

COPY requirements.txt /middlewarekam
RUN pip install -r requirements.txt

COPY app.py /middlewarekam
COPY production.ini /middlewarekam

CMD python app.py
