package main

import "sync"

type OctreeNode struct {
	Box      BoundingBox
	Depth    int
	IsLeaf   bool
	IsFilled bool
	Children [8]*OctreeNode
}

type Stats struct {
	TotalVoxels   int
	NodesAtDepth  map[int]int
	PrunedAtDepth map[int]int
	mu            sync.Mutex
}

func newStats() *Stats {
	return &Stats{
		NodesAtDepth:  make(map[int]int),
		PrunedAtDepth: make(map[int]int),
	}
}

func (s *Stats) addNode(depth int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.NodesAtDepth[depth]++
}

func (s *Stats) addPruned(depth int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.PrunedAtDepth[depth]++
}

func (s *Stats) addVoxel() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.TotalVoxels++
}

func buildOctree(node *OctreeNode, faces []Segitiga, vertices []Vector3, maxDepth int, stats *Stats) {
	stats.addNode(node.Depth)

	var intersectingFaces []Segitiga
	for _, face := range faces {
		v0 := vertices[face.V1]
		v1 := vertices[face.V2]
		v2 := vertices[face.V3]

		if checkIntersection(node.Box, v0, v1, v2) {
			intersectingFaces = append(intersectingFaces, face)
		}
	}

	if len(intersectingFaces) == 0 {
		stats.addPruned(node.Depth)
		return
	}

	if node.Depth == maxDepth {
		node.IsLeaf = true
		node.IsFilled = true
		stats.addVoxel()
		return
	}

	h := node.Box.HalfDimension / 2.0
	c := node.Box.Center

	offsets := [8]Vector3{
		{X: c.X - h, Y: c.Y - h, Z: c.Z - h}, {X: c.X + h, Y: c.Y - h, Z: c.Z - h},
		{X: c.X - h, Y: c.Y + h, Z: c.Z - h}, {X: c.X + h, Y: c.Y + h, Z: c.Z - h},
		{X: c.X - h, Y: c.Y - h, Z: c.Z + h}, {X: c.X + h, Y: c.Y - h, Z: c.Z + h},
		{X: c.X - h, Y: c.Y + h, Z: c.Z + h}, {X: c.X + h, Y: c.Y + h, Z: c.Z + h},
	}

	var wg sync.WaitGroup

	for i := 0; i < 8; i++ {
		childNode := &OctreeNode{
			Box:   BoundingBox{Center: offsets[i], HalfDimension: h},
			Depth: node.Depth + 1,
		}
		node.Children[i] = childNode
		if node.Depth < 3 {
			wg.Add(1)
			go func(child *OctreeNode) {
				defer wg.Done()
				buildOctree(child, intersectingFaces, vertices, maxDepth, stats)
			}(childNode)
		} else {
			buildOctree(childNode, intersectingFaces, vertices, maxDepth, stats)
		}
	}

	wg.Wait()
}
