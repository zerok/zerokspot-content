---
date: '2005-04-14T12:00:00-00:00'
language: en
tags:
- c-c--
- development
title: strtok() trap
---


Just something small I noticed while writing a CGI program in C: strtok() makes changes on the passed string. If you do for example this:

-------------------------------



<pre class="code">

next = strtok(mystring,"&");

</pre>



Don't expect, that mystring has the same content as before. It will only have the same content if your string couldn't be tokenized. But if there was another token found in this string, mystring will now have a \0 after the first occurance. 



So it's better to make a copy of the string (for example with strcpy() or strdup()) and let strtok() operate on this string instead of mystring.