---
date: '2006-03-15T12:00:00-00:00'
language: en
tags:
title: Executable check in Rubygem ...
---


Sometimes I think I don't learn anything from my mistakes. Everytime I write a new rubygem, at first my executable doesn't do anything no matter what commandline parameter I set etc. The reason is quite simple: Since I do a lot of Python scripting to basically every executable looks like this for me:

-------------------------------



<pre class="code">if __FILE__ == $0

	# ... code ...

end</pre>



This normally a good way to get a section of the script to be only executed when the file itself is called and not as part of a library inclusion. The problem is now, that if you're doing this with a rubygem executable it won't work. Rubygem automatically creates a wrapper script for each executable somewhere in your path (for example /usr/bin). Above comparision would now look like this:



<pre class="code">if "/usr/lib/ruby/gems/1.8/gems/mygem-1.0.0/bin/mygem_exec" == "/usr/bin/mygem_exec"

	# ... code ...

end</pre>



... which obviously can't work ;) So just as a small reminder esp. to myself: You don't need this check inside a rubygem since rubygems does the checking for you and you split the executable into a separated file anyway ;)