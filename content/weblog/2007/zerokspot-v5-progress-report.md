---
date: '2007-07-17T12:00:00-00:00'
language: en
tags:
- django
- lifestream
- pipes
- yahoo
- zerokspot
title: zerokspot v5 progress report
---


Work on the next iteration of this site is progressing quite well :) Thanks to the holidays I can finally invest a little bit more time into the site. In fact I already think about really pushing it out of the door this week. 



-------------------------------



But first I want to get a few things done:

1. Add a **lifestream** comparable to what [Manuela Hoffmann](http://www.manuela-hoffmann.info/stream/) and [Jeff Croft](http://www2.jeffcroft.com/stream/) have on their sites. The implementation of this is nearly complete and uses [Yahoo! Pipes](http://pipes.yahoo.com/) in combination with cPickle and Mark Pilgrim's [feedparser](http://feedparser.org/).
2. **Pinging support:** I think I already wrote that so I just have to activate it (Wouldn't have made much sense to ping technorati with localhost URLs, would it? ;-)).
3. **Recheck all feeds:** I always forget what the ordering of feed-items should be, so I better check all the feeds again before going live.

I think then everything should be more or less ready ... at least I hope so, because I barely can wait to finally post from this new engine.