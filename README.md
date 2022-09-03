## Apirest en Golang
###  Mvp
##### Requisitos Alcanzados
-Los eventos podrán gestionarse (alta, baja, modificación) sólo por un rol
administrador.
-Los eventos cuentan con un título, descripción corta, descripción larga, fecha y
hora, organizador, lugar, y un estado (borrador o publicada).
-Cuando una publicación tiene estado borrador, sólo los administradores pueden
visualizarla.
-Los usuarios (administradores o no) pueden visualizar los eventos activos
(publicados y con fecha-hora en el futuro) como así los eventos completados
(publicados pero con fecha y hora en el pasado). Este listado de eventos se debe
poder filtrar por fecha, estado, y título.
-A su vez, se contará con un endpoint en el cual se mostrarán todos los eventos
inscriptos, filtrándose por activos o completados
##### Requisitos no alcanzados
-Los usuarios pueden visualizar e inscribirse en eventos publicados, siempre y
cuando la fecha del evento sea futura.
-Los eventos cuya fecha y hora ya hayan transcurrido, pueden visualizarse pero no
inscribirse.
### Dependencias del proyecto
version de golang:  go1.19
##### Como puedo usar el proyecto
git clone https://github.com/AgustinBrandan/ApirestGolang.git
##### Iniciar el servidor
go run .\api.go
### ENDPOINT
https://github.com/AgustinBrandan/ApirestGolang/tree/main/imagenes
En las imagenes encontraremos las pruebas que hice en Postman 
#### Admin
###### POST /admin/login 
Se necesita iniciar sesion para ser admin,se crea un token que nos da permisos por un tiempo editar,eliminar y crear Eventos.
###### GET /admin/eventos/ 
Una vez logeado el admin tiene permiso para ver todo los eventos.
###### POST /admin/eventos
El admin tiene los permisos para crear un nuevo evento.
###### DELETE /admin/eventos/:titulo
El admin puede borrar un evento con el titulo del evento.
###### PUT /admin/eventos/:titulo 
El admin puede cambiar un registro a traves del titulo del evento.
###### Get /admin/eventos/estado/:estado
El admin puede ver los eventos por estado
#### Usuarios
###### GET /user/eventos
El usuario solo ve los eventos que estan en publicado
###### GET /user/eventos/:titulo
El usuario puede filtrar los eventos por titulo
##### GET /user/eventos/fecha/:fecha
El usuario puede filtrar los eventos por fecha




