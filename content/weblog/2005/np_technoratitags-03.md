---
date: '2005-02-14T12:00:00-00:00'
language: en
tags:
- nucleuscms
title: NP_TechnoratiTags 0.3
---


What's new in 0.3? There is now a switch for transforming "+"s into whitespaces in the tagdisplay. To get this working you also have to make a small change in your "Look of the tags" field:
Previously you probably had something like this:
<pre class="code">
&lt;li&gt;&lt;a href="http://technorati.com/tag/%t" rel="tag"&gt;%t&lt;/a&gt;&lt;/li&gt;
</pre>
Now you replace the %t for the link text with %d :-)

-------------------------------

## Changelog

<pre>
------------------------------------------------------------------------
r4 | zerok | 2005-02-13 17:39:16 +0100 (Sun, 13 Feb 2005) | 2 lines

Now really 0.3 :S

------------------------------------------------------------------------
r3 | zerok | 2005-02-13 17:38:25 +0100 (Sun, 13 Feb 2005) | 2 lines

Version 0.3

------------------------------------------------------------------------
r2 | zerok | 2005-02-12 12:54:07 +0100 (Sat, 12 Feb 2005) | 3 lines

* Added a new option to enable displaying of +s as whitespaces
* The item look has now the placeholder %d for the link text of a tag

------------------------------------------------------------------------
</pre>

<a href="http://www.zerokspot.com/uploads/NP_TechnoratiTags-0.3.tar.gz">Download</a>
