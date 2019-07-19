import json
import os
from subprocess import PIPE, Popen
import sys
import time

from typeof import isdict, islist

APP = os.path.abspath(os.path.join(os.path.dirname(__file__), "..", "..", "bin", "scloud"))

# Prints an error message and optionally exits.
def error(msg, exit_code=1):
    sys.stderr.write("error: ")
    sys.stderr.write(msg)
    sys.stderr.write('\n')
    if exit_code: sys.exit(exit_code)


def fatal(msg):
    error(msg, 1)


# Returns integer epoch time in seconds.
def epoch():
    return int(time.time())


# Ensures that all strings referenced by the given value are utf8 encoded.
def utf8(obj):
    if isdict(obj):
        return {utf8(k): utf8(v) for k, v in obj.iteritems()}
    if islist(obj):
        return [utf8(item) for item in obj]
    if isinstance(obj, unicode):
        return obj.encode("utf-8")
    return obj


def loads(sz):
    try:
        return utf8(json.loads(sz))
    except Exception:
        return sz


# Returns <exit_code>, <stdout>, <stderr>
def scloud(*args):
    args = [APP] + list(args)
    #print " ".join(args)
    p = Popen(args, stdout=PIPE, stderr=PIPE)
    sout, serr = p.communicate()
    sout, serr = sout.strip(), serr.strip()
    rerr = None
    if len(serr) > 0:
        if serr.startswith("error: "):
            serr = serr[7:]
        rerr = loads(serr)
    rout = None
    if len(sout) > 0:
        rout = loads(sout)
    return p.returncode, rout, rerr

def is400(obj):
    return obj.get("HTTPStatusCode", None) == 400

def is400servicex(obj):
    return obj.get("status", None) == 400

def is404(obj):
    return obj.get("HTTPStatusCode", None) == 404

def is404servicex(obj):
    return obj.get("status", None) == 404

def is409(obj):
    return obj.get("HTTPStatusCode", None) == 409

def is409servicex(obj):
    return obj.get("status", None) == 409

def has_valid_token():
    code, _, _ = scloud("identity", "validate-token")
    return code == 0


# Ensure the local scloud has a valid access token.
def ensure_token():
    if not has_valid_token():
        fatal("please run: scloud login")


# Ensure teh local scloud has a valid tenant setting.
def ensure_tenant():
    code, tname, _ = scloud("get", "tenant")
    if code or tname is None:
        fatal("please run: scloud set tenant <tenant-name>")
    code, tenant, _ = scloud("identity", "get-tenant", tname)
    if code:
        fatal("unable to find selected tenant")
    if tenant.get("status", "") != "ready":
        fatal("please select a tenant with status 'ready'")


def setup():
    ensure_token()
    ensure_tenant()


def main():
    setup()


if __name__ == "__main__":
    main()
