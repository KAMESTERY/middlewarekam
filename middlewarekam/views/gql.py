from pyramid.view import view_config

# from .schema import schema
from ..services.schema import execute_query

@view_config(
    route_name='graphiql',
    renderer='templates/graphiql.jinja2'
)
def gql_view(request):
    return {}

@view_config(
    route_name='gql_query',
    request_method='POST',
    renderer='json'
)
def gql_query(request):
    result = None
    post = request.json_body

    if 'query' in post:
        query = post.get('query')
        # result = schema.execute(query)
        result = execute_query(query)
    elif 'mutation' in post:
        mutation = post.get('mutation')
        # result = schema.execute(mutation)
        result = execute_query(mutation)
    else:
        result = dict(data=dict())

    return dict(data=result.data)
