---
date: '2004-12-30T12:00:00-00:00'
language: en
tags:
- ubuntulinux
title: New to Ubuntu
---


In the last days I played around with the thought of giving Ubuntu a try on my laptop (Acer TM803). Why? Reiser4 (not tweaked) doesn't worked out as I had anticipated on this machine so I wanted to reformated my main partition again and I also wanted to work with a binary distribution once again after about 1.5 years of Gentoo ;-) 

-------------------------------



So I went to <http://www.ubuntulinux.org> and read some feature lists as well as many posts in the english forum. I quite like their quite straight and tight binding to the Gnome project and also Gnome's release schedule so I downloaded the i386 iso image. About 20 minutes later (actually 20 minutes after starting the installation process ;-) ) I had a working Gnome 2.8 on my laptop. Nice: Most parts of my hardware were correctly identified and configured including things like the Synaptic touchpad, the wireless lan interface and X. I was and still am really impressed :-)

I'm now using XFS as primary FS although I was warned that probably grub would make problems here. Seems like my laptop is an exception as I've noticed no problems with XFS on the root-partition and grub as boot loader. So everything went really smooth, but then I noticed some things that are bugging at least a bit:
* only a release candidate or pre-release of Firefox is bundled with Warty and the final version is also not included in the update package
* wxGTK was built against GTK1 (which is recommended by ./configure script but I got used to wxGTK presenting itself in the smoother GTK2 look ;-) )

Solving the first "problem" wasn't really that hard. 
1. <code>apt-get remove mozilla-firefox</code> which also removes the ubuntu-desktop metapackage.
2. Download firefox 1.0 and install it
3. Add a link to the executable to your $PATH

I'm not really that happy that removing firefox also removed ubuntu-desktop but as it's just a metapackage I also see no real problem in this for now.

But then to the 2nd point: As far as I can tell I have 3 options here:
* Removing wxgtk completely and building my own one
* Using the source package and changing some settings of that package
* Updating to the unstable branch of Ubuntu and checking out what's in there ( I don't know if this branch includes a real GTK2 wxGTK2! )

Going with the first option would definitely produce more problems than anything else so ... let's see: Option 2. I downloaded the source using apt-get and checked out the rules file. Bad memories *g* For now I will keep this as a backup solution simply because I don't want to mess around with source packages right after getting a Debian system after about 1.5 years of Gentoo ;-)

Living on the fast lane of Gentoo (~x86 ;-) ) for the last year I'm now giving the fast lane of Ubuntu a try. This lane seems to be even more fast then the Gentoo-pendant because they also include development package of Gnome (et al.) into this branch. The step seems to be a little harsh for just wxgtk but this is a really fresh installation and reinstalling Ubuntu should only take 20 minutes if necessary :-)

Hmm.... now that think about it: There are some ways to mix branches in apt-get (also with priorities) so at least the Firefox issue could probably also be resolved by using a Firefox package from Hoary in Warty :-)