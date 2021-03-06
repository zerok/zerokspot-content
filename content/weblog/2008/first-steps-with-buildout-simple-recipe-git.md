---
date: '2008-10-06T12:00:00-00:00'
language: en
tags:
- buildout
- git
title: 'First steps with zc.buildout: Simple recipe for git'
---


Just a small plug for a project I'm currently working on. I'm playing a little bit around with [zc.buildout](http://pypi.python.org/pypi/zc.buildout/) right now and have come to a point where I want to have dependencies not only as subversion-repository, distutils or egg. I simply end up with too many Python packages on github for not having a way to use them. So while still learning my way around buildout I wrote a little recipe to do just that: clone a git-repo and make it accessible as a part. With this you can define first of all where the repository is, and also what branch or explicit revision you want to have for your project.

There is definitely still a ton of stuff missing that you'd find in any other zc.buildout recipe, but I see this mostly as a good way to finally find my way around :-) Anyway, you can find it on the [PyPI](http://pypi.python.org/pypi/zerokspot.recipe.git/) with a source package and an egg for Python 2.6. The code is also available on [github.com](http://github.com/zerok/zerokspot.gitrecipe/).