
import os
import inspect

from .query import Query
from gripql.util import AttrDict, Rate, process_url

class Compute(Query):

    def __init__(self, url, graph, user=None, password=None):
        self.query = []
        url = process_url(url)
        self.base_url = url
        self.url = url + "/v1/graph/" + graph + "/compute"
        self.graph = graph
        if user is None:
            user = os.getenv("GRIP_USER", None)
        self.user = user
        if password is None:
            password = os.getenv("GRIP_PASSWORD", None)
        self.password = password

    #I don't know why I had to copy this itself.
    #It wouldn't let me use the copy inherited by the superclass
    def __append(self, part):
        q = self.__class__(self.base_url, self.graph)
        q.query = self.query[:]
        q.query.append(part)
        return q

    def messageV(self, message, reduce):
        mSrc = inspect.getsource(message)
        rSrc = inspect.getsource(reduce)
        return self.__append({
            "messageV" : {
                "message": { "starlark" : mSrc },
                "reduce": { "starlark" : rSrc }
            }
        })

    def mapV(self, mapper):
        src = inspect.getsource(mapper)
        return self.__append({
            "mapV" : {
                "map": { "starlark" : src }
            }
        })
