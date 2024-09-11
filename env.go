package env

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Env struct {
	storage map[string]EnvEntry
}

type EnvEntry struct {
	Key   string
	Value string
}

func (e *EnvEntry) String() string {
	return e.Value
}

func (e *EnvEntry) Int() int {
	if val, err := strconv.Atoi(e.Value); err == nil {
		return val
	}
	return 0
}

func (e *EnvEntry) Bool() bool {
	if val, err := strconv.ParseBool(e.Value); err == nil {
		return val
	}
	return false
}

func (e *Env) OrElse(key string, fallback string) string {
	if entry, ok := e.storage[key]; ok {
		return entry.Value
	}
	return fallback
}

func (e *Env) OrElseInt(key string, fallback int) int {
	if entry, ok := e.Get(key); ok {
		return entry.Int()
	}
	return fallback
}

func (e *Env) OrElseBool(key string, fallback bool) bool {
	if entry, ok := e.Get(key); ok {
		return entry.Bool()
	}
	return fallback
}

func (e *Env) Get(key string) (EnvEntry, bool) {
	if entry, ok := e.storage[key]; ok {
		return entry, true
	}

	ee := os.Getenv(key)
	if ee != "" {
		entry := EnvEntry{Key: key, Value: ee}
		e.storage[key] = entry
		return entry, true
	}

	return EnvEntry{}, false
}

func NewEnv() *Env {
	env := &Env{storage: make(map[string]EnvEntry)}
	env.load(".env")
	return env
}

func (e *Env) load(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip comments or empty lines
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		// Split key and value by the first occurrence of '='
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove surrounding quotes if any
		value = strings.Trim(value, `"'`)

		// Store the key-value pair in the storage
		e.storage[key] = EnvEntry{Key: key, Value: value}
	}

	return scanner.Err()
}
