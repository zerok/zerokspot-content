---
date: '2006-05-26T12:00:00-00:00'
language: en
tags:
- del-icio-us
- ma-gnolia
- ruby
title: Ma.gnolia -> del.icio.us
---


ByeBye Ma.gnolia. It was a good start but that's all it was for me. I've now gone back to del.icio.us for my daily social bookmarking. For my taste the interface on Ma.gnolia is simply too cramped with adverticing and meta information which makes the whole site way too slow for normal use. Also the post-form's tag-autocompletion is in my opinion hardly usable.

So now I'm back with the service that started it all and all I'm really missing is the group-feature which I wanted to use but no one wanted to join me ;) OK, for example the ruby-group is nice, but you can do the same thing with tags.

So yesterday I wrote a small script to first of all give me a list of all the bookmarks that I had in my Ma.gnolia repository but that were missing on del.icio.us. I won't post the script here for now because I guess it was quite close to the terms of use for the del.icio.us API. I throttled it down to one call every 2 seconds but anyway :) 

So I will just post some problems I came across that are all basically difference between the export file of del.icio.us and the one of Ma.gnolia.



-------------------------------



1. Del.icio.us stores URLs in a HTML escaped format. So it's &amp;amp; instead of &amp;.
2. Tags are stored differently. Ma.gnolia uses a comma-separated list while del.icio.us separates the tags with spaces.
3. The del.icio.us API is behind HTTPS but I think their certificate is broken or something, so the open-uri library for Ruby was no longer an option. Now you have two options in my opinion: net/http or simply `curl`. I went with the later one. Since the URLs have GET parameters, you have to quote the URL.
4. The del.icio.us API requires a delay of at least 1 second between two queries. I went with 2 seconds. If you have the time, be nice to the guys and gals and wait even longer :)
5. If you want to keep the date when you originally added the link to Ma.gnolia also in del.icio.us, use the dt-parameter with an ISO8601 timestamp for UTC. In Ruby this would be something like this:
	<pre class="code">
		require 'time'
		date=Time.at(timestamp).utc.iso8601
	</pre>
	
I think that's all :) Now all I hope, is that del.icio.us will soon add a group functionality like the one Ma.gnolia has.