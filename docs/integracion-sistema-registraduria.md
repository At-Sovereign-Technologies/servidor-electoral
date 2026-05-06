# Guia de Integracion: Sistema Electoral → Modulo de Gestion Pre-Electoral

**Fecha:** 2026-05-06  
**Version:** 1.0  
**Autor:** Equipo Sello Legitimo - Modulo de Gestion  
**Destinatario:** Desarrollador Sistema Electoral  

---

## 1. Objetivo

Este documento describe como el **Sistema Electoral** debe consumir informacion de elecciones y tarjetones desde el **Modulo de Gestion Pre-Electoral** va gRPC y REST para poder renderizar los tarjetones electorales.

---

## 2. Arquitectura de Servicios

```
+-----------------------------------+
|  ConfiguracionEleccion-service    |  gRPC SERVER (port 9090)
|  EleccionService:                 |  REST (port 8081)
|    ListarElecciones               |
|    ObtenerEleccion                |
+--------+--------------------------+
         |
         | gRPC (EleccionServiceBlockingStub)
         |
+--------v--------------------------+     REST (port 8082)
|  GestionPreElectoral-service      |<------------------------+
|  (gRPC CLIENT de EleccionService) |                          |
|                                   |                          |
|  REST endpoints:                  |                         |
|    /api/candidaturas/...          |                          |
|    /api/censo/...                 |                          |
+-----------------------------------+                          |
                                                                |
+-----------------------------------+                          |
|  Sistema Electoral (TU SERVICIO)  |--------------------------+
|  - Necesitas leer elecciones      |  gRPC CLIENT a port 9090
|  - Necesitas leer tarjetones      |  REST CLIENT a port 8082
+-----------------------------------+
```

### Puertos y Hosts (Docker Compose)

| Servicio | REST Port | gRPC Port | Host Docker |
|---|---|---|---|
| ConfiguracionEleccion | 8081 | 9090 | `configuracion-eleccion` |
| GestionPreElectoral | 8082 | N/A (solo client) | `gestion-pre-electoral` |
| MockJurados | 8083 | 9091 (server) | `mock-jurados` |

### Variables de Entorno para Conectar

```bash
# Conexion gRPC a ConfiguracionEleccion
ELECCIONES_GRPC_HOST=configuracion-eleccion
ELECCIONES_GRPC_PORT=9090

# Conexion REST a GestionPreElectoral
GESTION_PRE_ELECTORAL_URL=http://gestion-pre-electoral:8082
```

---

## 3. Servicio gRPC: EleccionService

El servicio gRPC `EleccionService` corre en **ConfiguracionEleccion-service** (port 9090). Desde aqui obtienes los datos maestros de cada eleccion.

### 3.1. Proto Definition (`elecciones.proto`)

```protobuf
syntax = "proto3";

package elecciones;

// Opciones Java (si tu servicio es Java/Spring):
// option java_package = "com.selloLegitimo.grpc.elecciones";
// option java_outer_classname = "EleccionesProto";
// option java_multiple_files = true;

service EleccionService {
  rpc ListarElecciones (ListarEleccionesRequest) returns (ListarEleccionesResponse);
  rpc ObtenerEleccion (ObtenerEleccionRequest) returns (EleccionDetalle);
}

message ListarEleccionesRequest {}

message ListarEleccionesResponse {
  repeated EleccionResumen elecciones = 1;
}

message EleccionResumen {
  int64 id = 1;
  string nombre_oficial = 2;
  string estado = 3;
}

message ObtenerEleccionRequest {
  int64 id = 1;
}

message EleccionDetalle {
  int64 id = 1;
  string nombre_oficial = 2;
  string estado = 3;
  string fecha_inicio_jornada = 4;
  string fecha_cierre_jornada = 5;
  string documento_no_votable = 6;
  repeated string excenciones_habilitadas = 7;
  string fecha_inicio_mod_candidaturas = 8;
  string fecha_fin_mod_candidaturas = 9;
  string fecha_limite_reemplazo_candidaturas = 10;
  int32 edad_minima_candidatura = 11;
}
```

### 3.2. RPC: ListarElecciones

Retorna todas las elecciones registradas, con su estado actual.

**Request:** Vacio (`ListarEleccionesRequest`)

**Response:**
```json
{
  "elecciones": [
    { "id": 1, "nombre_oficial": "Elecciones Legislativas 2026", "estado": "PUBLICADA" },
    { "id": 2, "nombre_oficial": "Gobernador Antioquia 2027", "estado": "BORRADOR" },
    { "id": 3, "nombre_oficial": "Alcalde Bogota 2027", "estado": "EN_CURSO" }
  ]
}
```

