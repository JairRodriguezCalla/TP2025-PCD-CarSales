#define N 3  // Cantidad de procesos

chan sem = [1] of { bit };  // Semáforo binario, implementa exclusión mutua

bool critical = false;  // Recurso compartido para validar la exclusión mutua

proctype Process(byte id) {
    do
    ::  // Intentar entrar a la sección crítica
        sem ? _;  // P (wait) - Adquiere el semáforo
        
        // Sección Crítica
        assert(!critical);  // Verifica que nadie más está en la sección crítica
        critical = true;
        printf("Proceso %d ENTRANDO a sección crítica\n", id);

        // Simula trabajo en la sección crítica
        critical = false;
        printf("Proceso %d SALIENDO de sección crítica\n", id);

        sem ! 0;  // V (signal) - Libera el semáforo

        // Sección No Crítica (puedes simular trabajo aquí si deseas)
    od
}

init {
    // Inicializar el semáforo con un recurso disponible
    sem ! 0;

    atomic {
        run Process(1);
        run Process(2);
        run Process(3);
    }
}
