# Instance Finder

This is a small tool that will parse terraform tfstate files and present information relating to Instance names and IP addresses associated with them.

i.e.
```
+-----------------+-----------+-----------------+
|      NAME       | PRIVATEIP |    PUBLICIP     |
+-----------------+-----------+-----------------+
| azure-webserver | 10.0.0.24 | 51.234.12.162   |
| azure-centos02  | 10.0.0.68 |                 |
| azure-centos03  | 10.0.0.69 |                 |
| azure-centos01  | 10.0.0.70 |                 |
+-----------------+-----------+-----------------+
```
Currently this is working for AWS and Azure providers