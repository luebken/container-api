# Container-API

A POC of showing the API of a container. It's a simple wrapper around on label convetions and config data from Docker.


## Example:

The container image build from this [Dockerfile](https://github.com/luebken/currentweather/blob/master/Dockerfile) would produce the following output

```
$ container-api luebken/currentweather-nodejs
 -----------------------------------------------------------
| Image:  luebken/currentweather-nodejs
|-----------------------------------------------------------
| Author:   Matthias Luebken, matthias@luebken.com
| Size:     158 MB
| Created:  2016-03-01 15:59
|-----------------------------------------------------------
| Container API:
| * Mandatory ENVs to configure:
|   - ENV:           OPENWEATHERMAP_APIKEY
|   - Description:   APIKEY to access the OpenWeatherMap. Get one at http://openweathermap.org/appid
|   - Mandatory:     true
| * Optional ENVs to configure:
|     - < empty >
| * Available ports:
|     -  1337/tcp {}
| * Volumes:
 -----------------------------------------------------------
```

### Build and use

See the Makefile for building and usage. 

> Note: This is POC state!

### More

More on this topic can be found here: https://github.com/luebken/container-patterns/