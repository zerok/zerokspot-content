---
date: '2007-06-28T12:00:00-00:00'
language: en
tags:
- bzr
- darwinports
- dvcs
- fink
- hg
- macosx
- macports
- software
title: Some software infrastructure changes
---


After quite some time (and probably right before I have to send my Powerbook in for some repairs) I decided I wanted to have something new again, so I first of all thought that my current software management for things like Gimp, Python and Ruby is a complete mess.



-------------------------------

 

I have more or less 3 repositories for software:

1. Fink for basic libraries I use to compile other stuff
2. /opt/sw for ... well, that other stuff like Python, Ruby and Subversion
3. An old DarwinPorts installation that until yesterday wasn't even in my PATH or or LD_LIBRARY_PATH anymore

All of these had their reason d'etre, but ... well, as you can see: It's just a mess. So yesterday  I gave DarwinpPrts - or MacPorts as this project is called now - another look, liked what I saw (py25-readline ;-)) and decided to slowly migrate everything from the first 2 repositories into the 3rd and upgrade it to Macports. 

This also means, that I will have to re-compile some applications that I had outside these repositories (basically my Apache+PHP+MySQL system), but since I hardly do any PHP coding on my PowerBook anymore, this has a very low priority.

Just one word of advice: If you have an old gettext installation in Darwin/MacPorts, don't get the glorious idea, to update it. Esp. when it is a real version jump like from 0.14 to 0.16. I did that and now I spent the better part of the day recompiling everything I had already installed before updating gettext ;-)

The other change I at least consider is switching away from mercurial after I read in some discussion, that it requires a 3rd party merge tool for _any_ merge operation (which was also confirmed on #hg) ... Definitely makes telling people to use it who are quite reluctant to install _anything_.

So I looked at bzr again after quite some time and from what I can tell after about 2 days it comes down to: If you require UTF8 filenames, use Mercurial, if you don't need UTF8 filenames and want to work without external tools and have a built-in merger: bzr. 

I somehow also quite like some of the configuration options of bzr but on some operations it's really slow :-/ Neither of these is a problem for me though:

1. I will slap anyone who uses special characters (as in non-ASCII) in filenames around with a large Duden
2. Nothing against a small coffee break from time to time ;-) It's not that severe since my projects tend to be tiny most of the time.

But I guess for now it really comes down to these 2 criteria. I personally also like some other aspects of bzr more compared to hg, but this is absolutely bonus stuff and probably no one else will care about them so I won't even name them here ;-)

Hmm... it took long enough to write this article for MacPorts to finish the installation of what I had in my custom software repository ... time for some sleep then :)