---
date: '2010-01-06T12:00:00-00:00'
language: en
tags:
- django
- messaging
title: Django 1.2a1 has a message for you
---


Today, Django 1.2 alpha1 was released and it comes with a ton of new features compared to the 1.1 branch. You can read up about them in the&nbsp;<a href="http://docs.djangoproject.com/en/dev/releases/1.2-alpha-1/">release notes</a>&nbsp;but I want to take a quick look at probably my favorite addition: The new <a href="http://docs.djangoproject.com/en/dev/ref/contrib/messages/">messaging framework</a>. Why this and not the support for multiple databases? Well, I see myself using messaging in many more places than multiple databases ;-)

--------

Previously the <a href="http://docs.djangoproject.com/en/1.1/topics/auth/#messages">messaging system (&lt;= 1.1)</a> was was part of&nbsp;<em>django.contrib.auth</em> which is sometimes not what you want if all you need are messages. This has been deprecated now and replaced with a system that is rather similar to logging systems. If you want to display a message to the user you specify the message as well as a message level, which is internally represented as nothing more than an integer, so it can be filtered out depending on the system configuration by a simple greater-than comparison.

<pre class="code">from django.contrib import messages

messages.add_message(request, message.INFO, u&quot;This is a message with the level INFO=25&quot;)
messages.info(request, u&quot;Shortcut for the above&quot;) </pre>

How these messages are stored until they can be presented to the user is completely configurable. You can go with the session object or plain old cookies and with a combination of those thanks to the <a href="http://docs.djangoproject.com/en/dev/ref/contrib/messages/#storage-backends">MESSAGE_STORAGE</a> setting.

The messages then get through a <a href="http://docs.djangoproject.com/en/dev/ref/contrib/messages/#enabling-messages">middleware and a context processor</a> into your template, where they can be displayed using the &quot;messages&quot; variable (so be sure to not use this variable within your code or create your own little context processor that uses a different variable name). Each message comes with a set of tags (<code>{{ message.tags }}</code>), that can then be used for CSS classes et al.

As with the message store, also the <a href="http://docs.djangoproject.com/en/dev/ref/contrib/messages/#message-levels">levels</a> and their <a href="http://docs.djangoproject.com/en/dev/ref/contrib/messages/#message-tags">tags</a> can be easily configured depending on your needs.

A big &quot;Thank you&quot; to <a href="http://www.caktusgroup.com/blog/">Tobias McNulty</a> and everyone else involved in this package. I&#39;m really excited to use this in future projects. So far I only skimmed through the docs and the source code ... and it is hard to resist switching some of my projects over to Django 1.2 because of all the great stuff that is in there :-)
