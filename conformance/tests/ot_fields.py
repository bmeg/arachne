
def test_fields(O):
    errors = []

    O.addVertex("vertex1", "person", {"name": "han", "age": 35, "occupation": "smuggler", "starships": ["millennium falcon"]})

    expected = {
        u"gid": u"vertex1",
        u"label": u"person",
        u"data": {u"name": u"han"}
    }
    resp = O.query().V().fields(["name"]).execute()
    if resp[0] != expected:
        errors.append("vertex contains incorrect fields: \nexpected:%s\nresponse:%s" % (expected, resp))

    expected = {
        u"gid": u"vertex1",
        u"label": u"person",
        u"data": {}
    }
    resp = O.query().V().fields(["non-existent"]).execute()
    if resp[0] != expected:
        errors.append("vertex contains incorrect fields: \nexpected:%s\nresponse:%s" % (expected, resp))

    return errors
