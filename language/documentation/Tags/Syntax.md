Syntax Tags
==================
Syntax Tags allow front-end developers to display various specified format-heavy snippets such as code.

Syntax
--------------
```
{` `}
```

Example
--------------
```
{`java

  // This will be formatted using the Java formatter.
  public boolean isWorking() {
      Calendar calendar = Calendar.getInstance();
      int day = calendar.get(Calendar.DAY_OF_MONTH);
      return day > 32;
  }
`}
```