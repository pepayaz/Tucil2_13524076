package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bacaObj(filePath string) (*Model, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("gagal buka file: %v", err)
	}
	defer file.Close()

	model := &Model{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		if parts[0] == "v" && len(parts) >= 4 {
			x, errX := strconv.ParseFloat(parts[1], 64)
			y, errY := strconv.ParseFloat(parts[2], 64)
			z, errZ := strconv.ParseFloat(parts[3], 64)

			if errX == nil && errY == nil && errZ == nil {
				model.Vertices = append(model.Vertices, Vector3{X: x, Y: y, Z: z})
			}

		}

		if parts[0] == "f" && len(parts) >= 4 {
			v1Str := strings.Split(parts[1], "/")[0]
			v2Str := strings.Split(parts[2], "/")[0]
			v3Str := strings.Split(parts[3], "/")[0]

			v1, err1 := strconv.Atoi(v1Str)
			v2, err2 := strconv.Atoi(v2Str)
			v3, err3 := strconv.Atoi(v3Str)

			if err1 == nil && err2 == nil && err3 == nil {
				model.Faces = append(model.Faces, Segitiga{V1: v1 - 1, V2: v2 - 1, V3: v3 - 1})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error saat membaca file: %v", err)
	}

	if len(model.Vertices) == 0 || len(model.Faces) == 0 {
		return nil, fmt.Errorf("file input tidak valid: tidak ada data vertex atau face yang ditemukan")
	}

	return model, nil
}
