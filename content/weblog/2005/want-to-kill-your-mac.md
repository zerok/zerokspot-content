---
date: '2005-11-30T12:00:00-00:00'
language: en
tags:
title: Want to kill your Mac?
---


Today I've found a quite simple way of killing a running Mac system ;) Actually it doesn't really do anything all that great but simply prevents the whole system from working without a nice hard-reboot :) If you have the system firewall deactivated, never add a ipfw rule that denies connections via loopback (or to 127.0.0.0/8, am not sure here). I had the system firewall accidentally disabled and then ran one of my scripts to set some quite restrictive drop rules after the the initial rules normally set by the system firewall. Since these rules where missing now, I couldn't open any new windows, terminals, etc. Not even going to the standby mode worked anymore ;)

-------------------------------



Is there a shortcut for rebooting the whole system that makes it a little less hard then simply holding the power button for a few seconds? ;)