### 3.3. RPC: ObtenerEleccion

Retorna el detalle completo de una eleccion especifica por su ID.

**Request:**
```json
{ "id": 1 }
```

**Response:**
```json
{
  "id": 1,
  "nombre_oficial": "Elecciones Legislativas 2026",
  "estado": "PUBLICADA",
  "fecha_inicio_jornada": "2026-03-15T08:00:00",
  "fecha_cierre_jornada": "2026-03-15T16:00:00",
  "documento_no_votable": "PEP",
  "excenciones_habilitadas": ["indigenas", "fuerzas_militares"],
  "fecha_inicio_mod_candidaturas": "2025-11-01T00:00:00",
  "fecha_fin_mod_candidaturas": "2026-01-15T23:59:59",
  "fecha_limite_reemplazo_candidaturas": "2026-02-28T23:59:59",
  "edad_minima_candidatura": 30
}
```

### 3.4. Estados de Eleccion

| Estado | Significado |
|---|---|
| `BORRADOR` | Configuracion inicial, no visible publicamente |
| `PUBLICADA` | Visible, periodos de candidaturas abiertos o proximos |
| `EN_CURSO` | Jornada electoral activa (tarjetones deben mostrarse) |
| `CERRADA` | Jornada finalizada |
| `RESULTADOS_OFICIALES` | Resultados oficiales publicados |

**Para el Sistema Electoral, las elecciones relevantes para tarjetones son las que estan en estado `PUBLICADA` o `EN_CURSO`.**

### 3.5. Ejemplo de Cliente gRPC (Java)

```java
// Configuracion del canal
ManagedChannel channel = ManagedChannelBuilder
    .forAddress("configuracion-eleccion", 9090)
    .usePlaintext()
    .build();

EleccionServiceBlockingStub stub = EleccionServiceGrpc.newBlockingStub(channel);

// Listar elecciones
ListarEleccionesResponse response = stub.listarElecciones(ListarEleccionesRequest.newBuilder().build());
for (EleccionResumen e : response.getEleccionesList()) {
    System.out.println(e.getId() + ": " + e.getNombreOficial() + " [" + e.getEstado() + "]");
}

// Obtener detalle
EleccionDetalle detalle = stub.obtenerEleccion(
    ObtenerEleccionRequest.newBuilder().setId(1L).build());
System.out.println("Jornada: " + detalle.getFechaInicioJornada() + " - " + detalle.getFechaCierreJornada());

channel.shutdown();
```

### 3.6. Ejemplo de Cliente gRPC (Python)

```python
import elecciones_pb2
import elecciones_pb2_grpc
import grpc

channel = grpc.insecureChannel('configuracion-eleccion:9090')
stub = elecciones_pb2_grpc.EleccionServiceStub(channel)

# Listar elecciones
response = stub.ListarElecciones(elecciones_pb2.ListarEleccionesRequest())
for e in response.elecciones:
    print(f"{e.id}: {e.nombre_oficial} [{e.estado}]")

# Obtener detalle
detalle = stub.ObtenerEleccion(elecciones_pb2.ObtenerEleccionRequest(id=1))
print(f"Jornada: {detalle.fecha_inicio_jornada} - {detalle.fecha_cierre_jornada}")

channel.close()
```

### 3.7. Ejemplo de Cliente gRPC (Node.js / TypeScript)

```typescript
import { EleccionServiceClient } from './generated/elecciones_grpc_pb';
import { ListarEleccionesRequest, ObtenerEleccionRequest } from './generated/elecciones_pb';
import * as grpc from '@grpc/grpc-js';

const client = new EleccionServiceClient(
  'configuracion-eleccion:9090',
  grpc.credentials.createInsecure()
);

// Listar elecciones
client.listarElecciones(new ListarEleccionesRequest(), (err, response) => {
  if (err) throw err;
  response.getEleccionesList().forEach(e => {
    console.log(`${e.getId()}: ${e.getNombreOficial()} [${e.getEstado()}]`);
  });
});

// Obtener detalle
const req = new ObtenerEleccionRequest();
req.setId(1);
client.obtenerEleccion(req, (err, detalle) => {
  if (err) throw err;
  console.log(`Jornada: ${detalle.getFechaInicioJornada()} - ${detalle.getFechaCierreJornada()}`);
});
```

---

## 4. REST API: Tarjetones (GestionPreElectoral)

El Modulo de Gestion Pre-Electoral expone los tarjetones via **REST** en el puerto **8082**. No existe actualmente un servicio gRPC para tarjetones; se accede via HTTP.

