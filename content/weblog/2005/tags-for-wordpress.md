---
date: '2005-04-30T12:00:00-00:00'
language: en
tags:
- development
- zerokspot
title: Tags for Wordpress
---


During the last couple of days I thought about some ways how I could realize some basic tagging support for WordPress and I think I will go with following solution: Ajax in the admin panel ;-)

-------------------------------



The basic idea is this: I add a JavaScript generated form as a replacement for the category-listing in the post form. In this form the author enters the tags (they are in fact categories but not for me anymore after finishing this *g*). Then when hitting the submit button, a form checker script will check if there are tags listed in the text field, that are not yet represented in the categories table. If this is the case, the script asks the user if these new tags should be created and then creates them through a XMLHttpRequest call. This would require no database changes and only minimal changes in <em>one</em> wp-admin file.



I'm not yet sure, if this will really work, but I will definitively give it a try today :-)