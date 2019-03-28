#!/usr/bin/env python
# -*- coding: utf-8 -*-
import sys
if len(sys.argv) != 2:
    print "invalid number of args: ./inline_replace.py filename"
    exit(-1)
filename = sys.argv[1]

with open(filename, "r") as fd:
    lines = fd.readlines()
    if len(lines) == 0:
        exit(0)

    new_lines = []
    for index, line in enumerate(lines):
        if line.find("platform_quota") != -1:
            continue
        line = line.replace(",omitempty", "")
        lines[index] = line

with open(filename, "w") as fd:
    if lines[0].strip() != "//@generated":
        fd.write("//@generated\n")
    fd.writelines(lines)
