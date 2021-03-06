---
date: '2011-01-09T12:00:00-00:00'
language: en
tags:
- html5
- webforms2
- jquery
- jquerytools
- h5validate
title: HTML5 Form Validation Now!
---


OK, originally this post should have been about [h5Validate][h5v], a very simple and
lightweight plugin for jQuery that handles two specific parts of the whole
forms/webforms 2.0 part of HTML5: input patterns and the required-attribute. But
then I stumbled upon the new [form validation module of jQuery tools][val], one of
those plugins I'm close to add by default to every project I'm working on.
Oh boy ...

-------------------------------------------------------------------------------

## h5Validate

But first things first: [h5Validate][h5v] right now
handles two major form additions in HTML5: Allowing to mark fields as required
and offers client-side validation for input fields that have a [pattern][pat]
attached to them.  This is right now just a small subset of what browsers
will hopefully eventually support natively when it comes to form validation
but the plugin looks like a nice start. As with all the plugins I'm going to
mention here, h5Validate does not only try to stick to the standard but go
beyond it in some areas.

As an example: right now h5Validate only validates fields on focus-related
events (e.g. blur) while for instance [Firefox 4][ff4] does the validation
when the user tries to submit the form. You can also create a central place
for patterns and then apply them to multiple elements:
    
    <script type="text/javascript">
        $(function() {
            $.h5Validate.addPatterns({
                phone: /[0-9]{13}/
            });
            $('form').h5Validate();
        });
    </script>
    <form method="post">
    <input type="text" class="h5-phone" />
    </form>

The current version is 0.3.x so there is probably still quite a lot in the
future for this plugin, but especially when it gets submit-support and custom
error messages (which are both on the TODO list right now) it will really
become interesting :-) For now it seems like the field's title attribute is
the only way to customize the error message (similar to what FF4 has by
default). And you have to provide your own element to display the error
message in, compared to a custom widget handling that for you in the next
Firefox version.

Right now there are quite a few HTML5 form related plugins out there that all
only implement a subset of what HTML5 has to offer with regards to forms and
their validation. Besides h5Validate there is for instance [html5form][h5f] which
handles placeholders, required fields, autocomplete=off as well as some
non-pattern related input validation. I've also found [webforms2][wf2] which
provides support for a rather impressive subset compared to the other two but
doesn't handle placeholders among other things and seems no longer maintained.

## jQuery Tools Validator

Probably the biggest and most feature complete plugin for webforms 2.0
validation is the [Validator][val] module in [jQuery tools][jqt]. When I first saw it it
was like X-Mas all over again. According to the docs it not only supports
patterns, custom validators and required fields but also handles type=email,
type=url and type=number out of the box.  range and date input fields are
support by two other modules.

For rendering the error messages it does not require that you change the
markup but instead renders the required elements on its own. Thankfully it
also doesn't simply use the title attribute when rendering the error message
but goes for a custom attribute (data-message) and falls back to globally
defined ones. So while in other frameworks you would define the error message
like this:
    
    <input type="text" name="alphanum" pattern="[a-zA-Z0-9]" title="This field
        may only contain alpha-numeric characters" />

jQuery tools Validator keeps the title as an info-attribute and takes the
error message from [data-message][dm] like this:
    
    <input type="text" name="alphanum" pattern="[a-zA-Z0-9]" title="Enter
        something here" data-message="This field
        may only contain alpha-numeric characters" />

And if your app already uses data-message for something else, you're free to
use a different attribute:
    
    $('form').validator({
        'messageAttr': 'date-errormessage'
    });

But these are just the not-localized error messages. The module offers a whole
range of ways to also handle [localized error messages][jtl].

As with practically any part of jQuery tools also the form validator offers a
handful of custom events, in this case there are events for the whole form and
also for specific fields. Just a small example: Let's say you have a
registration form and a new users comes by and wants to create a new account.
Sadly her username is already taken and you want to offer her a few randomly
generated options based on some parts of her registration details.
    
    $('.registration-username').oninvalid(function(e, message) {
        // Fetch a list of names and render them.
        // ...
    });

The only downside I could find so far about this is that the Validator module
seems to have some issues with Firefox 4. But, given that FF4 is still in beta
I guess it's safe to assume that these issues will be fixed around the time
4.0 ships.

So, especially if you combine Validator with the [Dateinput][di] module it looks like a
truly awesome package that I'm really looking forward to use in some future
project :D OK, now I have to calm down, get a hot milk with honey or something
...

[h5v]: https://github.com/dilvie/h5Validate
[ff4]: http://blog.oldworld.fr/index.php?post/2010/11/17/HTML5-Forms-Validation-in-Firefox-4
[pat]: http://www.whatwg.org/specs/web-forms/current-work/#the-pattern
[h5f]: http://www.matiasmancini.com.ar/jquery-plugin-ajax-form-validation-html5.html
[wf2]: http://code.google.com/p/jquery-webforms-2/
[jqt]: http://flowplayer.org/tools/index.html
[di]: http://flowplayer.org/tools/dateinput/index.html
[dm]: http://flowplayer.org/tools/validator/index.html#input_message
[val]: http://flowplayer.org/tools/validator/index.html
[jtl]: http://flowplayer.org/tools/validator/index.html#localization
