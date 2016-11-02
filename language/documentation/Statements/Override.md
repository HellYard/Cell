Override Statement
==================
The Override Statement lets your override parts of a template file that you inherit from.

Syntax
--------------
```
Index.cell
<p class="test">Hello</p>
<p class="test2">World</p>

Home.cell
{% copy Index.cell %}
{% override %}
<p class="test">Bye</p>
{* Output: <p>Bye</p><p>World</p> *}
```