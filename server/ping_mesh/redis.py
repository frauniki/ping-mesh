import json

import redis

from .config import get_config


client = redis.from_url(get_config().redis_url)


def get_all_values() -> [dict]:
    keys = client.keys('*')
    res = []

    for key in keys:
        res.append(json.loads(client.get(key)))

    return sorted(res, key=lambda x: x['timestamp'], reverse=True)


def search_by_hostname(hostname: str) -> [dict]:
    res: [dict] = []
    for value in get_all_values():
        if value.get('agent_hostname') == hostname:
            res.append(value)
    return res