### 4.1. Obtener el Ultimo Tarjeton Generado

```
GET /api/candidaturas/elecciones/{eleccionId}/tarjeton
```

Retorna el ultimo tarjeton generado para una eleccion. Si no existe ningun tarjeton en la auditoria, regenera uno con ordenamiento alfabetico como fallback.

**Response 200:**
```json
{
  "eleccionId": 1,
  "circunscripcion": null,
  "fechaGeneracion": "2026-05-05T14:30:00",
  "semillaUsada": 1714920600000,
  "entradas": [
    {
      "orden": 1,
      "nombreCandidato": "Maria Garcia",
      "partido": "Partido Liberal",
      "fotoUrl": "/fotos/maria-garcia.jpg",
      "tipo": "CANDIDATO"
    },
    {
      "orden": 2,
      "nombreCandidato": "Carlos Rodriguez",
      "partido": "Conservador",
      "fotoUrl": "/fotos/carlos-rodriguez.jpg",
      "tipo": "CANDIDATO"
    },
    {
      "orden": 3,
      "nombreCandidato": "Voto en Blanco",
      "partido": null,
      "fotoUrl": null,
      "tipo": "VOTO_BLANCO"
    }
  ]
}
```

### 4.2. Generar un Nuevo Tarjeton

```
POST /api/candidaturas/elecciones/{eleccionId}/tarjeton
Content-Type: application/json
```

**Request body:**
```json
{
  "circunscripcion": null,
  "tipoOrdenamiento": "ALEATORIO_AUDITADO",
  "semillaAleatoria": null,
  "actor": "sistema-electoral"
}
```

| Campo | Tipo | Requerido | Descripcion |
|---|---|---|---|
| `circunscripcion` | string | No | Filtrar por circunscripcion (null = todas) |
| `tipoOrdenamiento` | string | Si | `ALFABETICO` o `ALEATORIO_AUDITADO` |
| `semillaAleatoria` | int64 | No | Semilla para orden aleatorio reproducible (si es null, se genera automaticamente) |
| `actor` | string | Si | Identificador del sistema/usuario que solicita |

**Response 200:** Mismo formato que GET, pero con el tarjeton recien generado.

### 4.3. Listar Candidaturas de una Eleccion

```
GET /api/candidaturas/elecciones/{eleccionId}
```

Retorna todas las candidaturas de una eleccion (incluidas las que no estan en el tarjeton).

**Response 200:**
```json
[
  {
    "id": 1,
    "eleccionId": 1,
    "nombreCandidato": "Maria Garcia",
    "documento": "1000000001",
    "partido": "Partido Liberal",
    "circunscripcion": "Bogota",
    "fotoUrl": "/fotos/maria-garcia.jpg",
    "estado": "APROBADO",
    "candidaturaReemplazadaId": null,
    "motivoReemplazo": null,
    "justificacionReemplazo": null,
    "actorUltimaModificacion": "admin",
    "fechaInscripcion": "2026-01-10T10:00:00",
    "fechaActualizacion": "2026-01-10T10:00:00",
    "version": 1
  }
]
```

### 4.4. Estados de Candidatura Relevantes para Tarjeton

Solo las candidaturas en estado **APROBADO** o **PUBLICADO** se incluyen en el tarjeton.

| Estado | En tarjeton? |
|---|---|
| `BORRADOR` | No |
| `POSTULADO` | No |
| `EN_VALIDACION` | No |
| `APROBADO` | **Si** |
| `PUBLICADO` | **Si** |
| `RECHAZADO` | No |
| `BLOQUEADO` | No |
| `REEMPLAZADA` | No |
| `REVOCADA` | No |

---

## 5. Estructura de Datos del Tarjeton

### 5.1. Modelo `TarjetonRespuestaDto`

| Campo | Tipo | Descripcion |
|---|---|---|
| `eleccionId` | int64 | ID de la eleccion a la que pertenece el tarjeton |
| `circunscripcion` | string \| null | Circunscripcion filtrada (null = todas) |
| `fechaGeneracion` | datetime (ISO-8601) | Fecha/hora en que se genero el tarjeton |
| `semillaUsada` | int64 \| null | Semilla del orden aleatorio (null si es alfabetico) |
| `entradas` | list\<TarjetonEntradaDto\> | Lista ordenada de entradas del tarjeton |

### 5.2. Modelo `TarjetonEntradaDto`

