# Container-API

A POC of showing the API of a container. It's a simple wrapper around on label convetions and config data from Docker.


## Example:

The container image build from this [Dockerfile](https://github.com/luebken/currentweather/blob/master/Dockerfile) would produce the following output

```
$ container-api luebken/currentweather-nodejs
 -----------------------------------------------------------
| Image:  luebken/currentweather-nodejs:latest
|-----------------------------------------------------------
| Author:   matthias.luebken@gmail.com
| Size:     158 MB
| Created:  2016-03-04 11:24
|-----------------------------------------------------------
| Container API:
|
| * Expected Links:
|   - Name:          redis:latest
|   - Port:          1337
|   - Description:   Needed for requests caching
|   - Mandatory:     true
|
| * Expected ENVs:
|   - ENV:           OPENWEATHERMAP_APIKEY
|   - Description:   APIKEY to access the OpenWeatherMap. Get one at http://openweathermap.org/appid
|   - Mandatory:     true
|
| * Expected args:
|   - Arg:           -q QUERY
|   - Description:   The query for openweathermap.
|
| * Available ports:
|     -  1337/tcp {}
|
| * Volumes:
 -----------------------------------------------------------

```

### Build and use

See the Makefile for building and usage. 

> Note: This is POC state!

### More

More on this topic can be found here: https://github.com/luebken/container-patterns/