# ApiGen
Permite generar datos de prueba de manera infinita y procedural. Cuenta con una variedad de *funciones* que permiten darnos datos generados según la necesidad que tengas.

## Uso
Su uso es tan sencillo como mandar peticiones comunes y por medio del *body* darle la estructura de la respuesta con las funciones que quieres que utilize.

#### Request
```json
{
    "user": "rand:username",
    "name": "rand:name",
    "phone": "rand:phone"
}
```
#### Response
```json
{
    "user": "nature_moon",
    "name": "Samuel Smith",
    "phone": "+1 56486555"
}
```
# Funciones
|  Nombre  |     Comando     |                     Detalles                      |
| -------- | --------------- | ------------------------------------------------- |
| Name     | `rand:name`     | Genera nombres y apellidos                        |
| Username | `rand:user`     | Genera nombres de usuarios                        |
| Email    | `rand:email`    | Genera correos electrónicos                       |
| Phone    | `rand:phone`    | Genera números de teléfono de cualquier país      |
| Score    | `rand:score`    | Genera una puntuación de 0 al 5                   |
| Priority | `rand:priority` | Genera una prioridad entre `low`,`meddium`,`high` |

