package libretro // Core is an instance of a dynamically loaded libretro core

import "unsafe"

// Core is an instance of a dynamically loaded libretro core
type Core struct {
	handle DlHandle

	SymRetroInit                    unsafe.Pointer
	SymRetroDeinit                  unsafe.Pointer
	SymRetroAPIVersion              unsafe.Pointer
	SymRetroGetSystemInfo           unsafe.Pointer
	SymRetroGetSystemAVInfo         unsafe.Pointer
	SymRetroSetEnvironment          unsafe.Pointer
	SymRetroSetVideoRefresh         unsafe.Pointer
	SymRetroSetControllerPortDevice unsafe.Pointer
	SymRetroSetInputPoll            unsafe.Pointer
	SymRetroSetInputState           unsafe.Pointer
	SymRetroSetAudioSample          unsafe.Pointer
	SymRetroSetAudioSampleBatch     unsafe.Pointer
	SymRetroRun                     unsafe.Pointer
	SymRetroReset                   unsafe.Pointer
	SymRetroLoadGame                unsafe.Pointer
	SymRetroUnloadGame              unsafe.Pointer
	SymRetroSerializeSize           unsafe.Pointer
	SymRetroSerialize               unsafe.Pointer
	SymRetroUnserialize             unsafe.Pointer
	SymRetroGetMemorySize           unsafe.Pointer
	SymRetroGetMemoryData           unsafe.Pointer

	AudioCallback       *AudioCallback
	FrameTimeCallback   *FrameTimeCallback
	DiskControlCallback *DiskControlCallback

	MemoryMap []MemoryDescriptor
}
