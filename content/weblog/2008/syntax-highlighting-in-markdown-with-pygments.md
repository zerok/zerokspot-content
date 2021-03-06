---
date: '2008-06-18T12:00:00-00:00'
language: en
tags:
- markdown
- pygments
- python
title: Syntax Highlighting in Markdown with Pygments
---


If you want to get syntax highlighting using Pygments in Markdown texts, you have more or less 2 options. For one you can work on the HTML output of the Markdown2HTML converter of your choice, or you can try extending Markdown to offer special syntax for stuff you want to pass over to Pygments.

A good idea if you want to go the first route is the use of some kind of HTML parsing library like BeautifulSoup as shown in [this snippet](http://www.djangosnippets.org/snippets/360/) and [this article](http://www.unessa.net/en/hoyci/2006/11/highlighting-code-using-pygments-and-beautiful-soup/).

This post will focus on the 2nd approach, though, since it is a bit cleaner :-)

-------------------------------

Both options rely heavily on what Markdown library you're using. At least some of them will have problems when you mix Markdown with plain old HTML elements to separate your source code or produce some other not really expected results. This problem is one of the main reasons why it's probably better to extend the library you're using to offer code highlighting. 

With [Python-Markdown](http://www.freewisdom.org/projects/python-markdown/) this is pretty straight forward and Pygments itself comes with a [small preprocessor](http://dev.pocoo.org/projects/pygments/browser/external/markdown-processor.py?rev=425%3A9bcc2cff451b) that's supposed to be all you need to integrate Pygments with Markdown. The problem is, though, that the preprocessor comes a little bit too late in the whole conversion process (namely *after* Markdown temporarily cleaned up the text and removed all HTML to insert it later again, if I understood it correctly). So whenever you want to write an extension that produces HTML itself, you should normally want it to be executed *before* that whole HTML-removal happens, in order to keep it unaffected by the rest of MD2HTML processing. So if you look at how Python-Markdown's extension system works, you don't want to write a Preprocessor (which is executed after the input was split into lines and the whole HTML removal took place), but a TextPreprocessor which operates on the original line structure instead of a list of lines.

If you don't do it, I most of the time ended up with additional paragraphs wrapped around my code snippets that I absolutely wanted to get rid of.

The code I'm using right now and which is heavily based on the markdown-processor bundled with Pygments looks like this:

<pre><code class="language-python">
from markdown import TextPreProcessor
from pygments.formatters import HtmlFormatter
from pygments.lexers import get_lexer_by_name, TextLexer
from pygments import highlight
import re

class CodeBlockPreprocessor(TextPreprocessor):
    pattern = re.compile(
        r'^\s*@@ (.+?) @@\s*(.+?)^\s*@@', re.M|re.S)

    formatter = HtmlFormatter(cssclass="codehl")

    def run(self, lines):
        def repl(m):
            try:
                lexer = get_lexer_by_name(m.group(1))
            except ValueError:
                lexer = TextLexer()
            code = m.group(2).replace('\t','    ')
            code = pygments.highlight(code, lexer, self.formatter)
            code = code.replace('\n\n', '\n&amp;nbsp;\n').replace('\n', '&lt;br /&gt;').replace('\\@','@')
            return '\n\n%s\n\n' % code
        return self.pattern.sub(repl, lines)
</code></pre>

As you can see, it only has some minor changes like the different parent class, a slightly different regular expression for the wrapping Markdown syntax, the extra replacement for new lines and tabs (I hate tabs :-P) and the extra cssclass attribute. 

To use this class, put it into a script like this:

<pre><code class="language-python">
from markdown import Markdown

someText = """
@\@ python @@
def test():
    print "lala"
@\@
"""

md = Markdown()
md.textPreprocessors.insert(0, CodeBlockPreprocessor())
print md.convert(someText)
</code></pre>



**Update:** Most of the changes mentioned here are now also [part of the markdown-processor.py](http://dev.pocoo.org/projects/pygments/changeset/6286fd13ce24) module in Pygments.
