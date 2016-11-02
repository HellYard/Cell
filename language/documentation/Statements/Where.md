Where Statement
==================
The Where Statement let's you perform checks in for statements without adding additional code inside the for loop.

Syntax
--------------
```
Backend:
<?php
   $arr = [
      0 => [
        "one",
        "two"
      ]
      1 => [
        "three",
        "four"
      ]
   ];
?>

Frontend:
{% for i in arr where i.0 is "one" %}
{% i.1 %}
{% endfor %}

{* Output: two *}
```