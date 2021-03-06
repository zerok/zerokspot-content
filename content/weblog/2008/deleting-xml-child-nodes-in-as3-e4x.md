---
date: '2008-03-10T12:00:00-00:00'
language: en
tags:
- actionscript
- actionscript3
- e4x
- xml
title: Deleting XML child nodes in ActionScript3/E4X
---


Yesterday I needed to rid an XML node of all child nodes using ActionScript3's E4X API. Since I'm still completely new to [AS3](http://labs.adobe.com/wiki/index.php/ActionScript_3) (and AS in general) as well as [E4X](http://www.ecma-international.org/publications/standards/Ecma-357.htm) I first thought the solution for this would be close to how you do it is done in DOM by just iterating over all children and executing <code>node.removeChild(child)</code>, but for some reason, doing it this way didn't always work out for me.

-------------------------------

<pre><code>for (var i:int=children.length ; i >= 0; i--){
    delete children[i];
}</code></pre>

Don't ask me why, it sometimes didn't work. But there is also a shorter way to do it, anyway: After taking a closer look at the E4X API in ActionScript3, I noticed the [setChildren](http://livedocs.adobe.com/flex/3/langref/XML.html#setChildren()) method of the XML class. First I tried stuff like <code>setChildren(null)</code> and <code>setChildren("")</code>, but all of that resulted in quite strange results like

<pre>&lt;node&gt;null&lt;/node&gt;</pre>

or 

<pre>&lt;node&gt;&lt;/node&gt;</pre>

At least the 2nd one is understandable since probably "" creates a new empty text-node, but anyway ... Another option is passing an empty [XMLList](http://livedocs.adobe.com/flex/3/langref/XMLList.html) to setChildren(), which finally produced what I wanted:

<pre>&lt;node /&gt;</pre>

And <code>setChildren(new XMLList())</code> is more beautiful than iterating over all children manually, anyway ;-)

In my opinion, the documentation in this area could be a little bit more complete, but I guess this was always a little bit of a problem with XML-related documentation: How to delete stuff is most of the time only mentioned in some obscure footnote ;-)

