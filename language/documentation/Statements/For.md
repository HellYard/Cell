For Statement
==================
The For Statement lets you iterate over the values of an array.

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
{% for i in arr %}
{% i.0 %}
{% endfor %}

{* Output: onethree *}
```