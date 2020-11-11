# go-meli-integration
Repositorio para la materia LAB III de la Universidad Blas Pascal
## Instrucciones de uso
### 1. Descarga de Api
Para descargar la Api se necesitara correr el siguiente comando en tu terminal de Go:
`go get -u github.com/MatiasCermak/go-meli-integration`
Con esto se descargara la Api en tu dispositivo.
### 2.  Obtener Access Token
Para obtenerlo se utiliza la siguiente llamada:
`https://auth.mercadolibre.com.ar/authorization?response_type=code&client_id=3589105139307129&redirect_uri=http://localhost:8080/auth`

Al ejecutarse la llamada, se recibirá un JSON  como el siguiente:

 [JavaScript]
{"Access_token":"APP_USR-3589105139307129-111122-5aa4c31191b3b9b761a94f36ade5675c-667732467","Token_type":"bearer","Expires_in":21600,"Scope":"offline_access read write","User_id":667732467,"Refresh_token":"TG-5fac6cd21c19cc00060ee065-667732467"}

<!--stackedit_data:
eyJoaXN0b3J5IjpbMTE4OTYwODQxOCwyMDExMjg3MjUwLDI1OT
g3NTAwN119
-->