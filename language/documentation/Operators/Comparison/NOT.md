NOT Operator
==================
The NOT Operator determines whether two values are not equal to each other, and returns true or false accordingly. If the
operator is used on two classes, the value returned will be whether the class on the left is not an instance of that
on the right hand side of the expression. The latter behaviour also works with a class and string, where the string is
the name of a class.

Syntax
--------------
```
Backend:
<?php
  $value = false;
  $cat = new Animal("Cat");
  $animal = new Animal("Generic");
?>

Frontend:
{% value not true %}
{* Outputs: true *}

{% cat not animal %}
{* Outputs: false *}

{% cat not Mammal %}
{* Outputs: true *}
```