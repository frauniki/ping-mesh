import json

from fastapi import FastAPI
from fastapi.encoders import jsonable_encoder

from .model import PingMeshResult
from .redis import (
    client as redis_client,
    get_all_values,
)
from .config import get_config


LATEST_TAG = 'latest'
DEFAUT_RESPONSE_COUNT = 100


app = FastAPI()


@app.post('/ping-mesh')
async def post_ping_mesh(req: PingMeshResult):
    req_dict = jsonable_encoder(req)

    key = req.timestamp.isoformat()
    redis_client.set(key, json.dumps(req_dict))
    redis_client.set(LATEST_TAG, json.dumps(req_dict))
    redis_client.expire(key, get_config().redis_expire_time)

    return "success"


@app.get('/ping-mesh')
async def get_ping_mesh():
    # Returns latest 100 results
    return get_all_values()[0:DEFAUT_RESPONSE_COUNT]


@app.get('/ping-mesh/all')
async def get_ping_mesh_all():
    return get_all_values()


@app.get('/ping-mesh/latest')
async def get_ping_mesh_latest():
    value = redis_client.get(LATEST_TAG)
    if not value:
        return None

    return json.loads(redis_client.get(LATEST_TAG))
