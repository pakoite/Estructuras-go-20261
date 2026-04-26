# Examen Práctico: Implementación de Juego UNO

## Objetivo

Desarrollar una versión simplificada del juego UNO utilizando el lenguaje **Go**, aplicando correctamente estructuras de datos (`struct`, slices), control de flujo y lógica de juego.

---

## Requerimientos del sistema

Debes implementar un programa que simule una partida básica de UNO entre **2 jugadores**.

### 1. Estructuras obligatorias

Debes definir al menos los siguientes `struct`:

* `Card` → color y valor
* `Player` → nombre y mano de cartas
* `Game` → estado general del juego (mazo, jugadores, turno, pila de descarte)

---

### 2. Funcionalidades mínimas

Tu programa debe incluir:

* Crear un mazo de cartas (mínimo números del 0 al 9 por color)
* Mezclar el mazo
* Repartir cartas a cada jugador (mínimo 5)
* Mantener una pila de descarte

---

### 3. Lógica del juego

En cada turno:

* Mostrar las cartas del jugador actual
* Permitir seleccionar una carta para jugar
* Validar que la carta sea válida:

  * mismo color **o**
  * mismo número

Si el jugador **no puede jugar**:

* Debe robar una carta

---

### 4. Flujo del juego

* Alternar turnos entre jugadores
* El juego termina cuando:

  * un jugador se queda sin cartas

---

### 5. Salida en consola

El programa debe mostrar:

* Turno actual
* Cartas del jugador
* Carta en la pila de descarte
* Acciones realizadas

---

## <span style="color:red">**Restricciones**
</span>

* No usar librerías externas

* No copiar código sin entenderlo

* Debes poder explicar tu implementación

---

## Criterios de evaluación

### 1. Uso de estructuras (30%)

* Uso correcto de `struct`
* Organización del código

### 2. Lógica del juego (30%)

* Validación correcta de jugadas
* Flujo de turnos funcional

### 3. Funcionalidad (20%)

* El programa corre sin errores
* Se puede completar una partida

### 4. Claridad del código (10%)

* Nombres claros
* Código legible

### 5. Defensa (10%)

* Capacidad de explicar el código
* Responder preguntas sobre su funcionamiento

---

## Puntos extra (opcional)

* Implementar cartas especiales (+2, reversa, salto)

* Más de 2 jugadores

* Interfaz mejorada en consola

---

## Nota importante

Puedes apoyarte en herramientas externas (incluyendo IA), pero **debes comprender completamente tu código**, ya que se te pedirá explicarlo y modificarlo durante la evaluación.

---

## Advertencia

Si no puedes explicar tu código o hacer modificaciones simples en clase, se considerará como **no acreditado**, independientemente de que el programa funcione.

---
