/*
1) Crear una tabla llamada GASTO para almacenar registros de gastos hogareños.
La misma debe considerar para un gasto: el concepto, el importe y la fecha.
*/
CREATE TABLE gasto (
    concepto CHAR(50) NOT NULL
    importe INT NOT NULL
    fecha DATE NOT NULL
);

/*
2) Realice mediante el lenguaje SQL la inserción de cuatro registros para la tabla anterior.
Un correspondiente a un gasto de importe 50 $, por compra de artículos, del día 10/10/2010,
otro de 165 $ por compra semanal de comida del 20/10/2010,
otro en concepto de reparación auto por 2400$ correspondiente al día 14/11/2010 y
otro por servicios de cable por 240$ correspondientes al mes de octubre
(Podría interpretarse como pagados el 10/11/2010).
*/
INSERT INTO gasto (concepto, importe, fecha)
VALUES
('compra de artículos', 50, '10/10/2010'),
('compra semanal de comida', 165, '20/10/2010'),
('reparación auto', 2400, '14/11/2010'),
('servicios de cable', 240, '10/11/2010');

/*
3) Realice una consulta para obtener todos los registros de la tabla mostrando todos los campos.
*/
SELECT * FROM gasto;

/*
4) Realice una consulta para obtener todos los registros cuya fecha de la tabla sea mayor a 1/11/2010
(Octubre en adelante) mostrando todos los campos.
*/
SELECT * FROM gasto
WHERE fecha >= '01/11/2010';

/*
5) Realice una consulta para obtener todos los registros de la tabla cuya fecha sea menor a 1/11/2010
(Septiembre para atrás) y su importe sea mayor a cien pesos mostrando todos los campos.
*/
SELECT * FROM gasto
WHERE fecha < '01/11/2010'
AND importe > 100;

/*
6) Realice mediante el lenguaje SQL una actualización de los que cuya fecha sea menor a 15/11/2010
modificando la fecha a la fecha actual.
*/
UPDATE gasto
SET fecha = '05/01/2024'
WHERE fecha < '15/11/2010';

/*
¿Una vez realizada la actualización, para los registros afectados en la actualización anterior podría modificar su importe?
¿Qué tendría que haber hecho?
*/
UPDATE gasto
SET fecha = '05/01/2024', importe = 1000
WHERE fecha < '15/11/2010';

/*
7) Realice mediante el lenguaje SQL un borrado de los registros cuyo importe sea mayor a 200$.
*/
DELETE FROM gasto
WHERE importe > 200;

/*
8) ¿Pueden existir dos registros totalmente iguales? ¿Cómo se podría diferenciarlos?
*/
-- Pueden existir mientras no se especifique la restricción UNIQUE, se podria diferenciarlos agregando un campo ID a cada registro.
