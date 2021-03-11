
from __future__ import absolute_import


def test_path_out_out_out(man):
    errors = []

    G = man.setGraph("swapi")

    count = 0
    for res in G.query().V("Film:1").out().outOpt("pilots"):
        print(res)

    return errors