| Campo | Tipo | Descripcion |
|---|---|---|
| `orden` | int | Posicion en el tarjeton (1-based, secuencial) |
| `nombreCandidato` | string | Nombre del candidato o "Voto en Blanco" |
| `partido` | string \| null | Partido politico (null para voto en blanco) |
| `fotoUrl` | string \| null | URL de la foto del candidato (null para voto en blanco) |
| `tipo` | string | `CANDIDATO` o `VOTO_BLANCO` |

### 5.3. Reglas de Negocio

1. **Voto en Blanco siempre va al final:** La entrada con `tipo = VOTO_BLANCO` siempre tiene el `orden` mas alto.
2. **Orden aleatorio auditado:** Cuando `tipoOrdenamiento = ALEATORIO_AUDITADO`, los candidatos se mezclan usando una semilla (`semillaUsada`). La misma semilla produce el mismo orden, garantizando reproducibilidad y auditabilidad.
3. **Solo APROBADO/PUBLICADO:** Solo candidaturas en estado `APROBADO` o `PUBLICADO` aparecen en el tarjeton.
4. **Filtro por circunscripcion:** Si se especifica `circunscripcion`, solo se incluyen candidaturas de esa circunscripcion.
5. **Semilla de auditoria:** El tarjeton generado se almacena con su lista completa de entradas en la tabla `auditoria_eventos` (tipo `TARJETON_GENERADO`). El endpoint GET recupera el ultimo generado; el POST genera uno nuevo.

---

## 6. Flujo de Integracion Recomendado

### Paso 1: Obtener Elecciones Disponibles (gRPC)

```
gRPC → ConfiguracionEleccion:9090 → ListarElecciones()
```

Filtra las elecciones con estado `PUBLICADA` o `EN_CURSO` para mostrar al jurado/votante.

### Paso 2: Obtener Detalle de la Eleccion (gRPC)

```
gRPC → ConfiguracionEleccion:9090 → ObtenerEleccion(id)
```

Obtiene fechas de jornada, tipo de eleccion, y otros datos maestros necesarios.

### Paso 3: Obtener el Tarjeton (REST)

```
GET http://gestion-pre-electoral:8082/api/candidaturas/elecciones/{eleccionId}/tarjeton
```

Retorna el tarjeton con las candidaturas ordenadas y el voto en blanco al final.

### Paso 4: Renderizar el Tarjeton

Con los datos del tarjeton, renderizar segun la normativa colombiana:
- Grilla de 2 columnas con barras de color del partido
- Iniciales del candidato en avatar circular
- Casillas de votacion preferencial
- Seccion de VOTO EN BLANCO al final
- Encabezado oficial con datos de la eleccion

### Diagrama de Secuencia

```
Sistema Electoral          ConfiguracionEleccion         GestionPreElectoral
      |                           |                              |
      |-- ListarElecciones() ---->|                              |
      |<-- [lista elecciones] ---|                              |
      |                           |                              |
      |-- ObtenerEleccion(1) --->|                              |
      |<-- EleccionDetalle ------|                              |
      |                           |                              |
      |-- GET /api/candidaturas/elecciones/1/tarjeton --------->|
      |<-- TarjetonRespuestaDto --------------------------------|
      |                           |                              |
      | [Renderizar tarjeton]     |                              |
```

---

## 7. Proposed gRPC Service: TarjetonService (For Future Implementation)

Currently, tarjetones are only available via REST. If pure gRPC integration is needed, the following service definition is proposed for GestionPreElectoral-service to implement. **This is NOT yet implemented** — it requires adding a gRPC server to GestionPreElectoral.

```protobuf
syntax = "proto3";

package elecciones;

service TarjetonService {
  rpc ObtenerTarjeton (ObtenerTarjetonRequest) returns (TarjetonRespuesta);
  rpc ObtenerUltimoTarjeton (ObtenerTarjetonRequest) returns (TarjetonRespuesta);
  rpc ListarCandidaturas (ListarCandidaturasRequest) returns (CandidaturasRespuesta);
}

message ObtenerTarjetonRequest {
  int64 eleccion_id = 1;
  string circunscripcion = 2;
  string tipo_ordenamiento = 3;   // ALFABETICO | ALEATORIO_AUDITADO
  int64 semilla_aleatoria = 4;
  string actor = 5;
}

message TarjetonEntrada {
  int32 orden = 1;
  string nombre_candidato = 2;
  string partido = 3;
  string foto_url = 4;
  string tipo = 5;                // CANDIDATO | VOTO_BLANCO
}

message TarjetonRespuesta {
  int64 eleccion_id = 1;
  string circunscripcion = 2;
  string fecha_generacion = 3;
  int64 semilla_usada = 4;
  repeated TarjetonEntrada entradas = 5;
}

message ListarCandidaturasRequest {
  int64 eleccion_id = 1;
}

message CandidaturaResumen {
  int64 id = 1;
  int64 eleccion_id = 2;
  string nombre_candidato = 3;
  string documento = 4;
  string partido = 5;
  string circunscripcion = 6;
  string foto_url = 7;
  string estado = 8;
}

message CandidaturasRespuesta {
  repeated CandidaturaResumen candidaturas = 1;
}
```

