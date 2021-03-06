---
date: '2005-02-12T12:00:00-00:00'
language: en
tags:
- csharp
title: Marking methods (etc.) as deprecated in CSharp/Mono
---


I think everyone who has ever coded knows or at least has seen this problem: You write a function/method and the after after that you begin to code another version of it and want to somehow let the world know, that yesterdays version should no longer be used. Right, I'm talking about deprecating stuff in your program.

-------------------------------


But how do you want to let everyone know about this change in your API? You could simply write it into the documentation which is like a friendly and soft reminder. But an even better and stronger way would be, to let the compiler or interpreter of your program give a warning if a deprecated function is called, a deprecated classobject is generated.

Some languages out there support this kind of strong but still gentle reminder: So far I've seen examples in Java 1.5 and C#/Mono. I would have expected C and C++ having similiar features but I could only find the documentation way of deprecating stuff. If someone knows of ways to mark methods etc. as deprecated in any other language, please let me know :-)

Here I just want to show a little example about how it works in C#/Mono. In C# you can annotate everything using predefined or custom attributes. The attributes are added in brackets in front of the rest of the declaration:

<pre class="code">
[ATTRIBUTE] public void Main();
</pre>

There are quite a few predefined attributes and attribute classes but only one is important for us right now: Obsolete. I've written a small sample code that sets to components as obsolete: A class variable and a class method.

<pre class="code">
using System;
public class Test{
	[Obsolete]static String TestString = "Hello";
	[Obsolete]public static void Whatever(){
		Console.WriteLine("{0} world",Test.TestString);
	}
	public static void Main(){
		Test.Whatever();
	}
}
</pre>

When compiling this piece of code, you will get warnings because of the obsolete state of these components:

<pre class="output">
zerok@intrepid:~/tmp $ mcs Test.cs
Test.cs(3) warning CS0612: 'Test.TestString' is obsolete
Test.cs(5) warning CS0612: 'Test.TestString' is obsolete
Test.cs(8) warning CS0612: 'Test.Whatever()' is obsolete
Compilation succeeded - 3 warning(s)
</pre>

Quite handy :-)

Again: If you know of ways to mark stuff deprecated in other languages please let me know :-)

[UPDATE] Oops, completely forgot to post some links:
<ul>
<li><a href="http://www.go-mono.org/monkeyguide/html/en/languages/attributes-basics.html">Monkeyguide on go-mono.org</a></li>
<li><a href="http://www.go-mono.com/docs/">Docs on go-mono.com</a></li>
</ul>