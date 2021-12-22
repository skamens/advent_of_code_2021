package main

import (
	"aoc21/aocutil"
	"fmt"
	"math"
)

type packet struct {
	version      int
	typeID       int
	value        int
	lengthTypeID int
	subPackets   []packet
}

func binaryToInt(bstr string) int {
	result := 0

	for _, b := range bstr {
		result = result << 1
		if b == '1' {
			result = result | 1
		}
	}
	return result
}

func parseLiteralValue(str *string) int {
	var binstr string

	for {
		indicator := (*str)[0]
		binstr = binstr + (*str)[1:5]
		(*str) = (*str)[5:]
		if indicator == '0' {
			break
		}
	}

	return binaryToInt(binstr)
}

func readPacket(binaryString *string) packet {
	var thePacket packet

	// First grab the version

	thePacket.version = binaryToInt((*binaryString)[0:3])
	*binaryString = (*binaryString)[3:]

	// Next get the type
	thePacket.typeID = binaryToInt((*binaryString)[0:3])
	*binaryString = (*binaryString)[3:]

	thePacket.value = -1

	if thePacket.typeID == 4 {
		thePacket.value = parseLiteralValue(binaryString)
		return thePacket
	}

	lengthTypeID := binaryToInt((*binaryString)[0:1])
	*binaryString = (*binaryString)[1:]

	if lengthTypeID == 0 {
		subPacketLength := binaryToInt((*binaryString)[0:15])
		*binaryString = (*binaryString)[15:]
		subPacket := (*binaryString)[0:subPacketLength]
		*binaryString = (*binaryString)[subPacketLength:]

		for len(subPacket) > 0 {
			thePacket.subPackets = append(thePacket.subPackets, readPacket(&subPacket))
			fmt.Printf("Version %d, type %d, literal %d\n",
				thePacket.subPackets[len(thePacket.subPackets)-1].version,
				thePacket.subPackets[len(thePacket.subPackets)-1].typeID,
				thePacket.subPackets[len(thePacket.subPackets)-1].value)
		}
	} else {
		// This means the next 11 bits tells us how many subpackets there are
		numSubPackets := binaryToInt((*binaryString)[0:11])
		*binaryString = (*binaryString)[11:]

		for i := 0; i < numSubPackets; i++ {
			thePacket.subPackets = append(thePacket.subPackets, readPacket(binaryString))
			fmt.Printf("Version %d, type %d, literal %d\n",
				thePacket.subPackets[i].version,
				thePacket.subPackets[i].typeID,
				thePacket.subPackets[i].value)
		}
	}

	return thePacket
}

func addVersions(p packet) int {

	total := p.version

	for _, sp := range p.subPackets {
		total += addVersions(sp)
	}
	fmt.Printf("addVersion: myversion=%d, total=%d\n", p.version, total)
	return total
}

func packetValue(p packet) int {
	switch p.typeID {
	case 4:
		return p.value
	case 0:
		p.value = 0
		for _, sp := range p.subPackets {
			p.value += packetValue(sp)
		}
	case 1:
		p.value = 1
		for _, sp := range p.subPackets {
			p.value *= packetValue(sp)
		}
		return p.value
	case 2:
		p.value = math.MaxInt64
		for _, sp := range p.subPackets {
			v := packetValue(sp)
			if v < p.value {
				p.value = v
			}
		}
	case 3:
		p.value = math.MinInt64
		for _, sp := range p.subPackets {
			v := packetValue(sp)
			if v > p.value {
				p.value = v
			}
		}
	case 5:
		if packetValue(p.subPackets[0]) > packetValue(p.subPackets[1]) {
			p.value = 1
		} else {
			p.value = 0
		}
	case 6:
		if packetValue(p.subPackets[0]) < packetValue(p.subPackets[1]) {
			p.value = 1
		} else {
			p.value = 0
		}
	case 7:
		if packetValue(p.subPackets[0]) == packetValue(p.subPackets[1]) {
			p.value = 1
		} else {
			p.value = 0
		}
	}

	return p.value
}

func main() {

	input := "input16.txt"

	lines := aocutil.LoadStringArray(input)

	hexMap := map[byte]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111"}

	var binaryString string

	// I know there's only one line, but this is pretty :-)
	for _, l := range lines {
		for b := range l {
			binaryString = binaryString + hexMap[l[b]]
		}
	}

	thePacket := readPacket(&binaryString)

	// Now add up all the version numbers
	fmt.Printf("PacketValue: %d\n", packetValue(thePacket))

}