**Si prefieres gRPC puro, haga saber al equipo de Gestion para que implementemos este servicio.**

---

## 8. Ejemplo Completo: Obtener y Renderizar un Tarjeton

```bash
# 1. Listar elecciones via gRPC (usando grpcurl)
grpcurl -plaintext configuracion-eleccion:9090 elecciones.EleccionService/ListarElecciones

# 2. Obtener detalle de eleccion 1
grpcurl -plaintext -d '{"id": 1}' configuracion-eleccion:9090 elecciones.EleccionService/ObtenerEleccion

# 3. Obtener el tarjeton de la eleccion 1 via REST
curl -s http://gestion-pre-electoral:8082/api/candidaturas/elecciones/1/tarjeton | jq .

# 4. Generar un nuevo tarjeton con orden aleatorio via REST
curl -s -X POST http://gestion-pre-electoral:8082/api/candidaturas/elecciones/1/tarjeton \
  -H "Content-Type: application/json" \
  -d '{"tipoOrdenamiento":"ALEATORIO_AUDITADO","actor":"sistema-electoral"}' | jq .
```

---

## 9. Datos de Prueba (Seed Data)

Para pruebas en el entorno Docker local, existen las siguientes elecciones y candidaturas precargadas:

### Elecciones (ConfiguracionEleccion)

| ID | Nombre | Estado | Tipo |
|---|---|---|---|
| 1 | Elecciones Legislativas 2026 | PUBLICADA | LEGISLATIVA |
| 2 | Gobernador Antioquia 2027 | BORRADOR | TERRITORIAL |
| 3 | Alcalde Bogota 2027 | EN_CURSO | TERRITORIAL |
| 4 | Alcalde Bogota 2027 | BORRADOR | TERRITORIAL |

### Candidaturas (GestionPreElectoral)

Se incluyen 17 candidaturas en estados `APROBADO` y `PUBLICADO` para las elecciones 1-4, con ciudadanos habilitados en el censo electoral.

---

## 10. Consideraciones

1. **Autenticacion:** Los endpoints REST actualmente no requieren autenticacion interna entre servicios. El gateway (Caddy) maneja Authelia OIDC para el frontend. Para comunicacion servicio-a-servicio, conectar directamente al puerto REST del servicio (8082).

2. **Versionado del tarjeton:** Cada generacion de tarjeton queda registrada en la tabla `auditoria_eventos` con tipo `TARJETON_GENERADO`. El GET siempre retorna el ultimo generado. Para追溯 historial, se debe consultar directamente la tabla de auditoria.

3. **Reproducibilidad:** Si se usa `ALEATORIO_AUDITADO` con una `semillaAleatoria` fija, el orden sera siempre el mismo. Esto es critico para auditar que el tarjeton mostrado al votante coincide con el certificado por la autoridad electoral.

4. **Circunscripcion:** Si la eleccion tiene tipo de circunscripcion, se puede filtrar el tarjeton pasando el parametro `circunscripcion` al generar.

5. **Campso de fecha:** Todos los campos de fecha/fechaHora usan formato ISO-8601 (e.g., `2026-03-15T08:00:00`).

6. **Puerto gRPC:** El servidor gRPC usa canal plaintext (sin TLS) dentro de la red Docker. Para produccion con Ingress/Gateway externo, se debe configurar TLS.

---

## 11. Protocol Buffer Files Location

Los archivos `.proto` se encuentran en el repositorio del proyecto en:

| Archivo | Ruta | Servicio |
|---|---|---|
| `elecciones.proto` | `ConfiguracionEleccion-service/src/main/proto/elecciones.proto` | Definicion canonica |
| `elecciones.proto` | `GestionPreElectoral-service/src/main/proto/elecciones.proto` | Copia para cliente |
| `elecciones.proto` | `MockJurados-service/src/main/proto/elecciones.proto` | Copia para cliente |
| `jurados.proto` | `MockJurados-service/src/main/proto/jurados.proto` | Servicio de jurados |

**Recomendacion:** Copia `elecciones.proto` a tu proyecto y genera los stubs para tu lenguaje usando `protoc`.