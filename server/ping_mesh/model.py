from datetime import datetime

from pydantic import BaseModel
from typing import (
    List,
    Optional,
)


class IPAddress(BaseModel):
    ip: str
    zone: str


class Statistics(BaseModel):
    packet_receive: int
    packet_sent: int
    packet_loss: float
    ip_address: IPAddress
    address: str
    rtts: Optional[List[str]] = None
    min_rtt: str
    max_rtt: str
    avg_rtt: str
    standard_deviation_rtt: str


class Host(BaseModel):
    name: str
    host: str
    result: Optional[Statistics] = None


class PingMeshResult(BaseModel):
    agent_hostname: str
    timestamp: datetime
    hosts: List[Host]
