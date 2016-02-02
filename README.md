
Commands:
---------

`LISTTITLE` is full list title.

`LISTNAME` is prefix of it (case-insensitive)

  - Show list of lists:
    ```
    $ wundercli list
    ```

  - Show list items:
    ```
    $ wundercli show [LISTNAME]
    ```

  - Create new list:
    ```
    $ wundercli create [LISTTITLE]
    ```

  - Remove the list:
    ```
    $ wundercli remove [LISTNAME]
    ```

  - Add item to a list:
    ```
    $ wundercli append [LISTNAME [ITEMTEXT]]
    ```

  - Mark item checked:
    ```
    $ wundercli check [LISTNAME]
    ```

  - Edit item:
    ```
    $ wundercli edit [LISTNAME]
    ```

