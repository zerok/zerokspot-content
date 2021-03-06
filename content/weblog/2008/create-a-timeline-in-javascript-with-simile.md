---
date: '2008-06-25T12:00:00-00:00'
language: en
tags:
- javascript
- simile
- timeline
title: Create a Timeline in JavaScript with SIMILE
---


From time to time you get to a point, where you want to present information in a way that makes it easy to see and understand their temporal relationship. For these occasions, timelines are a pretty neat tool, yet how do you do something like this within a website in a way that keeps the whole thing interactive? 

Besides the usual suspects (the 1000th Flash/Flex implementation), there are also some JavaScript based solutions, that get the job done quite well. For a project that is part of my master thesis right now, we were looking for a simple timeline tool that could easily be embedded and extended to our needs and found the [SIMILE Project](http://simile.mit.edu/) of the MIT. Part of this project is also a [timeline](http://simile.mit.edu/timeline) implementation in JavaScript that is remarkably flexible. 


-------------------------------


<div class="figure"><img src="http://img.skitch.com/20080625-ph2buh5q11txi4rf2jemijmxp5.png" alt="" /><p class="caption">Multiple bands (Screenshot from <a href="http://simile.mit.edu/timeline/">simile.mit.edu/timeline/</a>)</p></div>

Among other cool ideas is a way to have multiple bands in your timeline. For instance you have one band that shows events happening on a day, below that you have have a narrower band showing the same events on a monthly base and so on. 

## Small Example

As an example let's just build a small timeline of the recent history of [Django](http://www.djangoproject.com), including the major branches and releases. 

## Load the Timeline API

Therefor first get the latest source:

    $ svn checkout http://simile-widgets.googlecode.com/svn/timeline/trunk timeline

Then place the new timeline folder for instance right into the folder containing a django\_timeline.html, which will be our main playground for now. The project itself comes with tons of separated JavaScript files, but also delivers a simple AJAX-loader that we will integrate into our `django_timeline.html` file like this:

```
<head>
    <!-- ... -->
    <script src="timeline/src/webapp/api/timeline-api.js" 
        type="text/javascript" charset="utf-8"></script>
    <!-- ... -->
</head>
```

Next we need to define a container for our timeline and initialize it using the onload hook (or whatever event you want to trigger the initialization). Therefor we create a simple div that we give the ID "timeline".

```
<style type="text/css" media="screen">
    #timeline {width: 1000px; height: 400px;}
</style>
<script type="text/javascript" charset="utf-8">
    function init(){};
</script>
<!-- ... -->
<body onload="init()">
    <div id="timeline"></div>
</body>
```

Note that setting at least the height of the timeline is required and that you have to use fixed values here, so no 100% for you. If you don't specify the width, it will default to 100%.

Now we just have to tell the Timeline API, how we want our timeline to be structured. This is done by creating an Array with one or more bandInfo objects in it and pass it to an initialization function. Let's say, we want 2 bands: One representing months, the other years, then our onload-handler would look like this:

```
function init(){
    var eventSource = new Timeline.DefaultEventSource(0);
    var bandInfos = [
        Timeline.createBandInfo({
            width: '70%',
            intervalUnit: Timeline.DateTime.MONTH,
            intervalPixels: 100,
            eventSource: eventSource
        }),
        Timeline.createBandInfo({
            width: '30%',
            intervalUnit: Timeline.DateTime.YEAR,
            intervalPixels: 400,
            eventSource: eventSource
        })
    ];
    bandInfos[1].syncWith = 0;
    bandInfos[1].highlight = true;
    timeline = Timeline.create(document.getElementById('timeline'),
        bandInfos);
}
```

This function basically does 3 things:

1. It defines the two bands mentioned above with different interval units.
2. It synchronizes the movements between the two bands so that if you scroll in one band, the other band moves as well.
3. It creates a new "Timeline"-object using these two bands and binds it to our container div.


## Loading some Data

A timeline is nice. A timeline with some data in it is much better, though. To load the data into our timeline, you can use 3 formats: XML, JSON and I assume they mean with "SPARQL" the [SPARQL Query Results XML](http://www.w3.org/TR/rdf-sparql-XMLres/) format. For this example we will simply use JSON and put the whole Django timeline into a file named "django.json" within the same folder as our "django\_timeline.html". 

The structure of such a JSON datasource is pretty simple. Basically, Timeline looks for a top-level element named "entries" which is a list of events. A small excerpt from the Django timeline would look like this:

```
{
    'entries': [
        /* ... */
        {
            'title': 'UnicodeBranch',
            'description': 'Unicode branch of Django',
            'start': 'April 07 2007 00:15:27 GMT',
            'end': 'July 04 2007 04:28:29 GMT',
            'isDuration': false
        },
        /* ... */
    ]
}
```

Since we want to have the data available right away, we have to append this to our init() function:

```
Timeline.loadJSON("django.json", function(data, url){
    eventSource.loadJSON(data, url);
});
```

The Timeline's own loadJSON-method will then simply eval the data and pass thet struct over to the EventSource's loadJSON-method, which will the build Event-objects from it.

Well, that's more or less all you have to do to get a timeline with some data. You can see the current state of the timeline [here](/uploads/timeline_demo/django_timeline.html).

## Styling

Nice, but there is definitely some Django-flair missing here. SIMILE Timeline also has something to offer to those of you, how want everything flashy and colorful.

It comes with a small theming system that basically consists of a class holding all the style information used for rendering the timeline. You can find this class definition in the timeline/src/webapp/api/scripts/themes.js file. When you have your own theme class, just set it as the new default theme by using Timeline's setDefaultTheme() method. Sadly, for now the implementation there seems to be quite incomplete. The version I'm working with here (r1416) has still quite a lot of code commented out -- not only in the ClassicTheme class, but also in the painters that should use these settings.  For example the event.tape.impreciseColor setting is passed to the Timeline.OriginalEventPainter.\_paintEventTape function, but never applied to the tape itself there.

So if you don't want to start hacking the JavaScript files themselves, I guess your best option for now is to use CSS. For example to add some Django-green to the duration-bars, add following CSS definition:

```
.timeline-event-tape {background:#275132;}
```

For more general styling, scan the timeline using for instance [Firebug](http://getfirebug.com/), which will reveal classes for basically every component of the timeline. If you want to apply a special style to a certain event you can do so by adding "textColor", "classname" or "color" to the actual event:

```
{
    'entries': [
        {
            'title': 'UnicodeBranch',
            'description': 'Unicode branch of Django',
            'start': 'April 07 2007 00:15:27 GMT',
            'end': 'July 04 2007 04:28:29 GMT',
            'isDuration': false,
            'classname': 'w00t'
        }
    ]
}
```

The sad thing again is that currently this only affects the labels of these events, but since the whole project appears to be quite active, I guess, it is only a matter of time until you can see this class on every event-related element.

If you can live with this (in my opinion minor) problems and/or don't mind extending the JS files yourself, SIMILE Timeline is in my opinion really a great tool for generating timelines.

Be sure to check out the [official documentation](http://simile.mit.edu/timeline/docs/) for further details.
