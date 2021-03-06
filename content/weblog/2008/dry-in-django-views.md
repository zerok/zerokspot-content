---
date: '2008-04-19T12:00:00-00:00'
language: en
tags:
- django
- dry
title: DRY in Django views
---


A couple of days ago, [a topic](http://groups.google.com/group/django-users/browse_thread/thread/476d53ffea24b809#) came up on the django-users mailinglist where someone asked how to avoid code repetition in views if the views mostly do the same (for instance do some processing on the request parameters or check a certain session variable etc.).

There is one question you have to ask yourself first, though: Does this common processing require request-specific data like session-data or request parameters? 

-------------------------------

In either case, you have multiple options. For instance you could write a function or decorator to do this common processing. This only makes sense, though, if most of the processing results end up in for instance the request object itself again. Otherwise you'd simply have way too many return values to keep the code maintainable.

Another option would be to convert your view into a class and put shared processing for instance into the constructor.

If you now want to for example handle a couple of request-parameters in the same way in multiple views, option 2 is actually pretty simple. You could, for instance, get something quite flexible while requiring only 3 components: (1) the actual view class, (2) a small factory function to convert such a class into a usable view and (3) the view function that comes out of the factory function
    
    def create_view(klass):
        def _func(request, *args, **kwargs):
            after = '__after__'
            o = klass(request, *args, **kwargs)
            r = o(request, *args, **kwargs)
            if hasattr(o, after):
                return getattr(o, after)(r)
            return r
        return _func
    
    class BaseView(object):
        def __init__(self, request, *args, **kwargs):
            # Do some generic stuff here
            pass
        def __call__(self, *args, **kwargs):
            raise RuntimeError, "Do not call the BaseView"
        def __after__(self, response):
            return response
    
    class MyView(BaseView):
        def __call__(self, request, *args, **kwargs):
            return HttpResponse("Hello World")
    
    myview = create_view(MyView)

With a structure like this you'd put all the functionality that is shared by your views into the \_\_init\_\_-method of the BaseView and then put the view-specific stuff into the \_\_call\_\_-method. And if you want to do some post-processing, just overwrite the \_\_after\_\_-method.

If you want to use decorators like django.contrib.auth's login_required, you can do so by decorating the generated view function:
    
    myview = login_required(create_view(MyView))

If you don't do any request-specific processing that you want to aggregate into one function, you could also do something like this (similar to how the new [django.contrib.formtools.wizard](http://code.djangoproject.com/browser/django/trunk/django/contrib/formtools/wizard.py?rev=7294) module does it):
    
    #--- views.py ---
    
    class MyView(object):
        def __init__(self, *args, **kwargs):
            # Do something
            pass
        
        def __call__(self, request, *args, **kwargs):
            # Actual view with response generation etc.
            return HttpResponse("Hello World")
    
    #--- urls.py ---
    
    from django.conf.urls.defaults import *
    from .views import MyView

    urlpatterns = patterns('',
        (r'^/?', MyView()),
    )

Here you more or less only use the constructor to configure your view while in the previous approach you can also use it to handle request parameters.

In 2006 Rob Wolfe [posted a similar approach](http://mail.python.org/pipermail/python-list/2006-August/399781.html) without putting a class instance into the URL configuration, which has the advantage, that you don't have to actually import the view module upfront. Here the actual view object is created when the views.py is first imported, which -- as in the approach above -- does the whole common processing as rarely as possible.