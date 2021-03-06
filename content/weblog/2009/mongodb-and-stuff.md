---
date: '2009-09-26T12:00:00-00:00'
language: en
tags:
- django
- mongodb
- zerokspot
title: MongoDB & Stuff
---


For quite some time I wanted to write something with a document-based database system again. My last trip into this area was in Feb. 2009 when I needed a small feed-aggregator for a local Barcamp and wanted to provide a live-ticker for all the tweets and photos related to that event.&nbsp;Back then I went with&nbsp;<a href="http://couchdb.apache.org/">CouchDB</a>&nbsp;with its slick web interface and HTTP API, but it is by far not the only attractive system on the block.

-----------

A few weeks ago I suddenly wanted to work with such a system again and remembered something that I wanted to write for my site for ages: A place to keep lists of games and books that I own, so that I can easily link them from blog posts. Since both these categories have only a small subset of properties in common (title, slug, createdAt-timestamp, review) and in the future I maybe want to add even more categories and esp. properties (like a reference to some <a href="http://memory-alpha.org/">Memory Alpha</a>&nbsp;page for example) I wanted to keep the database backend as flexible as possible. So I decided to use a document-centric DB for this section. But which one? It&#39;s not like I will update the database multiple times per second, so CouchDB would be a candidate. Something else that I didn&#39;t really like about CouchDB, though, was the lack of composite unique keys. All you can do here, as far as I know, is to encode the value of each property that should be part of that unique-property into the main key (<i>doc_id</i>) of a document which ends up in duplication of information.

So I went looking around a bit and stumbled upon <a href="http://www.mongodb.org/">MongoDB</a>. MongoDB is similar in many regards to CouchDB but in some takes a slightly different approach. For once it doesn&#39;t have a generic HTTP interface but relies on <a href="http://www.mongodb.org/display/DOCS/Drivers">platform-dependent interfaces</a> (similar to MySQL, PostgreSQL and friends). It also lets you organize documents not only into databases but also collections. Within these collections you can then define <a href="http://www.mongodb.org/display/DOCS/Indexes">custom indices</a> that also allow for composite unique keys (or &quot;compound key indices&quot; as they are called in the documentation). This is especially handy if you want to have a slug be unique only for a certain content type (i.e. book or game).

The integration with Python (and in my case django) is also pretty straight forward. You need to install the <a href="http://pypi.python.org/pypi/pymongo/">pymongo</a> library and that&#39;s mostly it. <a href="http://api.mongodb.org/python/0.16/pymongo.connection.Connection-class.html">pymongo&#39;s connection object</a> already acts like a connection pool so you just need to have such an object as a module-global variable in your <a href="http://www.djangoproject.com">Django</a> application.

<pre class="code">from pymongo.connection import Connection
from django.conf import settings

connection = Connection(settings.MONGODB_HOST, settings.MONGODB_PORT)
database = connection[settings.MONGODB_NAME]

def index(request):
    games = database.collection.find({&#39;type&#39;:&#39;game&#39;})
    # ...</pre>

For those of you who prefer seeing a complete example, take a look at Mike Dirolf&#39;s <a href="http://github.com/mdirolf/DjanMon">sample django+mongodb project</a>.

There is one small problem, though, with pymongo&#39;s <a href="http://api.mongodb.org/python/0.16/pymongo.cursor.Cursor-class.html">Cursor</a> object: You can&#39;t use it out of the box with django&#39;s Paginator. The one missing method to work there is Cursor.__getitem__ with support for slices. Depending on how you use the paginator, you should be able to monkey-patch it into the Cursor with a combination of <a href="http://api.mongodb.org/python/0.16/pymongo.cursor.Cursor-class.html#skip">Cursor.skip</a> and <a href="http://api.mongodb.org/python/0.16/pymongo.cursor.Cursor-class.html#limit">Cursor.limit</a>, though.

The rest was easy: Writing some small forms for adding and modifying books and games and presenting them with some images. While writing the presentation code, I noticed another small issue: Sorting over multiple fields with one of them being optional doesn&#39;t really work all that well in version 1.0. Something like that is interesting if you&#39;re working with books that are optionally part of a series and you want them to be listed with the series first and then the main title. Luckily this seems to have already been fixed in the <a href="http://jira.mongodb.org/browse/SERVER-282">development version</a>.

In general, I really enjoyed this small trip into the field of document-based databases again. They are definitely not the solution for every storage problem out there, but for small stuff and esp. problems where the data-structure has to be highly flexible I can definitely see myself using them more and more :D
