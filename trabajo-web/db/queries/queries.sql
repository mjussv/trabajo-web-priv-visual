-- ===== USUARIOS =====
INSERT INTO usuarios (nombre, email) VALUES ($1, $2) RETURNING *;
SELECT * FROM usuarios WHERE id = $1;
SELECT * FROM usuarios ORDER BY nombre;

-- ===== PRENDAS ====
INSERT INTO prendas (usuario_id, nombre, categoria) VALUES ($1, $2, $3) RETURNING *;
SELECT * FROM prendas WHERE id = $1;
SELECT * FROM prendas WHERE usuario_id = $1 ORDER BY nombre;
DELETE FROM prendas WHERE id = $1;
DELETE FROM prendas WHERE usuario_id = $1;

-- ===== LOOKS (CRUD principal) =====
INSERT INTO looks (
    usuario_id, nombre, descripcion, ocasion, temporada, prenda_estrella_id, url_imagen, estilo
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

SELECT
    l.id, l.nombre, l.descripcion, l.ocasion, l.temporada, l.url_imagen, l.estilo, l.fecha_creacion,
    u.nombre AS autora,
    p.nombre AS prenda_estrella
FROM looks l
JOIN usuarios u ON l.usuario_id = u.id
LEFT JOIN prendas p ON l.prenda_estrella_id = p.id
WHERE l.id = $1;

SELECT
    l.id, l.nombre, l.descripcion, l.ocasion, l.temporada, l.url_imagen, l.estilo, l.fecha_creacion,
    u.nombre AS autora,
    p.nombre AS prenda_estrella
FROM looks l
JOIN usuarios u ON l.usuario_id = u.id
LEFT JOIN prendas p ON l.prenda_estrella_id = p.id
ORDER BY l.fecha_creacion DESC;

SELECT
    l.id, l.nombre, l.descripcion, l.ocasion, l.temporada, l.url_imagen, l.estilo, l.fecha_creacion,
    u.nombre AS autora,
    p.nombre AS prenda_estrella
FROM looks l
JOIN usuarios u ON l.usuario_id = u.id
LEFT JOIN prendas p ON l.prenda_estrella_id = p.id
WHERE l.usuario_id = $1
ORDER BY l.fecha_creacion DESC;

SELECT * FROM looks WHERE ocasion = $1 ORDER BY fecha_creacion DESC;

SELECT l.*
FROM looks l
JOIN prendas p ON l.prenda_estrella_id = p.id
WHERE p.nombre ILIKE '%' || $1 || '%'
ORDER BY l.fecha_creacion DESC;

UPDATE looks
SET nombre = $2,
    descripcion = $3,
    ocasion = $4,
    temporada = $5,
    prenda_estrella_id = $6,
    url_imagen = $7,
    estilo = $8
WHERE id = $1
RETURNING *;

DELETE FROM looks WHERE id = $1;

-- ===== ME GUSTA (like tipo Instagram) =====
INSERT INTO me_gusta_looks (usuario_id, look_id) VALUES ($1, $2)
ON CONFLICT DO NOTHING;

DELETE FROM me_gusta_looks WHERE usuario_id = $1 AND look_id = $2;

SELECT COUNT(*) FROM me_gusta_looks WHERE look_id = $1;

SELECT EXISTS (
    SELECT 1 FROM me_gusta_looks WHERE usuario_id = $1 AND look_id = $2
);
