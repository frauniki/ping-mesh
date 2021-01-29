FROM tiangolo/uvicorn-gunicorn-fastapi:python3.8

WORKDIR /app
COPY ./server ./server
COPY ./setup.py .

RUN pip install .

CMD [ "uvicorn", "ping_mesh:app", "--host", "0.0.0.0", "--port", "8000" ]
