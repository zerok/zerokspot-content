---
date: '2012-01-15T12:00:00-00:00'
language: en
tags:
- python
title: When setup.py check fails
---


Last night I was about to release a new version of django-flatblocks, when I ran
into a weird problem:

<pre class="code">
$ python setup.py check -r                                                                                                                                                                           [develop] 13:58+0100
running check
Traceback (most recent call last):
  File "setup.py", line 32, in <module>
    zip_safe = False,
  File ".../distutils/core.py", line 152, in setup
    dist.run_commands()
  File ".../distutils/dist.py", line 953, in run_commands
    self.run_command(cmd)
  File ".../distutils/dist.py", line 972, in run_command
    cmd_obj.run()
  File ".../distutils/command/check.py", line 69, in run
    self.check_restructuredtext()
  File ".../distutils/command/check.py", line 111, in check_restructuredtext
    for warning in self._check_rst_data(data):
  File ".../distutils/command/check.py", line 138, in _check_rst_data
    parser.parse(data, document)
  File ".../docutils/parsers/rst/__init__.py", line 157, in parse
    self.statemachine.run(inputlines, document, inliner=self.inliner)

    .....

  File ".../docutils/parsers/rst/states.py", line 991, in implicit_inline
    return [nodes.Text(unescape(text), rawsource=unescape(text, 1))]
  File ".../docutils/nodes.py", line 331, in __new__
    return reprunicode.__new__(cls, data)
UnicodeDecodeError: 'ascii' codec can't decode byte 0xc3 in position 26: ordinal not in range(128)
</pre>

The cause of this issue was a bit unexcected: A string in the release notes of
a version I had already released about a year ago. So what changed? Probably
the Python version but also most definitely the version of [docutils][3] I had
installed in order to preview the release notes before every release.

So I started downgrading from 0.8.1 all the way to 0.5.x to finally make the
check command not die on me anymore.

There is also an issue in [Python's issue tracker][1] which has a [patch][2]
attached (and committed) that solves this issue. So if you don't want to
downgrade docutils in your virtualenv, this is probably the way to go :-)


[1]: http://bugs.python.org/issue13114
[2]: https://bitbucket.org/mirror/cpython/changeset/8d837bd8148a
[3]: http://docutils.sourceforge.net/
