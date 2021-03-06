---
date: '2006-10-24T12:00:00-00:00'
language: en
tags:
- customizing
- extensions
- thunderbird
title: 'Thunderbird: The way I like it'
---


During my recent moving around between Mail.app and Thunderbird I discovered some extensions for the Mozilla mail. Only one of the extensions I'm currently using in Thunderbird has been in my list for quite some time, the others are new to me, but could become the features that will hold me with Thunderbird for the next couple of years ;)



-------------------------------



* [Enigmail](https://addons.mozilla.org/thunderbird/71/) is an extension to allow the encryption/description and signing/verifying of emails using GnuPG. It also offers a simple method to import public keys from various servers out there. This is the one extension I've mentioned above which I've been using basically since I know Thunderbird/Mozilla Mail.
* [Nostalgy](https://addons.mozilla.org/thunderbird/2487/) provides you with a way to move and copy emails into different folders using the keyboard alone.
* [Copy Sent to Current](https://addons.mozilla.org/thunderbird/2561/): One of the nice features of Gmail is, that it tries to bundle mails into conversations. This is also possible with Thunderbird out of the box using the threaded view, but if you don't BCC the mails to yourself, your own mails won't become part of these threads. And here is where this extension comes in handy: It lets you store a copy of your mail into an arbitrary folder _and_ into your Sent folder. 
* [Growl New Message Notification](https://addons.mozilla.org/thunderbird/3448/) simply integrates Thunderbird into Growl. Whenever you get a new mail it will pop up.

Another small thing that I like about Mail.app, but that I never really know how to get in Thunderbird was a larger sidebar. Not wider, but simply with a larger font in there to make dragging and dropping much easier. Getting something like this is actually very easy the moment you find out about the userChrome.css, a configuration css file for Thunderbird's GUI. [Here](http://forums.mozillazine.org/viewtopic.php?t=477746)'s the topic I started on mozillazine.org to get the whole thing right, and here comes the result:

<a class="figure thickbox" href="http://zerokspot.com/uploads/thunderbird_bigsidebar.png" title="Click to zoom"><img src="http://zerokspot.com/uploads/thunderbird_bigsidebar.mini.png" alt="Thunderbird with a larger sidebar"/></a>

With all this I now finally feel at home with Thunderbird once again :)

__Update:__ Just a small addition to make the whole "getting the font larger" is explained a little bit ;-)

Create a new file in your profile's _chrome_ folder (on MacOSX this would be ~/Library/Thunderbird/Profiles/xxxxxxx.xxxxxxx/chrome) with the name userChrome.css and put following into it:

<pre class="code css">
#folderTree &gt; treechildren::-moz-tree-cell-text{
 font-size:11pt !important;
}
#folderTree &gt; treechildren::-moz-tree-row {
  height: 30px !important; 
}
</pre>

Save it and restart Thunderbird.