---
date: '2008-06-16T12:00:00-00:00'
language: en
tags:
- fireeagle
- geo
- yahoo
title: Finally a Fire Eagle Invite
---


<img src="http://img.skitch.com/20080616-kpd778am7mr9dgwj65pwcfn86a.png" alt="" class="left" />I guess there is now a new wave of invites for Yahoo's location data dispatcher [Fire Eagle](http://fireeagle.yahoo.net/) out, because I finally received an invite tonight, 3 months after the first wave hit.

For those of you who haven't heard yet of this service (which is not all that surprising considering how low the buzz around this service has been so far), Fire Eagle tries to offer a unified architecture for sharing your current location with web applications and web sites. 

-------------------------------

This way you'd only have to link for instance your GPS-enabled mobile device with Fire Eagle and sites that you want to use this data, can access it. For each application you can then define, how much detail you want to share with it.

<div class="figure">
    <img src="http://img.skitch.com/20080616-fc1ky5r4jqq1ie1r9jj3spym.png" alt="" />
    <p class="caption">How much do you want to share with an application?</p>
</div>

If you want to drop off the radar for a while, there is also a button in the privacy settings, that prevents any data forwarding from Fire Eagle to applications for the time being.

On the developer side the service's API offers the functionality you'd expect: Handlers for user and location queries as well as update functions to update a user's location. Authentication as far as I can tell works more or less exclusively using OAuth. Since this service is at least slightly oriented towards generic mobile applications, the location struct can also hold information about the cell towers that I guess where used to triangulate a user's position. I have only very limited knowledge in this area, but the data provided there appears to be usable enough also to determine how reliable "exact" information provided by Fire Eagle might be.

Besides these functions, Fire Eagle also offers 2 functions for finding users. One for receiving a list of people who have recently updated their location and the other for finding people around a certain location.

I have to say, that I have quite mixed feelings about the use of this service, though. On the one hand I think it is a good idea to have a single API to update your location information, otherwise you'd have to connect every single geo application with your mobile device. On the other hand a centralized service that basically hosts all your information as always is first of all a bottleneck and second also a bit of a privacy dilemma. I guess only time will tell, how this all turns out, but Fire Eagle definitely looks interesting enough :-)