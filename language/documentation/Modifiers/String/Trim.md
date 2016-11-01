Trim Modifier
==================
The Trim Modifier returns the string with all leading and trailing whitespace characters trimmed.

Syntax
--------------
```
Backend:
<?php
  $str = "  Example string  "
?>

Frontend:
{% str %}
{* Outputs:   Example string   *}

{% str trim %}
{* Outputs: Example string *}
```