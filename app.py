import os.path
import site

# add `lib` subdirectory as a site packages directory, so our `main` module can load
# third-party libraries.
site.addsitedir(os.path.join(os.path.dirname(__file__), 'lib'))

import logging
import os
from paste.deploy import loadapp
from pyramid.paster import (
    setup_logging,
    get_app
)
from asgiref.wsgi import WsgiToAsgi

log = logging.getLogger(__name__)

def bootstrap():
    log.info("Bootstrapping Application...")
    # app = loadapp(
    #     'config:production.ini',
    #     relative_to=os.path.dirname(
    #         os.path.realpath(__file__)
    #     )
    # )
    setup_logging(
        f'{os.path.dirname(os.path.realpath(__file__))}/production.ini'
    )
    app = get_app(
        f'{os.path.dirname(os.path.realpath(__file__))}/production.ini'
    )
    return WsgiToAsgi(app)

app = bootstrap()

def main():
    import os
    import uvicorn
    from multiprocessing import cpu_count
    uvicorn.run(app, host='0.0.0.0', port=int(os.environ.get("PORT", 8000)), workers=cpu_count(), reload=True)

if __name__ == "__main__":
    main()
