---
date: '2006-04-21T12:00:00-00:00'
language: en
tags:
title: htmltoc.rb
---


I'm currently writing on a small paper for a course about web accessibility by [Giorgio Brajnik](http://www.dimi.uniud.it/giorgio/) of the University of Udine and the requirements for the paper are XHTML. Since every paper in my opinion has to have a table of contents, this also counts for the XHTML paper ;) But updating the TOC everytime you make a change manually?! Nah. It's one of these times when I happy that I can code ;)



-------------------------------



The resulting script can be found here. It's a small Ruby program parsing the XHTML document for h[2-6] tags and building a small toc out of them. Just to give a small example:

<pre class="code">
&lt;html&gt;
&lt;body&gt;
&lt;h1&gt;My little article&lt;/h1&gt;
&lt;!-- &lt;%= toc %&gt; --&gt;
&lt;h2&gt;1&lt;/h2&gt;
&lt;h2&gt;2&lt;/h2&gt;
&lt;h3&gt;2.1&lt;/h3&gt;
&lt;/body&gt;
&lt;/html&gt;
</pre>

After calling `htmltoc.rb -i input.html -o output.html` the <%= toc %> call will have been replaced with following listing:

<pre class="code">
&lt;ul id=&quot;_maintoc&quot;&gt;
	&lt;li&gt;1&lt;/li&gt;
	&lt;li&gt;2
		&lt;ul&gt;
			&lt;li&gt;2.1&lt;/li&gt;
		&lt;/ul&gt;
	&lt;/li&gt;
&lt;/ul&gt;
</pre>

(In fact the tabbing was just added by me manually to make it better readable ;) )

If you give each header and id, the TOC entry will be a link pointing to that id in the document.

Another neat feature is, that you can tell the parser to skip certain chapters (incl. subchapters) by simply assigning the "notoc" class to the respective header element.

I hope you like it :) 

<a href="http://zerokspot.com/uploads/htmltoc.rb" class="download" title="Download htmltoc.rb">Download</a>
