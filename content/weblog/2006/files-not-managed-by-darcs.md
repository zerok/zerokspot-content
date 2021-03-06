---
date: '2006-05-22T12:00:00-00:00'
language: en
tags:
- darcs
- ruby
title: Files not managed by Darcs
---


Want to know what files are not included in a darcs repository? Well, me too :) From what I can tell, darcs itself doesn't offer something like that and if it does, it was at least a good excuse for some Ruby'ing ;)



-------------------------------



<pre class="code">
#!/usr/bin/env ruby
require &apos;find&apos;
require &apos;pathname&apos;
require &apos;fileutils&apos;

# Define what should not be included in the listing
IGN_EXTENSIONS = %w{.aux .log .toc .blg .bbl .lof .out .o .pyc .class .bak}
IGN_LEAVES = %w{.svn _darcs .DS_Store}
def ignore?(file)
  return true if IGN_LEAVES.include?(File.basename(file))
  return true if IGN_EXTENSIONS.include?(File.extname(file))
end

# Get a list of all managed files
versioned_files = []
`darcs query manifest`.each_line{|f| versioned_files &lt;&lt; f.strip}
exit(2) unless $?.success?

# Go to the folder with the _darcs folder in it
orig_folder = File.expand_path(&apos;.&apos;)
while !File.exists?(&apos;_darcs&apos;)
  cur_folder = File.expand_path(&apos;.&apos;)
  raise &quot;You aren&apos;t in a darcs managed folder&quot; if cur_folder==&apos;/&apos;
  FileUtils.cd(File.dirname(cur_folder))
end

# Do the comparision
Find.find(&apos;.&apos;) do |f|
  next if f == &apos;.&apos;
  Find.prune if ignore?(f)
	puts f unless versioned_files.include?(f)
end

# Go back to where you&apos;ve come from
FileUtils.cd(orig_folder)
</pre>