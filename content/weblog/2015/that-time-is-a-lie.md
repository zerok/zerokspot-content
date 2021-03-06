---
date: '2015-02-15T10:49:50+01:00'
language: en
tags:
- python
title: That Time is a Lie
---


While skimming through the [release notes][1] for the upcoming Python 3.5
version I noticed a reference to a bug in the datetime module. Since I use that
basically whenever I do *anything* in Python I got curious.

Turns out that up until 3.5 there was a `datetime.time` instance that got
evaluated to False in a boolean context ([issue13936][2]):

```
> bool(datetime.time(0, 0, 0))
False
```

So basically whenever you have a time instance, make sure to compare it against
another one of None, but never do the implicit boolean check!

While this behaviour has been [documented][3] ...

> in Boolean contexts, a time object is considered to be true if and only if,
> after converting it to minutes and subtracting utcoffset() (or 0 if that’s
> None), the result is non-zero.

... that doesn't change that it is highly un-intuitive. The reasons for this
implementation have been laid out in the issue but I'm really glad that with 3.5
all time instance will evaluate to true!

[1]: https://docs.python.org/dev/whatsnew/3.5.html

[2]: https://bugs.python.org/issue13936

[3]: https://docs.python.org/2/library/datetime.html#datetime-time
