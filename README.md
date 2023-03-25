## **SFILE DOCUMENT**
### basic function,[**official site**](https://brotherhoodhk.org/codelab/sfile)
* **put file into filesystem(if file exists in filesystem,it will update)**
    ```
    sfile add filename
    ```
* **remove file from virtual filesystem**
    ```
    sfile rm filename
    ```
* **pull file from virtual filesystem**
    ```
    sfile get filename
    ```
* **list all files in virtual filesystem**
    ```
    sfile list
    ```
* **clear local virtual filesystem**
    ```
    sfile clear
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
### other function
* **search file in remote filesystem**
    ```shell
    sfile search filename
    ```
* **configure your remote information**
    ```shell
    sfile config user:authkey@hostadd
    #example
    sfile config jake:jake123@brotherhoodhk.org:2999
    ```
* **show your remote information**
    ```shell
    sfile config show remoteinfo
    ```