import json
from datetime import datetime

from fastapi import FastAPI
from fastapi.encoders import jsonable_encoder
from pydantic import BaseModel
from typing import List


LOCAL_SAVE_FILE_PATH = './output.json'


app = FastAPI()


class IPAddress(BaseModel):
    ip: str
    zone: str


class Statistics(BaseModel):
    packet_receive: int
    packet_sent: int
    packet_loss: float
    ip_address: IPAddress
    address: str
    rtts: List[str]
    min_rtt: str
    max_rtt: str
    avg_rtt: str
    standard_deviation_rtt: str


class Host(BaseModel):
    name: str
    host: str
    result: Statistics


class PingMeshResult(BaseModel):
    timestamp: datetime
    hosts: List[Host]


def json_local_save(payload: dict):
    with open(LOCAL_SAVE_FILE_PATH, 'w') as f:
        json.dump(payload, f, indent=4)


@app.post('/ping-mesh')
async def ping_mesh(req: PingMeshResult):
    req_dict = jsonable_encoder(req)
    json_local_save(req_dict)

    return "success"
