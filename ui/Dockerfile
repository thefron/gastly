FROM python:3.6

RUN apt-get update && apt-get upgrade -y && apt-get autoremove && apt-get autoclean

WORKDIR /app
COPY manage.py requirements.txt /app/
RUN pip install -r requirements.txt
COPY . /app/

EXPOSE 8001
CMD ["python", "manage.py", "runserver", "0.0.0.0:8001"]
