---
date: '2007-04-24T12:00:00-00:00'
language: en
tags:
- django
- middleware
- openid
- session
title: OpenID Session Middleware for Django
---


Simon Willison just did something that will probably save me a week or so on my way to django@zerokspot: He released an [OpenID consumer app for django](http://simonwillison.net/2007/Apr/24/openidconsumer/).

From what I've read so far, this one doesn't (yet!) integrate itself into Django's auth subsystem, but provides a session middleware.


-------------------------------


I guess now I can rewrite the commenting app again, this time to support OpenID and remove the captcha for those people, who are logged in. My implementation probably won't really use Simon's apart from being an inspiration, though, since I would like to really bind a username to an OpenID in order to make name-binding possible in comments. Now I just need the time for the whole project again. Perhaps this Saturday I will have a chance. Or Friday afternoon ;-)

As with many other Django related projects, also this one can be found on [Google Code](http://code.google.com/p/django-openid/)... Sourceforge is really starting to die, isn't it?