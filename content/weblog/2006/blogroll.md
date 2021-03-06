---
date: '2006-01-10T12:00:00-00:00'
language: en
tags:
- development
- zerokspot
title: Blogroll
---


As <a href="http://weblog.zerokspot.com/posts/508/">promised quite some time ago</a> I've now finally uploaded a <strong><a href="http://www.zerokspot.com/blogroll/">HTML representation of my blogroll</a></strong> as well as an OPML file which also includes the actual feed addresses. I'm still working on the automation of the generation of this but I'm currently having some small problems with the AppleScripting for [Vienna](http://opencommunity.co.uk/vienna2.html).

-------------------------------

 I have following little script to export my feedlist:

<pre class="code">
	osascript -e "tell application \"Vienna\" to export subscriptions to \"${PWD}/feeds.opml\"" \
</pre>

... quite tight. So for those of you (like myself) who prefer the "Better 1 line too much but then readable" approach to coding:

<pre class="code">
	tell application "Vienna"
		export subscriptions to "feeds.opml"
	end tell
</pre>

(Add the $PWD path in your mind, please ;) ) 

Now my problems with this:

1. I somehow don't get the same feedlist I'd get when exporting using "File -&gt; Export Subscriptions". Somehow everything is just flying around :-?
2. The script triggers a dialog box no matter if it succeeds or not. Quite disturbing IMO if you plan to cron that thing ;)

For now, please don't ask me why. Somehow I can't make an anonymous checkout from [Sourceforge](http://sourceforge.net) for at least 2 days now, but the moment I get a snapshot I will try to find my way through all this ObjC code. I'm not really good in ObjC ... actually I have seen it for the first time in my life just a few weeks ago and haven't done anything with it yet. But perhaps I still find a solution for this.

In the meantime: If anyone has some suggestions about this or could recommend another decent news reader (free) that also allows exporting the feedlist as OPML (no matter how), I'm all ear :)