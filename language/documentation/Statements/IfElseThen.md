If-Else-Then Statement
==================
The If-Else-Then Statement lets you perform conditional checks in a structured order. If first condition is false, then
each elseif condition is evaluated in order. In the event all elseif conditions result, and an else condition is present,
the else condition is ran.

Syntax
--------------
```
{% if false %}
Hi
{% elseif 1 > 2 %}
Yo
{% else %}
Hello
{% endif %}

{* Output: Hello *}
```