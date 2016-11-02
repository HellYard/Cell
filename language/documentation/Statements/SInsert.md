SInsert Statement
==================
The SInsert Statement lets your include external template files, and snippets in your template file, without letting the
included item perform various statements, operators, or modifiers.

Syntax
 --------------
 ```
 Index.cell
 <p>test</p>
 {% throw "Unwanted exception!" %}

 Home.cell
 <p>Hello</p>
 {% sinsert Index.cell %}
 {* Output:
  * <p>Hello</p>
  * <p>test</p>
  * {% throw "Unwanted exception!" %}
  *}
 ```