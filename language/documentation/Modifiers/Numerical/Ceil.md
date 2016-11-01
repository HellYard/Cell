Ceil Modifier
==================
The Ceil Modifier transforms a numerical precision value into the smallest integer that is greater than or
equal to the value being modified.

Syntax
--------------
```
Backend:
<?php
  $num = 1.5345114;
?>

Frontend:
{% num ceil %}
{* Outputs: 2.0 *}
```