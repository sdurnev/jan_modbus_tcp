# jan-modbus-tcp


Read modbus arguments from 19000 address, and returns a json object.

Programm flags:

-ip - janitza ip address (defaut value "localhost");

-port - janitza modbus tcp port (defaut value 502);

-id - janitza modbus slave ID (defaut value 1);

-q - quantity of janitza modbus arguments, value range 1 - 61 (defaut value 61).

Example:

`jan_modbus_tcp -ip=192.168.10.10 -port=1502 -id=2 -q=10`
