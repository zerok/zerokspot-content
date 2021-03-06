---
date: '2005-06-06T12:00:00-00:00'
language: en
tags:
- development
title: document.images[i].complete?
---


While messing around with JavaScript to build a small frontend for my book collection I came to the point where I wanted a fallback solution if a book doesn't have an image associated with it. Sure, I could simply set a flag in the XML-file including my book collection, but I wanted to solve this with JavaScript. So I found the attribute <em>complete</em> of document.images[i].

-------------------------------



<pre class="code">

for(i = 0 ; i < document.images.length ; i++){

   if(!document.images[i].complete){

      document.images[i].src="noimage.png";

   }

}

</pre>



With this code you normally should be able (according to JS 1.1) to find broken images and replace them with a default image. This also works without a problem in Firefox as can be seen in here:



<img src="http://www.zerokspot.com/uploads/documentimagescompleteff.png" alt="document.images[i].complete in Firefox"/>



Opera8 on the other has a quite strange behavior: It works without a problem in static HTML pages where the code can be executed for example on the "onLoad" event but it fails when the DOM-tree is manipulated and images are added on-the-fly with executing the image-fallback after the last image is added to the tree. I couldn't find any real solution for this yet. The only thing that seems to work in both browsers is defining a background-color via CSS for the images that should be completely covered by the images. If the image doesn't appear, at least the background would be rendered. Background-images seem to fail in Opera8 too.