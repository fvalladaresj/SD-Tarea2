# Sistemas Distrubuidos - Tarea 2

## Integrantes
* Vicente Hoare Silva - 201773050-1
* Florencia Sofía Valladares Jerez - 201773007-2 

# Instrucciones
## En cada máquina iniciando sesion como el usuario 'alumno'
```go
export GOPATH="$HOME/go"
export PATH=$PATH:"$HOME/go/bin" 
```
## Iniciar procesos
En cada maquina cuando se inicia sesion como 'alumno' hay que dirigirse a:

```
 /home/alumno/SD-Tarea2 
```

luego en esta carpeta se pueden ejecutar los comandos correspondientes a cada maquina listados a continuacion:  
Iniciar los procesos en su máquina correspondiente
### maquina: dist133 - Jugador
> make jugador
### maquina: dist134 - Lider
> make lider
### maquina: dist135 - NameNode
> make namenode
### maquina: dist136 - Pozo
> make pozo

## Jugar
1. En el **Jugador** ingresar **participar** para unirse al juego
2. En el **Lider** ingresar 1 para iniciar etapa o 2 para ver jugadas de un jugador en específico
3. En el **Jugador** ingresar jugadas según corresponda
4. Al finalizar cada etapa, en el **Jugador** se debe ingresar 1 para continuar a la siguiente, o 2 para ver el monto en el Pozo (solo es posible tras finalizar la etapa 1 o 2) 
5. Mediante el **Lider** se debe dar inicio a la siguiente etapa ingresando 1.

## Consultar jugadas de un jugador
En el **Lider** en vez de dar inicio a la etapa, se debe ingresar 2 para ver las jugadas de un jugador. Luego ingresar la id del jugador (numerados de 0 a 15, siendo 0 el jugador "humano")

## Jugador "humano" ha muerto
En caso de que el jugador "humano" pierda, se continuará automáticamente hasta finalizar la etapa y luego mediante el **Lider** se volvera a preguntar si desea dar inicio a la siguiente etapa o consultar las jugadas de un jugador.

## Juego finalizado 
Al terminar un juego completamente, usar CTRL+C [^1] para terminar los procesos y luego ingresar el siguiente comando en cada máquina para limpiar los archivos generados.
> make clean 

[^1]: Si bien el proceso del jugador termina automáticamente al ganar o ser eliminado, el proceso del data node debe seguir ejecutándose y debe ser terminado manualmente.