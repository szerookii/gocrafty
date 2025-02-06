package chunk

import (
	"fmt"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol/packet/packets/play"
	"sync"
)

func GenerateFlatChunk(x, z int32) *Chunk {
	c := NewChunk(x, z)

	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			c.SetBlock(i, 30, j, 7) // bedrock
			c.SetBlock(i, 31, j, 1) // stone
			c.SetBlock(i, 32, j, 1) // stone
			c.SetBlock(i, 33, j, 1) // stone
			c.SetBlock(i, 34, j, 1) // stone
			c.SetBlock(i, 35, j, 2) // grass
		}
	}

	return c
}

type ChunkSection struct {
	blocks [16 * 16 * 16]uint16
}

type Chunk struct {
	mu       *sync.RWMutex
	sections [16]*ChunkSection
	X        int32
	Z        int32

	IsEmpty bool
}

func NewChunk(x int32, z int32) *Chunk {
	return &Chunk{
		mu:      &sync.RWMutex{},
		X:       x,
		Z:       z,
		IsEmpty: true,
	}
}

func getBlockIndex(x int, y int, z int) int {
	if x < 0 || z < 0 || x >= 16 || z >= 16 {
		panic("Coords (x=" + fmt.Sprint(x) + ",z=" + fmt.Sprint(z) + ") out of section bounds")
	}
	return ((y & 0xf) << 8) | (z << 4) | x
}

func (c *Chunk) getSection(y int) *ChunkSection {
	c.mu.RLock()
	sectionId := y >> 4
	section := c.sections[sectionId]
	c.mu.RUnlock()
	if section == nil {
		c.mu.Lock()
		section = &ChunkSection{}
		c.sections[sectionId] = section
		c.mu.Unlock()
	}
	return section
}

func (c *Chunk) SetBlock(x int, y int, z int, blockType uint8) *Chunk {
	section := c.getSection(y)

	c.mu.Lock()
	section.blocks[getBlockIndex(x, y, z)] = uint16(blockType) << 4
	c.IsEmpty = false
	c.mu.Unlock()

	return c
}

func (c *Chunk) SetState(x int, y int, z int, state uint8) {
	if state < 0 || state >= 16 {
		panic(fmt.Sprint(state) + " is not a valid state, should be 0-15.")
	}
	section := c.getSection(y)

	c.mu.Lock()
	blockIndex := getBlockIndex(x, y, z)
	currentBlockData := section.blocks[blockIndex]
	section.blocks[blockIndex] = uint16((currentBlockData & 0xfff0) | uint16(state))
	c.mu.Unlock()
}

func (c *Chunk) Marshal() *play.ChunkData {
	c.mu.RLock()
	chunkData := &play.ChunkData{
		GroundUpContinuous: true,
	}

	var bitMask uint16 = 0
	dataSize := 256
	for i := len(c.sections) - 1; i >= 0; i-- {
		if c.sections[i] == nil {
			bitMask <<= 1
		} else {
			bitMask <<= 1
			bitMask += 1

			BLOCKS_IN_SECTION := 16 * 16 * 16
			dataSize += BLOCKS_IN_SECTION * 5 / 2
			dataSize += BLOCKS_IN_SECTION / 2
		}
	}
	chunkData.PrimaryBitMask = bitMask

	chunkData.ChunkX = c.X
	chunkData.ChunkZ = c.Z

	data := make([]byte, dataSize)
	pos := 0

	for _, section := range c.sections {
		if section == nil {
			continue
		}
		for _, block := range section.blocks {
			data[pos] = byte(block & 0xff)
			data[pos+1] = byte(block >> 8)
			pos += 2
		}

		for i := 0; i < 16*16*16; i++ {
			data[pos+i] = 0xaa // 10 in each 4 bit segment
		}
	}

	chunkData.Data = data

	c.mu.RUnlock()
	return chunkData
}
