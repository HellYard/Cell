Escape Modifier
==================
The Escape Modifier escapes a string variable type to be query safe.

Syntax
--------------
```
Backend:
<?php
  $str = "Example string"
?>

Frontend:
{% str escape %}
{* Outputs: Example string *}
```