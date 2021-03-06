---
date: '2006-07-21T12:00:00-00:00'
language: en
tags:
- drupal
- switch
- wordpress
- zerokspot
title: Why the move to Drupal?
---


First of all: I didn't move away from WordPress because I think it is a bad system. In fact IMO WordPress is an excellent tool for people who want to have a simple way to publish posts. But think about following situation: From time to time I want to write a product review, for example about a book. How would you do something like that in WordPress? By using the custom fields. One for the author and one for the isbn as well as one for the rating itself. In Drupal you can do the same thing in a fast and easy and a slow and quite custom way: flexinode or by writing your own module for let's say a bookreview. Such a module could then also offer you a custom form for writing book reviews. Don't get me wrong: This is possible with WordPress, too, as plugins like StructuredBlogging have shown, but IMO it's quite hacky.



-------------------------------



And this is were Drupal in my opinion really shines: The clean API and database structure, something that might not be important for the normal users out there, who just want to have a blog. But it's important for me, as I always want to have my own stuff. 

I also prefer Drupal's appoach of putting as much as possible into cronjobs instead of running them on a request base. Think for example about search indexes etc.

Well, this was all nice and kind of fuzzy since I'm not yet really an expert on Drupal, but at least from my work during the last couple of weeks, the Drupal API simply looks better to me than the WordPress API. This might also be, because WP has a lot of legacy code and structure in it. You see it quite a lot in the database were there are not just a few fields that are no longer used or were their names are not really following only one convention :) These things simply started  to annoy me that much, that I  considered alternatives.

The first one was TextPattern, which also offers one of the main reasons for me, to switch away from WP: Comment Preview. Everytime I went to someone else's blog and wanted to comment, I found that little preview button down there quite handy :) WordPress also has a plugin for this, but ... more about installing plugins and WordPress' plugins in general perhaps later if I'm still motivated ;)

Anyway: TextPattern, while being a great tool too, has some big drawbacks for me: Limited plugin API, no XMLRPC. For a few weeks now there's a XMLRPC implementation [available](http://textpattern.com/weblog/201/xml-rpc-support-for-textpattern-403-is-here) which had a really minor problem with the HTTP content type sent by the server in the beginning which broke support with any client out there that actually expected an XMLRPC server to speak HTTP ;)

More or less at the same time Drupal released their version [4.7.0](http://drupal.org/drupal-4.7.0) which got me interested because of features like Free Tagging and the better RSS integration. I've been using Drupal before for [gamerslog.com](http://www.gamerslog.com) were I was quite pleased with the system but I was simply missing one little thing: It was not possible to store drafts on a post to edit and publish them later. This small requirement died for me a few months ago were I started using TextMate for all my writing and then simply RPC'd it up to the server or copy'n pasted it, so I could finally really consider Drupal once again :)

Well, during the testing I came to like the way everything is done in Drupal and esp. found the API (as mentioned above) and the phptemplate engine very nice. In WordPress everything about the theme is in one or at least a couple of php files. phptemplate uses the same approach but offers something else too. Here you can also theme the modules' output by simply writing themename\_template functions that are used to override theme\_template functions defined by a module to render content. Very useful :)

The move to Drupal also resulted in a completely new URL structure which is not an accident. For the last few years I had multiple sites where I posted my content too. One for gaming, one for books and my main weblog. With Drupal I think I can finally realize one site for all of my content thanks to its flexible module system.

So to sum this probably a little bit too long and definitely too confusing post up: Why Drupal?

* IMO cleaner and more powerful API and cleaner database structure
* phptemplate engine &gt; WordPress' templating
* Put whatever you can into cronjobs
* Different node types
* Taxonomy system is way more flexible than WordPress' categories
* Comment preview
* some other things I've forgotten :)

Don't get me wrong :) Wordpress might be the right tool for you, it still isn't it for me right now :)