# Parrot

![Parrot](resources/Parrot.ico)

## Application to store and search minecraft coordinates

Parrot attempts to allow users on a minecraft server to save interesting locations. The inspiration comes from finding unique landscapes scattered through the procedurally generated world. Originally I wrote down the locations of these coordinates on a sticky note and then shared them through chat to fellow server players. This application avoids the desk full of chicken scratch sticky notes.

The application consists of three main services: 
[plugin](#plugin), 
[webapp](#webapp), 
and [api](#api)

## Services

### Plugin
-----
The plugin allows the user to mark and recall interesting locations while on the server. 

### Webapp
-----
The webapp allows user to view and add coordinates from a web interface.

### API
-----
The backend api is responsible for storing the coordinates. The api provides a REST interface to allow data transactions from the front end web application and minecraft plugin. The supported operations are documented in the [open api specification](api/openapi.yml) file

**Build and run**

In _api/src_

`go run index.go`

**Port**

The default port is set to `8080`