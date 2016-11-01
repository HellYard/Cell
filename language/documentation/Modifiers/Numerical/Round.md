Round Modifier
==================
The Round Modifier rounds a numerical value. If used on a non-precision values the variable is rounded to the
nearest five, otherwise it's rounded to the nearest tenths place by default.

Arguments
--------------
Places from decimal - 1 - Default: 1(tenths)

Syntax
--------------
```
Backend:
<?php
  $num = 13;
  $numPrecise = 1.5345114;
?>

Frontend:
{% num round %}
{* Outputs: 15 *}

{% numPrecise round=2 %}
{* Outputs: 1.53 *}
```