---
date: '2005-04-18T12:00:00-00:00'
language: en
tags:
- development
- zerokspot
title: Opera and image redirections?
---


Is it just me or has Opera8 beta3 really a small problem with redirected images? I mean with URLs that redirect to any image. I'm using a normal HTTP redirect and Mozilla Firefox seems to handle it quite well. An example for it would be the "Latest book review" image in the sidebar on the main site. Another one would be <a href="http://www.gamerslog.com/node/view/66/">my <cite>Prince of Persia: The Sands of Time</cite> review on gamerslog.com</a>.

-------------------------------



My redirect should send a 301 status code ("Moved permanently"), but I've also tried it without the custom code which should then make it a 302.



Perhaps I will find a different solution, but some kind of redirect is necessary, since the gallery for the uploaded images here is not yet on it's final destination. The "final destination" just currently redirects everything to the temporary location of the gallery.



Has perhaps anyone an idea, what could make this work for Gecko AND Opera?