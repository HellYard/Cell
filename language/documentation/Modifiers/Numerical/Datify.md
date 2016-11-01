Datify Modifier
==================
The Datify Modifier transforms an integer that servers as the day of a month into an integer with the proper suffix.

Syntax
--------------
```
Backend:
<?php
  $day = 16;
  $dayTwo = 21
?>

Frontend:
{% day datify %}
{* Outputs: 16th *}

{% dayTwo datify %}
{* Outputs: 21st *}
```