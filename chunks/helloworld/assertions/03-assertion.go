package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
import (
    "log"
    "net"
    "time"
)

var psgDataTypes = []string{
    "BOOLEAN",
    "FLOAT",
    "DECIMAL",
    "NUMERIC",
    "DOUBLE PRECISION",
    "REAL",
    "INT",
    "INTEGER",
    "SMALLINT",
    "BIGINT",
    "SERIAL",
    "BIGSERIAL",
    "VARCHAR",
    "TEXT",
    "CHAR",
    "DATE",
    "TIME",
    "TIMESTAMP",
    "TIMESTAMPTZ",
    "INET",
}

func ConvertDatabaseValue(columnType string, val interface{}, col string) (interface{}, error) {
    if val == nil {
        return nil, nil // Возвращаем nil, если значение NULL
    }

    switch columnType {
    case "INT", "INTEGER", "SMALLINT", "BIGINT", "SERIAL", "BIGSERIAL":
        if v, ok := val.(int64); ok {
            return v, nil
        }
        log.Printf("unexpected type for column %s: expected int64, got %T", col, val)
        return 0, nil // Значение по умолчанию для целых чисел

    case "FLOAT", "DOUBLE PRECISION", "REAL", "NUMERIC", "DECIMAL":
        if v, ok := val.(float64); ok {
            return v, nil
        }
        log.Printf("unexpected type for column %s: expected float64, got %T", col, val)
        return 0.0, nil // Значение по умолчанию для чисел с плавающей запятой

    case "BOOLEAN":
        if v, ok := val.(bool); ok {
            return v, nil
        }
        log.Printf("unexpected type for column %s: expected bool, got %T", col, val)
        return false, nil // Значение по умолчанию для булевых значений

    case "VARCHAR", "TEXT", "CHAR":
        if v, ok := val.([]byte); ok {
            return string(v), nil
        }
        log.Printf("unexpected type for column %s: expected []byte, got %T", col, val)
        return "", nil // Значение по умолчанию для строк

    case "DATE", "TIME", "TIMESTAMP", "TIMESTAMPTZ":
        if v, ok := val.(time.Time); ok {
            return v.Format(time.RFC3339), nil
        }
        log.Printf("unexpected type for column %s: expected time.Time, got %T", col, val)
        return time.Time{}.Format(time.RFC3339), nil // Значение по умолчанию для времени

    case "INET", "CIDR":
        if v, ok := val.([]byte); ok {
            return net.IP(v).String(), nil
        }
        log.Printf("unexpected type for column %s: expected []byte, got %T", col, val)
        return "", nil // Значение по умолчанию для IP-адресов

    default:
        if v, ok := val.([]byte); ok {
            return string(v), nil
        }
        log.Printf("unexpected type for column %s: expected []byte, got %T", col, val)
        return "", nil // Значение по умолчанию для неизвестных типов
    }
}
