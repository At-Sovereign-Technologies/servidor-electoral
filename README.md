# Servidor Electoral

## Descripción General

El **Servidor Electoral** es una aplicación desarrollada en Go que actúa como coordinador central para la gestión de procesos electorales activos. Este sistema se encarga de coordinar todas las elecciones activas, configurar nodos de votación, activar puntos de votación y manejar la seguridad mediante sistemas de autenticación y firmas criptográficas.

## Funcionalidades Principales

- **Coordinación de Elecciones Activas**: Actuador encargado de coordinar y gestionar todas las elecciones en proceso.
- **Configuración de Nodos de Votación**: Realiza la configuración de nodos de votación activa en tiempo real.
- **Activación de Puntos de Votación**: Gestiona la activación de puntos de votación y manejo de llaves de acceso.
- **Trazabilidad de Votos**: Guarda y procesa todos los votos enviados desde la votación activa, manteniendo un registro auditado.
- **Comunicación Segura**: Proporciona endpoints para la comunicación segura entre los distintos puntos de votación.

## Tecnologías Utilizadas

- **Lenguaje**: Go 1.26.0
- **Base de Datos**: PostgreSQL (via GORM)
- **Framework Web**: Echo v4
- **Criptografía**: Ed25519 para firmas digitales
- **ORM**: GORM
- **Contenedorización**: Docker y Docker Compose

## Estructura de Datos

### Entidades Principales

El sistema gestiona las siguientes entidades:

- **Elecciones**: Evento electoral (presidencial, legislativa, territorial)
- **Candidatos**: Candidatos registrados por elección
- **Votantes**: Ciudadanos habilitados para votar (censo electoral)
- **Jurados**: Supervisores de puntos de votación
- **Puntos de Votación**: Ubicaciones físicas donde se realiza la votación
- **Terminales de Votación**: Máquinas de votación dentro de un punto
- **Nodos de Votación**: Servidores que coordinan la votación activa

## API Principal

### Configuración del Nodo

**Endpoint**: `GET /nodo`  
**Headers**: `Authorization: Bearer <secreto>`

Retorna la configuración completa del nodo incluyendo:
- Información de la elección actual
- Lista de candidatos
- Configuración de puntos de votación
- Jurados y sus credenciales
- Terminales de votación y votantes asociados

### Configuración del Punto de Votación

**Endpoint**: `GET /puesto`  
**Headers**: `Authorization: Bearer <jwt>`

Retorna la configuración del punto de votación específico con:
- Información de la elección
- Candidatos disponibles
- Jurados asignados
- Terminales de votación

### Confirmación de Identidad

**Endpoint**: `GET /votante/[documento]`

Verifica si un votante ya ha emitido su voto.

### Registro de Voto

**Endpoint**: `GET /votar`  
**Headers**: `Authorization: Bearer <jwt>`

Procesa el voto emitido por un votante, incluyendo:
- ID de terminal
- ID de votante
- ID de candidato
- Firma digital Ed25519 del voto

## Archivos de Configuración

### Configuración del Nodo de Votación Activa

```json
{
    "id": 1,
    "secreto": "<jwt>",
    "server_url": "https://electoral.sello-legitimo.site/",
    "database_url": "postgres://postgres@postgres:5432/votacion",
    "queue_url": "amqps://user:password@host:port/vhost"
}
```

### Configuración de Terminal de Votación

```json
{
    "id": 1,
    "secreto": "<jwt>",
    "clave_privada": "<secreto_ed25519>",
    "cluster_url": "https://nodo-votacion.local/"
}
```

### Configuración de Terminal de Jurado

```json
{
    "id": 1,
    "secreto": "<jwt>",
    "parent_url": "https://maquina-jurado.local/"
}
```

## Seguridad

El sistema implementa múltiples capas de seguridad:

- **Autenticación JWT**: Tokens Bearer para acceso a endpoints
- **Firmas Digitales Ed25519**: Cada voto es firmado digitalmente
- **Contraseñas Argon2**: Hashing seguro de contraseñas de jurados
- **Autorización por Rol**: Diferentes permisos para jurados, terminales y nodos

## Instalación y Configuración

### Requisitos Previos

- Go 1.26.0 o superior
- PostgreSQL 12+
- Docker y Docker Compose (opcional)

### Inicialización

1. Clonar el repositorio:
```bash
git clone https://github.com/At-Sovereign-Technologies/servidor-electoral.git
cd servidor-electoral
```

2. Instalar dependencias:
```bash
go mod download
```

3. Configurar variables de entorno:
```bash
export DATABASE_URL="postgres://postgres@localhost:5432/votacion"
export ELECTORAL_SECRET="<jwt_secreto>"
```

### Ejecución con Docker Compose

```bash
docker-compose up
```

Esto iniciará:
- La aplicación del servidor electoral
- Una instancia de PostgreSQL
- Los servicios necesarios para la votación activa

## Estructura del Proyecto

```
servidor-electoral/
├── internal/
│   ├── drivers/
│   │   └── database/
│   │       └── gormstore/
│   ├── models/
│   ├── services/
│   └── store/
├── docs/
├── Dockerfile
├── compose.yml
└── go.mod
```

### Componentes Clave

- **Models**: Definición de estructuras de datos (Candidato, Votante, Elección, etc.)
- **Services**: Lógica de negocio (servicios de candidato, votante, jurado, elección)
- **Store**: Interfaces y implementaciones de persistencia de datos
- **Drivers**: Implementaciones específicas (GORM para PostgreSQL)

## Flujo de Votación

1. **Inicialización**: Se configura el nodo con la información de la elección
2. **Identificación**: El votante se identifica en el punto de votación
3. **Votación**: El votante selecciona un candidato en la terminal
4. **Firma**: El voto se firma digitalmente con la clave Ed25519
5. **Registro**: El voto se registra en la base de datos con trazabilidad completa
6. **Confirmación**: Se marca al votante como votado para evitar doble votación

## Consideraciones Legales y Éticas

Este sistema fue desarrollado para procesos electorales legales y democráticos. Cumple con principios de:

- **Confidencialidad del voto**: Los votos se almacenan de manera segura
- **Integridad electoral**: Firma digital de votos para garantizar no alteración
- **Trazabilidad**: Registro completo de auditoría de todos los eventos
