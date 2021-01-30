from functools import lru_cache

from pydantic import BaseSettings


class Config(BaseSettings):
    redis_url: str = 'redis://localhost:6379/0'
    redis_expire_time: int = 86400


@lru_cache()
def get_config():
    return Config()
