---
date: '2005-12-15T12:00:00-00:00'
language: en
tags:
- macosx
- software
title: Erasing AudioCDRWs on MacOSX 10.4.3
---


For some reason I wasn't able to erase an AudioCDRW I burned a few days ago using the Disk Utilities bundled with MacOSX. All I got were greyed out buttons :-? Since I'm still not really mac-savvy yet, it took me quite some time to find a commandline utility which I hoped would still be able to erase that disc. First I thought about going with cdrecord since I already know this tool. But it's still marked as unstable in fink for Tiger, so I kept looking ;)

-------------------------------



After five minutes more of googling I found the man page of drutil which is included in the base system. It's basically a commandline tool for interacting with the CD/DVD device and can be used for extracting for example the cdinfo or similiar things but is also able to erase CD/DVD*RWs



<pre class="command">drutil erase full</pre>



Quite cool ;)