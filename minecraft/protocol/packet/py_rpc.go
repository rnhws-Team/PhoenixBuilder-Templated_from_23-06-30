package packet

import (
	//"bytes"
	//"fmt"
	"phoenixbuilder/minecraft/protocol"
)


type PyRpc struct {
	Content []byte
}

// ID ...
func (*PyRpc) ID() uint32 {
	return IDPyRpc
}

// Marshal ...
func (pk *PyRpc) Marshal(w *protocol.Writer) {
	w.ByteSlice(&pk.Content)
	w.Bytes(&[]byte{0xae,0x23,0xdb,0x05})
	//fmt.Printf("%d\n",len(pk.Content))
	//fmt.Printf("%X\n",buf.Bytes())
	//_ = protocol.WriteByteSlice(buf, pk.Content)
	//_ = protocol.WriteVaruint32(buf, pk.DisallowBatching)
	//buf.Write([]byte{0xae,0x23,0xdb,0x05})
	//fmt.Printf("%X\n\n",buf.Bytes())
	//var outuint32 uint32
	//protocol.Varuint32(bytes.NewBuffer([]byte{0xae,0x23,0xdb,0x05}),&outuint32)
	//fmt.Printf("%d\n",outuint32)
	
}

// Unmarshal ...
func (pk *PyRpc) Unmarshal(r *protocol.Reader) {
	r.ByteSlice(&pk.Content)
	/*var bt byte
	var bt2 byte
	var bt3 byte
	var bt4 byte
	r.Uint8(&bt)
	r.Uint8(&bt2)
	r.Uint8(&bt3)
	r.Uint8(&bt4)
	fmt.Printf("Addi:%X %X %X %X %X\n", bt, bt2, bt3, bt4)*/
}