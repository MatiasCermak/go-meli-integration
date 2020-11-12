# go-meli-integration
######	 Matias Cermak - Jeremias Fernandez
Repositorio para la materia LAB III de la Universidad Blas Pascal.
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

Tras lo cual, se recibirá en el body un JSON  como el siguiente:

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
Aquí tendremos nuestro Access Token y nuestro UserId, que serán utilizados para las diferentes llamadas posteriores a nuestra Api.
### 3. Endpoints

`/items/all?token=$ACCESS_TOKEN&userid=$USER_ID`
Este endpoint devuelve todos los items con sus respectivas preguntas de un vendedor y las ventas concretadas.
Trayendo un JSON  como el siguiente:

``` [JSON] 
{
	"Items": [
	{
		"Id": "MLA896805185",
		"Title": "Item De Prueba - Por Favor, No Ofertar",
		"Price": 3500,
		"Quantity": 1,
		"SoldQuantity": 0,
		"Picture": "http://http2.mlstatic.com/D_651089-MLA43985576583_112020-O.jpg",
		"Question":	[
		{
			"date_created": "2020-11-05T17:33:18.673-04:00",
			"item_id": "MLA896805185",
			"status": "UNANSWERED",
			"text": "Cuánto sale, boeeeeh",
			"id": 11585807777,
			"answer": ""
		},
		{
			"date_created": "2020-11-11T19:09:28.063-04:00",
			"item_id": "MLA896805185",
			"status": "UNANSWERED",
			"text": "Lo tenes en rojo?",
			"id": 11597339942,
			"answer": ""
		}
				]
},
{
"Id": "MLA896801928",
"Title": "Item De Test - No Ofertar",
"Price": 80000,
"Quantity": 1,
"SoldQuantity": 0,
"Picture": "http://http2.mlstatic.com/D_996968-MLA44005467265_112020-O.jpg",
"Question": [
{
"date_created": "2020-11-11T18:49:22.795-04:00",
"item_id": "MLA896801928",
"status": "UNANSWERED",
"text": "Permutas?",
"id": 11597301371,
"answer": ""
},
{
"date_created": "2020-11-11T18:50:13.728-04:00",
"item_id": "MLA896801928",
"status": "UNANSWERED",
"text": "Vi una publicacion igual mas barata",
"id": 11597300952,
"answer": ""
},
{
"date_created": "2020-11-11T18:51:50.574-04:00",
"item_id": "MLA896801928",
"status": "UNANSWERED",
"text": "Lo tenes en azul?",
"id": 11597306025,
"answer": ""
},
{
"date_created": "2020-11-11T18:52:28.623-04:00",
"item_id": "MLA896801928",
"status": "UNANSWERED",
"text": "Envias a domicilio?",
"id": 11597306462,
"answer": ""
},
{
"date_created": "2020-11-11T19:06:41.032-04:00",
"item_id": "MLA896801928",
"status": "UNANSWERED",
"text": "Envias a Domicilio?",
"id": 11597336065,
"answer": ""
},
{
"date_created": "2020-11-11T19:06:55.255-04:00",
"item_id": "MLA896801928",
"status": "UNANSWERED",
"text": "Aceptas ofertas?",
"id": 11597337179,
"answer": ""
}
]
}
],
"Sales": [
{
"Id": 12227685685,
"Title": "Item De Prueba - Por Favor, No Ofertar",
"Date": "2020-11-10T01:16:16.000-04:00",
"Price": 3500,
"PriceTotal": 4174.99
},
{
"Id": 12249026115,
"Title": "Item De Test - No Ofertar",
"Date": "2020-11-11T18:57:39.000-04:00",
"Price": 80000,
"PriceTotal": 80662.49
},
{
"Id": 12249166880,
"Title": "Item De Prueba - Por Favor, No Ofertar",
"Date": "2020-11-11T19:10:28.000-04:00",
"Price": 3500,
"PriceTotal": 4287.49
}
]
}
```
<!--stackedit_data:
eyJoaXN0b3J5IjpbMTAwMTAzNDMwNCwtNjc5NTc0OTMwLDE2Mj
k1MDA4OTQsMjE0NDMzMTExMCwyMDExMjg3MjUwLDI1OTg3NTAw
N119
-->