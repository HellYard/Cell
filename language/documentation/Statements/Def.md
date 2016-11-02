Def Statement
==================
The Def Statement lets your create a temporary variable that'll be removed after the code block has been executed.

Syntax
--------------
```
Home.cell
{% if true %}
{% def temp as 'Hello World' %}
{% temp %}
{% endif %}
{* Output: Hello World *}
{% temp %}
{* Output: null *}
```