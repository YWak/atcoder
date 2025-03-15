#!/usr/bin/env python

"""
online-judge-toolsの履歴からURLを取得します。
"""
import json
import sys

with open(sys.argv[1] + '/download-history.jsonl') as f:
    info = json.load(f)

print(info["url"])
