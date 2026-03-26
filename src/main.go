package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Penggunaan: ./voxelizer <path_file.obj> <max_depth>")
		os.Exit(1)
	}

	inputPath := os.Args[1]
	maxDepth, err := strconv.Atoi(os.Args[2])
	if err != nil || maxDepth < 0 {
		log.Fatal("max_depth harus angka positif.")
	}

	model, err := bacaObj(inputPath)
	if err != nil {
		log.Fatalf("gagal baca file: %v", err)
	}

	fmt.Printf("Memproses %d faces...\n", len(model.Faces))

	startTime := time.Now()

	rootBox := hitungRootBox(model.Vertices)
	rootNode := &OctreeNode{
		Box: rootBox, Depth: 0,
	}

	stats := newStats()

	buildOctree(rootNode, model.Faces, model.Vertices, maxDepth, stats)

	var voxels []BoundingBox
	collectVoxels(rootNode, &voxels)

	ext := filepath.Ext(inputPath)
	base := inputPath[0 : len(inputPath)-len(ext)]
	outputPath := base + "-voxelized.obj"

	nVert, nFace, err := tulisObj(voxels, outputPath)
	if err != nil {
		log.Fatalf("gagal tulis output: %v", err)
	}

	elapsedTime := time.Since(startTime)

	fmt.Println("========================================")
	fmt.Printf("Jumlah voxel  : %d\n", stats.TotalVoxels)
	fmt.Printf("Jumlah vertex : %d\n", nVert)
	fmt.Printf("Jumlah face   : %d\n", nFace)

	fmt.Println("\nStatistik node terbentuk pada depth:")
	printDepthStats(stats.NodesAtDepth, maxDepth)

	fmt.Println("\nStatistik node yang tidak perlu ditelusuri:")
	printDepthStats(stats.PrunedAtDepth, maxDepth)

	fmt.Printf("\nMax depth        : %d\n", maxDepth)
	fmt.Printf("Waktu eksekusi     : %v\n", elapsedTime)
	fmt.Printf("Output disimpam di : %s\n", outputPath)
	fmt.Println("========================================")
}

func printDepthStats(m map[int]int, maxDepth int) {
	var depths []int
	for d := range m {
		if d > 0 {
			depths = append(depths, d)
		}
	}

	sort.Ints(depths)
	for _, d := range depths {
		fmt.Printf("%d : %d\n", d, m[d])
	}
}
