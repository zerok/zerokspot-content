---
date: '2009-12-22T12:00:00-00:00'
language: en
tags:
- accessiblity
- javascript
- keyboard
- lightbox
- overlay
- quicktip
title: 'QuickTip: JQueryTools Overlays and Accessibility'
---


During the latest minor design changes here (no sidebar anymore, yay!) I placed the information that was previously provided within a sidebar next to every entry into a couple of lightbox elements in combination with a small dropdown menu. Both were implemented so that non-javascript-enabled browsers should still be able to at least access the content, but I wasn't yet certain how to add the keyboard navigation to it.

Today I browsed a bit around and found a really nice article by Roger Johansson about &quot;<a href="http://www.456bereastreet.com/archive/200910/lightboxes_and_keyboard_accessibility/">Lightboxes and keyboard accessiblity</a>&quot;. He basically sums up all my gripes I normally have with lightbox implementations out there; and just to pick one here: They lose focus. Most already come with a keybinding for closing them again, but they normally out of the box don&#39;t restore your focus from before entering them.&nbsp;<a href="http://flowplayer.org/tools/index.html">jQuery Tools</a>&#39; overlays seem not to be an exception here, but it is rather simple to give an overlay some memory in this regard :-)

---------

You have to trigger the overlays programmatically then then more or less stow away the information about the last focused element:

<pre class="code">var overlayElement = jQuery(&#39;#someOverlay&#39;).overlay({&#39;api&#39;:true&#39;});
overlayElement.bind(&#39;onLoad onClose&#39;, function(evt){
  if (evt.type == &#39;onLoad&#39;){
    this._previousFocus = document.activeElement;
  } else {
    this._previousFocus.focus();
    this._previousFocus = null;
  }
});
overlayElement.load();</pre>

Note that document.activeElement is not available in every browser (e.g. Safari 3.x) but all current browsers should support it.

To keep the flow of the keyboard navigation going you should also focus on the first focusable element within the lightbox:

<pre class="code">var focusable = jQuery(&#39;a, input, button, textarea&#39;, overlayElement);
if (focusable.length &gt; 0) focusable[0].focus();</pre>

I&#39;ve placed this just right after I stored the previously focused element away.

That&#39;s basically it. Now if you open a lightbox via keyboard, your focus jumps right to the first focusable element within the lightbox (i.e. a link or form element) and you can navigate in there. Once you close the lightbox (normally via the escape button), the focus jumps right back to where you left it before entering the overlay.

One thing I haven&#39;t figured out yet is how I can kind of restrict the list of focusable elements to those within the current overlay, but this is just on my nice-to-have list.
