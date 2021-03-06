---
date: '2007-08-17T12:00:00-00:00'
language: en
tags:
- django
- i18n
- symfony
title: i18n URLs in Symfony and Django
---


For the last two days I've been now playing a little bit around with the Symfony framework, which is somehow comparable to what Rails does in the Ruby world, while being written in PHP. I especially like how i18n is handled there ...


-------------------------------

... and immediately start porting at least some ideas over to Django. Nothing major really, but something that I need for a small project I'm currently working on. 

Symfony offers you a way to define a parameter in URLs to act as placeholder for the language (or culture) of the user. For example if you have an international site that should support English and German, you'd have 2 sets of URLs:

1. http://domain.com/en/whatever/
2. http://domain.com/de/whatever/

In Symfony there is now a built-in way to handle this:

    url:   /:sf_culture/:module/:action/*
    
Here :sf_culture acts as this placeholder. But Symfony doesn't only just dispatch requests correctly based on the user's culture, but also automatically rewrites all the URLs currectly if you use the respective helper functions.

For now I just played a little bit around with this concept and ended up with a simple middleware for Django that also takes its own dedicated parameter called "dj_culture" (I had my innovative phase this morning) and loads the respective language pack in the background. It also removes this view kwarg after doing this so that you don't have to rewrite every single view you want to be able to use this middleware with.

For those of you interested in the result of this little experiment, I've put the code onto [djangosnippets](http://www.djangosnippets.org/snippets/371/). The next step will probably be to write an extension of the UrlNode class for the {% url %} template tag in order to also get the 2nd part of the functionality mentioned above.