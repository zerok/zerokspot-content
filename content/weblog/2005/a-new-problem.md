---
date: '2005-05-02T12:00:00-00:00'
language: en
tags:
title: A new problem
---


While trying to compile and install PHP and Apache I've noticed a problem: All the module installation commands (should be libtool calls) where blank. Not I at least have a feeling what's going on there. In the libtool wrapper script provided by Apache (which is also used for PHP later on) the actual libtool calls are stored in the variable $CMD. The problem is, that this variable seems to be reserved, so it can't be used for normal assignments. After replacing the variable with $MYCMD (no cookie for me tonight) `make install` worked as was to be expected of an installation label.

-------------------------------



Now I at least know a work-around, but I'd be far more interested in what's causing this. I'm using GNU Bash 3.00.16(1) and I'd really love to learn of a permanent solution for this problem :-(