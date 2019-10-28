def includeme(config):
    config.add_static_view('static', 'static', cache_max_age=3600)
    config.add_route('home', '/')
    config.add_route('gql_query_options', '/gql_query')
    config.add_route('gql_query', '/gql_query')
    config.add_route('graphiql', '/graphiql')
    config.add_route('media', '/media/{media_id}')
