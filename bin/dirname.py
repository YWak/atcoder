#!/usr/bin/env python3

"""
online-judge-toolsの履歴からディレクトリ名を生成します。
"""
from datetime import datetime
import json
import os

historyJson = os.environ.get('HOME') + '/.cache/online-judge-tools/download-history.jsonl'
with open(historyJson) as f:
    info = json.load(f)

ts = datetime.fromtimestamp(info["timestamp"])
url = info["url"]
path = url[url.rfind("/")+1:]

print(ts.strftime("%Y/%m/%d/") + path)
