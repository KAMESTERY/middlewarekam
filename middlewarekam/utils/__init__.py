import base64

def to_encoded_str(vstr: str) -> str:
    # encdata = base64.b64encode(vstr).replace('/', '!')
    encdata = vstr.replace('/', '!')
    return encdata

def from_encoded_str(b64str: str) -> str:
    # v = base64.b64decode(b64str.replace('!', '/'))
    v = b64str.replace('!', '/')
    return v
