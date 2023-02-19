## **SFILE DOCUMENT**
### basic function
* **put file into filesystem(if file exists in filesystem,it will update)**
    ```
    sfile add filename
    ```
* **remove file from filesystem**
    ```
    sfile rm filename
    ```
* **pull file from filesystem**
    ```
    sfile get filename
    ```
* **list all files in filesystem**
    ```
    sfile list
    ```
### remote function
* **upload a single file to host**
    ```
    sfile upload filename
    sfile upload --private dirname/filename
    ```
* **get a single file from host**
    ```
    sfile pull filename
    sfile pull --private dirname/filename
    ```
* **remove a single file from host**
    ```
    sfile clean filename
    sfile clean --private dirname/filename
    ```
* **remove host's private dir**
    ```
    sfile clean -r dirname
    ```
* **make private dir**
    ```
    sfile mkdir dirname
    ```