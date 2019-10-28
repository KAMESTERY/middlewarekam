from pyramid.view import view_config
from pyramid.response import Response

# from .schema import schema
from ..services.schema import execute_query

@view_config(
    route_name='graphiql',
    renderer='templates/graphiql.jinja2'
)
def gql_view(request):
    return {}

@view_config(
    route_name='gql_query_options',
    request_method='OPTIONS',
    renderer='json'
)
def gql_query_options(request):
    response = Response(body={})
    response.headers.update({
        'Access-Control-Allow-Origin', '*',  # Callable from ANY domains #TODO: Revisit this!!!!
        'Access-Control-Allow-Methods', 'OPTIONS, POST',
        'Access-Control-Allow-Headers', 'Authorization',
        'Access-Control-Allow-Credentials', 'true',
        'Content-Type', 'application/json'
    })

    return response

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

    response = Response(body=dict(data=result.data))
    response.headers.update({
        'Access-Control-Allow-Origin', '*', # Callable from ANY domains #TODO: Revisit this!!!!
        'Access-Control-Allow-Methods', 'OPTIONS, POST',
        'Access-Control-Allow-Headers', 'Authorization',
        'Access-Control-Allow-Credentials', 'true',
        'Content-Type', 'application/json'
    })

    return response
