---
date: '2007-05-10T12:00:00-00:00'
language: en
tags:
- flickr
title: Popular photos on Flickr
---


On nearly Web2.0 site you see it. It's not the login field, it's not the advertising. It's the popular links where you get to a site that miraculously shows you the most popular item posted or referred to by the users of this particular site. 

Well, one famous exception of this rule is a not completely unknown photo community called [Flickr](http://flickr.com). You might have heard of it. Flickr has something different: They call it "interesting" ... interesting.



-------------------------------



<blockquote><p>There are lots of things that make a photo 'interesting' (or not) in the Flickr. Where the clickthroughs are coming from; who comments on it and when; who marks it as a favorite; its tags and many more things which are constantly changing. Interestingness changes over time, as more and more fantastic photos and stories are added to Flickr.</p><cite><a href="http://flickr.com/explore/interesting/">flickr.com</a></cite></blockquote>

I guess this might get into history books as one of the most fuzziest description ever. So it's not just the number of views or the number of comments and so on, but a combination of all these. . The most funny thing about it is, though, that Yahoo! has [filed a patent](http://www.boingboing.net/2006/11/07/flickr_files_a_paten.html) for this interstingness in Feb 2006. As a small addition to this it would have been nice to also have a ranking of the most-commented-on photo or simply a ranking based on one single criterium...

Anyway: Flickr offers pages for such interesting photos with listings per day, last 7 days and so on. But nothing like: "Most intersting photos ever". At least not in the public web interface. Thanks to David Wilkinson's [Greasemonkey script](http://www.dopiaza.org/flickr/greasemonkey/howinteresting/index.php) for determining how popular a photo is, I found out about the interstingness-(desc|asc) criterium for the [search function](http://flickr.com/services/api/flickr.photos.search.htm) in Flickr's API.

... would have been really nice. Just to bad that if you really want to abuse the search function this way, you get something like this:

<blockquote>Parameterless searches have been disabled. Please use flickr.photos.getRecent instead.</blockquote>

So you have to be at least a little bit more specific. You can for example get some results if you specify the license you want to filter by.

Why do I want this? Well, I was naive enough to select a task for a course that involves analyzing popular photos on Flickr regarding a certain kind of distribution. First of all: I'm not really into the mathematical background of all this and for now just want to mess a little bit around with all this stuff. So I'm not really sure how Flickr's intestringness algorithm influences the distribution of these photos depending on the time they were uploaded.

So I guess, I will settle with analyzing the interestingness-ranking for specific dates and also the overall ranking for specific licenses since this seems to be the broadest ranking I can get out of Flickr for now.

P.S.: I will really try to write more often again. It's currently quite stressful here, but I'll really try :-)