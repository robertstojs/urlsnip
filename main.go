package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "log/syslog"
    "net/http"
    "regexp"
)

type URLMapping struct {
    ShortURL     string `json:"shortURL"`
    OriginalURL  string `json:"originalURL"`
    RegexPattern string `json:"regexPattern,omitempty"`
}

var (
    urlMappings []URLMapping
    sysLogger   *syslog.Writer
    err         error
)

func initLogger() {
    // Setup syslog for logging.
    sysLogger, err = syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, "urlsnip")
    if err != nil {
        log.Fatalf("Failed to initialize syslog: %v", err)
    }
}

func loadConfig(filePath string) {
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        logMessage := fmt.Sprintf("Failed to read config file: %v", err)
        log.Fatal(logMessage)
        sysLogger.Err(logMessage)
    }

    err = json.Unmarshal(data, &urlMappings)
    if err != nil {
        logMessage := fmt.Sprintf("Failed to parse config file: %v", err)
        log.Fatal(logMessage)
        sysLogger.Err(logMessage)
    }

    log.Println("Configuration file loaded successfully")
    sysLogger.Info("Configuration file loaded successfully")
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
    requestedURL := r.URL.Path[1:]

    for _, mapping := range urlMappings {
        if mapping.ShortURL == requestedURL {
            http.Redirect(w, r, mapping.OriginalURL, http.StatusFound)
            logMessage := fmt.Sprintf("Redirected '%s' to '%s'", requestedURL, mapping.OriginalURL)
            log.Println(logMessage)
            sysLogger.Info(logMessage)
            return
        }

        if mapping.RegexPattern != "" {
            matched, regexErr := regexp.MatchString(mapping.RegexPattern, requestedURL)
            if regexErr != nil {
                logMessage := fmt.Sprintf("Regex error for pattern '%s': %v", mapping.RegexPattern, regexErr)
                log.Println(logMessage)
                sysLogger.Err(logMessage)
                continue
            }
            if matched {
                http.Redirect(w, r, mapping.OriginalURL, http.StatusFound)
                logMessage := fmt.Sprintf("Redirected via regex '%s' to '%s' using pattern '%s'", requestedURL, mapping.OriginalURL, mapping.RegexPattern)
                log.Println(logMessage)
                sysLogger.Info(logMessage)
                return
            }
        }
    }

    logMessage := fmt.Sprintf("No redirect mapping found for url '%s'", requestedURL)
    log.Println(logMessage)
    sysLogger.Warning(logMessage)
    http.NotFound(w, r)
}

func main() {
    configFilePath := flag.String("config", "config.json", "Path to the configuration file")
    port := flag.Int("port", 8080, "Port on which the server will run")
    flag.Parse()

    initLogger()
    defer sysLogger.Close()

    loadConfig(*configFilePath)

    http.HandleFunc("/", redirectHandler)

    serverAddress := fmt.Sprintf(":%d", *port)
    log.Printf("Server starting on %s", serverAddress)
    sysLogger.Info(fmt.Sprintf("Server starting on %s", serverAddress))

    err = http.ListenAndServe(serverAddress, nil)
    if err != nil {
        sysLogger.Err(fmt.Sprintf("Server failed to start: %v", err))
        log.Fatalf("Server failed to start: %v", err)
    }
}
