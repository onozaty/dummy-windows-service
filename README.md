# dummy-windows-service

It is a Windows service that does nothing.  
Used to confirm service start and stop operations.

## Usage

You can download the executable file from the following link.

* https://github.com/onozaty/dummy-windows-service/releases/latest

Register as a service with `dummy-windows-service.exe install`.

```
C:\dummy-windows-service>dummy-windows-service.exe install
2022/04/10 16:41:27 Action(install) succeeded.
```

The service name is `DummyWindowsService`.

You can start the service with `start` and stop it with `stop`.

```
C:\dummy-windows-service>dummy-windows-service.exe start
2022/04/10 16:47:10 Action(start) succeeded.

C:\dummy-windows-service>dummy-windows-service.exe stop
2022/04/10 16:50:07 Action(stop) succeeded.
```

`uninstall` uninstalls the service.

```
C:\dummy-windows-service>dummy-windows-service.exe uninstall
2022/04/10 16:52:28 Action(uninstall) succeeded.
```