
# curl -O http://snap.stanford.edu/data/bigdata/amazon/amazon-meta.txt.gz

import sys
import gzip
import re
import ophion

def line_read(line):
    res = re.search(r'^\s*([\w]+):\s+(.*)\s*$', line.rstrip())
    if res is None:
        return None
    return ( res.group(1), res.group(2) )


class Writer:
    def __init__(self, host):
        self.o = ophion.Ophion(host)
        self.record_count = 0

    def add_record(self, rec):
        q = self.o.query().addV(rec['ASIN'])
        
        for i in ["Id", "group", "title", "salesrank"]:
            if i in rec:
                q = q.property(i, rec[i])
        q.execute()
        
        for i in rec.get('similar', []):
            self.o.query().V(rec['ASIN']).addE("similar").to(i).execute()
        
        self.record_count += 1
        if self.record_count % 1000 == 0:
            print "%s vertices written" % (self.record_count)
        
    
writer = Writer(sys.argv[2])

with gzip.GzipFile(sys.argv[1]) as handle:
    
    record = None
    for line in handle:
        if len(line.rstrip()) == 0:
            if record is not None:
                writer.add_record(record)
                record = None
        else:
            e = line_read(line)
            if e:
                if e[0] == "Id":
                    record = {"Id" : e[1]}
                elif e[0] == "ASIN":
                    record["ASIN"] = e[1]
                elif e[0] in ["group", "title", "salesrank"]:
                    record[e[0]] = e[1]
                elif e[0] == "similar":
                    record['similar'] = e[1].split("  ")[1:]
                elif e[0] == "categories":
                    for i in range(int(e[1])):
                        l = handle.readline()
                        #print e
                elif e[0] == "reviews":
                    #print "review", e[1].split("  ")
                    record['reviews'] = []
                    f = line_read(e[1].split("  ")[1])
                    for i in range(int(f[1])):
                        l = handle.readline().strip()
                        eres = re.search("([^\s]+)\s+cutomer:\s+(\w+)\s+rating:\s+(.)\s+votes:\s+(\d+)\s+helpful:\s+(\d+)", l)
                        if eres is None:
                            print "review", e[1]
                            print "review", l
                        date, customer, rating, votes, helpful = eres.groups()
                        record['reviews'].append({
                            "date" : date,
                            "customer" : customer,
                            "rating" : rating,
                            "votes" : votes,
                            "helpful" : helpful
                        })
                else:
                    print e
                #print e[0]
            