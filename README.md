# Circula Server

API mínima em Go para validar a comunicação com o aplicativo Android do Circula Bazar, mantido no repositório `pelesoftwork-ship-it/circula`.

## Executar

Requer Go 1.23 ou superior. Ambiente validado com Go 1.27.3 no Windows.

```powershell
go run ./cmd/api
```

O servidor escuta somente em `127.0.0.1:8080`. A API usa a biblioteca padrão do Go e não requer banco de dados ou serviços externos nesta etapa.

## Endpoint de validação

`GET /health` retorna HTTP 200, `Content-Type: application/json; charset=utf-8` e:

```json
{"service":"circula-api","status":"ok"}
```

Para conferir localmente no PowerShell:

```powershell
Invoke-RestMethod http://127.0.0.1:8080/health
```

Para conectar o celular físico por USB, com depuração autorizada:

```powershell
adb devices -l
adb -s SERIAL_DO_CELULAR reverse tcp:8080 tcp:8080
```

Manter a API em execução e tocar em **Testar conexão** no aplicativo. O redirecionamento USB deve ser refeito ao reconectar o aparelho. O uso de HTTP local se destina à validação do ambiente; o aplicativo Android o permite somente na variante debug.

## Verificação e compilação

```powershell
go vet ./...
go build -o bin/circula-api.exe ./cmd/api
```

Esta versão contém apenas o endpoint de saúde. Autenticação, persistência e funcionalidades do marketplace serão implementadas posteriormente.
