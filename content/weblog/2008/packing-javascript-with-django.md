---
date: '2008-01-17T12:00:00-00:00'
language: en
tags:
- django
- javascript
title: Packing JavaScript with Django
---


Since I started messing around with JavaScript a little more for work and also finally used this opportunity for getting some knowledge when it comes to YUI, I was about to look into ways to efficiently deploy separated JavaScript files (or at least I planned to get into this topic this weekend). I guess for Django, there is now - thanks to Pedro Vale Lima -  a really nice [management command](http://pedro.valelima.com/blog/2008/jan/17/deploying-compacted-javascript-django/) that makes packaging and compressing multiple JavaScripts easy.

Naturally you could also write an external script using one of the many packers that are already out there, but with the integration right into Django it's just much cleaner :-)