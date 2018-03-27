#!/usr/bin/env python
# -*- coding: utf-8 -*-

import datetime
import json

def moredata(topic, data, srv=None):
    payload = json.loads(data['payload'])   # payload is the origin JSON

    t = datetime.datetime.fromtimestamp(payload['tst']).strftime('%Y-%m-%d %H:%M:%S')
    payload['tstamp'] = t
    return payload

