#!/bin/bash

# Configuración
CONTAINER_NAME="dimisa"
BACKUP_FILE="/home/informatica/DIMISA/DIMISA_API/backup_mysql.sql.gz"
LOG_FILE="/home/informatica/DIMISA/DIMISA_API/backup.log"

# Registrar inicio
echo "=== Respaldo iniciado: $(date) ===" >> $LOG_FILE

# Realizar respaldo (sobreescribe el archivo anterior)
docker exec $CONTAINER_NAME mysqldump -u root -p'hgmaza25' --all-databases 2>> $LOG_FILE | gzip > $BACKUP_FILE

# Verificar resultado
if [ $? -eq 0 ]; then
    echo "Respaldo completado exitosamente: $BACKUP_FILE" >> $LOG_FILE
    echo "Tamaño: $(du -h $BACKUP_FILE | cut -f1)" >> $LOG_FILE
else
    echo "ERROR: Falló el respaldo" >> $LOG_FILE
fi

echo "" >> $LOG_FILE
