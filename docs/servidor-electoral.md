# Servidor Electoral

## Funcionalidades
- Actuador encargado de coordinar todas las
elecciones activas.
- Se encarga de configurar nodos de votacion
activa. 
- Activacion de puntos de votacion 
y manejo de llaves.

## Datos a utilizar
- Elecciones Activas
- Candidatos (por Eleccion)
- Votantes (por Eleccion)
- Jurados (por Eleccion)
- Puntos de Votacion
- Terminales de Votacion (por Punto de Votacion)
- Terminales de Jurado (por Punto de Jurado)

## Queries y Mutaciones

### Sistema Registraduria

#### Listar Elecciones
- **Endpoint**: GET /api/elecciones

```json
[{
  "id": 1,
  "nombreOficial": "Elección Presidencial 2026",
  "pais": "Colombia",
  "tipoEleccion": "PRESIDENCIAL",
  "codigoMetodoElectoral": "ME_02",
  "fechaInicioJornada": "2026-05-31T08:00:00",
  "fechaCierreJornada": "2026-05-31T16:00:00",
  "modalidadHabilitada": "AMBAS",
  "tipoCircunscripcion": "NACIONAL",
  "documentoNoVotable": "CÉDULA DE EXTRANJERÍA",
  "numeroCurules": null,
  "formulaCifraRepartidora": null,
  "condicionVictoria": "SUPERA_UMBRAL_DE_PRIMERA_VUELTA_SOBRE_VOTOS_VALIDOS",
  "estado": "PUBLICADA",
  "modelosCandidatura": ["UNICO"],
  "excencionesHabilitadas": [
    "PERSONAL ACTIVO DE FUERZAS MILITARES Y POLICIA",
    "DISCAPACIDAD PERMANENTE CERTIFICADA"
  ],
  "configuracionSenado": null,
  "configuracionCamara": null,
  "configuracionCamaraEspeciales": null
}]
```

#### Listar Candidatos
- **Endpoint**: GET /api/elecciones/[id]

```json
[{
    "id": 1,
    "eleccionId": 1,
    "nombreCandidato": "Augusto Pedicino",
    "documento": "12345678",
    "partido": "Pacto Historico",
    "circunscripcion": "no se que sea esta mierda",
    "fotoUrl": "base64/algo",
    "estado": "No se",
    ...
}]
```

#### Listar Votantes
- **Endpoint:** GET /api/censo/elecciones/[eleccionId]/registros

**Nota:** Falta emular los puestos de votacion por votante.

```json
{
  "contenido": [
    {
      "id": 101,
      "eleccionId": 1,
      "tipoDocumento": "CC",
      "numeroDocumento": "1020304050",
      "nombres": "Juan Carlos",
      "apellidos": "Pérez Gómez",
      "fechaNacimiento": "1990-04-18",
      "departamento": "Antioquia",
      "municipio": "Medellín",
      "estado": "HABILITADO",
      "causalEstado": "NINGUNA",
      "observacion": "Registro válido",
      "actorUltimaModificacion": "sistema",
      "fechaActualizacion": "2026-05-16T10:25:30"
    }
  ],
  "totalElementos": 2,
  "totalPaginas": 1,
  "numeroPagina": 0,
  "tamanoPagina": 10
}
```

#### Listar Jurados
- **Endpoint:** GET /
  
```json
NO ESTA
```

#### Listar Puntos de Votacion
- **Endpoint:** GET /
  
```json
NO ESTA
```

#### Listar Terminales/Mesas de Votacion
- **Endpoint:** GET / 

```json
{
    "id": 1,
    "mesaCode": "MESA-101",
    "validVotes": 0,
    "blankVotes": 0,
    "nullVotes": 0,
    "unmarkedVotes": 0,
    "status": "OPEN"
}
```

#### Listar Terminales de Jurados de Votacion
- **Endpoint:** GET /
  
```json
NO ESTA
```

---

### Nodos de Votacion Activa

#### Configuracion Nodo
- **Endpoint:** GET /nodo
- **Headers:**
    - **Authorization:** Bearer (Secreto)

```json
{
    "eleccion": {
        "id": 1,
        "nombre": "Congreso 2026",
        "tipo_eleccion": "presidencial", // presidencial, legislativa o territorial
        "fecha_inicio": 1778979739, // unix timestamp
        "fecha_fin": 1778979739, // unix timestamp
    },
    
    "candidatos": [{
        "id": 1,
        "nombre": "Augusto Pedicino",
        "documento": "12345678",
        "partido": "Pacto Historico",
        "foto_url": "data:image/png;base64,...",
    }],
    
    "puntos": [{
        "id": 1,
        "nombre": "Puesto Avenida Chile",
        "latitud": 10.99,
        "longitud": 44.34,
        "activo": true,
        "secreto": "(jwt)", // Esto es el bearer token

        "jurados": [{
            "id": 1,
            "nombre": "Augusto Pedicino",
            "documento": "12345678",
            "usuario": "augusto",
            "hash": "(contraseña argon2)",
        }],

        "terminales": [{
            "id": 1,
            "secreto": "(jwt)", // Esto es el bearer token
            "clave_publica": "(clave Ed25519)", // Esto es el secreto del puesto para firmar Ed25519
            "activo": true,
            "votantes": [{
                "id": 1,
                "nombre": "Augusto Pedicino",
                "documento": "12345678",
            }],

        }],
    }]
}
```

## Archivos de Configuracion
### Nodo de Votacion Activa
```json
{
    "id": 1,
    "secreto": "(jwt)",
    "server_url": "https://electoral.sello-legitimo.site/",
    "database_url": "postgres://postgres@postgres:5432/votacion",
    "queue_url": "amqps://user:password@host:port/vhost"
}
```

### Terminal de Votacion
```json
{
    "id": 1,
    "secreto": "(jwt)",
    "clave_privada": "(secreto)",
    "cluster_url": "https://nodo-votacion.local/"
}
```

### Terminal de Jurado
```json
{
    "id": 1,
    "secreto": "(jwt)",
    "parent_url": "https://maquina-jurado.local/"
}
```