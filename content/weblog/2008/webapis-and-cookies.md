---
date: '2008-03-18T12:00:00-00:00'
language: en
tags:
- api
- cookies
- http
- webapi
title: Web APIs and Cookies?
---


Is there perhaps some space left between complex authentication systems like [OAuth](http://oauth.net/) and the bare minimum like HTTP Basic Auth? Basically everyone uses cookies for normal user connections, why not also support them for application developers?

-------------------------------

First of all, please let me state, that is is only views at things from a standalone-/desktop-application view. For real mash-ups or service integrations like for instance that you can print your photos on [Flickr](http://flickr.com) using [Moo.com](http://moo.com/)'s service, going the cookie-way isn't a good idea, since its potential overhead in those situations (having to do cookie management instead of just normal query string management) and also potentially [opens some phishing doors](http://blog.monstuff.com/archives/000296.html). But when you take "normal" applications into consideration, it makes sense. This post is also very focused on services like [Twitter](http://twitter.com) and [Pownce](http://pownce.com) given my current attempts to integrate some apps with either of them.

Normally you have a couple of ways to communicate on an application level with the various APIs for websites out there:

1.  Using HTTP Basic authentication, which adds the user's username and password into every request
2.  Negotiating a session token with the service which normally involves an "application password" and an application key and then add the user's username and password indirectly by redirecting the user to a 3rd party login page in order to give the user a more transparent view on what's going on in the background.

The 2nd option is esp. interesting when you write a web mash-up, where you definitely want to give the user the extra benefit of redirecting her to a page that she already trusts (since she has an account there) in order to authenticate herself there and then just get redirected back to that new mash-up. 

The downside here is, that it's quite complex. You need quite a few subsequent connections and redirects that are also visible to the user. This has the extra benefit of giving the service provider some additional information about what application a user is using to communicate with their service and also offers the end-user more transparency. 

Let's ignore the first option for now given its obvious disadvantage with the username and password included in all requests, but is there perhaps a middle ground? 

More or less all sites, that offer an API in one way or another also already have some other way to authenticate normal web users through a login form. Through this you create a session cookie that is then used to make sure that you've authenticated yourself with this service for subsequent requests. Using this for APIs would be quite logical and many people already do it as indicated by this sentence in the [Twitter API docs](http://groups.google.com/group/twitter-development-talk/web/api-documentation):

> Session cookies and parameter-based login are known to work but are not officially supported.

and this [paragraph on the unofficial Twitter wiki](http://twitter.pbwiki.com/API+Docs#LoginAuthorization):

> If you prefer, you can use cookie management to login via the website auth. Send a POST request to http://twitter.com/login with the 'username\_or\_email' and 'password' parameters set as appropriate. The headers of the response will include a '\_twitter\_session' cookie. Present this on all future requests.

So why not just offer an official cookie based auth system? As can be seen on services like Twitter and Pownce, the overhead for the service provider should be minimal (with just just changing the login URL and the response headers on a successful login every few seconds and at least telling everyone when a change happens there) and for application developers it would offer a simple alternative for session handling compared to complex token negotiations like OAuth. 

Please don't get me wrong, though: OAuth is a good idea and its great that there finally exists such a movement to unify authentication between login systems (not to say that signing requests (or at least their parameters) is good), but if you offer something like HTTP Basic authentication for your webservice, you already give people a way to easily connect on an application level without having complete control over what application is used there (since no signing), yet with the disadvantage of including the auth tokens (in this case username and password) directly in each request. *Officially* supporting cookie based session handling wouldn't produce an overhead on the service end yet make keeping authed without having to add the username and password with each an every request.

So I guess the reason for not officially supporting cookie based session handling is primarily a way to discourage it given its problems when it comes to cross-domain communication. Another reason for it could be that the service providers have the problem that they have first of all to trust user credentials, but also only a subset of application and those potentially even in different levels (like Flickr for instance stores what an application is allowed to do with users' data). At least the latter point is quite week, though, considering that Pownce *and* Twitter offer HTTP Basic authentication.