# Beer Test 🚀
Este es un Web API diseñado para ayudar a los usuarios a identificar los precios de la cerveza, agregar nuevos precios y calcular el valor de las cajas de cerveza. El API está equipado con Unit Tests, lo que garantiza la calidad del software. Además, está construido utilizando la arquitectura limpia, lo que significa que está estructurado de manera modular para mantener un alto nivel de mantenibilidad y escalabilidad. El API también utiliza una base de datos postgresSQL para el almacenamiento de datos, lo que garantiza una gestión de datos segura y eficiente. Para facilitar la implementación y la gestión del software, todo el sistema está diseñado para cargarse automáticamente a través de Docker Compose. En resumen, este Web API es una solución robusta y confiable para identificar precios de cerveza y calcular el valor de las cajas de cerveza.


## Endpoints 📋

``` GET - /beers ```
 
``` GET - /beers/:beerID ```

``` POST - /beers ```

``` GET - /beers/:beerID/boxprice - currency  -quantity ```


Requisitos 📋
Para poder utilizar este software, se necesitan las siguientes herramientas:

Docker

Docker-Compose

Estas herramientas son necesarias para garantizar una implementación exitosa del software y una experiencia de usuario sin problemas. Si aún no tienes estas herramientas instaladas en tu sistema, puedes descargarlas de los sitios web oficiales de Docker y Docker-Compose.

# ID of beers created in migration

``` 2bcaaa15-a7c0-4d57-bafa-b5d6c2cf410d```

``` 5f21ed03-513a-4731-9323-59d46c1d739b```

``` 56dfa6bc-5b13-4401-a4d0-15e1fff3a784```

# Instructions to start the Web API 


Para iniciar el Web API, simplemente abre una terminal y navega hasta el directorio raíz del proyecto. Luego, ejecuta el siguiente comando en la terminal:

``` docker-compose up ```

Este comando se encargará de cargar y ejecutar todos los contenedores necesarios para el Web API, incluyendo la base de datos y el servidor web. Una vez que los contenedores se hayan cargado correctamente, podrás acceder al Web API a través de tu navegador o herramienta de cliente HTTP.

Recuerda que es importante tener Docker y Docker-Compose instalados en tu sistema antes de ejecutar este comando. Si aún no los tienes instalados, puedes descargarlos de los sitios web oficiales de Docker y Docker-Compose.