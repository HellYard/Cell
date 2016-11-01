Floor Modifier
==================
The Floor Modifier transforms a numerical(double, or float) variable into the largest integer that is less than or
equal to the variable's value.

Syntax
--------------
```
Backend:
<?php
  $num = 1.5345114;
?>

Frontend:
{% num floor %}
{* Outputs: 1.0 *}
```