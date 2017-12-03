IPCache Variable
==================
The IPCache variable is identical to the Cache variable except it will cache parsed files on a ip basis. This means that
the cached data may be different for each ip that visits the site, which is ideal for sites with account systems that
wish to utilize the cache system.

Syntax
--------------
```
{{ ipcache [age(in minutes)] }}
```

Example
--------------
```
{* We want to cache a page that prints a user's IP, so using the normal cache variable will show one IP to every user.
   To fix this, we need to use the ipcache variable. In this example, we have two visitors with the IPs of
   161.228.194.230 and 194.218.212.39. We also have a defined variable of ip, which is the user's ip address. *}
{{ ipcache 300 }}
<p>Welcome unknown user with IP of: {% ip %}</p>
{* This will save the parsed content of index.cell to a file called index.pcell, and indexed by the user's ip. The next
   time the user's visit the site they'll see a cached version of the site with their own ip displayed. *}
```