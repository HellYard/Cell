Pad Modifier
==================
The Pad Modifier pads a numerical value with placeholder zeros. If you pad a double or float value that does not include
a decimal point, padding it will automatically place a decimal character in the opposite direction of that which is being
padded.

Arguments
--------------
Amount - Integer Value - Default: 1
Direction - Left/Right - Default: Left

Syntax
--------------
```
Backend:
<?php
  $num = 1;
?>

Frontend:
{% num pad %}
{* Outputs: 01 *}

{% num pad=2,right %}
{* Outputs: 100 *}
```