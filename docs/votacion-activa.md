# Servidor Electoral

## Funcionalidades
- Encargados de guardar y procesar todos los
votos enviados desde la votacion activa y mantener trazabilidad.

- Endpoints para la comunicacion de los distintos puntos de votacion.
- 
## Datos a utilizar
- Eleccion
- Candidatos (por Eleccion)
- Votantes (por Eleccion)
- Jurados (por Eleccion)
- Puntos de Votacion
- Terminales de Votacion (por Punto de Votacion)
- Terminales de Jurado (por Punto de Jurado)

## Queries y Mutaciones

### Sistema Registraduria

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
        "jwt": "(jwt)", // Esto es el bearer token

        "jurados": [{
            "id": 1,
            "nombre": "Augusto Pedicino",
            "documento": "12345678",
            "usuario": "augusto",
            "hash": "(contraseña argon2)",
        }],

        "terminales": [{
            "id": 1,
            "jwt": "(jwt)", // Esto es el bearer token
            "activo": true,
            "secreto": "(secreto)", // Esto es el secreto del puesto para firmar Ed25519
            "votantes": [{
                "id": 1,
                "nombre": "Augusto Pedicino",
                "documento": "12345678",
            }],

        }],
    }]
}
```

---

### Punto de Votacion

#### Configuracion Punto
- **Endpoint:** GET /puesto
- **Headers:**
    - **Authorization:** Bearer (JWT)

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
    
    "punto": {
        "id": 1,
        "nombre": "Puesto Avenida Chile",
        "latitud": 10.99,
        "longitud": 44.34,

        "jurados": [{
            "id": 1,
            "nombre": "Augusto Pedicino",
            "documento": "12345678",
            "usuario": "augusto",
            "hash": "(contraseña argon2)",
        }],

        "terminales": [{
            "id": 1,
            "votantes": [{
                "id": 1,
                "nombre": "Augusto Pedicino",
                "documento": "12345678",
            }],
        }],
    }
}
```

#### Confirmar Identidad
- **Endpoint:** GET /votante/[documento]

```json
{
    "votado": true,
}
```

#### Voto
- **Endpoint:** GET /votar
- **Headers:**
    - **Authorization:** Bearer (JWT)

```json
{
    "voto": {
        "terminal": 1,
        "votante": 1,
        "candidato": 1,
    },
    "firma": "(firma del voto en Ed25519)"
}
```