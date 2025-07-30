package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/karalabe/hid"
)

func getCPUTemp() int {

	paths := []string{
		"/sys/class/thermal/thermal_zone0/temp",
		"/sys/class/hwmon/hwmon0/temp1_input",
	}

	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err == nil {
			raw := strings.TrimSpace(string(data))
			val, err := strconv.Atoi(raw)
			if err == nil {
				return val / 1000
			}
		}
	}

	log.Println("Não foi possível ler temperatura")
	return 0
}

func main() {
	const vendorID = 0xaa88
	const productID = 0x8666

	devices := hid.Enumerate(0, 0)
	fmt.Printf("Encontrados %d dispositivos HID\n", len(devices))

	var targetDevice *hid.DeviceInfo
	for i, dev := range devices {
		if dev.VendorID == vendorID && dev.ProductID == productID {
			targetDevice = &devices[i]
		}
	}

	if targetDevice == nil {
		log.Fatalf("Dispositivo HID com ID %04x:%04x não encontrado", vendorID, productID)
	}

	device, err := targetDevice.Open()
	if err != nil {
		log.Fatalf("Erro ao abrir o dispositivo HID: %v", err)
	}
	defer device.Close()

	fmt.Println("Conectado ao dispositivo HID.")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		temp := getCPUTemp()
		buf := []byte{0x00, byte(temp)}

		_, err = device.Write(buf)
		if err != nil {
			log.Printf("Erro ao escrever: %v", err)
		}
	}
}
