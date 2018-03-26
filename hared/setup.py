from setuptools import setup

def readme():
    with open('README.rst') as f:
        return f.read()

setup(name='hared',
    version='0.9',
    description='hare daemon',
    long_description=readme(),
    lassifiers=[
      'Development Status :: 3 - Alpha',
      'License :: OSI Approved :: MIT License',
      'Programming Language :: Python :: 2.7',
      'Operating System :: POSIX',
      'Topic :: Communications',
      'Topic :: Internet',
    ],
    keywords='UDP MQTT daemon PAM login SSH',
    url='https://jpmens.net/2018/03/25/alerting-on-ssh-logins/',
    author='Jan-Piet Mens',
    author_email='jp@mens.de',
    license='MIT',
    include_package_data=True,
    packages=['hared'],
    scripts=['bin/hared'],
    install_requires=[
        'paho-mqtt',
    ],
    zip_safe=False)
