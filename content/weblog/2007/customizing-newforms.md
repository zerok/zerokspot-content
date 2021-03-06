---
date: '2007-06-01T12:00:00-00:00'
language: en
tags:
- django
- newforms
- tutorial
title: Customizing newforms
---


You like newforms but are a little bit uncertain how to style your forms using CSS with form.as_ul? Well, if you want more CSS classes to style upon, here a simple way on how to get them the easy way.

This mini-tutorial will show you how to add CSS classes to the help text and position the errorlist differently for fields in a newform-Form rendered with as\_p() (although this also should apply to all the other rendering methods).



-------------------------------



First of all a short look at the current implementation of the as\_p() method of the BaseForm class as defined in django.newforms.forms:

<pre class="code">
def as_p(self):
    &quot;Returns this form rendered as HTML &lt;p&gt;s.&quot;
    return self._html_output(u&apos;&lt;p&gt;%(label)s %(field)s%(help_text)s&lt;/p&gt;&apos;,
		u&apos;&lt;p&gt;%s&lt;/p&gt;&apos;, &apos;&lt;/p&gt;&apos;, u&apos; %s&apos;, True)
</pre>

So if we want to change this, just create a subclass of the Form class (which itself is a subclass to BaseForm) somewhere in your code - let's say in myproj.myapp.forms for now:

<pre class="code">
import django.newforms as forms

class BaseForm(forms.Form):
	def as_p(self):
		return self._html_output(u&apos;&lt;p&gt;%(label)s %(field)s%(help_text)s&lt;/p&gt;&apos;,
			u&apos;&lt;p&gt;%s&lt;/p&gt;&apos;, &apos;&lt;/p&gt;&apos;, u&apos; %s&apos;, True)
</pre>

To get more detailed output, all we have to do, is change the parameters passed to self.\_html\_output(..), so I will skip the rest from now on. 

This method offers parameters for everything we will need from now on:

<dl>
	<dt>normal_row</dt>
	<dd>Here you can set the base format for the field rendering.</dd>
	<dt>error_row</dt>
	<dd>This argument allows you to style the "wrapper" for the errorlist and will only get used if really an error affecting the current field appears and it will only be used if the errorlist should get its own row (as can be set with the last argument).</dd>
	<dt>row_ender</dt>
	<dd>This is only used internally, so we will skip this.</dd>
	<dt>help_text_html</dt>
	<dd>Formating the help output which only gets used if there is any help_text to apply it on ;-)</dd>
	<dt>errors_on_separate_row</dt>
	<dd>Remember when I said error_row is a little bit misleading? Well, this only really applied when you set this parameter to False ;-) Then you will have to put "%(errors)s" somewhere into your normal_row format in order to get any error rendering at all :-)</dd>
</dl>

So let's apply this intel: First of all, it would be handy, if the help_text of each form field had it's own class, so we could easily ban it to its own line using display:block if we want to:

<pre class="code">
return self._html_output(u&apos;&lt;p&gt;%(label)s %(field)s&lt;span class=&quot;help&quot;&gt;%(help_text)s&lt;/span&gt;&lt;/p&gt;&apos;,
	u&apos;&lt;p&gt;%s&lt;/p&gt;&apos;, &apos;&lt;/p&gt;&apos;, u&apos; %s&apos;, True)
</pre>

This is the dirty way of doing it, since it will create that span even if the actual help text is empty. To get around this, you can change the help\_text\_html argument:

<pre class="code">
return self._html_output(u&apos;&lt;p&gt;%(label)s %(field)s%(help_text)s&lt;/p&gt;&apos;,
	u&apos;&lt;p&gt;%s&lt;/p&gt;&apos;, &apos;&lt;/p&gt;&apos;, u&apos;&lt;span class=&quot;help&quot;&gt;%s&lt;/span&gt;&apos;, True)
</pre>


So this will make styling the help text way easier. But one big problem for me still remained: The error messages for each field were flying around like crazy. That's because the argument that is set to True here (errors\_on\_separate\_row) actually tells the method, to put the error messages to a field on their own rows, or in this case in their own paragraphs. Since I don't like that, I will just switch it to _False_, which will put the error list right into the same paragraph with the rest of the field's output. If you do this, you will also have to add the errors to the first argument:

<pre class="code">
u&apos;&lt;p&gt;%(label)s %(field)s%(errors)s%(help_text)s&lt;/p&gt;&apos;, u&apos;&lt;p&gt;%s&lt;/p&gt;&apos;
</pre>

The error\_row content can stay the same, since the error list already comes with the "errorlist" class by default :-)

