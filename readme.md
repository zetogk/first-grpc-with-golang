## Requisitos
* Instalar Golang (v1.6+).
* Instalar https://github.com/google/protobuf/releases y asegurarte que el archivo binario quede dentro del PATH del sistema.
* Instalar plugin de protoc para Go (Para este caso de ejemplo): `go get -u github.com/golang/protobuf/protoc-gen-go`.

## Estructura de carpetas

La siguiente será nuestra estructura de carpetas

* **definicionservicio**
    * **notificador.proto** Tenemos definido nuestra definición del sericio

* **servidor**
    * **main.go** Usa el archivo generador por protoc para implementar el servidor

* **cliente**
    * **main.go** Usa el archivo generador por protoc para consumir el servicio
    * **notificador.pb.go** // Archivo generado por protoc, será usado por el servidor y el cliente

## Pasos

*Debes poner el repo en tu GOPATH/src/ejemplo_servicio*

### Generar archivo

Ejecuta el siguiente comando dentro de la carpeta raíz del repositorio:

```bash
protoc -I definicionservicio/ definicionservicio/notificador.proto --go_out=plugins=grpc:definicionservicio
```

*protoc* nos generará un archivo llamado *notificador.pb.go* el cual debe ser usado tanto por el servidor como por el cliente.


### Construimos nuestro servidor

Entra a la carpeta **servidor** donde se encuentra el archivo **main.go**, allí están los comentarios de como se hace la implementación.


### Construimos nuestro cliente

Entra a la carpeta **cliente** donde se encuentra el archivo **main.go**, allí están los comentarios para consumir el cliente y hacer tu petición RPC.


### Listo para usar

* Abre una terminal donde correrás `go run ./servidor/main.go` allí se expondrá el servicio que estará a la espera de llamadas.
* Abre una terminal adicional donde correrás `go run ./cliente/main.go`, desde donde se generará el llamado RPC.