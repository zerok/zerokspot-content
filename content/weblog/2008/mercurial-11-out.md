---
date: '2008-12-03T12:00:00-00:00'
language: en
tags:
- hg
- mercurial
title: Mercurial 1.1 is out
---


Finally. [Mercurial 1.1][] is out with a [tons of great new features and fixes][] and finally *without* deprecation warnings if you're using Python 2.6 ;-) Also most of the extensions got some new features like the pager-extension finally having an option to use it only on a subset of the available commands using the ``attend``-flag in the configuration:
    
    [pager]
    attend = diff,log
    
On the web-front there are some new and improved themes as well as improved WSGI support, which I really have to give a try again after all these months with git ;-) In general, this release looks really great and it (as well as [bitbucket.org][]) will probably make me want to use Mercurial a little bit more again in the future.

[tons of great new features and fixes]: http://www.selenic.com/mercurial/wiki/index.cgi/WhatsNew#head-b1d1f9a535adb686d6e0a490e049261313f10d7d
[Mercurial 1.1]: http://selenic.com/pipermail/mercurial/2008-December/022670.html
[bitbucket.org]: http://bitbucket.org