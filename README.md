
Commands:
---------

`LISTTITLE` is full list title.

`LISTNAME` is prefix of it (case-insensitive)

  - Show list of lists:
    ```
    $ wundercli list all
    ```

  - Show list items:
    ```
    $ wundercli list show [LISTNAME]
    ```

  - Create new list:
    ```
    $ wundercli list create [LISTTITLE]
    ```

  - Remove the list:
    ```
    $ wundercli list remove [LISTNAME]
    ```

  - Add task to a list:
    ```
    $ wundercli task create [LISTNAME [ITEMTEXT]]
    ```

  - Mark task checked:
    ```
    $ wundercli task check [LISTNAME]
    ```

  - Edit task:
    ```
    $ wundercli task edit [LISTNAME]
    ```

