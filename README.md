# Web form to upload files to GCS

[<img src="https://img.shields.io/badge/Last%20Update-18%2F03%2F23-orange">](#web-form-to-upload-files-to-gcs)
[<img src="https://img.shields.io/badge/Version-1.0.0-blue">](#web-form-to-upload-files-to-gcs)
[<img src="https://img.shields.io/badge/Status-OK-green">](#web-form-to-upload-files-to-gcs)

Un ejemplo de formulario web simple para cargar arhivos en Google Cloud Storage.

<details id="contentTable">
    <summary>Tabla de contenidos</summary>
    <ol>
        <li>
            <a href="#requisitos-previos">Requisitos previos</a>
        </li>
        <li>
            <a href="#setup">Setup</a>
        </li>
        <li>
            <a href="#ejecución">Ejecución</a>
        </li>
        <li>
            <a href="#valores-de-configuración">Valores de configuración</a>
        </li>
    </ol>
</details>

## Requisitos previos
+ [Go](https://golang.org/dl/)
+ [Docker](https://docs.docker.com/install/)
+ [make](https://www.gnu.org/software/make/) (opcional)

## Setup
1. Crear un bucket en GCS
2. Descargar el archivo de credenciales de GCP en la carpeta `credentials` y renombrarlo a [`keys.json`](credentials/keys.json).
3. Modificar los valores del archivo [`config.json`](config.json).
   - [`Descripción de los valores de configuración`](#valores-de-configuración).

## Ejecución
1. Ejecutar el comando `make docker-br` para construir la imagen y crear el contenedor de Docker. Si no tienes make instalado no te preocupes, sigue los pasos del siguiente punto.
2. Crea la imagen de Docker.
```bash
docker build -t gcs-upload .
```
3. Crea un contenedor con la imagen creada en el paso anterior y expón el puerto 8080.
```bash
docker run --name gcs-upload_prod -dp 8080:8080 gcs-upload
```
4. Ejecuta el contenedor.
```bash
docker start gcs-upload_prod
```
5. Por último, abre la dirección [`http://localhost:8080`](http://localhost:8080) en el navegador y prueba subir un archivo.

## Valores de configuración
| Nombre | Tipo | Descripción |
| --- | --- | --- |
| `webHost` | String | El host en el que se ejecutará el servidor. |
| `webPort` | String | El puerto en el que se ejecutará el servidor. |
| `bucketName` | String | El nombre del archivo del bucket. |
| `bucketFolder` | String | El nombre de la carpeta a usar en el bucket. |
| `contentType` | String | Tipo de contenido del archivo. |
| `ceredentialsFile` | String | Ruta del archivo de credenciales. |

&emsp;
<p align="center">Copyright &copy; 2023 | <a href="https://github.com/selmansem" target="_blank">Souhaib EM</a></p>

<p align="center"><a href="#contentTable">&#8593; Volver a la tabla de contenido &#8593;</a></p>
