---
date: '2010-07-11T12:00:00-00:00'
language: en
tags:
- javascript
- jsdoc
- documentation
- api
title: Javadoc for JavaScript?
---


For personal but also work-related projects I'm currently looking into
different coding styles for JavaScript and especially into how to document the
source files. As someone who knows Java  I was kind of looking for something
similar to javadoc.
[Wikipedia](http://en.wikipedia.org/wiki/Comparison_of_documentation_generators)
provides a great overview of documentation tools, alby not a complete one. For
instance, I couldn't find YUI Doc or pdoc in there.

--------------

So far I've found following tools being mentioned without hatred on the web:

* [JSDoc Toolkit](http://code.google.com/p/jsdoc-toolkit/)
* [YUI Doc](http://developer.yahoo.com/yui/yuidoc/)
* [pdoc](http://pdoc.org/)

One Java- and JavaScript-based, one Python-based and one Ruby-based and all
appear to be mainained. Of these only JSDoc Toolkit seems to be without any
framework-bias. I was a bit surprised to see that YUI Doc was written in
Python given YUI Compressor being a Java project. Anyway, so far JSDoc Toolkit
looks like the solution I will probably look more into thanks to it having no
external dependencies except for Java and not being tied to any particular
framework, but I'm still kind of missing the *one* solution.

Something that also all the big libraries use and you therefor easily can use
to produce a single large API documentation for a whole project, instead of
having to look for documentation in as many places as you have dependencies. I
guess, one can dream :-)

Have I missed a project? What are you using?
