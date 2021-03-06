---
date: '2007-12-08T12:00:00-00:00'
language: en
tags:
- apps
- domain
- google
title: Google Apps for your domain?
---


For about a week Dreamhost's e-mail forwarding to gmail addresses was down. I
don't really care who was really responsible for this (to me it seems like
Google caused it) made me look around for an alternative solution
for domain-related email addresses and I stumbled upon [Google's domain apps](http://www.google.com/a/).


-------------------------------

With this service Google offers anyone who has a domain to directly use for example GMail for their
domain-emails on a DNS level. So no simple forwarding but real email hosting.
So far this looks definitely interesting, but on a Wiki page in the 
respective google group I found [a list of limitations](http://groups.google.com/group/apps-discuss/web/features-gmail-has-but-the-paying-customers-of-google-apps-are-still-waiting-for):

>    * Google Groups integration 
>    * Color Labels
>    * AOL Chat integration
>    * Creating filer from this message feature
>    * Update contact profiles
>    * Integration of accounts with rest of Google service
>    * Gmail Mobile App for Google Apps

I went on to create a domain account for one of my sites just to see if these 
features were really missing ... and yeah, I was really really surprised
to see that there are for instance really no colored labels. BUMMER!

<div class="figure">
    <img src="/media/2007/googleapps_nolabels.png" />
    <p class="caption">There are really no labels.</p>
</div>

On the other hand the whole process of setting Google Apps up was pleasantly 
simple and straight forward. All I had to do (besides filling out a couple of
forms) was to upload a file to my domain and edit the MX-records of my domain.
For the latter Google also offers detailed tutorials for some of the
more well-known hosting companies out there.

Another problem I've faced so far were forwarders: There is no "Create a forwarder"
link anywhere in the email settings. From what I've seen so far, there are 
2 options to still get them to work:

1.  If the email should be forwarded to an address within the hosted domain, 
    just create a normal user account for it or add a "nickname" to the user
    that should receive that email.
2.  If the email should be forwarded to an outside email address, setting
    up a mailing list with that address as single recipient should also get
    this job done.
    
In my opinion the help section is lacking a bit in this regard ;-) So to 
sum up my experience with it so far (about 1h of use):

**PROS:**

*   Very easy setup
*   Free (if you choose the standard edition)
*   More or less complete Google services
*   Easy administration

**CONS:**

*   The "more or less" above. Without the new features for example in GMail
    the mailing is really lacking some important stuff.
*   Google Support is automatically receiving all mail sent to postmaster@domain.com
    and abuse@domain.com. Sure, you can add yourself also to this list, but
    I don't like it that Google gets these mails by default.
*   Missing documentation for forwarders.
