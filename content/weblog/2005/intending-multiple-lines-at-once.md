---
date: '2005-02-24T12:00:00-00:00'
language: en
tags:
- vim
title: Intending multiple lines at once
---


I just noticed something quite funny. I often have to intend multiple lines (blocks, whatever) of code because I added for example an if-condition somewhere in my code or something like that. Now you could intend every single line manually which is quite a big pain when the new block is quite long. 

-------------------------------



VIM has a very nice alternative: Simply select the block you want to intend in the visual mode (hit "v" then select the lines with your cursor keys) and then hit ">" to intend the whole selection :-) This also works the other way around with "<" ;-)