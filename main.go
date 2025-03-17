package streamer

func main() {
	core := NewCore()
	streamer := NewStreamer(core)

	streamer.StartAutomaticUpdate()
}
