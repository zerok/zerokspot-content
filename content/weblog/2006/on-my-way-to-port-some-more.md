---
date: '2006-11-04T12:00:00-00:00'
language: en
tags:
- 5-0
- drupal
- modules
title: On my way to port some more ...
---


As I've [written on Thursday](http://zerokspot.com/node/787), I've started porting the custom modules I've written for zerokspot.com in order to move ASAP over to Drupal 5.0 as soon as there is an RC out there. My main motivation behind this is the integration of jQuery into Drupal core, which should make some of the hacks I had to perform quite obsolete :) 

To make this transition possible I also had to one "foreign" modules which are still &lt;= 4.7 only in the Drupal repository. I guess when I'm done with all that, quite a few will have also entered this list. Probably the biggest fish in the pond will be the Akismet module which I'll probably not dare to touch :-) I just hope that [Markus](http://www.phpmix.org/) finds some time updating it before RC1 hits or at least before Drupal 5.0 final is released :)

-------------------------------



On the other side the [progress](http://drupal.org/node/82257) people are making to port core modules over to Drupal 5.0 is looking really good. With things like Pathauto and Views already working (at least according to this list) not to mention CCK should make everything quite simple also and esp. when it comes to [Ufdi.net](http://ufdi.net) :-)

There are sill a few of the modules I'm using here missing in that list though, but nothing that should be too hard porting on my own. I'm not sure though, if I should also post the patches for this into the issue tracker, since I'm IMO still too new to Drupal and I'm not sure if I'll have the time to really make the patches nice enough to be ready for the world, but if I have the time, I could also contribute them back ... although by the time I'm done with them, probably 10 other people will already have suggested their own patches ;-)