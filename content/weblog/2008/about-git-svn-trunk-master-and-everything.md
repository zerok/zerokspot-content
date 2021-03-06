---
date: '2008-08-08T12:00:00-00:00'
language: en
tags:
- git
- git-svn
title: About git-svn, trunk, master and everything
---


I guess this can be categorized as "Shooting yourself into the foot with a smile". For the last couple of days I've been working on a patch for bringing some [gettext](http://docs.python.org/lib/module-gettext.html) into [Sphinx](http://sphinx.pocoo.org/). First I started the work in the most recent release, just to see how hard it would be. Then I made a checkout from the SVN repo using [Git](http://git.or.cz/) in order to rewrite the patch for trunk. Everything fine, so far. Yesterday night then I wanted to go through the patch one more time before sending it through the ether and noticed something rather stupid: When checking out from SVN using git-svn, Git, for some reason, didn't make "trunk" the new master-branch, but instead, I guess, used the branch with the most recent commit.

When then moving to the "real" trunk and applying my patch, at least most of the hunks (and all the big ones) passed, so it only took me about 20 minutes to get the patch ready again.

So from now on, the first thing I will do when cloning an SVN repository is a `git reset --hard remotes/trunk`  as suggested by the [manual](http://www.kernel.org/pub/software/scm/git/docs/git-svn.html) (deep down in the basic examples) just to be on the safe side ;-)
