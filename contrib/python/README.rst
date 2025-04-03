Hared
-----

To use (with caution), simply run::

    hared

The program reads a configuration file at ``/usr/local/etc/hared.ini``
or one specified as ``$HARED_INI``. This INI file looks like this with
the following defaults:

::


    [defaults]
    verbose = False
    listenhost = localhost
    listenport = 8035
    mqtthost = 127.0.0.1
    mqttport = 1883
    topic = logging/hare

FreeBSD:

::

    pkg install py36-pip

You are encouraged to use our Go version of ``hared``
