Subtraction Operator
==================
The Subtraction Operator subtracts one value from another. When using the subtraction operator on two strings,
it'll subtract the string on the right hand on the expression from the one on the left. However, when the operator
is used on a string and integral value the result will be the string with the character places(integral value) removed
from the end.

Syntax
--------------
```
{% 5 - 2 %}
{* Outputs: 3 *}

{% "Hello" - "lo" %}
{* Outputs: Hel *}

{% "Hello" - 1 %}
{* Outputs: Hell *}
```