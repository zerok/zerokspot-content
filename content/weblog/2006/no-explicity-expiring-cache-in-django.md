---
date: '2006-08-28T12:00:00-00:00'
language: en
tags:
- caching
- development
- django
- python
- rubyonrails
title: No explicity expiring cache in Django?
---


For the last couple of days I've been playing around a little bit with Django, a RubyOnRails like webframework written in Python. Sure, you can compare it to RoR if you want, but in my opinion it has enough unique features to set it apart from the Ruby framework. One thing I'm so far missing though from Django is some of the caching functionality from RoR. 

Django offers 3 types of caching (from what I've learned so far):

1. Caching of a whole site
2. Caching of single views
3. Caching of single variables

All nice and good, and I really like how this seems to be done, but I'm somehow missing a way, to explicitly expire a cached view. Think about following scenario: You have a weblog with posts on it, your frontpage holds a list of the latest posts in a teaser-like view which also includes the number of comments made to each post. This count at the end of the day isn't really all that important, so you simply cache the whole page for let's say 10 minutes. 

-------------------------------



On the other end, you have a view for each single post, where you also list all the comments made to that post. When now a guest comments on that post, you want to invalidate the cache for the view showing this single post so that all the other users can see the comment right away without having to wait for our 10 minutes to pass.

In RoR you have the option, to tell from within an action, to expire the cache for a different action (for those of you who haven't played with Django: a controller's action in MVC/RoR is a view in Django) using something like this:

<pre class="code">expire_action(:controller=&gt;&apos;posts&apos;,:action=&gt;&apos;view&apos;,:id=&gt;123)</pre>

You would normally call this, when the post's author has changed something on post #123 or when someone has commented on it.

Now I'm just wondering, if someone in the Django community has already come up with a similiar method for the view-caching in the Python framework. It's not really a big issue, since you could always use the variable-based caching but it would definitely be a very handy feature and [I don't seem to be the only one](http://groups.google.com/group/django-users/browse_thread/thread/8a6acb40963b3611/75a020f6fbb3c598?tvc=2&q=expire#75a020f6fbb3c598) searching for something like this :-(