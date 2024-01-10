-- Ejercicio 1
-- Dada una base de datos con las siguientes tablas
CREATE TABLE department (
    id SERIAL PRIMARY KEY,
    name CHAR(50)
);

INSERT INTO department (id, name)
VALUES
(1, 'Producción'),
(2, 'Ventas'),
(3, 'RRHH'),
(4, 'Contabilidad');

CREATE TABLE employee (
    id SERIAL PRIMARY KEY,
    name CHAR(50),
    address CHAR(50),
    department_id INT,
    FOREIGN KEY (department_id) REFERENCES department(id)
);

INSERT INTO employee (id, name, address, department_id)
VALUES
(1, 'Carlos', 'Direcc1', 1),
(2, 'Pepe', 'Direcc2', 2),
(3, 'Susana', 'Direcc3', 2),
(4, 'Graciela', 'Direcc4', 3);

CREATE TABLE project (
    id SERIAL PRIMARY KEY,
    name CHAR(50)
);

INSERT INTO project (id, name)
VALUES
(1, 'Desarrollo'),
(2, 'Sistema de Gestión'),
(3, 'Capacitación');

CREATE TABLE assignment (
    employee_id INT,
    project_id INT,
    FOREIGN KEY (employee_id) REFERENCES employee(id),
    FOREIGN KEY (project_id) REFERENCES project(id)
);

INSERT INTO assignment (employee_id, project_id)
VALUES
(1, 3),
(2, 1),
(3, 2);

--1) ¿Se podría eliminar la tabla employee? ¿Por qué? Establezca cuales tablas pueden ser eliminadas y cuáles no.
/*
Si se intenta eliminar una tabla que esta siendo referenciada por una clave foranea en otra tabla, PostgreSQL no lo permitira
y aparecera un mensaje de error, por ejemplo, las tablas 'employee', 'department' y 'project'.
Por otro lado las tablas que no estan siendo referenciadas en otra tabla, como es el caso de 'assigment',
pueden ser eliminadas facilmente.
*/

--2) Crear una tabla hobby, en donde se pondrá el hobby de cada empleado.
CREATE TABLE hobby (
    employee_id INT,
    name CHAR(50),
    FOREIGN KEY (employee_id) REFERENCES employee(id)
);

INSERT INTO hobby (employee_id, name)
VALUES
(1, 'Pesca'),
(2, 'Senderismo'),
(3, 'Baile'),
(4, 'Ajedrez');

--3) Inserte en la tabla assignment, el registro correspondiente a Graciela, que trabaja en el proyecto de Capacitación.
INSERT INTO assignment (employee_id, project_id)
VALUES (4,3);

--4) Supongamos que ingresa un nuevo empleado a la empresa.
--Inserte todos los registros correspondientes para que queden registrados todos sus datos (incluídos el proyecto donde trabaja).
INSERT INTO employee (id, name, address, department_id)
VALUES (5, 'Tomas', 'Direcc5', 4);

INSERT INTO assignment (employee_id, project_id)
VALUES (5,1);

INSERT INTO hobby (employee_id, name)
VALUES (5, 'Juegos de video');

-- Ejercicio 2
-- 1) Crea una tabla llamada movie con los siguientes campos y restricciones.
CREATE TABLE movie (
    id SERIAL PRIMARY KEY NOT NULL,
    name CHAR(50) NOT NULL,
    acquired DATE,
    stock INT DEFAULT 10,
    price DECIMAL(2) NOT NULL
);

-- 2) Indica si las siguientes sentencias INSERT son correctas. En el caso de que no sean correctas indica el porqué:

-- a)
INSERT INTO movie (id, name, acquired, stock, price) VALUES (1,1704398539,20,60)
/*
Es incorrecta porque no pasa los valores correspondientes a los campos correctamente, ni pasa los tipos de datos correctos.
*/

-- b)
INSERT INTO movie VALUES ('2','El señor de los anillos',20);
/*
Es incorrecta porque no especifica los campos para los valores, intenta pasar un INTEGER en la columna de tipo DATE,
y ademas viola la restriccion NOT NUL del campo 'price'.
*/

-- c)
INSERT INTO movie(id, name, price) VALUES (3,'Las dos torres',240000);
/*
Es incorrecta porque intenta insertar en el campo 'price' un numero con mas decimales de los especificados,
generando un error de tipo 'numeric field overflow'.
*/

-- d)
INSERT INTO movie(id, name, acquired) VALUES (4,'El retorno del rey',1600000000);
/*
Es incorrecta porque intenta pasar un INTEGER en la columna de tipo DATE,
ademas viola la restriccion NOT NUL del campo 'price'.
*/
