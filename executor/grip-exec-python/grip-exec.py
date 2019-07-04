#!/usr/bin/env python

import sys
import grpc
import time
import json
import logging

import exec_pb2
import exec_pb2_grpc
from concurrent import futures


_ONE_DAY_IN_SECONDS = 60 * 60 * 24

logging.basicConfig(filename='grip-exec.log', level=logging.INFO)

class CallableCode:
    def __init__(self, funcName, code):
        self.name = funcName
        self.code = code
        self.env = {}
        exec(code, self.env)

    def call(self, values):
        return self.env[self.name](*values)


class PyGripExec:

    def __init__(self):
        self.code = {}
        self.code_num = 0

    def Compile(self, request, context):
        logging.info("Compile: %s", request.code)
        c = CallableCode(request.function, request.code)
        self.code[self.code_num] = c
        out = exec_pb2.CompileResult()
        out.id = self.code_num
        self.code_num += 1
        return out

    def Process(self, request_iterator, context):
        logging.info("Calling Processor")
        for req in request_iterator:
            if req.code in self.code:
                c = self.code[req.code]
                try:
                    logging.info("calling %s on %s" % (req.code, req.data))
                    data = json.loads(req.data)
                    value = c.call(data)
                    o = exec_pb2.Result()
                    o.data = json.dumps(value)
                    yield o
                except Exception as e:
                    o = exec_pb2.Result()
                    o.error = str(e)
                    logging.info("ExecError: %s" % (o.error))
                    yield o



if __name__ == "__main__":

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    exec_pb2_grpc.add_ExecutorServicer_to_server(
      PyGripExec(), server)
    port = 50000
    while True:
        new_port = server.add_insecure_port('[::]:%s' % port)
        if new_port != 0:
            break
        port += 1
    port = new_port
    server.start()
    print(port, flush=True)
    logging.info("Server started on port %d" % (port))

    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)
