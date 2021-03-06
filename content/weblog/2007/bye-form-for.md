---
date: '2007-12-31T12:00:00-00:00'
language: en
tags:
- django
- modelform
- newforms
title: Bye form_for_*
---


Hmm... just noticed a big change in Django's newforms API which happened
in [r6844][] (so basically on Dec 2nd): ``form_for_model`` and ``form_for_instance`` 
got deprecated in favor of the the new ModelForm class. The motivation
behind this move is quite well documented on the [django-dev mailinglist][], 
but for me personally it seems to boil down to replace `form_for_*` with
a much more declarative syntax (which is a good thing).

So if you for instance have a model ``Game`` like this (if I understood it correctly)
    
    class Game(models.Model):
        name = models.CharField(max_length=150)
        ...
    
you could easily create a form class for it like this:
    
    class GameForm(forms.ModelForm):
        class Meta:
            model = Game

I really can't wait to actually play around with this since it just makes
everything all so much more consistent :-)

For some details checkout the new [modelforms][] documentation page.


[django-dev mailinglist]: http://groups.google.com/group/django-developers/browse_thread/thread/44a782308af81796
[r6844]: http://code.djangoproject.com/changeset/6844
[modelforms]: http://www.djangoproject.com/documentation/modelforms/
