# Ichtaca: Microservicio de Autenticación mTLS para Servidores APT
Ichtaca (del náhuatl: "secreto" o "privado") es un microservicio de autenticación de vanguardia diseñado para asegurar y gestionar de forma robusta el acceso a servidores APT (Advanced Package Tool). Desarrollado en Go, Ichtaca implementa la autenticación mutua TLS (mTLS) para establecer un canal de comunicación seguro y verificar criptográficamente la identidad tanto del cliente como del servidor.

## Propósito y Misión
En entornos donde la seguridad y la integridad de la cadena de suministro de software son críticas, como en infraestructuras gubernamentales o de alta seguridad, la distribución de paquetes APT debe ser impenetrable. Ichtaca nace con la misión de ser el guardián de tus repositorios APT, garantizando que solo las entidades autenticadas y autorizadas puedan acceder y consumir los paquetes.

## Características Clave
- Autenticación Mutua TLS (mTLS): Establece una verificación de identidad bidireccional, asegurando que tanto el cliente que solicita acceso como el servidor APT son quienes dicen ser. Esto previene ataques de suplantación y accesos no autorizados.
- Diseño como Microservicio: Construido en Go, Ichtaca es ligero, eficiente y fácil de desplegar, integrar y escalar. Su arquitectura de microservicio garantiza que la capa de autenticación sea independiente, resiliente y modular.
- Control de Acceso Fino: Permite definir políticas de acceso granulares para diferentes clientes, asegurando que cada entidad tenga solo los permisos necesarios para interactuar con los repositorios.
- Integración con Servidores APT: Diseñado para interponerse de forma transparente entre los clientes y tu servidor APT existente, añadiendo una capa de seguridad esencial sin reconfiguraciones complejas del lado del repositorio.
- Orientado a Entornos Sensibles: Ideal para distribuciones personalizadas de Debian utilizadas en sectores gubernamentales o corporativos, donde la seguridad de la cadena de suministro de software es primordial.
- Código Abierto bajo GPLv3: Ichtaca se distribuye bajo la licencia GNU General Public License v3, promoviendo la libertad de usar, estudiar, compartir y modificar el software, a la vez que protege estas libertades para todas las obras derivadas.

## ¿Por qué Ichtaca?
Elegimos Ichtaca como nombre por su significado profundo y su resonancia con la función principal del servicio: proteger los secretos de tu infraestructura de software. Es el garante de que solo los "conocedores del secreto" (las entidades autenticadas) puedan acceder a los recursos vitales.

Con Ichtaca, tu servidor APT estará fortificado con una capa de seguridad avanzada, brindando tranquilidad en la gestión y distribución de software crítico.