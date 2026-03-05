//go:build gomock || generate

package quic

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/MahamShakir/quic-go-patch -destination mock_send_conn_test.go github.com/MahamShakir/quic-go-patch SendConn"
type SendConn = sendConn

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/MahamShakir/quic-go-patch -destination mock_raw_conn_test.go github.com/MahamShakir/quic-go-patch RawConn"
type RawConn = rawConn

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/MahamShakir/quic-go-patch -destination mock_sender_test.go github.com/MahamShakir/quic-go-patch Sender"
type Sender = sender

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/MahamShakir/quic-go-patch -destination mock_stream_sender_test.go github.com/MahamShakir/quic-go-patch StreamSender"
type StreamSender = streamSender

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/MahamShakir/quic-go-patch -destination mock_stream_control_frame_getter_test.go github.com/MahamShakir/quic-go-patch StreamControlFrameGetter"
type StreamControlFrameGetter = streamControlFrameGetter

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/MahamShakir/quic-go-patch -destination mock_stream_frame_getter_test.go github.com/MahamShakir/quic-go-patch StreamFrameGetter"
type StreamFrameGetter = streamFrameGetter

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/MahamShakir/quic-go-patch -destination mock_frame_source_test.go github.com/MahamShakir/quic-go-patch FrameSource"
type FrameSource = frameSource

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/MahamShakir/quic-go-patch -destination mock_ack_frame_source_test.go github.com/MahamShakir/quic-go-patch AckFrameSource"
type AckFrameSource = ackFrameSource

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/MahamShakir/quic-go-patch -destination mock_sealing_manager_test.go github.com/MahamShakir/quic-go-patch SealingManager"
type SealingManager = sealingManager

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/MahamShakir/quic-go-patch -destination mock_unpacker_test.go github.com/MahamShakir/quic-go-patch Unpacker"
type Unpacker = unpacker

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/MahamShakir/quic-go-patch -destination mock_packer_test.go github.com/MahamShakir/quic-go-patch Packer"
type Packer = packer

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/MahamShakir/quic-go-patch -destination mock_mtu_discoverer_test.go github.com/MahamShakir/quic-go-patch MTUDiscoverer"
type MTUDiscoverer = mtuDiscoverer

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/MahamShakir/quic-go-patch -destination mock_conn_runner_test.go github.com/MahamShakir/quic-go-patch ConnRunner"
type ConnRunner = connRunner

//go:generate sh -c "go tool mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/MahamShakir/quic-go-patch -destination mock_packet_handler_test.go github.com/MahamShakir/quic-go-patch PacketHandler"
type PacketHandler = packetHandler

//go:generate sh -c "go tool mockgen -typed -package quic -self_package github.com/MahamShakir/quic-go-patch -self_package github.com/MahamShakir/quic-go-patch -destination mock_packetconn_test.go net PacketConn"
