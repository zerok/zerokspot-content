---
date: '2008-10-21T12:00:00-00:00'
language: en
tags:
- buildout
- distutils
- python
- recipe
title: distutils-recipe for buildout
---


Today I released something I've planned to write for quite some time now: A very minimal recipe for zc.buildout that can install an arbitrary source distribution of a Python package. Sure, there is already [collective.recipe.distutils](http://pypi.python.org/pypi/collective.recipe.distutils/0.1) out there and big kudos to Kevin Teague for his work there, but I had some minor problems with it. I also wanted to get some more practice when it comes to writing recipes for buildout. As always, you can find the result  on [PyPI](http://pypi.python.org/pypi/zerokspot.recipe.distutils/) and [github](http://github.com/zerok/zerokspot.recipe.distutils/).

-------------------------------

So first, what's its purpose? Quite simple: You give it a list of URLs to tarballs with packages you want to have installed (basically source distributions produced either with setuptools or distutils but not published on PyPI) and it will download and install them into the part's directory. A quick example: I wanted to have an environment that includes the PIL. Since it isn't really available on PyPI nor is there an egg for it, I had to reply on  distutils to do its magic without any automated help:
    
    [buildout]
    parts = django pil
    
    [django]
    recipe = djangorecipe
    version = 1.0
    project = myproject
    settings = settings
    extra-paths = 
        ${pil:extra-path}
    
    [pil]
    recipe = zerokspot.recipe.distutils
    urls = 
        http://effbot.org/downloads/Imaging-1.1.6.tar.gz
        
I combined this example with the djangorecipe since it also shows you the "extra-path" variable provided by the recipe, which points directly to the respective "site-packages" folder. So in the end you will have a folder "pil" within the parts-directory which includes your usual folder-structure for Python-prefix. `${pil:location}` will point to that while `${pil:extra-path}` points to `parts/pil/lib/python2.6/site-packages`

There are now a couple of differences to Keven Teague's recipe:

1. It uses the "--prefix" option for the setup.py-execution instead of the "--install-{pure,plat}lib" arguments which also forces scripts to be installed into the predefined directory.
2. Every part has its own directory instead of putting everything into `${buildout:parts-directory}/site-packages`
3. It supports multiple packages within a single part
4. It is by far not as configurable as collective.recipe.distutils ;-)

So for most pure libraries both recipes should be equally well suited, but if you have a package that also includes script (which you don't want but are simply there) probably zerokspot.recipe.distutils is better. That said, those additional scripts are not really usable since the PYTHONPATH isn't coded into them as it would have been for instance for the manage.py-wrapper provided by djangorecipe, but normally you won't need them if all you need are the libraries themselves. Let's see, perhaps there is also a solution for this, although it is very low on my priority-list right now.
    