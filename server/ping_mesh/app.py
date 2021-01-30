import json

from fastapi import FastAPI
from fastapi.encoders import jsonable_encoder

from .model import PingMeshResult
from .redis import (
    client as redis_client,
    get_all_values,
    search_by_hostname,
)
from .config import get_config


LATEST_TAG = 'latest'
DEFAUT_RESPONSE_COUNT = 100


app = FastAPI()


def get_results(agent_hostname: str = ''):
    return search_by_hostname(agent_hostname) if agent_hostname else get_all_values()


@app.post('/ping-mesh')
async def post_ping_mesh(req: PingMeshResult):
    req_dict = jsonable_encoder(req)

    key = req.timestamp.isoformat()
    redis_client.set(key, json.dumps(req_dict))
    redis_client.hset(LATEST_TAG, req.agent_hostname, json.dumps(req_dict))
    redis_client.expire(key, get_config().redis_expire_time)

    return "success"


@app.get('/ping-mesh')
async def get_ping_mesh(agent_hostname: str = ''):
    # Returns latest 100 results
    return get_results(agent_hostname)[0:DEFAUT_RESPONSE_COUNT]


@app.get('/ping-mesh/all')
async def get_ping_mesh_all(agent_hostname: str = ''):
    return get_results(agent_hostname)


@app.get('/ping-mesh/latest')
async def get_ping_mesh_latest(agent_hostname: str = ''):
    if not agent_hostname:
        values = sorted(
            [
                json.loads(val.decode())
                for val in redis_client.hgetall(LATEST_TAG).values()
            ],
            key=lambda x: str(x.get('agent_hostname', '')),
            reverse=True,
        )
        if not values:
            return None
        else:
            return values[0]

    value = redis_client.hget(LATEST_TAG, agent_hostname)
    if not value:
        return None
    return json.loads(redis_client.get(LATEST_TAG))
