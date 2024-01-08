Prácticas: Servers

# Ejercicio 1

Asumiendo que realizó los ejercicios para sumar los primeros n números pares y los primeros n números impares. Les voy a pedir que hagan lo siguiente:

1. Crear un proyecto golang donde las funciones sean parte de un paquete que no sea el main (Por ejemplo: puede llamarlo _numbers)_ y que tenga un module path que sea un URL apuntando a un repositorio accesible.
2. Crearse otro proyecto golang para escribir un servidor que tenga dos endpoints
   1. /api/v1/sumEvens: Recibe como parámetro un número n y suma los primeros “n” números pares y devuelve en el cuerpo de la respuesta el resultado de dicha suma.
   2. /api/v1/sumOdds: Recibe como parámetro un número n y suma los primeros “n” números impares y devuelve en el cuerpo de la respuesta el resultado de dicha suma.

Se pide intencionalmente que NO se escriban las rutinas de código dentro del proyecto del servidor, sino que se importen usando los mecanismo de dependencias que tiene el lenguaje.

1. Se pide hacer ambos endpoints para recibir peticiones GET y POST, entender que los parámetros de la petición “viajan” de distinta forma si es un GET o si es un POST. Si es un GET son parte de la URL de la petición y si es un POST es parte del cuerpo de la petición.

Si te parece más sencillo podés tener 4 endpoints, 2 para GET y 2 para POST.

# Ejercicio 2

Escribir un programa que levante un servidor en el puerto 9090. El programa debe mantener en memoria una variable entera definida a nivel “global”. Además configurar el servidor para tener dos endpoints.

1. POST /api/v1/inc: Incrementa la variable entera en uno
2. POST /api/v1/dec: Decrementa la variable entera en uno
3. GET /api/v1/curr: Devuelve el valor de la variable entera.

Una vez escrito este programa servidor, vas a escribir un programa que juegue como cliente y tenga dos gorutinas. En una vas a enviar 10000 peticiones para incrementar y otra 1000 peticiones para decrementar. Finalmente vas a mandar una petición para obtener el valor de la variable entera una vez que terminen las dos gorutinas.

Se pide ejecutar el programa servidor y luego :

1. Ejecutando una vez el programa cliente
2. Ejecuta varias instancias del mismo programa cliente, es decir, ejecútalo varias veces en simultáneo.

Por último, se podrá apreciar que en el segundo caso para algún cliente el resultado que se le devuelve no es cero. ¿Esto a qué se debe? ¿ que hay que modificar para que a ambos clientes la petición “final” de obtener el valor de la variable entera devuelva un cero?

**Resoluciones posibles** _(if ❌🧠 then !👀, prohibido mirar sin antes pensar)_:

[int_server.go](https://go.dev/play/p/4AMPOqaM7pr)

[int_cliente.go](https://go.dev/play/p/DZR_R88SprX)

# Ejercicio 3

Desarrollar una API en Go que maneje saludos personalizados y mantenga un registro de los nombres ya saludados. La API deberá ofrecer respuestas específicas basadas en si el nombre ha sido previamente saludado o no.

### Endpoints de la API

1. **POST /hello**:
   - Funcionalidad: Saluda al nombre proporcionado en el cuerpo de la solicitud.
   - Comportamiento: Si el nombre ya ha sido saludado antes, la API debe ofrecer una respuesta de "hola de nuevo" en lugar de un simple saludo.
2. **GET /hello**:
   - Funcionalidad: Devuelve una lista de todos los nombres que han sido saludados.

### Requisitos Específicos

1. **Manejo de Estado**:
   - La API debe ser capaz de recordar los nombres que ya ha saludado. Considera el uso de una estructura de datos adecuada para almacenar estos nombres entre solicitudes.
2. **Validaciones**:
   - Asegúrate de que la API maneje adecuadamente las entradas inválidas y proporcione mensajes de error claros.
3. **Paquetes**
   - Se pide intencionalmente que NO se escriban las rutinas de código dentro del proyecto del servidor, sino que se importen usando los mecanismo de dependencias que tiene el lenguaje.

<aside>
💡 Ten especial cuidado con las condiciones de carrera al acceder o modificar el registro de nombres. Asegúrate de que las operaciones sobre la estructura de datos sean seguras.

</aside>
