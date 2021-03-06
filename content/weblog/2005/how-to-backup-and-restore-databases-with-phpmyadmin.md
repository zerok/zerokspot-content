---
date: '2005-06-23T12:00:00-00:00'
language: en
tags:
- phpmyadmin
- tutorial
title: How to backup and restore databases with phpMyAdmin
---


<div class="warning note">
	
	<strong>Warning note:</strong>: This article was written a long time ago and therefor is based on a really old version of phpMyAdmin. The basic ideas should be the same, but I will still try to update this article within the next week (today is the 22nd of July 2006) to use 2.8.x.

</div>


## Where can I get phpMyAdmin?

You can get phpMyAdmin here: [http://www.phpmyadmin.net](http://www.phpmyadmin.net) The package includes a Documentation.txt that explains how to install it on your server. (Sorry, but this isn't meant as a tutorial how to install phpMyAdmin as this is already well documented.)

-------------------------------


Backup procedure
Note
This tutorial is based on phpMyAdmin 2.5.4. The GUI of this tool changes from version to version but the basics have stayed the same for quite some time now.

   1. Open phpMyAdmin with your webbrowser and enter your database account data to access your database.
   2. Now you can see a drop down menu in the left frame of phpMyAdmin where you select your database. <
      <div class="figure"><img src="http://zerokspot.com/uploads/select_db.png" alt="Select a database"/></div>
   3. In the header of the right frame you should now see Export. Click it :-)
      <div class="figure"><img src="http://zerokspot.com/uploads/select_export.png" alt="Click Export"/></div>
   4. Now you see a multiline selection field where you can select tables to be backed up. If you select no table or all tables then the backup stores the whole database.
   5. Select on the right side the "Structure" and "Data" option and below the selection field "Add 'drop table'" , "Complete inserts" , " Extended inserts" , " Enclose table and field names with backquotes" and "Save as file". If you have also "zipped" and "gzipped" there you can get a smaller backup file. In the left field select "SQL" if it's not already selected.
      <div class="figure"><img src="http://zerokspot.com/uploads/export_form.png" alt="The "dump"-dialog"/></div>
   6. When you now click on Go a download popup should appear for your backup :-) 

##Restoration procedure
<div class="note">
	
<strong>Note:</strong>
This tutorial is based on phpMyAdmin 2.5.4. The GUI of this tool changes from version to version but the basics have stayed the same for quite some time now.

</div>

1. Open phpMyAdmin with your webbrowser and enter your database account data to access your database.
2. Now you can see a drop down menu in the left frame of phpMyAdmin where you select your database.
   <div class="figure"><img src="http://zerokspot.com/uploads/select_db.png" alt="Select a database"/></div>
3. In the header of the right frame there should be a button "SQL". Click it :-)
   <div class="figure"><img alt="SQL button" src="http://zerokspot.com/uploads/select_sql.png"/></div>
4. Below a text field there should be a smaller field with a "Browse..." button on its right side. When you click it appears a file select popup. Here you select your database backup file and finally click on "Go".
   <div class="figure"><img src="http://zerokspot.com/uploads/query_field.png" alt="Query field"/></div>
5. If there are no error messages that should have done it :-) 

## My database is bigger than 1MB
<div class="note">
	
	<strong>Note:</strong>
1MB is just a number. Every host can have set the upload limit to a different size. Some have 1MB, some 2MB and so on. If you notice any timeout errors or things that somehow look like something like that, this section is for you ;-)

</div>

If your backup file is bigger than 1MB it is possible that it is too big to be uploaded in one piece with phpMyAdmin. The upload limit is different on every host, so better ask your host before making too many chunks of the backup If your file is bigger than the allowed limit you have to split it up into smaller chunks. There are some tools that can do it for you but we want to do it here the hard way

1. Open the backup file in a text editor that supports bigger files. For example VIM.
2. Nearly every line ends with a ;. These are the points where you can split the file. Simply copy/paste everything after a specific ; into another file.
3. Be sure to give those files useful names so that you can insert them in the right order into the database later.