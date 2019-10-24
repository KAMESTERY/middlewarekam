import asyncio

import graphene

from graphql.execution.executors.asyncio import AsyncioExecutor

from .media import (
    MediaQuery,
)

class Query(
    graphene.ObjectType,
    MediaQuery,
):
    pass

schema = graphene.Schema(
    query=Query
)

def execute_query(query):
    res = schema.execute(
            query,
            executor=AsyncioExecutor(loop=asyncio.new_event_loop())
        )
    return res
