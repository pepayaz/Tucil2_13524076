package main

import (
	"bufio"
	"fmt"
	"os"
)

func collectVoxels(node *OctreeNode, voxels *[]BoundingBox) {
	if node == nil {
		return
	}
	if node.IsLeaf && node.IsFilled {
		*voxels = append(*voxels, node.Box)
		return
	}
	for i := 0; i < 8; i++ {
		if node.Children[i] != nil {
			collectVoxels(node.Children[i], voxels)
		}
	}
}

func tulisObj(voxels []BoundingBox, outputPath string) (int, int, error) {
	file, err := os.Create(outputPath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	vertexCount := 0
	faceCount := 0

	for _, voxel := range voxels {
		h := voxel.HalfDimension
		c := voxel.Center

		vertices := [8]Vector3{
			{X: c.X - h, Y: c.Y - h, Z: c.Z - h},
			{X: c.X + h, Y: c.Y - h, Z: c.Z - h},
			{X: c.X + h, Y: c.Y + h, Z: c.Z - h},
			{X: c.X - h, Y: c.Y + h, Z: c.Z - h},
			{X: c.X - h, Y: c.Y - h, Z: c.Z + h},
			{X: c.X + h, Y: c.Y - h, Z: c.Z + h},
			{X: c.X + h, Y: c.Y + h, Z: c.Z + h},
			{X: c.X - h, Y: c.Y + h, Z: c.Z + h},
		}

		for _, v := range vertices {
			fmt.Fprintf(writer, "v %f %f %f\n", v.X, v.Y, v.Z)
		}

		offset := vertexCount

		faces := [12][3]int{
			{1, 3, 2}, {1, 4, 3}, // Bawah
			{5, 6, 7}, {5, 7, 8}, // Atas
			{1, 2, 6}, {1, 6, 5}, // Depan
			{4, 7, 3}, {4, 8, 7}, // Belakang
			{1, 5, 8}, {1, 8, 4}, // Kiri
			{2, 3, 7}, {2, 7, 6}, // Kanan
		}

		for _, f := range faces {
			fmt.Fprintf(writer, "f %d %d %d\n", f[0]+offset, f[1]+offset, f[2]+offset)
		}

		vertexCount += 8
		faceCount += 12
	}

	writer.Flush()
	return vertexCount, faceCount, nil
}
