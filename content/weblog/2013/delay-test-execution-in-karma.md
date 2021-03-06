---
date: '2013-07-12T20:33:42+02:00'
language: en
tags:
- karma
- javascript
- testing
title: Delay test execution in Karma
---


Over the last days I've started to dig a little bit into [Karma][] for running some of our unittests. Everything worked surprisingly well right from the get-go but I had one small problem: Some of our tests should only be executed once the whole JavaScript app has been successfully initialized, which is relevant - among others - for [ExtJS][] and [SenchaTouch][] applications.

The problem here is, that by default Karma launches the test framework right when all the files have been loaded which might not be the right approach if your app requires some additional setup. But luckily, there is a little hook you can override.

-------------

First, you have to prevent Karma from running your tests at the earliest moment possible. To do this, add this assignment into some kind of setup.js right after your adapter has been loaded:

<pre><code class="language-javascript">window.__karma__.loaded = function() {};</code></pre>

This way, <code>window.\_\_karma\_\_.start</code> is not launched after the last file being loaded, which is also where the Mocha adapter has [its mocha.run()](https://github.com/karma-runner/karma/blob/v0.8.6/adapter/mocha.wrapper#L6) call.

Once the rest of the app has been initialized, I now execute <code>window.\_\_karma\_\_.start();</code> and also the tests depending on a fully initialized app are executed correctly :-)

Normally, this hook is used by adapters precisely for the purpose of starting the test suite when everything is supposed to be ready, so I don't feel bad about "abusing" it for my own definition of "ready" ;-)

Big thanks to Vojta Jína for [describing this on the mailinglist](https://groups.google.com/forum/#!msg/karma-users/XpEEuvCRdGc/K2Nxj1ACSl4J).


[karma]: http://karma-runner.github.io
[extjs]: http://www.sencha.com/products/extjs
[senchatouch]: http://www.sencha.com/products/touch/
