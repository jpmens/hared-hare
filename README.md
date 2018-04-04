# hared-hare

This is the C program (`hare`) and the Python daemon (`hared`) for [this story](https://jpmens.net/2018/03/25/alerting-on-ssh-logins/).

_hare_ is a small utility which is installed in a PAM configuration (e.g. for `sshd`) in order to log when a successful login is attempted, e.g. to alert on machines which are seldom visited or otherwise monitored.

_hare_ transmits a JSON string over a UDP datagram. The JSON looks like this:

```json
{
  "tty": "tty1",
  "service": "login",
  "hostname": "zabb01",
  "user": "jjolie",
  "tst": 1522154553,
  "rhost": "<unknown>",
  "remote" : "10.0.12.1"
}
```

The values for `user`, `rhost`, `tty`, and `service` are set from PAM from their `PAM_` equivalents, and `hostname` will contain the _gethostname(3)_ result as determined by _hare_. `remote` is the IP address of the _hare_ client as seen by _hared_.

Python `hared` is also installable via [https://pypi.python.org/pypi/hared/](https://pypi.python.org/pypi/hared/)

## More

* [hare](https://svnweb.freebsd.org/ports/head/sysutils/hare/) for FreeBSD
* [hared](https://svnweb.freebsd.org/ports/head/sysutils/hared/) for FreeBSD
* [py-hared](https://svnweb.freebsd.org/ports/head/sysutils/py-hared/) for FreeBSD
