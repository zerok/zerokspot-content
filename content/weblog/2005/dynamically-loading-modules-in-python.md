---
date: '2005-05-03T12:00:00-00:00'
language: en
tags:
- development
- python
title: Dynamically loading modules in Python
---


For a small fetch script I wanted to be able to simply put new modules into a certain directory and let the script use it at the text start. Python offers with the __import__(...) function a really nice tool for something like that. Here a small example:

-------------------------------



<pre class="code">

modules = []

for f in os.listdir("modules"):

	if f.startswith("mod_"):

		if f.endswith(".pyc"):

			minus = 4

		else:

			minus = 3

		module = "modules."+f[:-minus]

		if module in doneModules:

			continue

		doneModules.append(module)

		print "Importing %s"%(module)

		modules.append(__import__(module,globals(),locals(),["modules"]))

</pre>



The only really interesting part is the last line in the loop. Here a new module is loaded an appended to the modules list. Since the modules are in the "modules" package the last parameter of the __import__ call is no longer optional. Otherwise you'd only do a `import modules.mod_whatever` instead of a `from modules import mod_whatever`.



Another note: The first parameter has to be in this case "modules.mod_whatever" and not only "mod_whatever". If you forget this you will get a "module not found" error.