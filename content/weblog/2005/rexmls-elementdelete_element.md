---
date: '2005-11-08T12:00:00-00:00'
language: en
tags:
title: REXML's ElementNo.delete_element
---


Normallly you can use this method for deleting specific elements of a DOM-tree by passing their XPath or Element instance to this method. Since I thought: It's in the tree so I will simply delete it based from the respective Document#root ... not good ;) This will result in quite strange behaviour which alters for example the nodename of the parent element of the element to be removed. In my case I worked on a OPML file with following section:

-------------------------------



<pre class="code">&lt;outline ...&gt;

   &lt;outline id="someElement" ..../&gt;

&lt;/outline&gt;</pre>



Removing the element now using <code>document.root.delete_element(e)</code> caused absolutely nothing. So I had to try something else:



<pre class="code">document.root.delete_element(e.xpath)</pre>



This actually removed the element from the try but resulted in the parent element to be renamed to "outline[5]" where 5 was probably for being the 5th outline of the parent element ;)



Stupid, but if I had read the manualy before playing around with it I would have seen following statement:

<blockquote>This means that any parent can remove any descendant.<cite><a href="http://www.germane-software.com/software/XML/rexml/doc/classes/REXML/Element.html">REXML Documentation</a></cite></blockquote>



While this IMO also applies for the Document#root element it should at least given me a hint ;) 



So what did I learn from this? 

<code>e.remove</code> is short and good ;)

