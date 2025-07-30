# CPU Temperature Writer / Escritor de Temperatura da CPU

## English

### Description

This Go application reads the CPU temperature from the system and sends it to a TGT cooler device via HID (Human Interface Device) communication. The program continuously monitors the CPU temperature and transmits the data to the cooler every second, allowing for real-time temperature-based cooling control.

### Features

- **CPU Temperature Reading**: Reads temperature from multiple system paths (`/sys/class/thermal/thermal_zone0/temp` and `/sys/class/hwmon/hwmon0/temp1_input`)
- **HID Communication**: Connects to TGT cooler device using vendor ID `0xaa88` and product ID `0x8666`
- **Real-time Monitoring**: Updates temperature data every second
- **Error Handling**: Graceful error handling for device connection and temperature reading failures
- **System Service**: Can be configured as a systemd service for automatic startup

### How it Works

1. The application enumerates all HID devices on the system
2. Searches for the specific TGT cooler device using its vendor and product IDs
3. Opens a connection to the device
4. Continuously reads CPU temperature from system thermal sensors
5. Sends temperature data to the cooler device every second

### Installation and Setup

#### 1. Build the Application

```bash
go build -o cpu-temp-writer
sudo mv cpu-temp-writer /usr/local/bin/
```

#### 2. Create System Service

Create the systemd service file:

```bash
sudo nano /etc/systemd/system/cpu-temp-writer.service
```

Add the following content:

```ini
[Unit]
Description=Sends CPU temperature to TGT cooler
After=multi-user.target

[Service]
ExecStart=/usr/local/bin/cpu-temp-writer
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

#### 3. Enable and Start the Service

```bash
sudo systemctl daemon-reload
sudo systemctl enable cpu-temp-writer.service
sudo systemctl start cpu-temp-writer.service
```

#### 4. Check Service Status

```bash
systemctl status cpu-temp-writer.service
```

### Requirements

- Go 1.16 or higher
- Linux system with thermal sensors
- TGT cooler device connected via USB
- Root privileges for HID device access

---

## Português

### Descrição

Esta aplicação Go lê a temperatura da CPU do sistema e a envia para um dispositivo cooler TGT via comunicação HID (Human Interface Device). O programa monitora continuamente a temperatura da CPU e transmite os dados para o cooler a cada segundo, permitindo controle de resfriamento baseado na temperatura em tempo real.

### Funcionalidades

- **Leitura de Temperatura da CPU**: Lê temperatura de múltiplos caminhos do sistema (`/sys/class/thermal/thermal_zone0/temp` e `/sys/class/hwmon/hwmon0/temp1_input`)
- **Comunicação HID**: Conecta ao dispositivo cooler TGT usando vendor ID `0xaa88` e product ID `0x8666`
- **Monitoramento em Tempo Real**: Atualiza dados de temperatura a cada segundo
- **Tratamento de Erros**: Tratamento gracioso de erros para falhas de conexão do dispositivo e leitura de temperatura
- **Serviço do Sistema**: Pode ser configurado como serviço systemd para inicialização automática

### Como Funciona

1. A aplicação enumera todos os dispositivos HID no sistema
2. Procura pelo dispositivo cooler TGT específico usando seus IDs de vendor e produto
3. Abre uma conexão com o dispositivo
4. Lê continuamente a temperatura da CPU dos sensores térmicos do sistema
5. Envia dados de temperatura para o dispositivo cooler a cada segundo

### Instalação e Configuração

#### 1. Compilar a Aplicação

```bash
go build -o cpu-temp-writer
sudo mv cpu-temp-writer /usr/local/bin/
```

#### 2. Criar Serviço do Sistema

Crie o arquivo de serviço systemd:

```bash
sudo nano /etc/systemd/system/cpu-temp-writer.service
```

Adicione o seguinte conteúdo:

```ini
[Unit]
Description=Envia a temperatura da CPU para o cooler TGT
After=multi-user.target

[Service]
ExecStart=/usr/local/bin/cpu-temp-writer
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

#### 3. Ativar e Iniciar o Serviço

```bash
sudo systemctl daemon-reload
sudo systemctl enable cpu-temp-writer.service
sudo systemctl start cpu-temp-writer.service
```

#### 4. Verificar Status do Serviço

```bash
systemctl status cpu-temp-writer.service
```

### Requisitos

- Go 1.16 ou superior
- Sistema Linux com sensores térmicos
- Dispositivo cooler TGT conectado via USB
- Privilégios de root para acesso ao dispositivo HID

### Estrutura do Código

- `getCPUTemp()`: Função que lê a temperatura da CPU dos sensores do sistema
- `main()`: Função principal que gerencia a conexão HID e o loop de monitoramento
- Comunicação via HID usando a biblioteca `github.com/karalabe/hid`
- Timer de 1 segundo para atualizações regulares de temperatura

### Troubleshooting

- Certifique-se de que o dispositivo TGT está conectado e reconhecido pelo sistema
- Verifique se os sensores térmicos estão disponíveis nos caminhos especificados
- Execute com privilégios de administrador para acesso aos dispositivos HID
- Verifique os logs do serviço com `journalctl -u cpu-temp-writer.service`
