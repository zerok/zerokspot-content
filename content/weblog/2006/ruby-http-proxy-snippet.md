---
date: '2006-06-17T12:00:00-00:00'
language: en
tags:
- http
- ruby
- snippets
title: Ruby HTTP Proxy snippet
---


Just a small pattern is use all the time for writing HTTP requests that work

-------------------------------

with an without a proxy. 



<pre class="code">#!/usr/bin/env ruby

require &apos;net/http&apos;

require &apos;ostruct&apos;



proxy = ENV[&apos;http_proxy&apos;] ? URI.parse(ENV[&apos;http_proxy&apos;]) : OpenStruct.new

Net::HTTP::Proxy(proxy.host,proxy.port,proxy.user,proxy.password).start(&apos;www.apple.com&apos;) do |http|

  puts http.get(&apos;/&apos;).body

end</pre>



I use here the OpenStruct class which actually works like a more generic Struct where fields can added at will. Comes quite handy since You normally want to have "host", "port", "username" and "password" accessors for your proxy URL and Net::HTTP::Proxy is kind enough to accept nil values if they are not required :)