---
date: '2013-07-10T22:27:22+02:00'
language: en
tags:
- python
- europython
- ep2013
- pytest
- testing
title: pytest is awesome
---


[pytest][] has been on my radar for quite a while and I tentatively used it in some project as a simple replacement for Python's built-in unittest-module mostly because of awesome assertion output and because I hate to remember 100 different assertXYZ methods. But it became much more during the last couple of days.

Last Saturday, as part of the sprints at [EuroPython 2013][ep2013], [Holger Krekel][hk] gave those of us new to the framework a set of slides and helped us to get to know it better. Big thanks for that! Since then I've become convinced that py.test is basically the only testing framework I want to use for all my future Python projects (well, until something even more awesome comes around).

--------

As I said, I've only used the basics before and never got to stuff like fixtures and markers, but since last Saturday I have to wonder, how no other framework I've found so far has the -m flag (for selecting tests to run based on the markers they've got) and the -k flag for generally selecting tests by name.
    
    > cat test_markers.py
    import pytest


    @pytest.mark.integration
    def test_a():
        assert 1 == 1


    @pytest.mark.unit
    def test_b():
        assert 1 == 1

    > py.test -m "unit" --collectonly
    === test session starts =======================
    platform darwin -- Python 2.7.4 -- pytest-2.3.5
    plugins: cov
    collected 2 items
    <Module 'test_markers.py'>
      <Function 'test_b'>

    === 1 tests deselected by "-m 'unit'" =========
    === 1 deselected in 0.01 seconds ==============

As you can see in the example about, once test has been marked as "unit", you can select it with `-m "unit"`.

Also, injecting preconditions and helpers with fixtures just feels so much more natural than "beforeEach"/"setUp"-style methods.

<pre><code class="language-python">
@pytest.fixture
def mock_service():
    # ...
    return service

def test_feature_a(mock_service):
    assert mock_service.do_something() == 'as_expected'
</code></pre>


Digging into this stuff and also the ["scope"-setting][scope] for fixtures was tons of fun and I really miss all these little things whenever I have to work with other testing frameworks (as much as I've come to love [mocha][mo] + [chai][] on the JavaScript side of things).

Next, I will probably try to learn more about how to integrate pytest better with Django using [pytest-django][pdj], which was also demonstrated at this year's EuroPython in Florence by [Andreas Pelme][ap].

[pytest]: http://pytest.org/latest/
[hk]: http://holgerkrekel.net/
[ep2013]: https://ep2013.europython.eu/
[ap]: https://twitter.com/andreaspelme
[pdj]: https://pypi.python.org/pypi/pytest-django
[scope]: http://pytest.org/latest/fixture.html#working-with-a-module-shared-fixture
[mo]: http://visionmedia.github.io/mocha/
[chai]: http://chaijs.com
