---
date: '2005-10-04T12:00:00-00:00'
language: en
tags:
title: zk.bloglines.addFix.user.js
---


Since I'm not the only one having <a href="http://my.opera.com/community/forums/topic.dml?id=94922">a small problem</a> with the subscription form of bloglines using Opera 8.5 I've started playing around with various elements of this form to find the element which is messing with Opera's rendering engine. After a few days I've finally found the problem and fixed it using following UserJavaScript:

-------------------------------



<pre class="code">

// ==UserScript==

// @name Bloglines subscription bookmarklet folder fix

// @include http://www.bloglines.com/sub/*

// @namespace http://zerokspot.com

// @description Fixes a small problem with the bloglines bookmarklet/

//              subscription form that prevents people from selecting

//              the folder where the feed should be stored.

// ==/UserScript==

/*

 * This script is granted to the Public Domain

 */

(function(){

	var select = document.forms["subform"].folder;

	select.options[0].parentNode.removeChild(select.options[0]);

	var newDefault = document.createElement("option");

	newDefault.appendChild(document.createTextNode("_TOP_"));

	newDefault.setAttribute("value",0);

	select.insertBefore(newDefault,select.options[0]);

})();</pre>



I can't guarantee that it really works, but at least it does so for me :)