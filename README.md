# go-meli-integration
Repositorio para la materia LAB III de la Universidad Blas Pascal
## Instrucciones de uso
### 1. Descarga de Api
Para descargar la Api se necesitara correr el siguiente comando en tu terminal de Go:
`go get -u github.com/MatiasCermak/go-meli-integration`
Con esto se descargara la Api en tu dispositivo.
### 2.  Autenticación
Para obtenerlo se utiliza la siguiente llamada:
`https://auth.mercadolibre.com.ar/authorization?response_type=code&client_id=3589105139307129&redirect_uri=http://localhost:8080/auth`

Al ejecutarse la llamada, se te redireccionará a la siguiente url:
`http://localhost:8080/auth?code=TG-5fac6cd21c19cc00060ee064-667732467`
donde el parámetro "code" representa al código de autorización enviado por Mercado Libre.

se recibirá un JSON  como el siguiente:

``` [JSON] 
{
	"Access_token":"APP_USR-3589105139307129-111122-5aa4c31191b3b9b761a94f36ade5675c-667732467",
	"Token_type":"bearer",
	"Expires_in":21600,
	"Scope":"offline_access read write",
	"User_id":667732467,
	"Refresh_token":"TG-5fac6cd21c19cc00060ee065-667732467"
}
```

### 3. Endpoints


<!--stackedit_data:
eyJoaXN0b3J5IjpbMTQ1NzI1MzkxMywyMDExMjg3MjUwLDI1OT
g3NTAwN119
-->