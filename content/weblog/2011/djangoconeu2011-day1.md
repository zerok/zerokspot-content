---
date: '2011-06-07T12:00:00-00:00'
language: en
tags:
- django
- djangocon
- djangoconeu2011
- travel
title: 'Djangocon.eu 2011: Day 1'
---


DjangoCons are awesome, and this year's DjangoCon Europe 2011 is about to fit
into this category quite nicely :-) Monday morning started with an ...

-------------------------

## Opening Talk

<figure>
<img src="http://photos.h10n.me/Conferences/DjangoCon-Europe-2011/i-PkdWrg8/2/M/DSC0085-M.jpg" alt="" />
</figure>

... by Remco Wendt about some usual organisational topics as well as showing
how far DjangoCon's have come over the years. A big "Thank you" to all the
volunteers and sponsors who made and are still making this event possible. And
also to the one who got the idea of having a dedicated open-space area
throughout the whole conference.

## From static to real-time

<figure>
<img src="http://photos.h10n.me/Conferences/DjangoCon-Europe-2011/i-QBtbSrn/1/M/DSC0087-M.jpg" alt="" />
</figure>

Next was Eric Florenzano with his keynotes describing the evolution of a
simple idea through multiple iterations to a highly complex web application.
Something that is facilitated by Django's rather flexible architecture. He
also highlighted that techniques such as WebSockets may be the new best thing
since sliced bread, they are often overkill. For many problems even
short-polling is still the right solution.

Also, a shout went out to help port Socket.IO to other platforms or improve
the already existing ports. Socket.IO is a JavaScript abstraction layer of
WebSockets and similiar techniques implemented in Flash et al. originally
written in JavaScript for the server as well as the client side.

## Large Django sites at Mozilla

<figure>
<img src="http://photos.h10n.me/Conferences/DjangoCon-Europe-2011/i-z2XDFBw/0/M/DSC0119-M.jpg" alt="" />
</figure>

After a short break it was Andy McCay's turn to tell us all about the cool
sites and tools Mozilla is been working on of the last couple of years in
their effort to port their frontend infrastructure (like Mozilla Add-Ons) from
PHP (through CakePHP) over to Django. Having a site that generated more than
150 million requests per months must be an awesome test-bed ;-) Since writing
about all the tools mentioned by Andy would probably fill multiple posts on
their own, I'll just list them:

* [cache-machine caches](https://github.com/jbalogh/django-cache-machine) Django ORM objects and querysets
* [queryset-transform](https://github.com/simonw/django-queryset-transform) (by Simon Wilison)
* [timing-ware](https://github.com/jbalogh/zamboni/blob/master/apps/amo/middleware.py#L162)
* [bleach](https://github.com/jsocol/bleach) strips unwanted HTML out of user input
* [flake8](http://pypi.python.org/pypi/flake8/) checks Python source code for PEP8 violations
* [django-qunit](http://pypi.python.org/pypi/django-qunit/) integrated qunit into Django's testing infrastructure.
* [arecibo](http://www.areciboapp.com/) collects e-mails generated by multiple nodes in order to annoy the
  admin only once per error.
* [mozilla-playdoh](https://github.com/mozilla/playdoh) is a base-template for
  Django applications.
* [django-secure](https://github.com/carljm/django-secure) provides a
  middleware and helpers for checking your django project against a set of
  security guidelines.

## How I Learned to Stop Worrying and Love Python Packaging

<figure>
<img src="http://photos.h10n.me/Conferences/DjangoCon-Europe-2011/i-p8SMMFf/0/M/DSC0124-M.jpg" alt="" />
</figure>

After this whole bunch of new toys to play with, Jannis Leidl gave an overview
and history lesson of the parts of Python that make adding them to your
project possible: Python's packaging system. But he went beyond what is in the
standard lib and also described tools like [pip](http://pypi.python.org/pypi/pip) and [virtualenv](http://pypi.python.org/pypi/virtualenv) as well as gave some tips for how to better use them?

But this whole infrastucture is in flux right now with
[distutils2](https://bitbucket.org/tarek/distutils2) being on the horizon with
its move away from the rather annoying setup.py and on to a setup.cfg file.
That bundled with a standardized API for resolving dependencies and a backport
to Python 2.x makes it something I'm really looking forward to :-)

## Django and PyPy

<figure>
<img src="http://photos.h10n.me/Conferences/DjangoCon-Europe-2011/i-f7cxjTn/0/M/DSC0180-M.jpg" alt="" />
</figure>

The successor to Psyco, PyPy, is now at the level of CPYthon 2.7.1 and native
Python applications should by now work out of the box. The problem is still
the integration of c-extensions, but there are a couple of solutions out
there. Alax Gaynor gave a nice overview of them and described the state of
database driver support with PyPy with psycopg2 probably still being the best
option right now.

And I won't go into benchmarks here ;-)

## 3 CMS in 45 minutes

The next session was split into 3 parts where 3 CMS systems were presented by
their respective contributors:

* [django-fiber](https://github.com/ridethepony/django-fiber) presented by Dennis Bunskoek 
* [django-cms](https://www.django-cms.org/) presented by Jonas Obrist
* [merengue](http://www.merengueproject.org/) presented by Manuel Saelices

## One size fits all

<figure>
<img src="http://photos.h10n.me/Conferences/DjangoCon-Europe-2011/i-x4dKFkX/0/M/DSC0200-M.jpg" alt="" />
</figure>

After a break it was time for something completely different: A frontend talk
by Idan Gazit about using Compass/Sass in combination with other frameworks
like Less (not the Sass competitor but the grid framework ;-)). But the main
focus was on responsive webdesign and how to get a project in that direction
with the mentioned tools.

## Lightning talks

A set of lightning talks marked the end of day 1:

* Tom Christie introduced his take on a [REST framework for Django](http://django-rest-framework.org/).
* GSoC student Gregor Müllegger presented the current state of his Django
  forms refactoring as part of the summer of code.
* Nate Aune gave a short introduction into
  [DjangoZoom](http://djangozoom.com/)
* Roald de Vries presented a simple form creator written in Django.
* [JetBrains](http://www.jetbrains.com) gave out personal licenses of [PyCharm](http://www.jetbrains.com/pycharm/) and Dmitry Jemerov gave a
  small feature tour.
* Last but not least Peter Kese was looking for people wanting to help with
  running multiple sites using a single Django instance and improving the
  already existing support available for that in Django.

