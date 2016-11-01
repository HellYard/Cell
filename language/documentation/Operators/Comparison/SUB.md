SUB Operator
==================
The SUB Operator determines whether the class on the left is a child of that on the right hand side of the expression.
The latter behaviour also works with a class and string, where the string is the name of a class.

Syntax
--------------
```
Backend:
<?php
  /*
   * in this cat extends animal
   */
  $cat = new Cat("Cat");
  $animal = new Animal("Generic");
?>

Frontend:

{% cat sub animal %}
{* Outputs: true *}

{% cat sub Cat %}
{* Outputs: false *}
```