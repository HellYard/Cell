Floor Modifier
==================
The Floor Modifier transforms a numerical precision value into the largest integer that is less than or
equal to the value being modified.

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