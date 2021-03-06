---
date: '2008-08-24T12:00:00-00:00'
language: en
tags:
- django
- documentation
- sphinx
title: Django Documentation Refactor has landed
---


A long time coming the documentation refactoring has [finally landed](http://code.djangoproject.com/changeset/8506) on [docs.djangoproject.com](http://docs.djangoproject.com/) and in trunk. Compared to the previous format this one now uses Georg Brandl's [Sphinx](http://sphinx.pocoo.org/) documentation system, which also finally gives you as a user a simple way to have the documentation bundled with Django itself to be converted to HTML with one command (if everything works) :-). To do this, simply go to the docs folder of your Django checkout and run `make html` (naturally after installing Sphinx) and if everything works you should end up with a whole bunch of HTML files. Great work [Jacob](http://www.jacobian.org/) et al. :-)

<s>So far for me this isn't working yet, probably because I'm living on the bleeding edge with Sphinx for a couple of reasons, but perhaps I can find some time to dig into this.</s> If you get a syntax error line layout.html line 10 when building the documentation, make sure that you've installed Jinja 1.2 (instead of something older like 1.1). 

The structure is new too and it will definitely take some time getting used too, but the separation between primary content (the one on the left side of the index page) and additional articles (in the boxes on the right side) definitely makes sense :-) For now it isn't linked yet in the main menu on the [site](http://www.djangoproject.com), probably because it's still in testing stage.
