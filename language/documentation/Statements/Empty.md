Empty Statement
==================
The Empty Statement lets you define a value that should be outputted when the array being iterated over is empty.

Syntax
--------------
```
Backend:
<?php
   $arr = [];
?>

Frontend:
{% for i in arr %}
{% i.0 %}
{% empty "No items in array." %}
{% endfor %}

{* Output: No items in array. *}
```