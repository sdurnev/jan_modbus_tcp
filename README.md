# jan-modbus-tcp


Read modbus arguments from 19000 address, and returns a json object.

Programm flags:

-ip - janitaza ip address (defaut value "localhost");

-port - janitaza modbus tcp port (defaut value 502);

-id - janitaza modbus slave ID (defaut value 1);

-q - quantity of janitaza modbus arguments, value range 1 - 61 (defaut value 61).

Example:

jan_modbus_ip_mac_0.01.2 -ip=192.168.10.10 -port=1502 -id=2 -q=10