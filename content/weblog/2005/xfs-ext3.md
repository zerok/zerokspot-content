---
date: '2005-01-10T12:00:00-00:00'
language: en
tags:
- ubuntulinux
title: XFS -> Ext3?
---


First I wanted to give XFS a try simply because I've so far only heard good things about it. Fast, stable, many tools. Well, right on the 2nd day the system froze and I had to xfs_repair the whole root partition. Bad luck, I thought.

-------------------------------



Last night I had another freeze while unmounting a network partition. To be on the safe side I started xfs_repair again. After starting Firefox this morning I noticed that it didn't liked the Mouse Gestures extention anymore all of a sudden. Ok.... then Firefox started and I had no bookmarks anymore. Not that a tragedy.

Then I started Thunderbird and was confronted with the "Create a new account" dialog... Ey.... 5000 mails can't be all gone :-( The mbox files were still in the ~/.thunderbird directory but somehow the whole thing seemed to be corrupted. 

Everything was working fine until that crash so all I can think of is, that the crash of xfs_repair corrupted the .mozilla and .thunderbird folders. Tonight (or this afternoon, depending on my motivation) I will move to Ext3 or ReiserFS3. I had good but also bad experience with both of them so I will probably throw a coin *g*