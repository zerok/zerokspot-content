---
date: '2005-11-11T12:00:00-00:00'
language: en
tags:
title: Firefox 1.5 Release Candidate 2
---


The Mozilla Foundation today released RC2 of the upcoming Firefox 1.5. Currently there is only the enlish version available which according to the release notes does fix some issues with the auto-update functinality. I haven't really tried it yet since I'm currently using Opera once again but the Gnome VFS integration seems also to work better now. The only problem I still have appears in following scenario:

-------------------------------



* Access flickr to upload a photo

* Open the "Open File" dialog and mount the camera from within this dialog

* Upload the picture

* Try to unmount the camera after having finished uploading the picture.



Unmounting wont work. You have to close the browser before the device is free to be unmounted.



I've now -- after searching and not finding something related to this  -- made a <a href="https://bugzilla.mozilla.org/show_bug.cgi?id=316031">bugreport on bugzilla.mozilla.org</a>.



For more details about FF1.5RC2 please read the <a href="http://www.mozilla.org/products/firefox/releases/1.5.html">Release Notes</a> :)