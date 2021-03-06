---
date: '2008-07-26T12:00:00-00:00'
language: en
tags:
- conversion
- django
- mysql
- postgresql
title: 'Move a Django site to PostgreSQL: check'
---


When I moved over from Dreamhost to [Slicehost](http://www.slicehost.com), one of the (minor) motivators was the option to use something different than MySQL. Don't get me wrong, though: MySQL is a nice system, fast, relatively memory-efficient and so on. But what drove and still drives me nuts is that not every storage engine within [MySQL](http://www.mysql.com) supports transactions. MyISAM can be as fast as it gets, if it lacks transaction support, I won't use it for at least some of my tables. And having to manually check that every time I write a new app really started to annoy me. But how to move a site like this, that is using [Django](http://www.djangoproject.com), from one database system to another (or to be specific, [PostgreSQL](http://www.postgresql.org) in this scenario) as simply as possible?


-------------------------------

Note that this is just a rough description of the process I used. The principle should work in general, but you might face some additional difficulties that I didn't notice or haven't noticed yet. And don't forget to make a backup of your old data before doing anything described here.

The whole process is remarkably simple once you know what you have to do. Django helps a lot in this regard thanks to its fixture-system, that is luckily DBMS oblivious. Basically the whole move works like this:

1. Dump the database into a fixture using `python manage.py dumpdata --indent=4 > dump.js`
2. Change your settings module to point to a new PostgreSQL database (which you should create before the next step)
3. Run `python manage.py syncdb` to create the tables within your PostgreSQL database
4. Clean this database from the initial values with `python manage.py sqlflush | psql yourdatabase`
5. Run `python manage.py loaddata dump.js`
6. Make sure that the sequences in the new database are up to date

For step 5 I wrote a small script that goes through each sequence in the database (since this database is used for one Django site and one site only this works :P ) and resets its value to the highest id of the associated table:

<pre><code>
import os, sys
sys.path.insert(0, '../')
os.environ['DJANGO_SETTINGS_MODULE'] = 'mysite.settings'

from django import db
c = db.connection.cursor()
try:
    c.execute(r"""SELECT c.relname
        FROM pg_catalog.pg_class c
            LEFT JOIN pg_catalog.pg_namespace n ON n.oid = c.relnamespace
        WHERE c.relkind IN ('S','')
            AND n.nspname NOT IN ('pg_catalog', 'pg_toast')
            AND pg_catalog.pg_table_is_visible(c.oid)
        """)
    to_update = []
    for row in c:
        seq_name = row[0]
        rel_name = seq_name.split("_id_seq")[0]
        to_update.append((seq_name, rel_name,))
    for row in to_update:
        c.execute(r"SELECT setval('%s', max(id)) FROM %s"%row)
finally:
    c.close()
</code></pre>

Another problem that has to be solved before you do step 5 is that dumpdata has a small problem with BooleanFields in the models. If you're using any of those (you or any contrib module you're using), you have to do some cleaning up in the JSON dump. The problem is that the values of these BooleanFields are dumped as `0` or `1` instead of `false` or `true`, which confuses loaddata to no end. To make it easy for any processing script, I told `dumpdata` to nicely format the dump, which you can do with the `--indent` option, which comes in handy when trying to fix this small problem.

So I wrote a little Perl script (don't hurt me, please) that just goes over the dump and corrects it:

<pre><code>
while(<>){
    s/is_active": 0/is_active": false/;
    s/is_active": 1/is_active": true/;
    
    s/is_staff": 0/is_staff": false/;
    s/is_staff": 1/is_staff": true/;
    
    s/is_published": 0/is_published": false/;
    s/is_published": 1/is_published": true/;
    
    s/is_superuser": 0/is_superuser": false/;
    s/is_superuser": 1/is_superuser": true/;

    s/enable_comments": 0/enable_comments": false/;
    s/enable_comments": 1/enable_comments": true/;
    
    s/registration_required": 0/registration_required": false/;
    s/registration_required": 1/registration_required": true/;
    print $_;
}
</code></pre>

Naturally you'll have to correct the fieldnames for your own site :-) So basically before loading the dump again, run something like `perl convert.pl dump.json > dump2.json && mv dump2.json dump.json`.

As you can see, the process is really pretty straight forward, and, so far, it seems to work just fine :-)