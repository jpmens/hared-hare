# hared-hare

This is the C program (`hare`) and the Python daemon (`hared`) for [this story](https://jpmens.net/2018/03/25/alerting-on-ssh-logins/).

_hare_ is a small utility which is installed in a PAM configuration (e.g. for `sshd`) in order to log when a successful login is attempted, e.g. to alert on machines which are seldom visited or otherwise monitored.

_hare_ transmits a JSON string over a UDP datagram. The JSON looks like this:

```json
{
  "user": "jjolie",
  "rhost": "that.host",
  "service": "sshd",
  "hostname": "tiggr.ww.mens.de",
  "tst": 1522080746
}
```

The values for `user`, `rhost`, and `service` are set by PAM from their `PAM_` equivalents, and `hostname` will contain the _gethostname(3)_ result as determined by _hare_.

Python `hared` is also installable via [https://pypi.python.org/pypi/hared/](https://pypi.python.org/pypi/hared/)
