---
date: '2007-02-08T12:00:00-00:00'
language: en
tags:
- django
- dreamhostt
- python
title: Django working again
---


As a way to cope with stress I tend to start coding random stuff and also to simply mess with stuff I shouldn't. But this time I actually did 2 at least not catastrophical things: 

1. Teh return of the gradients on zerokspot.com. I hope you like them ;-)
2. I finally got Django to work again on my Dreamhost account.



-------------------------------



Esp. the 2nd item took me quite some time. About 2 weeks ago when [Django 0.95.1 got released](http://www.djangoproject.com/weblog/2007/jan/21/0951/) I immediately updated but in the process killed my Django site which isn't all that hard when you take the default steps to get Django running in the first place, as quite a few people can probably attest who still haven't got Django working on their Dreamhost accounts.

So yesterday I started playing around again since it annoys the hell out of me when I can't get something to work which should work out of the box. So I started once again tracing the problems and finding their roots. But everything looked okay when I ran the dispatch.fcgi from the commandline.

Then I created a small test.fcgi that should first of all as your usual WSGI "Hello World" script. Worked right away.

Then I simply wanted it to also give me the version of number of the Django installation. Not so much, anymore. A small site note about how I normally use Python libraries and also did so on Dreamhost: 

* I install everything for Python2.4 with ~/.python as prefix which puts 3rd party libraries into ~/.python/lib/python2.4/site-packages
* `export PYTHONPATH=~/.python/lib/python2.4/site-packages`

This _should_ normally work, but for some reason doesn't on Dreamhost. I think I found the problem with the Python installation on my server not understanding .pth files which are use by Python to manage libraries and are also used by Django. I'm not completely sure why it doesn't work on Dreamhost but it has perhaps something to do with the way I've configured by paths. 

But given that Django also didn't work when I tried it without actually installing it but just adjusting sys.path in the dispatch.fcgi to also include the actual library, I guess there might be something strange going on with my setup or with the whole server ...

Anyway: Easiest solution at least for me? I simply rolled out my own python installation where I can completely control all the default settings without that much extra effort. And everything seems to work so far :-)

If you're one of those people who still get 500-errors all the time while trying to get Django to work, perhaps this will also work for you :-) There's also a [page](http://wiki.dreamhost.com/index.php/Python) about this in the Dreamhost wiki, but I guess, if you've ever installed something under Linux/Unix/MacOSX from source, you know the procedure ;-)

The only thing that I think is a little bit strange there, is that the author uses $HOME as prefix. My totally person preference would be $HOME/local.