---
date: '2013-02-26T21:15:17+01:00'
language: en
tags:
- django
- sprint
- utrecht
- djangosprint
title: Django Sprints Utrecht 2013
---

    
In recent years I've been on quite a few sprints (various PyConDEs, EuroPythons and DjangoCon EUs) but I've so far never attended a sprint-only event. When [Florian][flo] told me there was one in Utrecht in February it was a good opportunity to get for a few days away from the displays at home and explore a place I've never been before but had been on my list for quite a while now.

---------------

## Getting there

<img src="http://photos.h10n.me/Events/Django-Sprint-2013-Utrecht/i-DwKdM2Z/0/M/DSC00507-M.jpg" alt="" style="float:left;margin:0 1em 1em 0;border:1px solid #DDD" />Fast forward to Utrecht: Getting there was very simple: Got on a plane in Graz, switched planes in Munich and get off in Amsterdam (by now I know Schiphol well enough to find the Starbucks and Burger King blind-folded) and onto a train to Utrecht. Thirty minutes and 8.30€ later I arrived in Utrecht into a big construction site. Turns out that every major city has to modernize their central station when I visit and pass through.

Luckily, all these efforts were limited to the area around the train station and they didn't really interfere with me finding my way around it, out of it and into the hotel (although it took me a couple of minutes to notice that the exit I had taken had been closer to the hotel than expected).


## Utrecht

Since it was still quite early (11:00 am with my plane having left Graz at 06:05 am) the room at my hotel wasn't ready yet and so I had about 4 hours to kill in town. So off I went with my backpack full of electronics and a camera to explore the city of Utrecht!

<img src="http://photos.h10n.me/Events/Django-Sprint-2013-Utrecht/i-9ZWTWWm/0/M/DSC00521-M.jpg"  style="float:right;margin:0 0 1em 1em;border:1px solid #DDD" alt=""/>After just a few minutes (and not being killed by any bicyclist, thank you very much) I reached the [Dom Church][dom]. Sadly, as with apparently every major site nowadays, the tower itself was undergoing renovations. Having seen probably more than enough clerical buildings from the inside in my life I opted to just look at the Dom and its tower from the outside. Considering how tightly every building is packed here I was really amazed about the openness around the tower and the actual church.

From there I basically just went strolling through the southern part of the inner city (since I messed up my inner compass considering that I actually wanted to go to the west). Utrecht is just a beautiful city (at least the old-town parts I was walking through) with all its water ways and large bike-streets. It also seems to be an exceptionally "young" city with lots of students probably thanks to the local university.

<figure>
    <img src="http://photos.h10n.me/Events/Django-Sprint-2013-Utrecht/i-PFK6LkF/0/XL/DSC00540_1_2_tonemapped-2-XL.jpg" alt="" />
</figure>

Eventually reaching the crossing of Venuslaan and Albatrosstraat I noticed that I was probably way off from where I actually wanted to be and so went back to the central station and the hotel to get some sleep.

## The Sprint

After a nice 13-hours sleep I made it (a little early) to the sprint location: The [Dutch Game Garden][dgg], a shared office space supported by some EU project and therefor required to also offer some space for "public" events like the Django sprint :-D

<figure>
<img src="http://photos.h10n.me/Events/Django-Sprint-2013-Utrecht/i-PrXrfRK/0/L/DSC00580-L.jpg" alt="" />
<figcaption><p>Everyone being hard at work in Utrecht</p></figcaption>
</figure>

If you've never been to a coding sprint just think about it as a big room full of laptops and smaller tables for people to collaborate on project-related issues. At the start there is a short introduction how to select a ticket from the issue-tracker, what flags to change on it and when and what you should do once you consider the issue being resolved.

At official Django sprints you also always have at least one core-developer on site who helps you with open questions and can eventually also get your code into the project's main repository. For this one we even had three: Florian Apolloner, Aymeric Augustin and Jannis Leidel. Sadly, Jannis got sick after the first day but even then we managed to get (according to Florian) around 100 commits into the master branch over the course of the weekend.

Combine that with the commits of all the other sprint locations (Cracow, Istanbul, Kansas City, Córdoba, San Franciso and Tokyo) I'm pretty sure we got a lot of work done in those two days (and only very rarely broke the build ;-)). Heck, even I got three tickets resolved:

* [#19758][19758] (Information disclosure within the password reset process)
* [#11971][11971] (Missing documentation of serialization formats)
* [#17320][17320] (Prevent users from accidentally adding whitespace characters to domain names in the Site object)

Big thanks to all the organizers for having us :-)

[19758]: https://code.djangoproject.com/ticket/19758
[11971]: https://code.djangoproject.com/ticket/11971
[17320]: https://code.djangoproject.com/ticket/17320


## Time between Coding

<img src="http://photos.h10n.me/Events/Django-Sprint-2013-Utrecht/i-JnmRPwK/0/M/DSC00586-M.jpg" style="float:left;margin:0 1em 1em 0;border:1px solid #DDD" alt=""/>For my stay I chose a [NH Hotel][hotel] right next to the main station in order to make leaving on Monday that much easier :-) I have never been to a NH hotel before but I always considered them to be around the 3-4 star range. My room definitely matched the 3-star category, albeit of the slightly older kind. The room has definitely seen better days but it was still nice esp. thanks to being in the 17th floor. Just some minor annoyances like really bad TV-reception, a weirdly structured bath room and a *really old* air conditioning system that has probably been cleaned the last time when it had been installed in the 1970s. Oh, and 128kbps WiFi. Seriously?!

That said, the breakfast more than simply made up for this. This has been the first time that I got scrambled eggs in the Netherlands that didn't look strange (due to having a slightly to light color) and you could add things like bacon, ham, tomatoes and other fine-cut ingredients to them :-) 

## Souvenirs

Originally I had planned to look out for one of these town-specific rubber ducks. NYC has one resembling the Statue of Liberty, Salzburg has a Mozart-duck, so I was curious what the Utrecht-version would look like. Due to long sprinting and sleeping far too well at the Hotel I never made it to a souvenir store, though, so I just got myself a couple of Starbucks mugs and the sprint organizers let me keep two packages of *Hagelslag* which I grew quite fond of during my three days in the Netherlands :-)

<figure>
<img src="http://photos.h10n.me/Events/Django-Sprint-2013-Utrecht/i-zfqRpwR/0/L/DSC00599-L.jpg" alt="" />
<figcaption><p>My souvenirs from Utrecht and Schiphol Airport</p></figcaption>
</figure>

If you see me the next time, expect me to have considerably more problems getting through doors :-P

[flo]: https://github.com/apollo13
[dom]: http://en.wikipedia.org/wiki/Dom_Tower_of_Utrecht
[dgg]: http://www.dutchgamegarden.nl/
[hotel]: http://www.nh-hotels.com/nh/en/hotels/the-netherlands/utrecht/nh-utrecht.html
