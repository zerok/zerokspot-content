---
date: '2010-05-26T12:00:00-00:00'
language: en
tags:
- djangocon
- django
- i18n
- projectalexandria
title: International documentation for Django - An idea
---


International documentation has always been on of those ponies for Django that
might help bring in more people but never got fed enough oat to develop all
the magical powers. So far most international usergroups (on a national or
language level) have some sort of translation for the documentation. Sadly all
these efforts never really produced a *complete* translation. So perhaps we need
a new approach.

-------------

At least from my experience with the [German translation
project](http://github.com/zerok/django-docs-de) there are a couple of reasons
why no complete doc-translation is out there:

* The documentation is huge
* Many articles in the documentation have reference-character, meaning that they are 1000+ lines long
* It is still a moving target. A slowly one with a relative fixed release-cycle but still.
* Many people realize only after saying that they want to help that translating a document like contributing.txt can take up to a whole week and they "disappear" ;-)

Because of at least the first two reasons, translating the whole documentation
is impractical if only volunteers are involved and probably unaffordable if
professionals should do it. From what I've seen so far, there is *no* complete
translation of the official documentation out there.
[Djangobrasil](http://docs.djangobrasil.org/) is the most complete I could
find and even they are stuck with 1.0. There has to be some way to strike a
balance by providing translations for only core documents like the
installation guide, the tutorials and the internals/contributing.txt file as
well as offering something different.

The "different" here IMO should be some collection of short articles focusing
on one rather specific topic. Think about "A List Apart meets Django Advent"
but with a bit more structure in order to make it easier for people to find
these articles or articles on specific topics.

Meet *Project Alexandria* !

On the technical side, the goal of this project is to provide first of all a
common framework for such an article-based knowledge base. Unless someone
comes up with a better idea, I currently have following things in mind for the
German part of Project Alexandria:

* Blog-style with some integration to an internal search engine in order to provide a listing of related articles
* RST as documentation format
* DVCS based with one or two gatekeepers who are responsible for pulling new or updated articles from contributors
* No exclusive to core-Django documentation but also open to highly application specific topics
* Provide a Django app that implements a base-framework for all of this

On the social side of things the project should eventually act as an umbrella
for language- or even country-specific projects to combine their efforts and
come up with a common framework for all of us where we can easily share
articles.

Well, these are only some first ideas. I'm probably going to give a short
lightning talk this afternoon at [Djangocon.eu](http://djangocon.eu) about the
the motivation for this. If you are interested in helping out with the
base-project or with the German subproject or if you think that this idea is
totally stupid, please tell me :-)

P.S.: I totally forgot to mention this, but with localized documentation I
don't necessarily mean translations. There are also some things that are only
relevant for local communities with regards to customs and how stuff is just
done around there and how you can adopt that with Django. The classic
documentation is not a place for something like that but Project Alexandria
should be.

So to clarify my proposal a bit more: Let's take Django Advent and make it an
openblog-like platform for specific topics but also for documents like really
basic tutorials and build a format that can be used by local communities to
provide relevant information on top of that for their users.
