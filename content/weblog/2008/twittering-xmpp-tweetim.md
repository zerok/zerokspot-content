---
date: '2008-12-17T12:00:00-00:00'
language: en
tags:
- jabber
- processone
- tweetim
- twitter
- xmpp
title: Twittering via XMPP via tweet.IM
---


<img src="http://zerokspot.com/uploads/tweetim.png" alt="" class="left" />During the last couple of days I started experimenting a little bit with XMPP (again) and also installed my own little ejabberd server. Right when I was finished with that, another project by [ProcessOne](http://www.process-one.net/) popped up on my radar: [Tweet.IM](http://tweet.im/). Remember when Twitter still had that nice XMPP-gateway where you could post using your Jabber client? Yeah, that one ... a long time ago. [Since November](http://www.process-one.net/en/blogs/article/tweetim_a_twitter_xmpp_gateway_service/) there is with Tweet.IM an alternative for that non-existant service which actually works pretty well. 

-------------------------------

You can post messages, send direct messages (although to all the people I want to talk to in private, I'd simply use their IM-account, to be honest) and receive the latest posts in an interval of, I guess, 5 minutes. 

So far I've noticed just one problem with it: If you're using Adium, the IM client tries a little bit too hard to interpret some of the data tweet.im sends it. The gateway sends some avatar information as XHTML-IM, to show you also the avatar of the user you received a tweet from. Adium doesn't really support XHTML-IM to that extent and shows you a broken image with each update. Not a big problem, but it becomes really annoying when it forwards that data to Growl ;-) There is currently [a discussion](http://www.process-one.net/en/forum/viewthread/151/) going on about this, so if you have this problem on the support forums. So, I guess, the more people thinking it's worth changing, the more likely something will change there. 

But otherwise, big thanks to ProcessOne for this service. 