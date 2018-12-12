
def mysum(x):
    return sum(x)

def test_map(O):
    errors = []

    O.addVertex("1", "Array", {"values": [1,2,3,4,5,6]})
    O.addVertex("2", "Array", {"values": [1.0,2.0,3.0]})

    q = O.query().V().hasLabel("Array").map(mysum)

    print(q.to_json())

    return errors
