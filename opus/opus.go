// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Sat, 02 Nov 2024 01:06:14 JST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package opus

/*
#cgo pkg-config: opus
#include "opus.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// EncoderGetSize function as declared in opus/opus.h:171
func EncoderGetSize(channels int32) int32 {
	cchannels, cchannelsAllocMap := (C.int)(channels), cgoAllocsUnknown
	__ret := C.opus_encoder_get_size(cchannels)
	runtime.KeepAlive(cchannelsAllocMap)
	__v := (int32)(__ret)
	return __v
}

// EncoderCreate function as declared in opus/opus.h:208
func EncoderCreate(fs int32, channels int32, application int32, error *int32) *Encoder {
	cfs, cfsAllocMap := (C.opus_int32)(fs), cgoAllocsUnknown
	cchannels, cchannelsAllocMap := (C.int)(channels), cgoAllocsUnknown
	capplication, capplicationAllocMap := (C.int)(application), cgoAllocsUnknown
	cerror, cerrorAllocMap := (*C.int)(unsafe.Pointer(error)), cgoAllocsUnknown
	__ret := C.opus_encoder_create(cfs, cchannels, capplication, cerror)
	runtime.KeepAlive(cerrorAllocMap)
	runtime.KeepAlive(capplicationAllocMap)
	runtime.KeepAlive(cchannelsAllocMap)
	runtime.KeepAlive(cfsAllocMap)
	__v := *(**Encoder)(unsafe.Pointer(&__ret))
	return __v
}

// EncoderInit function as declared in opus/opus.h:228
func EncoderInit(st *Encoder, fs int32, channels int32, application int32) int32 {
	cst, cstAllocMap := (*C.OpusEncoder)(unsafe.Pointer(st)), cgoAllocsUnknown
	cfs, cfsAllocMap := (C.opus_int32)(fs), cgoAllocsUnknown
	cchannels, cchannelsAllocMap := (C.int)(channels), cgoAllocsUnknown
	capplication, capplicationAllocMap := (C.int)(application), cgoAllocsUnknown
	__ret := C.opus_encoder_init(cst, cfs, cchannels, capplication)
	runtime.KeepAlive(capplicationAllocMap)
	runtime.KeepAlive(cchannelsAllocMap)
	runtime.KeepAlive(cfsAllocMap)
	runtime.KeepAlive(cstAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Encode function as declared in opus/opus.h:263
func Encode(st *Encoder, pcm []int16, frameSize int32, data []byte, maxDataBytes int32) int32 {
	cst, cstAllocMap := (*C.OpusEncoder)(unsafe.Pointer(st)), cgoAllocsUnknown
	cpcm, cpcmAllocMap := (*C.opus_int16)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pcm)).Data)), cgoAllocsUnknown
	cframeSize, cframeSizeAllocMap := (C.int)(frameSize), cgoAllocsUnknown
	cdata, cdataAllocMap := (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&data)).Data)), cgoAllocsUnknown
	cmaxDataBytes, cmaxDataBytesAllocMap := (C.opus_int32)(maxDataBytes), cgoAllocsUnknown
	__ret := C.opus_encode(cst, cpcm, cframeSize, cdata, cmaxDataBytes)
	runtime.KeepAlive(cmaxDataBytesAllocMap)
	runtime.KeepAlive(cdataAllocMap)
	runtime.KeepAlive(cframeSizeAllocMap)
	runtime.KeepAlive(cpcmAllocMap)
	runtime.KeepAlive(cstAllocMap)
	__v := (int32)(__ret)
	return __v
}

// EncodeFloat function as declared in opus/opus.h:304
func EncodeFloat(st *Encoder, pcm []float32, frameSize int32, data []byte, maxDataBytes int32) int32 {
	cst, cstAllocMap := (*C.OpusEncoder)(unsafe.Pointer(st)), cgoAllocsUnknown
	cpcm, cpcmAllocMap := (*C.float)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pcm)).Data)), cgoAllocsUnknown
	cframeSize, cframeSizeAllocMap := (C.int)(frameSize), cgoAllocsUnknown
	cdata, cdataAllocMap := (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&data)).Data)), cgoAllocsUnknown
	cmaxDataBytes, cmaxDataBytesAllocMap := (C.opus_int32)(maxDataBytes), cgoAllocsUnknown
	__ret := C.opus_encode_float(cst, cpcm, cframeSize, cdata, cmaxDataBytes)
	runtime.KeepAlive(cmaxDataBytesAllocMap)
	runtime.KeepAlive(cdataAllocMap)
	runtime.KeepAlive(cframeSizeAllocMap)
	runtime.KeepAlive(cpcmAllocMap)
	runtime.KeepAlive(cstAllocMap)
	__v := (int32)(__ret)
	return __v
}

// EncoderDestroy function as declared in opus/opus.h:315
func EncoderDestroy(st *Encoder) {
	cst, cstAllocMap := (*C.OpusEncoder)(unsafe.Pointer(st)), cgoAllocsUnknown
	C.opus_encoder_destroy(cst)
	runtime.KeepAlive(cstAllocMap)
}

// DecoderGetSize function as declared in opus/opus.h:406
func DecoderGetSize(channels int32) int32 {
	cchannels, cchannelsAllocMap := (C.int)(channels), cgoAllocsUnknown
	__ret := C.opus_decoder_get_size(cchannels)
	runtime.KeepAlive(cchannelsAllocMap)
	__v := (int32)(__ret)
	return __v
}

// DecoderCreate function as declared in opus/opus.h:423
func DecoderCreate(fs int32, channels int32, error *int32) *Decoder {
	cfs, cfsAllocMap := (C.opus_int32)(fs), cgoAllocsUnknown
	cchannels, cchannelsAllocMap := (C.int)(channels), cgoAllocsUnknown
	cerror, cerrorAllocMap := (*C.int)(unsafe.Pointer(error)), cgoAllocsUnknown
	__ret := C.opus_decoder_create(cfs, cchannels, cerror)
	runtime.KeepAlive(cerrorAllocMap)
	runtime.KeepAlive(cchannelsAllocMap)
	runtime.KeepAlive(cfsAllocMap)
	__v := *(**Decoder)(unsafe.Pointer(&__ret))
	return __v
}

// DecoderInit function as declared in opus/opus.h:440
func DecoderInit(st *Decoder, fs int32, channels int32) int32 {
	cst, cstAllocMap := (*C.OpusDecoder)(unsafe.Pointer(st)), cgoAllocsUnknown
	cfs, cfsAllocMap := (C.opus_int32)(fs), cgoAllocsUnknown
	cchannels, cchannelsAllocMap := (C.int)(channels), cgoAllocsUnknown
	__ret := C.opus_decoder_init(cst, cfs, cchannels)
	runtime.KeepAlive(cchannelsAllocMap)
	runtime.KeepAlive(cfsAllocMap)
	runtime.KeepAlive(cstAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Decode function as declared in opus/opus.h:462
func Decode(st *Decoder, data []byte, len int32, pcm []int16, frameSize int32, decodeFec int32) int32 {
	cst, cstAllocMap := (*C.OpusDecoder)(unsafe.Pointer(st)), cgoAllocsUnknown
	cdata, cdataAllocMap := (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&data)).Data)), cgoAllocsUnknown
	clen, clenAllocMap := (C.opus_int32)(len), cgoAllocsUnknown
	cpcm, cpcmAllocMap := (*C.opus_int16)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pcm)).Data)), cgoAllocsUnknown
	cframeSize, cframeSizeAllocMap := (C.int)(frameSize), cgoAllocsUnknown
	cdecodeFec, cdecodeFecAllocMap := (C.int)(decodeFec), cgoAllocsUnknown
	__ret := C.opus_decode(cst, cdata, clen, cpcm, cframeSize, cdecodeFec)
	runtime.KeepAlive(cdecodeFecAllocMap)
	runtime.KeepAlive(cframeSizeAllocMap)
	runtime.KeepAlive(cpcmAllocMap)
	runtime.KeepAlive(clenAllocMap)
	runtime.KeepAlive(cdataAllocMap)
	runtime.KeepAlive(cstAllocMap)
	__v := (int32)(__ret)
	return __v
}

// DecodeFloat function as declared in opus/opus.h:487
func DecodeFloat(st *Decoder, data []byte, len int32, pcm []float32, frameSize int32, decodeFec int32) int32 {
	cst, cstAllocMap := (*C.OpusDecoder)(unsafe.Pointer(st)), cgoAllocsUnknown
	cdata, cdataAllocMap := (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&data)).Data)), cgoAllocsUnknown
	clen, clenAllocMap := (C.opus_int32)(len), cgoAllocsUnknown
	cpcm, cpcmAllocMap := (*C.float)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pcm)).Data)), cgoAllocsUnknown
	cframeSize, cframeSizeAllocMap := (C.int)(frameSize), cgoAllocsUnknown
	cdecodeFec, cdecodeFecAllocMap := (C.int)(decodeFec), cgoAllocsUnknown
	__ret := C.opus_decode_float(cst, cdata, clen, cpcm, cframeSize, cdecodeFec)
	runtime.KeepAlive(cdecodeFecAllocMap)
	runtime.KeepAlive(cframeSizeAllocMap)
	runtime.KeepAlive(cpcmAllocMap)
	runtime.KeepAlive(clenAllocMap)
	runtime.KeepAlive(cdataAllocMap)
	runtime.KeepAlive(cstAllocMap)
	__v := (int32)(__ret)
	return __v
}

// DecoderDestroy function as declared in opus/opus.h:512
func DecoderDestroy(st *Decoder) {
	cst, cstAllocMap := (*C.OpusDecoder)(unsafe.Pointer(st)), cgoAllocsUnknown
	C.opus_decoder_destroy(cst)
	runtime.KeepAlive(cstAllocMap)
}

// PacketGetBandwidth function as declared in opus/opus.h:545
func PacketGetBandwidth(data *byte) int32 {
	cdata, cdataAllocMap := (*C.uchar)(unsafe.Pointer(data)), cgoAllocsUnknown
	__ret := C.opus_packet_get_bandwidth(cdata)
	runtime.KeepAlive(cdataAllocMap)
	__v := (int32)(__ret)
	return __v
}

// PacketGetSamplesPerFrame function as declared in opus/opus.h:556
func PacketGetSamplesPerFrame(data *byte, fs int32) int32 {
	cdata, cdataAllocMap := (*C.uchar)(unsafe.Pointer(data)), cgoAllocsUnknown
	cfs, cfsAllocMap := (C.opus_int32)(fs), cgoAllocsUnknown
	__ret := C.opus_packet_get_samples_per_frame(cdata, cfs)
	runtime.KeepAlive(cfsAllocMap)
	runtime.KeepAlive(cdataAllocMap)
	__v := (int32)(__ret)
	return __v
}

// PacketGetNbChannels function as declared in opus/opus.h:563
func PacketGetNbChannels(data *byte) int32 {
	cdata, cdataAllocMap := (*C.uchar)(unsafe.Pointer(data)), cgoAllocsUnknown
	__ret := C.opus_packet_get_nb_channels(cdata)
	runtime.KeepAlive(cdataAllocMap)
	__v := (int32)(__ret)
	return __v
}

// PacketGetNbFrames function as declared in opus/opus.h:572
func PacketGetNbFrames(packet *byte, len int32) int32 {
	cpacket, cpacketAllocMap := (*C.uchar)(unsafe.Pointer(packet)), cgoAllocsUnknown
	clen, clenAllocMap := (C.opus_int32)(len), cgoAllocsUnknown
	__ret := C.opus_packet_get_nb_frames(cpacket, clen)
	runtime.KeepAlive(clenAllocMap)
	runtime.KeepAlive(cpacketAllocMap)
	__v := (int32)(__ret)
	return __v
}

// PacketGetNbSamples function as declared in opus/opus.h:584
func PacketGetNbSamples(packet *byte, len int32, fs int32) int32 {
	cpacket, cpacketAllocMap := (*C.uchar)(unsafe.Pointer(packet)), cgoAllocsUnknown
	clen, clenAllocMap := (C.opus_int32)(len), cgoAllocsUnknown
	cfs, cfsAllocMap := (C.opus_int32)(fs), cgoAllocsUnknown
	__ret := C.opus_packet_get_nb_samples(cpacket, clen, cfs)
	runtime.KeepAlive(cfsAllocMap)
	runtime.KeepAlive(clenAllocMap)
	runtime.KeepAlive(cpacketAllocMap)
	__v := (int32)(__ret)
	return __v
}

// DecoderGetNbSamples function as declared in opus/opus.h:594
func DecoderGetNbSamples(dec *Decoder, packet []byte, len int32) int32 {
	cdec, cdecAllocMap := (*C.OpusDecoder)(unsafe.Pointer(dec)), cgoAllocsUnknown
	cpacket, cpacketAllocMap := (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&packet)).Data)), cgoAllocsUnknown
	clen, clenAllocMap := (C.opus_int32)(len), cgoAllocsUnknown
	__ret := C.opus_decoder_get_nb_samples(cdec, cpacket, clen)
	runtime.KeepAlive(clenAllocMap)
	runtime.KeepAlive(cpacketAllocMap)
	runtime.KeepAlive(cdecAllocMap)
	__v := (int32)(__ret)
	return __v
}

// PcmSoftClip function as declared in opus/opus.h:606
func PcmSoftClip(pcm []float32, frameSize int32, channels int32, softclipMem []float32) {
	cpcm, cpcmAllocMap := (*C.float)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pcm)).Data)), cgoAllocsUnknown
	cframeSize, cframeSizeAllocMap := (C.int)(frameSize), cgoAllocsUnknown
	cchannels, cchannelsAllocMap := (C.int)(channels), cgoAllocsUnknown
	csoftclipMem, csoftclipMemAllocMap := (*C.float)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&softclipMem)).Data)), cgoAllocsUnknown
	C.opus_pcm_soft_clip(cpcm, cframeSize, cchannels, csoftclipMem)
	runtime.KeepAlive(csoftclipMemAllocMap)
	runtime.KeepAlive(cchannelsAllocMap)
	runtime.KeepAlive(cframeSizeAllocMap)
	runtime.KeepAlive(cpcmAllocMap)
}

// RepacketizerGetSize function as declared in opus/opus.h:756
func RepacketizerGetSize() int32 {
	__ret := C.opus_repacketizer_get_size()
	__v := (int32)(__ret)
	return __v
}

// RepacketizerInit function as declared in opus/opus.h:775
func RepacketizerInit(rp *Repacketizer) *Repacketizer {
	crp, crpAllocMap := (*C.OpusRepacketizer)(unsafe.Pointer(rp)), cgoAllocsUnknown
	__ret := C.opus_repacketizer_init(crp)
	runtime.KeepAlive(crpAllocMap)
	__v := *(**Repacketizer)(unsafe.Pointer(&__ret))
	return __v
}

// RepacketizerCreate function as declared in opus/opus.h:780
func RepacketizerCreate() *Repacketizer {
	__ret := C.opus_repacketizer_create()
	__v := *(**Repacketizer)(unsafe.Pointer(&__ret))
	return __v
}

// RepacketizerDestroy function as declared in opus/opus.h:786
func RepacketizerDestroy(rp *Repacketizer) {
	crp, crpAllocMap := (*C.OpusRepacketizer)(unsafe.Pointer(rp)), cgoAllocsUnknown
	C.opus_repacketizer_destroy(crp)
	runtime.KeepAlive(crpAllocMap)
}

// RepacketizerCat function as declared in opus/opus.h:835
func RepacketizerCat(rp *Repacketizer, data []byte, len int32) int32 {
	crp, crpAllocMap := (*C.OpusRepacketizer)(unsafe.Pointer(rp)), cgoAllocsUnknown
	cdata, cdataAllocMap := (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&data)).Data)), cgoAllocsUnknown
	clen, clenAllocMap := (C.opus_int32)(len), cgoAllocsUnknown
	__ret := C.opus_repacketizer_cat(crp, cdata, clen)
	runtime.KeepAlive(clenAllocMap)
	runtime.KeepAlive(cdataAllocMap)
	runtime.KeepAlive(crpAllocMap)
	__v := (int32)(__ret)
	return __v
}

// RepacketizerOutRange function as declared in opus/opus.h:869
func RepacketizerOutRange(rp *Repacketizer, begin int32, end int32, data []byte, maxlen int32) int32 {
	crp, crpAllocMap := (*C.OpusRepacketizer)(unsafe.Pointer(rp)), cgoAllocsUnknown
	cbegin, cbeginAllocMap := (C.int)(begin), cgoAllocsUnknown
	cend, cendAllocMap := (C.int)(end), cgoAllocsUnknown
	cdata, cdataAllocMap := (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&data)).Data)), cgoAllocsUnknown
	cmaxlen, cmaxlenAllocMap := (C.opus_int32)(maxlen), cgoAllocsUnknown
	__ret := C.opus_repacketizer_out_range(crp, cbegin, cend, cdata, cmaxlen)
	runtime.KeepAlive(cmaxlenAllocMap)
	runtime.KeepAlive(cdataAllocMap)
	runtime.KeepAlive(cendAllocMap)
	runtime.KeepAlive(cbeginAllocMap)
	runtime.KeepAlive(crpAllocMap)
	__v := (int32)(__ret)
	return __v
}

// RepacketizerGetNbFrames function as declared in opus/opus.h:881
func RepacketizerGetNbFrames(rp *Repacketizer) int32 {
	crp, crpAllocMap := (*C.OpusRepacketizer)(unsafe.Pointer(rp)), cgoAllocsUnknown
	__ret := C.opus_repacketizer_get_nb_frames(crp)
	runtime.KeepAlive(crpAllocMap)
	__v := (int32)(__ret)
	return __v
}

// RepacketizerOut function as declared in opus/opus.h:912
func RepacketizerOut(rp *Repacketizer, data []byte, maxlen int32) int32 {
	crp, crpAllocMap := (*C.OpusRepacketizer)(unsafe.Pointer(rp)), cgoAllocsUnknown
	cdata, cdataAllocMap := (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&data)).Data)), cgoAllocsUnknown
	cmaxlen, cmaxlenAllocMap := (C.opus_int32)(maxlen), cgoAllocsUnknown
	__ret := C.opus_repacketizer_out(crp, cdata, cmaxlen)
	runtime.KeepAlive(cmaxlenAllocMap)
	runtime.KeepAlive(cdataAllocMap)
	runtime.KeepAlive(crpAllocMap)
	__v := (int32)(__ret)
	return __v
}

// PacketPad function as declared in opus/opus.h:926
func PacketPad(data []byte, len int32, newLen int32) int32 {
	cdata, cdataAllocMap := (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&data)).Data)), cgoAllocsUnknown
	clen, clenAllocMap := (C.opus_int32)(len), cgoAllocsUnknown
	cnewLen, cnewLenAllocMap := (C.opus_int32)(newLen), cgoAllocsUnknown
	__ret := C.opus_packet_pad(cdata, clen, cnewLen)
	runtime.KeepAlive(cnewLenAllocMap)
	runtime.KeepAlive(clenAllocMap)
	runtime.KeepAlive(cdataAllocMap)
	__v := (int32)(__ret)
	return __v
}

// PacketUnpad function as declared in opus/opus.h:939
func PacketUnpad(data []byte, len int32) int32 {
	cdata, cdataAllocMap := (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&data)).Data)), cgoAllocsUnknown
	clen, clenAllocMap := (C.opus_int32)(len), cgoAllocsUnknown
	__ret := C.opus_packet_unpad(cdata, clen)
	runtime.KeepAlive(clenAllocMap)
	runtime.KeepAlive(cdataAllocMap)
	__v := (int32)(__ret)
	return __v
}

// MultistreamPacketPad function as declared in opus/opus.h:955
func MultistreamPacketPad(data []byte, len int32, newLen int32, nbStreams int32) int32 {
	cdata, cdataAllocMap := (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&data)).Data)), cgoAllocsUnknown
	clen, clenAllocMap := (C.opus_int32)(len), cgoAllocsUnknown
	cnewLen, cnewLenAllocMap := (C.opus_int32)(newLen), cgoAllocsUnknown
	cnbStreams, cnbStreamsAllocMap := (C.int)(nbStreams), cgoAllocsUnknown
	__ret := C.opus_multistream_packet_pad(cdata, clen, cnewLen, cnbStreams)
	runtime.KeepAlive(cnbStreamsAllocMap)
	runtime.KeepAlive(cnewLenAllocMap)
	runtime.KeepAlive(clenAllocMap)
	runtime.KeepAlive(cdataAllocMap)
	__v := (int32)(__ret)
	return __v
}

// MultistreamPacketUnpad function as declared in opus/opus.h:970
func MultistreamPacketUnpad(data []byte, len int32, nbStreams int32) int32 {
	cdata, cdataAllocMap := (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&data)).Data)), cgoAllocsUnknown
	clen, clenAllocMap := (C.opus_int32)(len), cgoAllocsUnknown
	cnbStreams, cnbStreamsAllocMap := (C.int)(nbStreams), cgoAllocsUnknown
	__ret := C.opus_multistream_packet_unpad(cdata, clen, cnbStreams)
	runtime.KeepAlive(cnbStreamsAllocMap)
	runtime.KeepAlive(clenAllocMap)
	runtime.KeepAlive(cdataAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Strerror function as declared in opus/opus_defines.h:713
func Strerror(error int32) *byte {
	cerror, cerrorAllocMap := (C.int)(error), cgoAllocsUnknown
	__ret := C.opus_strerror(cerror)
	runtime.KeepAlive(cerrorAllocMap)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// GetVersionString function as declared in opus/opus_defines.h:719
func GetVersionString() *byte {
	__ret := C.opus_get_version_string()
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}
