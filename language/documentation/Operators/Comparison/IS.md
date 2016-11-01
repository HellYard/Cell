IS Operator
==================
The IS Operator determines whether two values are equal to each other, and returns true or false accordingly. If the
operator is used on two classes, the value returned will be whether or not the class on the left is an instance of that
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
{% value is true %}
{* Outputs: false *}

{% cat is animal %}
{* Outputs: true *}

{% cat is Animal %}
{* Outputs: true *}
```