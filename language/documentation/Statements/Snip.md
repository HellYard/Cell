Snip Statement
==================
The Snip Statement lets you add a template block to your snippets that you may include in your template files later on.
The snippet is named with the value after the snip statement.

Syntax
 --------------
 ```
 Index.cell
 <p>Hello</p>
 {% snip test %}
 <p>test</p>
 {% endsnip %}

 Home.cell
 {% insert Index.cell %}
 <p>World</p>
 {% insert test %}

 {* Output:
  * <p>Hello</p>
  * <p>World</p>
  * <p>test</p>
  *}
 ```